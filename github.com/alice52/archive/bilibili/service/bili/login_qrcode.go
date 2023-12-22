package bili

import (
	"encoding/json"
	"fmt"
	"github.com/alice52/archive/bilibili/errors"
	"github.com/alice52/archive/bilibili/service/system"
	"github.com/alice52/archive/common/global"
	"github.com/skip2/go-qrcode"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var emailService = system.SystemService.EmailService

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
	time.Sleep(30 * time.Second)
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

func GenerateAndEmail(content string, level qrcode.RecoveryLevel, writer io.Writer) error {

	code, _ := qrcode.New(content, level)
	_, err := writer.Write([]byte(code.ToSmallString(false)))
	if err != nil {
		return err
	}

	image := "qr.png"
	err = qrcode.WriteFile(content, qrcode.Medium, 256, image)
	if err != nil {
		fmt.Println("Error encoding PNG:", err)
		return err
	}

	defer func(im string) {
		if os.Remove(im) != nil {
			fmt.Println(err)
			return
		}
	}(image)
	err = emailService.SendAttach(global.CONFIG.Email.To, "Login QrCode", image)
	if err != nil {
		return err
	}

	return nil
}
