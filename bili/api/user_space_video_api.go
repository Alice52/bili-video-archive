package api

import (
	"encoding/json"
	"fmt"
	c "github.com/alice52/archive/bilibili/api/common"
	m "github.com/alice52/archive/bilibili/api/model"
)

const (
	mySpaceVideoUrl = "https://api.bilibili.com/x/polymer/web-dynamic/v1/feed/space?host_mid=%s"
)

// region response model

type SpaceVideo struct {
	HasMore        bool       `json:"has_more"`
	Items          []DataItem `json:"items"`
	Offset         string     `json:"offset"`
	UpdateBaseline string     `json:"update_baseline"`
	UpdateNum      int64      `json:"update_num"`
}

type DataItem struct {
	Basic Basic  `json:"basic"`
	IDStr string `json:"id_str"`
	// Modules Modules `json:"modules"`
	Orig    Orig   `json:"orig"`
	Type    string `json:"type"`
	Visible bool   `json:"visible"`
}

type Basic struct {
	CommentIDStr string         `json:"comment_id_str"`
	CommentType  int64          `json:"comment_type"`
	LikeIcon     PurpleLikeIcon `json:"like_icon"`
	RidStr       string         `json:"rid_str"`
}

type PurpleLikeIcon struct {
	ActionURL string `json:"action_url"`
	EndURL    string `json:"end_url"`
	ID        int64  `json:"id"`
	StartURL  string `json:"start_url"`
}

type Modules struct {
	ModuleAuthor  PurpleModuleAuthor  `json:"module_author"`
	ModuleDynamic PurpleModuleDynamic `json:"module_dynamic"`
	ModuleMore    ModuleMore          `json:"module_more"`
	ModuleStat    ModuleStat          `json:"module_stat"`
	ModuleTag     ModuleTag           `json:"module_tag"`
}

type PurpleModuleAuthor struct {
	Avatar          PurpleAvatar         `json:"avatar"`
	Face            string               `json:"face"`
	FaceNft         bool                 `json:"face_nft"`
	Following       interface{}          `json:"following"`
	JumpURL         string               `json:"jump_url"`
	Label           string               `json:"label"`
	Mid             int64                `json:"mid"`
	Name            string               `json:"name"`
	OfficialVerify  PurpleOfficialVerify `json:"official_verify"`
	Pendant         PurplePendant        `json:"pendant"`
	PubAction       string               `json:"pub_action"`
	PubLocationText string               `json:"pub_location_text"`
	PubTime         string               `json:"pub_time"`
	PubTs           int64                `json:"pub_ts"`
	Type            string               `json:"type"`
	Vip             PurpleVip            `json:"vip"`
}

type PurpleAvatar struct {
	ContainerSize  PurpleContainerSize  `json:"container_size"`
	FallbackLayers PurpleFallbackLayers `json:"fallback_layers"`
	Mid            string               `json:"mid"`
}

type PurpleContainerSize struct {
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
}

type PurpleFallbackLayers struct {
	IsCriticalGroup bool          `json:"is_critical_group"`
	Layers          []PurpleLayer `json:"layers"`
}

type PurpleLayer struct {
	LayerConfig PurpleLayerConfig `json:"layer_config"`
	Resource    PurpleResource    `json:"resource"`
	Visible     bool              `json:"visible"`
}

type PurpleRenderSpec struct {
	Opacity int64 `json:"opacity"`
}

type PurpleSizeSpec struct {
	Height *Height `json:"height"`
	Width  *Height `json:"width"`
}

type PurpleLayerConfig struct {
	IsCritical *bool      `json:"is_critical,omitempty"`
	Tags       PurpleTags `json:"tags"`
}

type PurpleTags struct {
	AvatarLayer map[string]interface{} `json:"AVATAR_LAYER,omitempty"`
	GeneralCFG  PurpleGENERALCFG       `json:"GENERAL_CFG"`
	IconLayer   map[string]interface{} `json:"ICON_LAYER"`
}

type PurpleGENERALCFG struct {
	ConfigType    int64               `json:"config_type"`
	GeneralConfig PurpleGeneralConfig `json:"general_config"`
}

type PurpleGeneralConfig struct {
	WebCSSStyle PurpleWebCSSStyle `json:"web_css_style"`
}

