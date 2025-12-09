package bing

import (
	"encoding/json"
	"github.com/samber/lo"
	"io"
	"net/http"
	"net/url"
)

var baseUrl = "https://api.bing.microsoft.com/v7.0/images/search"
var keyHeaderName = "Ocp-Apim-Subscription-Key"

type ImageSearchClient struct {
	Search func(query string) []ImageSearchResult
}

type Thumbnail struct {
	Url    string
	Width  int
	Height int
	Format string
}

type ImageSearchResult struct {
	Thumbnail Thumbnail
}

func CreateImageSearchClient(apiKey string) *ImageSearchClient {
	httpClient := &http.Client{}

	var search = func(query string) []ImageSearchResult {
		requestUrl := baseUrl + "?q=" + url.QueryEscape(query)
		req, _ := http.NewRequest("GET", requestUrl, nil)
		req.Header.Add(keyHeaderName, apiKey)
		resp, _ := httpClient.Do(req)
		bytes, _ := io.ReadAll(resp.Body)
		var dto ImageResultDto
		_ = json.Unmarshal(bytes, &dto)
		return lo.Map(dto.Value, func(result ImageResultValueDto, _ int) ImageSearchResult {
			return ImageSearchResult{
				Thumbnail: Thumbnail{
					Url:    result.ThumbnailUrl,
					Width:  result.Thumbnail.Width,
					Height: result.Thumbnail.Height,
					Format: result.EncodingFormat,
				},
			}
		})
	}

	return &ImageSearchClient{
		Search: search,
	}
}
