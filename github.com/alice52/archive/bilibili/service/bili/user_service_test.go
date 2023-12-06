package bili

import (
	"github.com/alice52/archive/bilibili/util"
	"os"
	"testing"
)

func TestClient_MySpaceInfo(t *testing.T) {
	client := &BClient{}
	resps, err := client.LoginWithQrCode(os.Stdout)
	if err != nil {
		return
	}
	for resp := range resps {
		if resp.LoginStatus == LoginSuccess {
			client.SetCookie(resp.Cookie)
		}
		continue
	}
	info, err := client.MySpaceInfo()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(util.MustMarshal(info))
}
