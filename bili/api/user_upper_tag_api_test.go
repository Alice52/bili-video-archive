package api

import (
	"fmt"
	"github.com/alice52/archive/bilibili/util"
	"testing"
)

func TestMyUppersTags(t *testing.T) {

	info, err := logonFunc().MyUppersTags()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(util.MustMarshal(info))
}
