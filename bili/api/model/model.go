package m

type BPCResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	TTL     int64  `json:"ttl"`
}

type BPResp[T any] struct {
	BPCResp
	Data []T `json:"data"`
}

type BResp[T any] struct {
	BPCResp
	Data T `json:"data"`
}

type BRList[T any] struct {
	HasMore bool  `json:"has_more"`
	Info    T     `json:"info"`
	Medias  []T   `json:"medias"`
	TTL     int64 `json:"ttl"`
}

type BRListOfItem[T any] struct {
	HasMore bool   `json:"has_more"`
	Items   []T    `json:"items"`
	Offset  string `json:"offset"`
}
