package vmmanagerapi

import (
	"encoding/json"
	"fmt"

	"github.com/gadost/go-vmmanager-api/types"
)

func (a *Api) Migrate(queryParams *ParamsQuery) (*types.MigrateFormResponse, error) {
	uri := fmt.Sprintf("/migrate%s", queryParams.Query)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	var f *types.MigrateFormResponse
	json.Unmarshal(bodyResp, &f)
	return f, err
}

func (a *Api) RandomNames(number int) ([]byte, error) {
	uri := fmt.Sprintf("/random_name/%d", number)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	return bodyResp, err
}

func (a *Api) TimeZone() ([]byte, error) {
	bodyResp, err := a.NewRequest(NilPayload, "/time_zone", requestTypeGet, DefaultService)

	return bodyResp, err
}