type PurpleWebCSSStyle struct {
	BackgroundColor string `json:"background-color"`
	Border          string `json:"border"`
	BorderRadius    string `json:"borderRadius"`
	BoxSizing       string `json:"boxSizing"`
}

type PurpleResource struct {
	ResImage PurpleResImage `json:"res_image"`
	ResType  int64          `json:"res_type"`
}

type PurpleResImage struct {
	ImageSrc PurpleImageSrc `json:"image_src"`
}

type PurpleImageSrc struct {
	Local       int64         `json:"local"`
	Placeholder *int64        `json:"placeholder,omitempty"`
	Remote      *PurpleRemote `json:"remote,omitempty"`
	SrcType     int64         `json:"src_type"`
}

type PurpleRemote struct {
	BFSStyle string `json:"bfs_style"`
	URL      string `json:"url"`
}

type PurpleOfficialVerify struct {
	Desc string `json:"desc"`
	Type int64  `json:"type"`
}

type PurplePendant struct {
	Expire            int64  `json:"expire"`
	Image             string `json:"image"`
	ImageEnhance      string `json:"image_enhance"`
	ImageEnhanceFrame string `json:"image_enhance_frame"`
	NPID              int64  `json:"n_pid"`
	Name              string `json:"name"`
	PID               int64  `json:"pid"`
}

type PurpleVip struct {
	AvatarSubscript    int64       `json:"avatar_subscript"`
	AvatarSubscriptURL string      `json:"avatar_subscript_url"`
	DueDate            int64       `json:"due_date"`
	Label              PurpleLabel `json:"label"`
	NicknameColor      string      `json:"nickname_color"`
	Status             int64       `json:"status"`
	ThemeType          int64       `json:"theme_type"`
	Type               int64       `json:"type"`
}

type PurpleLabel struct {
	BgColor               string `json:"bg_color"`
	BgStyle               int64  `json:"bg_style"`
	BorderColor           string `json:"border_color"`
	ImgLabelURIHans       string `json:"img_label_uri_hans"`
	ImgLabelURIHansStatic string `json:"img_label_uri_hans_static"`
	ImgLabelURIHant       string `json:"img_label_uri_hant"`
	ImgLabelURIHantStatic string `json:"img_label_uri_hant_static"`
	LabelTheme            string `json:"label_theme"`
	Path                  string `json:"path"`
	Text                  string `json:"text"`
	TextColor             string `json:"text_color"`
	UseImgLabel           bool   `json:"use_img_label"`
}

type PurpleModuleDynamic struct {
	Additional interface{} `json:"additional"`
	Desc       PurpleDesc  `json:"desc"`
	Major      interface{} `json:"major"`
	Topic      interface{} `json:"topic"`
}

type PurpleDesc struct {
	RichTextNodes []PurpleRichTextNode `json:"rich_text_nodes"`
	Text          string               `json:"text"`
}

type PurpleRichTextNode struct {
	OrigText *string `json:"orig_text,omitempty"`
	Text     *string `json:"text,omitempty"`
	Type     *string `json:"type,omitempty"`
}

type ModuleMore struct {
	ThreePointItems []ThreePointItem `json:"three_point_items"`
}

type ThreePointItem struct {
	Label  string  `json:"label"`
	Modal  *Modal  `json:"modal,omitempty"`
	Params *Params `json:"params,omitempty"`
	Type   string  `json:"type"`
}

type Modal struct {
	Cancel  string `json:"cancel"`
	Confirm string `json:"confirm"`
	Content string `json:"content"`
	Title   string `json:"title"`
}

type Params struct {
	DynamicID string `json:"dynamic_id"`
	Status    bool   `json:"status"`
}

type ModuleStat struct {
	Comment Comment `json:"comment"`
	Forward Forward `json:"forward"`
	Like    Like    `json:"like"`
}

type Comment struct {
	Count     int64 `json:"count"`
	Forbidden bool  `json:"forbidden"`
}

type Forward struct {
	Count     int64 `json:"count"`
	Forbidden bool  `json:"forbidden"`
}

type Like struct {
	Count     int64 `json:"count"`
	Forbidden bool  `json:"forbidden"`
	Status    bool  `json:"status"`
}

