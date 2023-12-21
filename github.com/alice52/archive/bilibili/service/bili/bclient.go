package bili

import (
	"io"
	"net/http"
	"time"
)

// LogonClient maintain logon status by scheduler job
var LogonClient = new(BClient)

type BClient struct {
	HttpClient *http.Client
	cookie     []string
}

func GetLogonClientV2() (*BClient, error) {
	err := LogonClient.doLoginIfNecessary()
	if err != nil {
		return nil, err
	}

	return LogonClient, nil
}

func GetLogonClient() (*BClient, error) {
	err := LogonClient.doLoginIfNecessary()
	if err == nil {
		return LogonClient, nil
	}

	time.Sleep(5 * time.Minute)
	return GetLogonClient()
}

func (client *BClient) newCookieRequest(method, url string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	for _, c := range client.cookie {
		request.Header.Add("Cookie", c)
	}
	return request, nil
}

func (client *BClient) SetCookie(cookie []string) {
	client.cookie = cookie
}
