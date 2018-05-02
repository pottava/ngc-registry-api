package registry

import (
	"fmt"

	"gopkg.in/resty.v1"
)

// Images returns NGC docker images
func Images(session, namespace, id string) ([]Image, error) {
	resty.SetCookie(cookie(session))
	resp, err := resty.R().
		SetResult(&ImagesResult{}).
		Get(fmt.Sprintf("/v2/org/%s/repos/%s/images", namespace, id))
	if err != nil {
		return nil, err
	}
	result, ok := resp.Result().(*ImagesResult)
	if !ok {
		return nil, fmt.Errorf("Parse error")
	}
	return result.Images, nil
}

// ImagesResult represents NGC docker images
type ImagesResult struct {
	RequestStatus requestStatus `json:"requestStatus"`
	Images        []Image       `json:"images"`
}

// Image represents NGC docker image
type Image struct {
	Tag         string `json:"tag"`
	Size        int64  `json:"size"`
	UpdatedDate string `json:"updatedDate"`
}
