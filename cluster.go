package vmmanagerapi

import (
	"encoding/json"
	"fmt"
)

type ClusterListResponse struct {
	LastNotify int `json:"last_notify"`
	List       []struct {
		VirtualizationType string `json:"virtualization_type"`
		OsStoragePath      string `json:"os_storage_path,omitempty"`
		ImageStoragePath   string `json:"image_storage_path,omitempty"`
		CPUNumber          struct {
			Total int `json:"total"`
			Used  int `json:"used"`
		} `json:"cpu_number,omitempty"`
		HostCount        int    `json:"host_count"`
		AccountHostCount int    `json:"account_host_count,omitempty"`
		HostPerNodeLimit int    `json:"host_per_node_limit,omitempty"`
		Name             string `json:"name"`
		Comment          string `json:"comment,omitempty"`
		RAMMib           struct {
			Total int `json:"total"`
			Used  int `json:"used"`
		} `json:"ram_mib,omitempty"`
		StorageMib struct {
			Total int `json:"total"`
			Used  int `json:"used"`
		} `json:"storage_mib,omitempty"`
		TimeZone               string `json:"time_zone,omitempty"`
		State                  string `json:"state,omitempty"`
		DataCenter             string `json:"data_center,omitempty"`
		HostDistributionPolicy string `json:"host_distribution_policy,omitempty"`
		HostFilter             []struct {
			Entity     string `json:"entity"`
			Expression string `json:"expression"`
		} `json:"host_filter,omitempty"`
		NodeCount       int `json:"node_count,omitempty"`
		ID              int `json:"id"`
		Overselling     int `json:"overselling,omitempty"`
		BackupLocations []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"backup_locations,omitempty"`
		DomainTemplate   string `json:"domain_template,omitempty"`
		DatacenterParams struct {
			Gateway      string `json:"gateway"`
			Mac          string `json:"mac"`
			BgpCommunity string `json:"bgp_community"`
			BgpAs        int    `json:"bgp_as"`
			BgpSessions  []struct {
				IP      string `json:"ip"`
				BgpAs   int    `json:"bgp_as"`
				Comment string `json:"comment"`
			} `json:"bgp_sessions"`
			BgpCommunityV6 string `json:"bgp_community_v6"`
			BgpAsV6        int    `json:"bgp_as_v6"`
			BgpSessionsV6  []struct {
				IP      string `json:"ip"`
				BgpAs   int    `json:"bgp_as"`
				Comment string `json:"comment"`
			} `json:"bgp_sessions_v6"`
		} `json:"datacenter_params,omitempty"`
		Storages []struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Type     string `json:"type"`
			Enabled  bool   `json:"enabled"`
			IsMain   bool   `json:"is_main"`
			VirtPool struct {
				VMStoragePath   string `json:"vm_storage_path"`
				VolumeGroup     string `json:"volume_group"`
				ZfsPoolName     string `json:"zfs_pool_name"`
				RbdPoolName     string `json:"rbd_pool_name"`
				RbdUser         string `json:"rbd_user"`
				PlacementGroups int    `json:"placement_groups"`
				NetworkDiskPath string `json:"network_disk_path"`
			} `json:"virt_pool"`
			Tags []int `json:"tags"`
		} `json:"storages,omitempty"`
		Storage struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Type     string `json:"type"`
			Enabled  bool   `json:"enabled"`
			IsMain   bool   `json:"is_main"`
			VirtPool struct {
				VMStoragePath   string `json:"vm_storage_path"`
				VolumeGroup     string `json:"volume_group"`
				ZfsPoolName     string `json:"zfs_pool_name"`
				RbdPoolName     string `json:"rbd_pool_name"`
				RbdUser         string `json:"rbd_user"`
				PlacementGroups int    `json:"placement_groups"`
				NetworkDiskPath string `json:"network_disk_path"`
			} `json:"virt_pool"`
			Tags []int `json:"tags"`
		} `json:"storage,omitempty"`
		DNSServers  []string `json:"dns_servers,omitempty"`
		DcNetworks  []int    `json:"dc_networks,omitempty"`
		NodeNetwork struct {
			Timeout int    `json:"timeout"`
			Gateway string `json:"gateway"`
		} `json:"node_network,omitempty"`
		VxlanMode     string `json:"vxlan_mode"`
		VxlanSettings string `json:"vxlan_settings"`
		Vxlans        []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"vxlans"`
		DatacenterType string `json:"datacenter_type,omitempty"`
	} `json:"list"`
}

