package api

import (
	"encoding/json"
	"errors"
	"fmt"
	c "github.com/alice52/archive/bili/api/common"
	"github.com/alice52/archive/bili/api/errs"
	m "github.com/alice52/archive/bili/api/model"
	"github.com/alice52/archive/common/global"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// LogonClient maintain logon status by scheduler job
var LogonClient = new(BClient)
var logonFunc = MustLogonClient

func init() {
	LogonClient, _ = GetLogonClient()
}

type BClient struct {
	HttpClient *http.Client
	cookie     []string
}

// Deprecated: use GetLogonClient
func GetLogonClientV2() (*BClient, error) {
	if err := LogonClient.doLoginIfNecessary(); err != nil {
		return nil, err
	}

	return LogonClient, nil
}

func MustLogonClient() *BClient {
	if client, err := GetLogonClient(); err != nil {
		panic(err)
	} else {
		return client
	}
}

func GetLogonClient() (*BClient, error) {
	if err := LogonClient.doLoginIfNecessary(); err == nil {
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

func (client *BClient) DoGet(url string) ([]byte, error) {
	request, err := logonFunc().newCookieRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errs.ErrUnexpectedStatusCode(resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}

func (client *BClient) GetList(url string) ([]byte, error) {

	if b, err := client.DoGet(url); err != nil {
		return nil, errors.New("read body error")
	} else {
		ss := &m.BPCResp{}
		if err = json.Unmarshal(b, ss); err != nil {
			return nil, err
		}
		if ss.Code != 0 {
			return nil, errs.StatusError{Code: ss.Code, Cause: ss.Message}
		}
		global.LOG.Info(string(b))

		return b, err
	}
}
func (client *BClient) GetListAll4SpaceVideo(url string) (*m.BResp[m.BRListOfItem[json.RawMessage]], error) {

	bpr := &m.BResp[m.BRListOfItem[json.RawMessage]]{}

	hasMore := true
	offset := ""
	for hasMore {
		var u string
		if len(offset) != 0 {
			u = fmt.Sprintf("%s&offset=%s", url, offset)
		} else {
			u = url
		}

		b, err := client.GetList(u)
		if err != nil {
			return bpr, err
		}
		ss := &m.BResp[m.BRListOfItem[json.RawMessage]]{}
		_ = json.Unmarshal(b, ss)
		if !ss.Data.HasMore {
			hasMore = false
		}

		bpr.Data.Items = append(bpr.Data.Items, ss.Data.Items...)
		offset = ss.Data.Offset
	}

	return bpr, nil
}

func (client *BClient) GetListAll4FavOfFolder(url string) (*m.BResp[m.BRList[json.RawMessage]], error) {

	bpr := &m.BResp[m.BRList[json.RawMessage]]{}

	hasMore := true
	pn := 1
	for hasMore {
		b, err := client.GetList(fmt.Sprintf("%s&pn=%d&ps=%d", url, pn, c.PageSizeMin))
		if err != nil {
			return bpr, err
		}
		ss := &m.BResp[m.BRList[json.RawMessage]]{}
		_ = json.Unmarshal(b, ss)
		if !ss.Data.HasMore {
			hasMore = false
		}

		bpr.Data.Medias = append(bpr.Data.Medias, ss.Data.Medias...)
		bpr.Data.Info = ss.Data.Info
		pn = pn + 1
	}

	return bpr, nil
}

func (client *BClient) GetP(url string) ([]byte, error) {

	if b, err := client.DoGet(fmt.Sprintf("%s&ps=%d", url, c.PageSizeMax)); err != nil {
		return nil, errors.New("read body error")
	} else {
		ss := &m.BPCResp{}
		if err = json.Unmarshal(b, ss); err != nil {
			return nil, err
		}
		if ss.Code != 0 {
			return nil, errs.StatusError{Code: ss.Code, Cause: ss.Message}
		}
		global.LOG.Info(string(b))

		return b, err
	}
}

func (client *BClient) GetAllP(url string) (*m.BPResp[json.RawMessage], error) {

	bpr := &m.BPResp[json.RawMessage]{}

	hasMore := true
	pn := 1
	for hasMore {
		b, err := client.GetP(fmt.Sprintf("%s&pn=%d", url, pn))
		if err != nil {
			return bpr, err
		}
		ss := &m.BPResp[json.RawMessage]{}
		_ = json.Unmarshal(b, ss)
		if len(ss.Data) != c.PageSizeMax {
			hasMore = false
		}

		bpr.Data = append(bpr.Data, ss.Data...)
		pn = pn + 1
	}

	return bpr, nil
}
