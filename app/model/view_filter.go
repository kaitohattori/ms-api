package model

type ViewFilter struct {
	VideoId *int    `json:"videoId,omitempty"`
	UserId  *string `json:"userId,omitempty"`
}

func NewViewFilter(videoId *int, userId *string) ViewFilter {
	return ViewFilter{VideoId: videoId, UserId: userId}
}
