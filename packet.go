package vmmanagerapi

import (
	"encoding/json"

	"github.com/gadost/go-vmmanager-api/types"
)

// LicensePacket set license key
func (a *Api) LicensePacket(params *types.License) (*types.ID, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		"/packet",
		requestTypePost,
		DefaultService)

	var t *types.ID
	json.Unmarshal(bodyResp, &t)
	return t, err
}