type ClusterResponse struct {
	VirtualizationType string `json:"virtualization_type"`
	OsStoragePath      string `json:"os_storage_path"`
	ImageStoragePath   string `json:"image_storage_path"`
	CPUNumber          struct {
		Total int `json:"total"`
		Used  int `json:"used"`
	} `json:"cpu_number"`
	HostCount        int    `json:"host_count"`
	AccountHostCount int    `json:"account_host_count"`
	HostPerNodeLimit int    `json:"host_per_node_limit"`
	Name             string `json:"name"`
	Comment          string `json:"comment"`
	RAMMib           struct {
		Total int `json:"total"`
		Used  int `json:"used"`
	} `json:"ram_mib"`
	StorageMib struct {
		Total int `json:"total"`
		Used  int `json:"used"`
	} `json:"storage_mib"`
	TimeZone               string `json:"time_zone"`
	State                  string `json:"state"`
	DataCenter             string `json:"data_center"`
	HostDistributionPolicy string `json:"host_distribution_policy"`
	HostFilter             []struct {
		Entity     string `json:"entity"`
		Expression string `json:"expression"`
	} `json:"host_filter"`
	NodeCount       int `json:"node_count"`
	ID              int `json:"id"`
	Overselling     int `json:"overselling"`
	BackupLocations []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"backup_locations"`
	DomainTemplate   string `json:"domain_template"`
	DatacenterParams struct {
		Gateway      string `json:"gateway"`
		Mac          string `json:"mac"`
		BgpCommunity string `json:"bgp_community"`
		BgpAs        int    `json:"bgp_as"`
		BgpSessions  []struct {
			IP      string `json:"ip"`
			BgpAs   int    `json:"bgp_as"`
			Comment string `json:"comment"`
		} `json:"bgp_sessions"`
		BgpCommunityV6 string `json:"bgp_community_v6"`
		BgpAsV6        int    `json:"bgp_as_v6"`
		BgpSessionsV6  []struct {
			IP      string `json:"ip"`
			BgpAs   int    `json:"bgp_as"`
			Comment string `json:"comment"`
		} `json:"bgp_sessions_v6"`
	} `json:"datacenter_params"`
	Storages []struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Type     string `json:"type"`
		Enabled  bool   `json:"enabled"`
		IsMain   bool   `json:"is_main"`
		VirtPool struct {
			VMStoragePath   string `json:"vm_storage_path"`
			VolumeGroup     string `json:"volume_group"`
			ZfsPoolName     string `json:"zfs_pool_name"`
			RbdPoolName     string `json:"rbd_pool_name"`
			RbdUser         string `json:"rbd_user"`
			PlacementGroups int    `json:"placement_groups"`
			NetworkDiskPath string `json:"network_disk_path"`
		} `json:"virt_pool"`
		Tags []int `json:"tags"`
	} `json:"storages"`
	Storage struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Type     string `json:"type"`
		Enabled  bool   `json:"enabled"`
		IsMain   bool   `json:"is_main"`
		VirtPool struct {
			VMStoragePath   string `json:"vm_storage_path"`
			VolumeGroup     string `json:"volume_group"`
			ZfsPoolName     string `json:"zfs_pool_name"`
			RbdPoolName     string `json:"rbd_pool_name"`
			RbdUser         string `json:"rbd_user"`
			PlacementGroups int    `json:"placement_groups"`
			NetworkDiskPath string `json:"network_disk_path"`
		} `json:"virt_pool"`
		Tags []int `json:"tags"`
	} `json:"storage"`
	DNSServers  []string `json:"dns_servers"`
	DcNetworks  []int    `json:"dc_networks"`
	NodeNetwork struct {
		Timeout int    `json:"timeout"`
		Gateway string `json:"gateway"`
	} `json:"node_network"`
	VxlanMode     string `json:"vxlan_mode"`
	VxlanSettings string `json:"vxlan_settings"`
	Vxlans        []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"vxlans"`
}

