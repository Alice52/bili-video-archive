package api

import (
	"encoding/json"
	"fmt"
	m "github.com/alice52/archive/bili/api/model"
)

var (
	tagUppers = "https://api.bilibili.com/x/relation/tag?tagid=%d" // &pn=%d
)

// region response model

type UpperInfo struct {
	Attribute      int64                  `json:"attribute"`
	ContractInfo   map[string]interface{} `json:"contract_info"`
	Face           string                 `json:"face"`
	FaceNft        int64                  `json:"face_nft"`
	Live           Live                   `json:"live"`
	Mid            int64                  `json:"mid"`
	NftIcon        string                 `json:"nft_icon"`
	OfficialVerify OfficialVerify         `json:"official_verify"`
	RecReason      string                 `json:"rec_reason"`
	Sign           string                 `json:"sign"`
	Special        int64                  `json:"special"`
	Tag            interface{}            `json:"tag"`
	TrackID        string                 `json:"track_id"`
	Uname          string                 `json:"uname"`
	Vip            Vip                    `json:"vip"`
}

type Live struct {
	JumpURL    string `json:"jump_url"`
	LiveStatus int64  `json:"live_status"`
}

type OfficialVerify struct {
	Desc string `json:"desc"`
	Type int64  `json:"type"`
}

type Vip struct {
	AccessStatus       int64  `json:"accessStatus"`
	AvatarSubscript    int64  `json:"avatar_subscript"`
	AvatarSubscriptURL string `json:"avatar_subscript_url"`
	DueRemark          string `json:"dueRemark"`
	Label              Label  `json:"label"`
	NicknameColor      string `json:"nickname_color"`
	ThemeType          int64  `json:"themeType"`
	VipDueDate         int64  `json:"vipDueDate"`
	VipStatus          int64  `json:"vipStatus"`
	VipStatusWarn      string `json:"vipStatusWarn"`
	VipType            int64  `json:"vipType"`
}

type Label struct {
	BgColor     string `json:"bg_color"`
	BgStyle     int64  `json:"bg_style"`
	BorderColor string `json:"border_color"`
	LabelTheme  string `json:"label_theme"`
	Path        string `json:"path"`
	Text        string `json:"text"`
	TextColor   string `json:"text_color"`
}

//endregion

// UppersOfTag 获取当前登录用户当前分组下的UP主列表
// https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/docs/user/relation.md#%E6%9F%A5%E8%AF%A2%E5%85%B3%E6%B3%A8%E5%88%86%E7%BB%84%E6%98%8E%E7%BB%86
func (client *BClient) UppersOfTag(tagid int64) (*m.BPResp[UpperInfo], error) {
	uppers := &m.BPResp[UpperInfo]{}
	if ss, err := client.GetAllP(fmt.Sprintf(tagUppers, tagid)); err != nil {
		return nil, err
	} else {
		for _, s := range ss.Data {
			u := &UpperInfo{}
			if err := json.Unmarshal(s, u); err != nil {
				return nil, err
			}
			uppers.Data = append(uppers.Data, *u)
		}

		return uppers, nil
	}
}
