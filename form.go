package vmmanagerapi

import (
	"encoding/json"
	"fmt"
)

type MigrateFormResponse struct {
	Nodes []struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		RAMMib struct {
			Unused int `json:"unused"`
		} `json:"ram_mib"`
		Storage struct {
			ID           int    `json:"id"`
			Name         string `json:"name"`
			Type         string `json:"type"`
			AvailableMib int    `json:"available_mib"`
		} `json:"storage"`
		Storages []struct {
			ID           int    `json:"id"`
			Name         string `json:"name"`
			Type         string `json:"type"`
			AvailableMib int    `json:"available_mib"`
		} `json:"storages"`
		Priority            int    `json:"priority"`
		State               string `json:"state"`
		Suitable            bool   `json:"suitable"`
		HostCreationBlocked bool   `json:"host_creation_blocked"`
		HostCount           int    `json:"host_count"`
		HostLimit           int    `json:"host_limit"`
	} `json:"nodes"`
}

func (a *Api) FormMigrate(queryParams *ParamsQuery) (*MigrateFormResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/migrate%s", queryParams.Query),
		requestTypeGet,
		DefaultService)

	var f *MigrateFormResponse
	json.Unmarshal(bodyResp, &f)
	return f, err
}

func (a *Api) FormRandomNames(number int) ([]byte, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/migrate%d", number),
		requestTypeGet,
		DefaultService)

	return bodyResp, err
}

func (a *Api) FormTimeZone() ([]byte, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		"/time_zone",
		requestTypeGet,
		DefaultService)

	return bodyResp, err
}