type ModuleTag struct {
	Text string `json:"text"`
}

type Orig struct {
	Modules OrigModules `json:"modules"`
}

type OrigBasic struct {
	CommentIDStr string         `json:"comment_id_str"`
	CommentType  int64          `json:"comment_type"`
	LikeIcon     FluffyLikeIcon `json:"like_icon"`
	RidStr       string         `json:"rid_str"`
}

type FluffyLikeIcon struct {
	ActionURL string `json:"action_url"`
	EndURL    string `json:"end_url"`
	ID        int64  `json:"id"`
	StartURL  string `json:"start_url"`
}

type OrigModules struct {
	ModuleDynamic FluffyModuleDynamic `json:"module_dynamic"`
}

type FluffyModuleAuthor struct {
	Avatar         FluffyAvatar         `json:"avatar"`
	Face           string               `json:"face"`
	FaceNft        bool                 `json:"face_nft"`
	Following      interface{}          `json:"following"`
	JumpURL        string               `json:"jump_url"`
	Label          string               `json:"label"`
	Mid            int64                `json:"mid"`
	Name           string               `json:"name"`
	OfficialVerify FluffyOfficialVerify `json:"official_verify"`
	Pendant        FluffyPendant        `json:"pendant"`
	PubAction      string               `json:"pub_action"`
	PubTime        string               `json:"pub_time"`
	PubTs          int64                `json:"pub_ts"`
	Type           string               `json:"type"`
	Vip            FluffyVip            `json:"vip"`
}

type FluffyAvatar struct {
	ContainerSize  FluffyContainerSize  `json:"container_size"`
	FallbackLayers FluffyFallbackLayers `json:"fallback_layers"`
	Mid            string               `json:"mid"`
}

type FluffyContainerSize struct {
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
}

type FluffyFallbackLayers struct {
	IsCriticalGroup bool          `json:"is_critical_group"`
	Layers          []FluffyLayer `json:"layers"`
}

type FluffyLayer struct {
	GeneralSpec FluffyGeneralSpec `json:"general_spec"`
	LayerConfig FluffyLayerConfig `json:"layer_config"`
	Resource    FluffyResource    `json:"resource"`
	Visible     bool              `json:"visible"`
}

type FluffyGeneralSpec struct {
	PosSpec    FluffyPosSpec    `json:"pos_spec"`
	RenderSpec FluffyRenderSpec `json:"render_spec"`
	SizeSpec   FluffySizeSpec   `json:"size_spec"`
}

type FluffyPosSpec struct {
	AxisX         float64 `json:"axis_x"`
	AxisY         float64 `json:"axis_y"`
	CoordinatePos int64   `json:"coordinate_pos"`
}

type FluffyRenderSpec struct {
	Opacity int64 `json:"opacity"`
}

type FluffySizeSpec struct {
	Height *Height `json:"height"`
	Width  *Height `json:"width"`
}

type FluffyLayerConfig struct {
	IsCritical *bool      `json:"is_critical,omitempty"`
	Tags       FluffyTags `json:"tags"`
}

type FluffyTags struct {
	AvatarLayer map[string]interface{} `json:"AVATAR_LAYER,omitempty"`
	GeneralCFG  FluffyGENERALCFG       `json:"GENERAL_CFG"`
	IconLayer   map[string]interface{} `json:"ICON_LAYER"`
}

type FluffyGENERALCFG struct {
	ConfigType    int64               `json:"config_type"`
	GeneralConfig FluffyGeneralConfig `json:"general_config"`
}

type FluffyGeneralConfig struct {
	WebCSSStyle FluffyWebCSSStyle `json:"web_css_style"`
}

type FluffyWebCSSStyle struct {
	BackgroundColor string `json:"background-color"`
	Border          string `json:"border"`
	BorderRadius    string `json:"borderRadius"`
	BoxSizing       string `json:"boxSizing"`
}

type FluffyResource struct {
	ResImage FluffyResImage `json:"res_image"`
	ResType  int64          `json:"res_type"`
}

type FluffyResImage struct {
	ImageSrc FluffyImageSrc `json:"image_src"`
}

type FluffyImageSrc struct {
	Local       int64         `json:"local"`
	Placeholder *int64        `json:"placeholder,omitempty"`
	Remote      *FluffyRemote `json:"remote,omitempty"`
	SrcType     int64         `json:"src_type"`
}

