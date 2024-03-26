package api

import (
	"encoding/json"
	"fmt"
	c "github.com/alice52/archive/bilibili/api/common"
	m "github.com/alice52/archive/bilibili/api/model"
)

var (
	coinedUrl = "https://api.bilibili.com/x/space/coin/video?vmid=%s"
)

// region response model

type CoinVideo struct {
	Aid        int64     `json:"aid"`
	Attribute  int64     `json:"attribute"`
	Bvid       string    `json:"bvid"`
	Cid        int64     `json:"cid"`
	Coins      int64     `json:"coins"`
	Copyright  int64     `json:"copyright"`
	Ctime      int64     `json:"ctime"`
	Desc       string    `json:"desc"`
	Dimension  Dimension `json:"dimension"`
	Duration   int64     `json:"duration"`
	Dynamic    string    `json:"dynamic"`
	InterVideo bool      `json:"inter_video"`
	IP         string    `json:"ip"`
	MissionID  int64     `json:"mission_id"`
	Owner      Owner     `json:"owner"`
	Pic        string    `json:"pic"`
	Pubdate    int64     `json:"pubdate"`
	Rights     Rights    `json:"rights"`
	Stat       Stat      `json:"stat"`
	State      int64     `json:"state"`
	Tid        int64     `json:"tid"`
	Time       int64     `json:"time"`
	Title      string    `json:"title"`
	Tname      string    `json:"tname"`
	Videos     int64     `json:"videos"`
}

//endregion

// UserCoined 查询用户最近投币视频
// https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/docs/user/space.md#%E6%9F%A5%E8%AF%A2%E7%94%A8%E6%88%B7%E6%9C%80%E8%BF%91%E6%8A%95%E5%B8%81%E8%A7%86%E9%A2%91web
func (client *BClient) UserCoined() (*m.BPResp[CoinVideo], error) {
	fav := &m.BPResp[CoinVideo]{}
	if bs, err := client.GetP(fmt.Sprintf(coinedUrl, c.SelfMid)); err != nil {
		return nil, err
	} else {
		return fav, json.Unmarshal(bs, &fav)
	}
}
