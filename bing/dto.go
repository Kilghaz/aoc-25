package bing

type ImageResultThumbnailDto struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type ImageResultValueDto struct {
	ThumbnailUrl   string                  `json:"thumbnailUrl"`
	EncodingFormat string                  `json:"encodingFormat"`
	Thumbnail      ImageResultThumbnailDto `json:"thumbnail"`
}

type ImageResultDto struct {
	TotalEstimatedMatches int                   `json:"totalEstimatedMatches"`
	Value                 []ImageResultValueDto `json:"value"`
}
