package vmmanagerapi

import (
	"encoding/json"
	"fmt"
)

type DiskListResponse struct {
	LastNotify int `json:"last_notify"`
	List       []struct {
		ID         int    `json:"id"`
		Bus        string `json:"bus"`
		Name       string `json:"name"`
		ExpandPart string `json:"expand_part"`
		TargetDev  string `json:"target_dev"`
		SizeMib    int    `json:"size_mib"`
		SizeMibNew int    `json:"size_mib_new"`
		Account    struct {
			ID    int    `json:"id"`
			Email string `json:"email"`
		} `json:"account"`
		Storage struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"storage"`
		Tags []int `json:"tags"`
		Host struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			BootOrder int    `json:"boot_order"`
			IsMain    bool   `json:"is_main"`
			Bus       string `json:"bus"`
		} `json:"host"`
		Node struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"node"`
		Cluster struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"cluster"`
	} `json:"list"`
}

type NewDiskRequest struct {
	Name       string `json:"name"`
	SizeMib    int    `json:"size_mib"`
	Node       int    `json:"node"`
	Account    int    `json:"account"`
	Storage    int    `json:"storage"`
	ExpandPart string `json:"expand_part"`
}

type UpdateDiskRequest struct {
	Name       string `json:"name"`
	Account    int    `json:"account"`
	SizeMib    int    `json:"size_mib"`
	ExpandPart string `json:"expand_part"`
	Pool       int    `json:"pool"`
	Defer      struct {
		Action string `json:"action"`
	} `json:"defer"`
}

type MigrateDiskRequest struct {
	Storage int  `json:"storage"`
	Node    int  `json:"node"`
	Plain   bool `json:"plain"`
}

type BackupIDRequest struct {
	Backup int `json:"backup"`
}

type UpdateBootOrderRequest struct {
	Disks []struct {
		ID        int `json:"id"`
		BootOrder int `json:"boot_order"`
	} `json:"disks"`
}

type DiskParams struct {
	IsMain bool   `json:"is_main"`
	Bus    string `json:"bus"`
}

func (a *Api) DiskList() (*DiskListResponse, error) {
	bodyResp, err := a.NewRequest([]byte(NilPayload), "/disk", requestTypeGet, DefaultService)
	var d *DiskListResponse
	json.Unmarshal(bodyResp, &d)
	return d, err
}

func (a *Api) DiskNew(params *NewDiskRequest) (*Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		"/disk",
		requestTypePost,
		DefaultService)
	var t *Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) DiskUpdate(did int, params *UpdateDiskRequest) (*Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/disk/%d", did),
		requestTypePost,
		DefaultService)
	var t *Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) DiskDelete(did int) (*Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/disk/%d", did),
		requestTypeDelete,
		DefaultService)
	var t *Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) DiskMigrate(did int, params *MigrateDiskRequest) (*Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/disk/%d/migrate", did),
		requestTypePost,
		DefaultService)
	var t *Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// DiskLocalStorageForMigrationGet No example outputs in official documentation
func (a *Api) DiskLocalStorageForMigrationGet(did int) ([]byte, error) {
	bodyResp, err := a.NewRequest([]byte(NilPayload),
		fmt.Sprintf("/disk/%d/migrate/local", did),
		requestTypeGet,
		DefaultService)

	return bodyResp, err
}

// DiskNetworkForMigrationGet No example outputs in official documentation
func (a *Api) DiskNetworkForMigrationGet(did int) ([]byte, error) {
	bodyResp, err := a.NewRequest([]byte(NilPayload),
		fmt.Sprintf("/disk/%d/migrate/network", did),
		requestTypeGet,
		DefaultService)

	return bodyResp, err
}

func (a *Api) DiskResore(did int, bid int) (*Task, error) {
	payload, _ := json.Marshal(&BackupIDRequest{
		Backup: bid,
	})
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/disk/%d/restore", did),
		requestTypePost,
		DefaultService)
	var t *Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) DiskBootOrderUpdate(hid int, params *UpdateBootOrderRequest) (*Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/host/%d/disk_order", hid),
		requestTypePost,
		DefaultService)
	var t *Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) DiskLink(hid int, did int, params *DiskParams) (*Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/host/%d/disk/%d", hid, did),
		requestTypePost,
		DefaultService)
	var t *Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) DiskUnlink(hid int, did int) (*Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/host/%d/disk/%d", hid, did),
		requestTypeDelete,
		DefaultService)
	var t *Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}
