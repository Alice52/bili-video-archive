package service

import (
	"github.com/alice52/archive/bili/api"
	"github.com/alice52/archive/bili/source/gen/dal"
	"github.com/alice52/archive/bili/source/gen/model"
	"github.com/alice52/archive/common/global"
	"go.uber.org/zap"
)

type UserUpperTagServiceIn struct{}

func (c *UserUpperTagServiceIn) SyncUpperTags() (err error) {
	tags, err := api.LogonClient.MyUppersTags()
	if err != nil {
		return err
	}

	for _, tag := range tags.Data {
		m := &model.ArchivedUpsTag{
			TagID:  tag.Tagid,
			Count_: &tag.Count,
			Name:   &tag.Name,
			Tip:    &tag.Tip,
		}

		if err = dal.Q.ArchivedUpsTag.Save(m); err != nil {
			global.LOG.Error("sync upper tags error", zap.Error(err))
		}
	}

	return err
}
