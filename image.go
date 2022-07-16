package vmmanagerapi

import (
	"encoding/json"
	"fmt"

	"github.com/gadost/go-vmmanager-api/types"
)

// ImageNew create new image
func (a *Api) ImageNew(hid int, params *types.Image) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/host/%d/image/create", hid),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// Images return images
func (a *Api) Images() (*types.ImagesListResponse, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		"/image",
		requestTypeGet,
		DefaultService)

	var h *types.ImagesListResponse
	json.Unmarshal(bodyResp, &h)
	return h, err
}

// Image return image params
func (a *Api) Image(imid int) (*types.ImageResponse, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/image/%d", imid),
		requestTypeGet,
		DefaultService)

	var h *types.ImageResponse
	json.Unmarshal(bodyResp, &h)
	return h, err
}

// ImageUpdate update image params
func (a *Api) ImageUpdate(imid int, params *types.Image) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/image/%d", imid),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// ImageDelete delete image
func (a *Api) ImageDelete(hid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/image/%d", hid),
		requestTypeDelete,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// ImageStateUpdate update image state
func (a *Api) ImageStateUpdate(imid int, params *types.ImageState) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/image/%d/change_state", imid),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// ImageCopy copy image
func (a *Api) ImageCopy(imid int, params *types.Node) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/image/%d/copy", imid),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// ImageMove move image
func (a *Api) ImageMove(imid int, params *types.Move) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/image/%d/move", imid),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// ImageNodeDelete delete image on node
func (a *Api) ImageNodeDelete(imid int, nid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/image/%d/node/%d", imid, nid),
		requestTypeDelete,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// ImageNodeEnable enable image on node
func (a *Api) ImageNodeEnable(imid int, nid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/image/%d/node/%d/enable", imid, nid),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// ImageNodeDisable disable image on node
func (a *Api) ImageNodeDisable(imid int, nid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/image/%d/node/%d/disable", imid, nid),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}
