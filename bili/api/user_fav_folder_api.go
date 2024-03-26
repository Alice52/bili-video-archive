package api

import (
	"encoding/json"
	"fmt"
	c "github.com/alice52/archive/bili/api/common"
	m "github.com/alice52/archive/bili/api/model"
)

var (
	favFoldersUrl = "https://api.bilibili.com/x/v3/fav/folder/created/list-all?up_mid=%s"
)

// region response model

type FavFolder struct {
	Count  int64       `json:"count"`
	List   []Folder    `json:"list"`
	Season interface{} `json:"season"`
}

type Folder struct {
	Attr       int64  `json:"attr"`
	FavState   int64  `json:"fav_state"`
	Fid        int64  `json:"fid"`
	ID         int64  `json:"id"`
	MediaCount int64  `json:"media_count"`
	Mid        int64  `json:"mid"`
	Title      string `json:"title"`
}

//endregion

// UserFavFolders 获取指定用户创建的所有收藏夹信息
// https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/docs/fav/info.md#%E8%8E%B7%E5%8F%96%E6%8C%87%E5%AE%9A%E7%94%A8%E6%88%B7%E5%88%9B%E5%BB%BA%E7%9A%84%E6%89%80%E6%9C%89%E6%94%B6%E8%97%8F%E5%A4%B9%E4%BF%A1%E6%81%AF
func (client *BClient) UserFavFolders() (*m.BResp[FavFolder], error) {
	folders := &m.BResp[FavFolder]{}
	if bs, err := client.GetList(fmt.Sprintf(favFoldersUrl, c.SelfMid)); err != nil {
		return nil, err
	} else {
		return folders, json.Unmarshal(bs, folders)
	}
}
