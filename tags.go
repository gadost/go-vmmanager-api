package vmmanagerapi

import (
	"encoding/json"
	"fmt"

	"github.com/gadost/go-vmmanager-api/types"
)

// Tags return all tags
func (a *Api) Tags() (*types.Tags, error) {
	bodyResp, err := a.NewRequest(NilPayload, "/tag", requestTypeGet, DefaultService)

	var r *types.Tags
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// TagDelete delete tag
func (a *Api) TagDelete(repoid int) (*types.DeletedResponse, error) {
	uri := fmt.Sprintf("/tag/%d", repoid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeDelete, DefaultService)

	var r *types.DeletedResponse
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// TagNew create new tag
func (a *Api) TagNew(params *types.TagNew) (*types.ID, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(payload, "/tag", requestTypePost, DefaultService)

	var r *types.ID
	json.Unmarshal(bodyResp, &r)
	return r, err
}
