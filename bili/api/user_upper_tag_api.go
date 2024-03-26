package api

import (
	"encoding/json"
	"github.com/alice52/archive/bili/api/errs"
	m "github.com/alice52/archive/bili/api/model"
)

var (
	upperTagUrl = "https://api.bilibili.com/x/relation/tags?"
)

//type UpperTag struct {
//	Code    int64   `json:"code"`
//	Data    []UpperTag `json:"data"`
//	Message string  `json:"message"`
//	TTL     int64   `json:"ttl"`
//}

type UpperTag struct {
	Tagid int64  `json:"tagid"`
	Name  string `json:"name"`
	Count int64  `json:"count"`
	Tip   string `json:"tip"`
}

// MyUppersTags 返回关注UP主的标签
// https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/docs/user/relation.md#%E6%9F%A5%E8%AF%A2%E5%85%B3%E6%B3%A8%E5%88%86%E7%BB%84%E5%88%97%E8%A1%A8
func (client *BClient) MyUppersTags() (*m.BPResp[UpperTag], error) {

	body, err := client.GetP(upperTagUrl)
	if err != nil {
		return nil, err
	}

	// respS := &UpperTag{}
	respS := &m.BPResp[UpperTag]{}
	if err = json.Unmarshal(body, respS); err != nil {
		return nil, err
	}

	if respS.Code != 0 {
		return nil, errs.StatusError{Code: respS.Code, Cause: respS.Message}
	}

	return respS, nil
}
