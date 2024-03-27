package service

import (
	"github.com/alice52/archive/bili/api"
	"github.com/alice52/archive/bili/source/gen/dal"
	"github.com/alice52/archive/bili/source/gen/model"
	"github.com/alice52/archive/common/global"
	"github.com/gookit/goutil/jsonutil"
	"go.uber.org/zap"
)

type UserLikedServiceIn struct{}

func (c *UserLikedServiceIn) SyncUserLiked() (err error) {
	items, err := api.LogonClient.UserLiked()
	if err != nil {
		return err
	}

	for _, item := range items.Data.List {
		m := &model.ArchivedLike{
			Bvid:     item.Bvid,
			Aid:      item.Aid,
			Cid:      item.Cid,
			Cover:    &item.Pic,
			Duration: item.Duration,
			LikeTime: item.Ctime,
			SeasonID: item.SeasonID,
			Intro:    &item.Desc,
			Title:    &item.Title,
			Type:     item.Tid,
		}

		up := jsonutil.MustString(item.Owner)
		m.Owner = &up
		cnt := jsonutil.MustString(item.Stat)
		m.CntInfo = &cnt
		resp := jsonutil.MustString(item)
		m.Resp = &resp

		if err = dal.Q.ArchivedLike.Save(m); err != nil {
			global.LOG.Error("sync user liked error", zap.Error(err))
		}
	}

	return err
}
