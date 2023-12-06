package bili

import (
	"encoding/json"
	"fmt"
	"github.com/alice52/archive/bilibili/errors"
	"github.com/alice52/archive/bilibili/service"
	"github.com/skip2/go-qrcode"
	"github.com/sourcegraph/conc/panics"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	generateQrCodeUrl = "https://passport.bilibili.com/x/passport-login/web/qrcode/generate"
	pollQrCodeUrl     = "https://passport.bilibili.com/x/passport-login/web/qrcode/poll"
	navInfoUrl        = "https://api.bilibili.com/x/web-interface/nav"
)

const (
	Low RecoveryLevel = iota
	Medium
	High
	Highest
)

var emailService = service.ServiceGroupApp.SystemServiceGroup.EmailService

type RecoveryLevel = qrcode.RecoveryLevel

type GenerateQrCodeResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		URL       string `json:"url"`
		QrcodeKey string `json:"qrcode_key"`
	} `json:"data"`
}

type PollQrCodeResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		URL          string `json:"url"`
		RefreshToken string `json:"refresh_token"`
		Timestamp    int    `json:"timestamp"`
		Code         int    `json:"code"`
		Message      string `json:"message"`
	} `json:"data"`
}

type NavInfoResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		IsLogin bool `json:"isLogin"`
	} `json:"data"`
}

// GenerateQrcode https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/login/login_action/QR.md
func (client *BClient) GenerateQrcode() (*GenerateQrCodeResp, error) {
	client.HttpClient = &http.Client{}
	resp, err := client.HttpClient.Get(generateQrCodeUrl)
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

	generateQrCodeResp := &GenerateQrCodeResp{}
	if err = json.Unmarshal(body, generateQrCodeResp); err != nil {
		return nil, err
	}

	if generateQrCodeResp.Code != 0 {
		return nil, errors.StatusError{Code: generateQrCodeResp.Code, Cause: generateQrCodeResp.Message}
	}
	return generateQrCodeResp, nil
}

// PollQrcode https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/login/login_action/QR.md
func (client *BClient) PollQrcode(qrcode string) (*PollQrCodeResp, http.Header, error) {
	client.HttpClient = &http.Client{}

	url := fmt.Sprintf("%s?qrcode_key=%s", pollQrCodeUrl, qrcode)
	resp, err := client.HttpClient.Get(url)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, nil, errors.ErrUnexpectedStatusCode(resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	pollQrCodeResp := &PollQrCodeResp{}
	if err = json.Unmarshal(body, pollQrCodeResp); err != nil {
		return nil, nil, err
	}

	if pollQrCodeResp.Code != 0 {
		return nil, nil, errors.StatusError{Code: pollQrCodeResp.Code, Cause: pollQrCodeResp.Message}
	}
	return pollQrCodeResp, resp.Header, nil
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

// LoginWithQrCode writer is where the qrcode be written
func (client *BClient) LoginWithQrCode(writer io.Writer) (<-chan LoginResp, error) {
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

func GenerateAndEmail(content string, level RecoveryLevel, writer io.Writer) error {
	var pc panics.Catcher

	pc.Try(func() {
		code, _ := qrcode.New(content, level)
		writer.Write([]byte(code.ToSmallString(false)))
		emailService.EmailTest()
	})

	return pc.Recovered().AsError()
}
