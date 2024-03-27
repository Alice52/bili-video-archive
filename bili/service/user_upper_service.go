package service

import (
	"github.com/alice52/archive/bili/api"
	"github.com/alice52/archive/bili/source/gen/dal"
	"github.com/alice52/archive/bili/source/gen/model"
	"github.com/alice52/archive/common/global"
	"github.com/gookit/goutil/jsonutil"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type UserUpperServiceIn struct{}

func (c *UserUpperServiceIn) SyncUppers() (err error) {
	upsTags, err := dal.Q.ArchivedUpsTag.Find()
	if err != nil {
		return err
	}

	for _, tag := range upsTags {
		if err = DoSyncUppers(tag.TagID); err != nil {
			global.LOG.Error("sync upper in tag"+cast.ToString(tag.TagID)+" error", zap.Error(err))
		}
	}

	return err
}

func DoSyncUppers(tagId int64) (err error) {
	uppers, err := api.LogonClient.UppersOfTag(tagId)
	if err != nil {
		return err
	}

	for _, upper := range uppers.Data {
		m := &model.ArchivedUp{
			TagID: tagId,
			Sign:  &upper.Sign,
			Uname: &upper.Uname,
			Mid:   upper.Mid,
		}
		resp := jsonutil.MustString(upper)
		m.Resp = &resp
		if err = dal.Q.ArchivedUp.Save(m); err != nil {
			global.LOG.Error("sync upper error", zap.Error(err))
		}
	}

	return err
}
