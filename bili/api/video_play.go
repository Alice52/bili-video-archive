package api

import (
	"encoding/json"
	"fmt"
	c "github.com/alice52/archive/bili/api/common"
	m "github.com/alice52/archive/bili/api/model"
)

const (
	playUrl = "https://api.bilibili.com/x/player/playurl?bvid=%s&cid=%d&qn=%d&fourk=1&fnval=%d"
)

type PlayUrl struct {
	From              string   `json:"from"`
	Result            string   `json:"result"`
	Message           string   `json:"message"`
	Quality           int      `json:"quality"`
	Format            string   `json:"format"`
	Timelength        int      `json:"timelength"`
	AcceptFormat      string   `json:"accept_format"`
	AcceptDescription []string `json:"accept_description"`
	AcceptQuality     []int    `json:"accept_quality"`
	VideoCodecid      int      `json:"video_codecid"`
	SeekParam         string   `json:"seek_param"`
	SeekType          string   `json:"seek_type"`
	Durl              []struct {
		Order     int      `json:"order"`
		Length    int      `json:"length"`
		Size      int      `json:"size"`
		Ahead     string   `json:"ahead"`
		Vhead     string   `json:"vhead"`
		URL       string   `json:"url"`
		BackupURL []string `json:"backup_url"`
	} `json:"durl"`
	Dash struct {
		Duration      int     `json:"duration"`
		MinBufferTime float64 `json:"min_buffer_time"`
		Video         []struct {
			ID           int      `json:"id"`
			BaseURL      string   `json:"base_url"`
			BackupURL    []string `json:"backup_url"`
			Bandwidth    int      `json:"bandwidth"`
			MimeType     string   `json:"mime_type"`
			Codecs       string   `json:"codecs"`
			Width        int      `json:"width"`
			Height       int      `json:"height"`
			FrameRate    string   `json:"frame_rate"`
			Sar          string   `json:"sar"`
			StartWithSap int      `json:"start_with_sap"`
			SegmentBase  struct {
				Initialization string `json:"initialization"`
				IndexRange     string `json:"index_range"`
			} `json:"segment_base"`
			Codecid int `json:"codecid"`
		} `json:"video"`
		Audio []struct {
			ID           int      `json:"id"`
			BaseURL      string   `json:"base_url"`
			BackupURL    []string `json:"backup_url"`
			Bandwidth    int      `json:"bandwidth"`
			MimeType     string   `json:"mime_type"`
			Codecs       string   `json:"codecs"`
			Width        int      `json:"width"`
			Height       int      `json:"height"`
			FrameRate    string   `json:"frame_rate"`
			Sar          string   `json:"sar"`
			StartWithSap int      `json:"start_with_sap"`
			SegmentBase  struct {
				Initialization string `json:"initialization"`
				IndexRange     string `json:"index_range"`
			} `json:"segment_base"`
			Codecid int `json:"codecid"`
		} `json:"audio"`
		Dolby struct {
			Type  int         `json:"type"`
			Audio interface{} `json:"audio"`
		} `json:"dolby"`
		Flac interface{} `json:"flac"`
	} `json:"dash"`
	SupportFormats []struct {
		Quality        int         `json:"quality"`
		Format         string      `json:"format"`
		NewDescription string      `json:"new_description"`
		DisplayDesc    string      `json:"display_desc"`
		Superscript    string      `json:"superscript"`
		Codecs         interface{} `json:"codecs"`
	} `json:"support_formats"`
	HighFormat   interface{} `json:"high_format"`
	LastPlayTime int         `json:"last_play_time"`
	LastPlayCid  int         `json:"last_play_cid"`
}

func (client *BClient) PlayUrl(bvid string, cid int64, qn c.Qn, fnval c.Fnval) (*m.BResp[PlayUrl], error) {
	playUrlResp := &m.BResp[PlayUrl]{}
	if ss, err := client.Get(fmt.Sprintf(playUrl, bvid, cid, qn, fnval)); err != nil {
		return nil, err
	} else {
		return playUrlResp, json.Unmarshal(ss, &playUrlResp)
	}
}
