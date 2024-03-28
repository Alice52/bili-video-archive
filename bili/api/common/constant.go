package c

const (
	SelfMid     = "316806518"
	PageSizeMax = 30
	PageSizeMin = 20
)

// Fnval https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/video/videostream_url.md#fnval%E8%A7%86%E9%A2%91%E6%B5%81%E6%A0%BC%E5%BC%8F%E6%A0%87%E8%AF%86
type Fnval int64

const (
	FnvalMP4  Fnval = 1
	FnvalDash Fnval = 16
	FnvalHDR  Fnval = 64
	Fnval4K   Fnval = 128

	FnvalAudio64K  Fnval = 30216
	FnvalAudio132K Fnval = 30232
	FnvalAudio192K Fnval = 30280
)

type Qn int64

func (qn Qn) String() string {
	switch qn {
	case Qn240P:
		return "240P"
	case Qn360P:
		return "360P"
	case Qn480P:
		return "480P"
	case Qn720P:
		return "720P"
	case Qn720P60:
		return "720P60"
	case Qn1080P:
		return "1080P"
	case Qn1080PPlus:
		return "1080P+"
	case Qn1080P60:
		return "1080P60"
	case Qn4k:
		return "4K"
	case QnAudio64K:
		return "64K"
	case QnAudio132K:
		return "132K"
	case QnAudio192K:
		return "192K"
	case QnAudioDolby:
		return "Dolby"
	case QnAudioHiRes:
		return "Hi-Res"
	default:
		return ""
	}
}

const (
	Qn240P      Qn = 6
	Qn360P      Qn = 16
	Qn480P      Qn = 32
	Qn720P      Qn = 64
	Qn720P60    Qn = 74
	Qn1080P     Qn = 80
	Qn1080PPlus Qn = 112
	Qn1080P60   Qn = 116
	Qn4k        Qn = 120

	QnAudio64K   Qn = 30216
	QnAudio132K  Qn = 30232
	QnAudio192K  Qn = 30280
	QnAudioDolby Qn = 30250
	QnAudioHiRes Qn = 30251
)