type CreateClusterRequest struct {
	VirtualizationType string `json:"virtualization_type"`
	DatacenterType     string `json:"datacenter_type"`
	Name               string `json:"name"`
	Comment            string `json:"comment"`
	IsoEnabled         bool   `json:"iso_enabled"`
	ManageDiskEnabled  bool   `json:"manage_disk_enabled"`
	TimeZone           string `json:"time_zone"`
	Interfaces         []struct {
		Interface int   `json:"interface"`
		Ippool    []int `json:"ippool"`
	} `json:"interfaces"`
	Os                     []int  `json:"os"`
	BackupLocations        []int  `json:"backup_locations"`
	OsStoragePath          string `json:"os_storage_path"`
	ImageStoragePath       string `json:"image_storage_path"`
	Overselling            int    `json:"overselling"`
	HostDistributionPolicy string `json:"host_distribution_policy"`
	HostFilter             []struct {
		Entity     string `json:"entity"`
		Expression string `json:"expression"`
	} `json:"host_filter"`
	DomainTemplate   string `json:"domain_template"`
	DatacenterParams struct {
		Gateway      string `json:"gateway"`
		Mac          string `json:"mac"`
		BgpCommunity string `json:"bgp_community"`
		BgpAs        int    `json:"bgp_as"`
		BgpSessions  []struct {
			IP      string `json:"ip"`
			BgpAs   int    `json:"bgp_as"`
			Comment string `json:"comment"`
		} `json:"bgp_sessions"`
		BgpCommunityV6 string `json:"bgp_community_v6"`
		BgpAsV6        int    `json:"bgp_as_v6"`
		BgpSessionsV6  []struct {
			IP      string `json:"ip"`
			BgpAs   int    `json:"bgp_as"`
			Comment string `json:"comment"`
		} `json:"bgp_sessions_v6"`
	} `json:"datacenter_params"`
	DomainChangeAllowed bool     `json:"domain_change_allowed"`
	ProxyEnabled        bool     `json:"proxy_enabled"`
	HostPerNodeLimit    int      `json:"host_per_node_limit"`
	DNSServers          []string `json:"dns_servers"`
	NodeNetwork         struct {
		Timeout int    `json:"timeout"`
		Gateway string `json:"gateway"`
	} `json:"node_network"`
	VxlanMode     string `json:"vxlan_mode"`
	VxlanSettings struct {
		BgpCommunity string `json:"bgp_community"`
		BgpAs        int    `json:"bgp_as"`
		BgpSessions  []struct {
			IP      string `json:"ip"`
			BgpAs   int    `json:"bgp_as"`
			Comment string `json:"comment"`
		} `json:"bgp_sessions"`
		BgpCommunityV6 string `json:"bgp_community_v6"`
		BgpAsV6        int    `json:"bgp_as_v6"`
		BgpSessionsV6  []struct {
			IP      string `json:"ip"`
			BgpAs   int    `json:"bgp_as"`
			Comment string `json:"comment"`
		} `json:"bgp_sessions_v6"`
	} `json:"vxlan_settings"`
}

type Tasks struct {
	ID    int   `json:"id"`
	Tasks []int `json:"tasks"`
}

type DeletedResponse struct {
	Deleted []struct {
		ID int `json:"id"`
	} `json:"deleted"`
}

type DCNetworksResponse struct {
	DcNetworks []int `json:"dc_networks"`
}

type QemuVersion struct {
	QemuVersion string `json:"qemu_version"`
}

type IPPoolRequest struct {
	Interfaces []struct {
		Interface int   `json:"interface"`
		Ippool    []int `json:"ippool"`
	} `json:"interfaces"`
}

type Name struct {
	Name string `json:"name"`
}

type SSHKeysRequest struct {
	SSHKeys []int `json:"ssh_keys"`
}

type AttachStorageToClusterRequest struct {
	VirtPool struct {
		VMStoragePath   string `json:"vm_storage_path"`
		VolumeGroup     string `json:"volume_group"`
		ZfsPoolName     string `json:"zfs_pool_name"`
		RbdPoolName     string `json:"rbd_pool_name"`
		RbdUser         string `json:"rbd_user"`
		PlacementGroups int    `json:"placement_groups"`
		NetworkDiskPath string `json:"network_disk_path"`
	} `json:"virt_pool"`
	IsMain         bool `json:"is_main"`
	HddOverselling int  `json:"hdd_overselling"`
	IgnoreChecks   bool `json:"ignore_checks"`
}

