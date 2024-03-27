package service

import (
	"github.com/alice52/archive/bili/api"
	"github.com/alice52/archive/bili/source/gen/dal"
	"github.com/alice52/archive/bili/source/gen/model"
	"github.com/alice52/archive/common/global"
	"github.com/gookit/goutil/jsonutil"
	"go.uber.org/zap"
)

type UserFavFolderServiceIn struct{}

func (c *UserFavFolderServiceIn) SyncUserFavFolders() (err error) {
	tags, err := api.LogonClient.UserFavFolders()
	if err != nil {
		return err
	}

	for _, item := range tags.Data.List {
		m := &model.ArchivedFavFolder{
			ID:         item.ID,
			Fid:        item.Fid,
			Mid:        item.Mid,
			MediaCount: item.MediaCount,
			Title:      &item.Title,
		}
		resp := jsonutil.MustString(item)
		m.Resp = &resp

		if err = dal.Q.ArchivedFavFolder.Save(m); err != nil {
			global.LOG.Error("sync user fav folders error", zap.Error(err))
		}
	}

	return err
}
