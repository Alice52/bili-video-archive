package api

import (
	"encoding/json"
	"fmt"
	m "github.com/alice52/archive/bilibili/api/model"
)

var (
	favInFolderUrl = "https://api.bilibili.com/x/v3/fav/resource/list?media_id=%s&platform=web"
)

// region response model

type Fav struct {
	HasMore bool    `json:"has_more"`
	Info    Info    `json:"info"`
	Medias  []Media `json:"medias"`
	TTL     int64   `json:"ttl"`
}

type Info struct {
	Attr       int64       `json:"attr"`
	CntInfo    InfoCntInfo `json:"cnt_info"`
	Cover      string      `json:"cover"`
	CoverType  int64       `json:"cover_type"`
	Ctime      int64       `json:"ctime"`
	FavState   int64       `json:"fav_state"`
	Fid        int64       `json:"fid"`
	ID         int64       `json:"id"`
	Intro      string      `json:"intro"`
	LikeState  int64       `json:"like_state"`
	MediaCount int64       `json:"media_count"`
	Mid        int64       `json:"mid"`
	Mtime      int64       `json:"mtime"`
	State      int64       `json:"state"`
	Title      string      `json:"title"`
	Type       int64       `json:"type"`
	Upper      InfoUpper   `json:"upper"`
}

type InfoCntInfo struct {
	Collect int64 `json:"collect"`
	Play    int64 `json:"play"`
	Share   int64 `json:"share"`
	ThumbUp int64 `json:"thumb_up"`
}

type InfoUpper struct {
	Face      string `json:"face"`
	Followed  bool   `json:"followed"`
	Mid       int64  `json:"mid"`
	Name      string `json:"name"`
	VipStatue int64  `json:"vip_statue"`
	VipType   int64  `json:"vip_type"`
}

type Media struct {
	Attr     int64        `json:"attr"`
	BvID     string       `json:"bv_id"`
	Bvid     string       `json:"bvid"`
	CntInfo  MediaCntInfo `json:"cnt_info"`
	Cover    string       `json:"cover"`
	Ctime    int64        `json:"ctime"`
	Duration int64        `json:"duration"`
	FavTime  int64        `json:"fav_time"`
	ID       int64        `json:"id"`
	Intro    string       `json:"intro"`
	Link     string       `json:"link"`
	Page     int64        `json:"page"`
	Pubtime  int64        `json:"pubtime"`
	Season   interface{}  `json:"season"`
	Title    string       `json:"title"`
	Type     int64        `json:"type"`
	Upper    MediaUpper   `json:"upper"`
}

type MediaCntInfo struct {
	Collect int64 `json:"collect"`
	Danmaku int64 `json:"danmaku"`
	Play    int64 `json:"play"`
}

type MediaUpper struct {
	Face string `json:"face"`
	Mid  int64  `json:"mid"`
	Name string `json:"name"`
}

//endregion

// UserFavOfFolder 获取收藏夹内容明细列表
// https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/docs/fav/list.md#%E8%8E%B7%E5%8F%96%E6%94%B6%E8%97%8F%E5%A4%B9%E5%86%85%E5%AE%B9%E6%98%8E%E7%BB%86%E5%88%97%E8%A1%A8
func (client *BClient) UserFavOfFolder(mediaId string) (*m.BResp[Fav], error) {
	fav := &m.BResp[Fav]{}
	if bs, err := client.GetListAll4FavOfFolder(fmt.Sprintf(favInFolderUrl, mediaId)); err != nil {
		return nil, err
	} else {
		if err := json.Unmarshal(bs.Data.Info, &fav.Data.Info); err != nil {
			return nil, err
		}

		for _, media := range bs.Data.Medias {
			me := &Media{}
			if err := json.Unmarshal(media, me); err != nil {
				return nil, err
			}
			fav.Data.Medias = append(fav.Data.Medias, *me)
		}
	}

	return fav, nil
}
