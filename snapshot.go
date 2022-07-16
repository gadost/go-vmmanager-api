package vmmanagerapi

import (
	"encoding/json"
	"fmt"

	"github.com/gadost/go-vmmanager-api/types"
)

func (a *Api) HostSnapshotDelete(hid int) (*types.Task, error) {
	uri := fmt.Sprintf("/host/%d/snapshot", hid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeDelete, DefaultService)

	var r *types.Task
	json.Unmarshal(bodyResp, &r)
	return r, err
}