type FluffyRemote struct {
	BFSStyle string `json:"bfs_style"`
	URL      string `json:"url"`
}

type FluffyOfficialVerify struct {
	Desc string `json:"desc"`
	Type int64  `json:"type"`
}

type FluffyPendant struct {
	Expire            int64  `json:"expire"`
	Image             string `json:"image"`
	ImageEnhance      string `json:"image_enhance"`
	ImageEnhanceFrame string `json:"image_enhance_frame"`
	NPID              int64  `json:"n_pid"`
	Name              string `json:"name"`
	PID               int64  `json:"pid"`
}

type FluffyVip struct {
	AvatarSubscript    int64       `json:"avatar_subscript"`
	AvatarSubscriptURL string      `json:"avatar_subscript_url"`
	DueDate            int64       `json:"due_date"`
	Label              FluffyLabel `json:"label"`
	NicknameColor      string      `json:"nickname_color"`
	Status             int64       `json:"status"`
	ThemeType          int64       `json:"theme_type"`
	Type               int64       `json:"type"`
}

type FluffyLabel struct {
	BgColor               string `json:"bg_color"`
	BgStyle               int64  `json:"bg_style"`
	BorderColor           string `json:"border_color"`
	ImgLabelURIHans       string `json:"img_label_uri_hans"`
	ImgLabelURIHansStatic string `json:"img_label_uri_hans_static"`
	ImgLabelURIHant       string `json:"img_label_uri_hant"`
	ImgLabelURIHantStatic string `json:"img_label_uri_hant_static"`
	LabelTheme            string `json:"label_theme"`
	Path                  string `json:"path"`
	Text                  string `json:"text"`
	TextColor             string `json:"text_color"`
	UseImgLabel           bool   `json:"use_img_label"`
}

type FluffyModuleDynamic struct {
	Additional interface{} `json:"additional"`
	Desc       FluffyDesc  `json:"desc"`
	Major      Major       `json:"major"`
	Topic      interface{} `json:"topic"`
}

type FluffyDesc struct {
	RichTextNodes []FluffyRichTextNode `json:"rich_text_nodes"`
	Text          string               `json:"text"`
}

type FluffyRichTextNode struct {
	JumpURL  string `json:"jump_url"`
	OrigText string `json:"orig_text"`
	Text     string `json:"text"`
	Type     string `json:"type"`
}

type Major struct {
	Archive Archive `json:"archive"`
	Type    string  `json:"type"`
}

type Archive struct {
	Aid            string    `json:"aid"`
	Badge          Badge     `json:"badge"`
	Bvid           string    `json:"bvid"`
	Cover          string    `json:"cover"`
	Desc           string    `json:"desc"`
	DisablePreview int64     `json:"disable_preview"`
	DurationText   string    `json:"duration_text"`
	JumpURL        string    `json:"jump_url"`
	Stat           VideoStat `json:"stat"`
	Title          string    `json:"title"`
	Type           int64     `json:"type"`
}

type Badge struct {
	BgColor string      `json:"bg_color"`
	Color   string      `json:"color"`
	IconURL interface{} `json:"icon_url"`
	Text    string      `json:"text"`
}

type VideoStat struct {
	Danmaku string `json:"danmaku"`
	Play    string `json:"play"`
}

type Height struct {
	Double  *float64
	Integer *int64
}

// endregion

// MySpaceVideoInfo https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/docs/dynamic/space.md#%E8%8E%B7%E5%8F%96%E7%94%A8%E6%88%B7%E7%A9%BA%E9%97%B4%E5%8A%A8%E6%80%81
func (client *BClient) MySpaceVideoInfo() (*m.BResp[SpaceVideo], error) {

	video := &m.BResp[SpaceVideo]{}
	if bs, err := client.GetListAll4SpaceVideo(fmt.Sprintf(mySpaceVideoUrl, c.SelfMid)); err != nil {
		return nil, err
	} else {
		for _, media := range bs.Data.Items {
			me := &DataItem{}
			if err := json.Unmarshal(media, me); err != nil {
				return nil, err
			}
			video.Data.Items = append(video.Data.Items, *me)
		}
	}

	return video, nil
}
