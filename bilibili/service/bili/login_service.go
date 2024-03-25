package bili

import (
	"encoding/json"
	"fmt"
	"github.com/alice52/archive/bilibili/errors"
	"github.com/alice52/archive/bilibili/util"
	"github.com/skip2/go-qrcode"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	generateQrCodeUrl = "https://passport.bilibili.com/x/passport-login/web/qrcode/generate"
	pollQrCodeUrl     = "https://passport.bilibili.com/x/passport-login/web/qrcode/poll"
	navInfoUrl        = "https://api.bilibili.com/x/web-interface/nav"
)

// LoginStatus https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/login/login_action/QR.md#%E6%89%AB%E7%A0%81%E7%99%BB%E5%BD%95web%E7%AB%AF
type LoginStatus int
type LoginResp struct {
	LoginStatus LoginStatus
	Cookie      []string
}

var (
	LoginSuccess           = LoginStatus(0)
	LoginNotScan           = LoginStatus(86101)
	LoginScanButNotConfirm = LoginStatus(86090)
	LoginExpired           = LoginStatus(86038)
)

func (client *BClient) doLoginIfNecessary() error {
	if client.isLogin() {
		return nil
	}

	resps, err := client.loginWithQrCode(os.Stdout)
	if err != nil {
		return err
	}

	for resp := range resps {
		switch resp.LoginStatus {
		case LoginSuccess:
			client.SetCookie(resp.Cookie)
			if err = util.SaveCookieFile(resp.Cookie); err != nil {
				return err
			}
		case LoginExpired:
			return fmt.Errorf("login qrcode expired")
		default:
			continue
		}
	}

	return fmt.Errorf("no login exception")
}

// LoginWithQrCode writer is where the qrcode be written
func (client *BClient) loginWithQrCode(writer io.Writer) (<-chan LoginResp, error) {
	generateQrCodeResp, err := client.GenerateQrcode()
	if err != nil {
		return nil, err
	}

	if err = GenerateAndEmail(generateQrCodeResp.Data.URL, qrcode.Low, writer); err != nil {
		return nil, err
	}

	var loginResp = make(chan LoginResp)
	go func() {
		defer close(loginResp)
		var (
			pollQrCodeResp *PollQrCodeResp
			respHeader     http.Header
		)
		for {
			pollQrCodeResp, respHeader, err = client.PollQrcode(generateQrCodeResp.Data.QrcodeKey)
			if err != nil {
				loginResp <- LoginResp{
					LoginStatus: LoginStatus(-1),
					Cookie:      nil,
				}
				return
			}
			loginResp <- LoginResp{
				LoginStatus: LoginStatus(pollQrCodeResp.Data.Code),
				Cookie:      respHeader.Values("Set-Cookie"),
			}
			switch pollQrCodeResp.Data.Code {
			case int(LoginSuccess), int(LoginExpired):
				return
			default:
				continue
			}
		}
	}()
	return loginResp, nil
}

func (client *BClient) isLogin() bool {
	cookie, err := util.ReadCookieFromFile()
	if err != nil {
		return false
	}
	client.SetCookie(cookie)
	info, err := client.NavInfo()
	if err != nil {
		return false
	}
	return info.Data.IsLogin
}

type NavInfoResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		IsLogin bool `json:"isLogin"`
	} `json:"data"`
}

func (client *BClient) NavInfo() (*NavInfoResp, error) {
	client.HttpClient = &http.Client{}
	request, err := client.newCookieRequest(http.MethodGet, navInfoUrl, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.ErrUnexpectedStatusCode(resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	navInfoResp := &NavInfoResp{}
	if err = json.Unmarshal(body, navInfoResp); err != nil {
		return nil, err
	}

	if navInfoResp.Code != 0 {
		return nil, errors.StatusError{Code: navInfoResp.Code, Cause: navInfoResp.Message}
	}

	return navInfoResp, nil
}
