package api

import (
	m "github.com/alice52/archive/bili/api/model"
)

var (
	historyUrl = "https://api.bilibili.com/x/web-interface/history/cursor?"
)

// region response model

//endregion

// History 获取当前登录用户当天的历史记录
// https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/docs/history&toview/history.md#%E8%8E%B7%E5%8F%96%E5%8E%86%E5%8F%B2%E8%AE%B0%E5%BD%95%E5%88%97%E8%A1%A8_web%E7%AB%AF
func (client *BClient) History() (*m.BPResp[any], error) {

	return nil, nil
}
