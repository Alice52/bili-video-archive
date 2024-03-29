package api

import (
	"encoding/json"
	"github.com/alice52/archive/bili/api/errs"
	"io/ioutil"
	"net/http"
)

const (
	mySpaceInfoUrl = "https://api.bilibili.com/x/space/myinfo"
)

type MySpaceInfoResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		Mid            int    `json:"mid"`
		Name           string `json:"name"`
		Sex            string `json:"sex"`
		Face           string `json:"face"`
		Sign           string `json:"sign"`
		Rank           int    `json:"rank"`
		Level          int    `json:"level"`
		Jointime       int    `json:"jointime"`
		Moral          int    `json:"moral"`
		Silence        int    `json:"silence"`
		EmailStatus    int    `json:"email_status"`
		TelStatus      int    `json:"tel_status"`
		Identification int    `json:"identification"`
		Vip            struct {
			Type       int   `json:"type"`
			Status     int   `json:"status"`
			DueDate    int64 `json:"due_date"`
			VipPayType int   `json:"vip_pay_type"`
			ThemeType  int   `json:"theme_type"`
			Label      struct {
				Path                  string `json:"path"`
				Text                  string `json:"text"`
				LabelTheme            string `json:"label_theme"`
				TextColor             string `json:"text_color"`
				BgStyle               int    `json:"bg_style"`
				BgColor               string `json:"bg_color"`
				BorderColor           string `json:"border_color"`
				UseImgLabel           bool   `json:"use_img_label"`
				ImgLabelURIHans       string `json:"img_label_uri_hans"`
				ImgLabelURIHant       string `json:"img_label_uri_hant"`
				ImgLabelURIHansStatic string `json:"img_label_uri_hans_static"`
				ImgLabelURIHantStatic string `json:"img_label_uri_hant_static"`
			} `json:"label"`
			AvatarSubscript    int    `json:"avatar_subscript"`
			NicknameColor      string `json:"nickname_color"`
			Role               int    `json:"role"`
			AvatarSubscriptURL string `json:"avatar_subscript_url"`
			TvVipStatus        int    `json:"tv_vip_status"`
			TvVipPayType       int    `json:"tv_vip_pay_type"`
		} `json:"vip"`
		Pendant struct {
			Pid               int    `json:"pid"`
			Name              string `json:"name"`
			Image             string `json:"image"`
			Expire            int    `json:"expire"`
			ImageEnhance      string `json:"image_enhance"`
			ImageEnhanceFrame string `json:"image_enhance_frame"`
		} `json:"pendant"`
		Nameplate struct {
			Nid        int    `json:"nid"`
			Name       string `json:"name"`
			Image      string `json:"image"`
			ImageSmall string `json:"image_small"`
			Level      string `json:"level"`
			Condition  string `json:"condition"`
		} `json:"nameplate"`
		Official struct {
			Role  int    `json:"role"`
			Title string `json:"title"`
			Desc  string `json:"desc"`
			Type  int    `json:"type"`
		} `json:"official"`
		Birthday      int  `json:"birthday"`
		IsTourist     int  `json:"is_tourist"`
		IsFakeAccount int  `json:"is_fake_account"`
		PinPrompting  int  `json:"pin_prompting"`
		IsDeleted     int  `json:"is_deleted"`
		InRegAudit    int  `json:"in_reg_audit"`
		IsRipUser     bool `json:"is_rip_user"`
		Profession    struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			ShowName    string `json:"show_name"`
			IsShow      int    `json:"is_show"`
			CategoryOne string `json:"category_one"`
			Realname    string `json:"realname"`
			Title       string `json:"title"`
			Department  string `json:"department"`
		} `json:"profession"`
		FaceNft        int `json:"face_nft"`
		FaceNftNew     int `json:"face_nft_new"`
		IsSeniorMember int `json:"is_senior_member"`
		Honours        struct {
			Mid    int `json:"mid"`
			Colour struct {
				Dark   string `json:"dark"`
				Normal string `json:"normal"`
			} `json:"colour"`
			Tags interface{} `json:"tags"`
		} `json:"honours"`
		DigitalID   string `json:"digital_id"`
		DigitalType int    `json:"digital_type"`
		LevelExp    struct {
			CurrentLevel int   `json:"current_level"`
			CurrentMin   int   `json:"current_min"`
			CurrentExp   int   `json:"current_exp"`
			NextExp      int   `json:"next_exp"`
			LevelUp      int64 `json:"level_up"`
		} `json:"level_exp"`
		Coins     float64 `json:"coins"`
		Following int     `json:"following"`
		Follower  int     `json:"follower"`
	} `json:"data"`
}

// MySpaceInfo https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/docs/user/info.md#%E7%99%BB%E5%BD%95%E7%94%A8%E6%88%B7%E7%A9%BA%E9%97%B4%E8%AF%A6%E7%BB%86%E4%BF%A1%E6%81%AF
func (client *BClient) MySpaceInfo() (*MySpaceInfoResp, error) {
	client.HttpClient = &http.Client{}
	request, err := client.newCookieRequest(http.MethodGet, mySpaceInfoUrl, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errs.ErrUnexpectedStatusCode(resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	mySpaceInfoResp := &MySpaceInfoResp{}
	if err = json.Unmarshal(body, mySpaceInfoResp); err != nil {
		return nil, err
	}

	if mySpaceInfoResp.Code != 0 {
		return nil, errs.StatusError{Code: mySpaceInfoResp.Code, Cause: mySpaceInfoResp.Message}
	}

	return mySpaceInfoResp, nil
}
