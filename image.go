package vmmanagerapi

import (
	"encoding/json"
	"fmt"

	"github.com/gadost/go-vmmanager-api/types"
)

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

func (a *Api) Images() (*types.ImagesListResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		"/image",
		requestTypeGet,
		DefaultService)

	var h *types.ImagesListResponse
	json.Unmarshal(bodyResp, &h)
	return h, err
}

func (a *Api) Image(imid int) (*types.ImageResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/image/%d", imid),
		requestTypeGet,
		DefaultService)

	var h *types.ImageResponse
	json.Unmarshal(bodyResp, &h)
	return h, err
}

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

func (a *Api) ImageDelete(hid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/image/%d", hid),
		requestTypeDelete,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

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

func (a *Api) ImageNodeDelete(imid int, nid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/image/%d/node/%d", imid, nid),
		requestTypeDelete,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) ImageNodeEnable(imid int, nid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/image/%d/node/%d/enable", imid, nid),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) ImageNodeDisable(imid int, nid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/image/%d/node/%d/disable", imid, nid),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}