type StoragesTasks struct {
	ID            int   `json:"id"`
	Task          int   `json:"task"`
	StoragesTasks []int `json:"storages_tasks"`
}

type HDDOversellingRequest struct {
	Value float32 `json:"value"`
}

type HDDOversellingResponse struct {
	Value float32 `json:"value"`
}

type VXLanResponse struct {
	LastNotify int `json:"last_notify"`
	List       []struct {
		Account struct {
			ID    int    `json:"id"`
			Email string `json:"email"`
		} `json:"account"`
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Tag   int    `json:"tag"`
		Hosts []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Node struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"node"`
		} `json:"hosts"`
		Ranges []string `json:"ranges"`
	} `json:"list"`
}

type NodesVXLanResponse struct {
	LastNotify int `json:"last_notify"`
	List       []struct {
		Node struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"node"`
		Vxlans []struct {
			ID    int    `json:"id"`
			Name  string `json:"name"`
			Tag   int    `json:"tag"`
			Hosts []struct {
				ID     int      `json:"id"`
				Name   string   `json:"name"`
				Ranges []string `json:"ranges"`
			} `json:"hosts"`
		} `json:"vxlans"`
	} `json:"list"`
}

type UsersVXLanResponse struct {
	LastNotify int `json:"last_notify"`
	List       []struct {
		Account struct {
			ID    int    `json:"id"`
			Email string `json:"email"`
		} `json:"account"`
		Hosts []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Node struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"node"`
			Vxlans []struct {
				ID     int      `json:"id"`
				Name   string   `json:"name"`
				Tag    int      `json:"tag"`
				Ranges []string `json:"ranges"`
			} `json:"vxlans"`
		} `json:"hosts"`
	} `json:"list"`
}

type SSHKeyResponse struct {
	LastNotify int `json:"last_notify"`
	List       []struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		SSHKey   string `json:"ssh_key"`
		Clusters []int  `json:"clusters"`
	} `json:"list"`
}

type SSHKeyRequest struct {
	Name     string `json:"name"`
	SSHKey   string `json:"ssh_key"`
	Clusters []int  `json:"clusters"`
}

func (a *Api) ClusterList() (*ClusterListResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		"/cluster",
		requestTypeGet,
		DefaultService)

	var c *ClusterListResponse
	json.Unmarshal(bodyResp, &c)

	return c, err
}

func (a *Api) ClusterNew(params *CreateClusterRequest) (*Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		"/cluster",
		requestTypePost,
		DefaultService)
	var t *Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) ClusterGet(cid int) (*ClusterResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/cluster/%d", cid),
		requestTypeGet,
		DefaultService)
	var c *ClusterResponse
	json.Unmarshal(bodyResp, &c)
	return c, err
}

func (a *Api) ClusterUpdate(cid int, params *CreateClusterRequest) (*Tasks, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/cluster/%d", cid),
		requestTypePost,
		DefaultService)
	var t *Tasks
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) ClusterDelete(cid int) (*DeletedResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/cluster/%d", cid),
		requestTypeDelete,
		DefaultService)
	var d *DeletedResponse
	json.Unmarshal(bodyResp, &d)
	return d, err
}

func (a *Api) ClusterConnectDCNetwork(cid int) (*DCNetworksResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/cluster/%d/dc_network", cid),
		requestTypePost,
		DefaultService)
	var d *DCNetworksResponse
	json.Unmarshal(bodyResp, &d)
	return d, err
}

func (a *Api) ClusterFixFRR(cid int) error {
	_, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/cluster/%d/fix_frr", cid),
		requestTypePost,
		DefaultService)

	return err
}

func (a *Api) ClusterHaAgentUpdate(cid int) (*Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/cluster/%d/ha_agent_update", cid),
		requestTypePost,
		DefaultService)
	var t *Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) ClusterInternalEditUpdate(cid int, params *QemuVersion) (*Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/cluster/%d/internal_edit", cid),
		requestTypePost,
		DefaultService)
	var t *Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) ClusterConnectIPPool(cid int, params *IPPoolRequest) error {
	payload, _ := json.Marshal(params)
	_, err := a.NewRequest(
		payload,
		fmt.Sprintf("/cluster/%d/ippool", cid),
		requestTypePost,
		DefaultService)

	return err
}

func (a *Api) ClusterLocalStorageGet(cid int) error {
	_, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/cluster/%d/local_storage", cid),
		requestTypeGet,
		DefaultService)

	return err
}

func (a *Api) ClusterSettingsUpdate(cid int, params *Name) error {
	payload, _ := json.Marshal(params)
	_, err := a.NewRequest(
		payload,
		fmt.Sprintf("/cluster/%d/settings", cid),
		requestTypePost,
		DefaultService)

	return err
}

func (a *Api) ClusterSSHKeysUpdate(cid int, params *SSHKeysRequest) error {
	payload, _ := json.Marshal(params)
	_, err := a.NewRequest(
		payload,
		fmt.Sprintf("/cluster/%d/ssh_key", cid),
		requestTypePost,
		DefaultService)

	return err
}

func (a *Api) ClusterStorageAttach(
	cid int, sid int, params *AttachStorageToClusterRequest) (
	*StoragesTasks, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/cluster/%d/storage/%d", cid, sid),
		requestTypePost,
		DefaultService)

	var r *StoragesTasks
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ClusterStorageDetach(cid int, sid int) (*DeletedResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/cluster/%d/storage/%d", cid, sid),
		requestTypeDelete,
		DefaultService)
	var d *DeletedResponse
	json.Unmarshal(bodyResp, &d)
	return d, err
}

func (a *Api) ClusterStorageCephConnect(
	cid int, sid int) (
	*StoragesTasks, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/cluster/%d/storage/%d/ceph_connect", cid, sid),
		requestTypePost,
		DefaultService)

	var r *StoragesTasks
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ClusterStorageCheck(cid int, sid int) (*Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/cluster/%d/storage/%d/check", cid, sid),
		requestTypePost,
		DefaultService)

	var r *Task
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ClusterStorageDisable(cid int, sid int) (*Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/cluster/%d/storage/%d/disable", cid, sid),
		requestTypePost,
		DefaultService)

	var r *Task
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ClusterStorageEnable(cid int, sid int) (*Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/cluster/%d/storage/%d/enable", cid, sid),
		requestTypePost,
		DefaultService)

	var r *Task
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ClusteHDDOversellingUpdate(
	cid int, sid int, params *HDDOversellingRequest) (
	*HDDOversellingResponse, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/cluster/%d/storage/%d/hdd_overselling", cid, sid),
		requestTypePost,
		DefaultService)

	var r *HDDOversellingResponse
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ClusterStorageMakeMain(cid int, sid int) (*Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/cluster/%d/storage/%d/make_main", cid, sid),
		requestTypePost,
		DefaultService)

	var r *Task
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ClusterVXLanGet(cid int) (*VXLanResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/cluster/%d/vxlan", cid),
		requestTypeGet,
		DefaultService)

	var r *VXLanResponse
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ClusterNodesVXLanGet(cid int) (*NodesVXLanResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/cluster/%d/vxlan/node", cid),
		requestTypeGet,
		DefaultService)

	var r *NodesVXLanResponse
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ClusterUsersVXLanGet(cid int) (*UsersVXLanResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/cluster/%d/vxlan/user", cid),
		requestTypeGet,
		DefaultService)

	var r *UsersVXLanResponse
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ClusterSSHKeyList() (*SSHKeyResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		"/ssh_key",
		requestTypeGet,
		DefaultService)

	var r *SSHKeyResponse
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ClusterAddSSHKey(
	params *SSHKeyRequest) (
	*Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		"/ssh_key",
		requestTypePost,
		DefaultService)

	var r *Task
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ClusterDeleteSSHKey(
	kid int) (
	*Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/ssh_key/%d", kid),
		requestTypeDelete,
		DefaultService)

	var r *Task
	json.Unmarshal(bodyResp, &r)
	return r, err
}
