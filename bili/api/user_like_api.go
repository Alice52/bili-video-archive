package api

import (
	"encoding/json"
	"fmt"
	c "github.com/alice52/archive/bili/api/common"
	m "github.com/alice52/archive/bili/api/model"
)

var (
	userLikeUrl = "https://api.bilibili.com/x/space/like/video?vmid=%s"
)

// region response model

type LikedVideo struct {
	List []Video `json:"list"`
}

type Video struct {
	Aid          int64     `json:"aid"`
	Bvid         string    `json:"bvid"`
	Cid          int64     `json:"cid"`
	Copyright    int64     `json:"copyright"`
	Cover43      string    `json:"cover43"`
	Ctime        int64     `json:"ctime"`
	Desc         string    `json:"desc"`
	Dimension    Dimension `json:"dimension"`
	Duration     int64     `json:"duration"`
	Dynamic      string    `json:"dynamic"`
	EnableVT     int64     `json:"enable_vt"`
	FirstFrame   string    `json:"first_frame"`
	InterVideo   bool      `json:"inter_video"`
	MissionID    int64     `json:"mission_id"`
	Owner        Owner     `json:"owner"`
	Pic          string    `json:"pic"`
	PubLocation  string    `json:"pub_location"`
	Pubdate      int64     `json:"pubdate"`
	ResourceType string    `json:"resource_type"`
	Rights       Rights    `json:"rights"`
	SeasonID     int64     `json:"season_id"`
	ShortLinkV2  string    `json:"short_link_v2"`
	Stat         Stat      `json:"stat"`
	State        int64     `json:"state"`
	Subtitle     string    `json:"subtitle"`
	Tid          int64     `json:"tid"`
	Title        string    `json:"title"`
	Tname        string    `json:"tname"`
	UpFromV2     int64     `json:"up_from_v2"`
	Videos       int64     `json:"videos"`
}

type Dimension struct {
	Height int64 `json:"height"`
	Rotate int64 `json:"rotate"`
	Width  int64 `json:"width"`
}

type Owner struct {
	Face string `json:"face"`
	Mid  int64  `json:"mid"`
	Name string `json:"name"`
}

type Rights struct {
	ArcPay        int64 `json:"arc_pay"`
	Autoplay      int64 `json:"autoplay"`
	Bp            int64 `json:"bp"`
	Download      int64 `json:"download"`
	Elec          int64 `json:"elec"`
	Hd5           int64 `json:"hd5"`
	IsCooperation int64 `json:"is_cooperation"`
	Movie         int64 `json:"movie"`
	NoBackground  int64 `json:"no_background"`
	NoReprint     int64 `json:"no_reprint"`
	Pay           int64 `json:"pay"`
	PayFreeWatch  int64 `json:"pay_free_watch"`
	UgcPay        int64 `json:"ugc_pay"`
	UgcPayPreview int64 `json:"ugc_pay_preview"`
}

type Stat struct {
	Aid      int64 `json:"aid"`
	Coin     int64 `json:"coin"`
	Danmaku  int64 `json:"danmaku"`
	Dislike  int64 `json:"dislike"`
	Favorite int64 `json:"favorite"`
	HisRank  int64 `json:"his_rank"`
	Like     int64 `json:"like"`
	NowRank  int64 `json:"now_rank"`
	Reply    int64 `json:"reply"`
	Share    int64 `json:"share"`
	View     int64 `json:"view"`
	VT       int64 `json:"vt"`
	Vv       int64 `json:"vv"`
}

//endregion

// UserLiked 查询用户最近点赞视频（
// https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/docs/user/space.md#%E6%9F%A5%E8%AF%A2%E7%94%A8%E6%88%B7%E6%9C%80%E8%BF%91%E7%82%B9%E8%B5%9E%E8%A7%86%E9%A2%91web
func (client *BClient) UserLiked() (*m.BResp[LikedVideo], error) {
	fav := &m.BResp[LikedVideo]{}
	if bs, err := client.Get(fmt.Sprintf(userLikeUrl, c.SelfMid)); err != nil {
		return nil, err
	} else {
		return fav, json.Unmarshal(bs, &fav)
	}
}
