package vmmanagerapi

import (
	"encoding/json"
	"fmt"

	"github.com/gadost/go-vmmanager-api/types"
)

// PlatformBackup get platform backup params
func (a *Api) PlatformBackup(bid int) ([]byte, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/platform/backup/%d", bid),
		requestTypeGet,
		DefaultService)

	return bodyResp, err
}

// PlatformBackupInfo get platform backup info
func (a *Api) PlatformBackupInfo(bid int) ([]byte, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/platform/backup/%d/info", bid),
		requestTypeGet,
		DefaultService)

	return bodyResp, err
}

// PlatformBackupSchedule return liost of scheduled platform backups
func (a *Api) PlatformBackupSchedule() (*types.PlatformBackupSchedule, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		"/platform/backup/schedule",
		requestTypeGet,
		DefaultService)

	var t *types.PlatformBackupSchedule
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// PlatformBackupScheduleNew schedule new platform backup
func (a *Api) PlatformBackupScheduleNew(
	params *types.PlatformBackupScheduleRequest) (*types.ID, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		"/platform/backup/schedule",
		requestTypePost,
		DefaultService)

	var t *types.ID
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// PlatformBackupScheduleUpdate update scheduled platform backup
func (a *Api) PlatformBackupScheduleUpdate(scid int,
	params *types.PlatformBackupScheduleRequest) (*types.ID, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/platform/backup/schedule/%d", scid),
		requestTypePost,
		DefaultService)

	var t *types.ID
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// PlatformBackupScheduleRun start scheduled platform backup
func (a *Api) PlatformBackupScheduleRun(scid int) (*types.ID, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/platform/backup/schedule/%d", scid),
		requestTypePost,
		DefaultService)

	var t *types.ID
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// PlatformBackupScheduleDelete delete scheduled platform backup
func (a *Api) PlatformBackupScheduleDelete(scid int) (*types.ID, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/platform/backup/schedule/%d", scid),
		requestTypeDelete,
		DefaultService)

	var t *types.ID
	json.Unmarshal(bodyResp, &t)
	return t, err
}
