package types

// Auth

type Password struct {
	Password string `json:"password"`
}

type InviteUser struct {
	Password        string `json:"password"`
	Lang            string `json:"lang"`
	AdditionalProp1 struct {
	} `json:"additionalProp1"`
}

type TokenLifetime struct {
	ExpiresAt   string `json:"expires_at"`
	Description string `json:"description"`
}

//Backups

type BackupList struct {
	LastNotify int `json:"last_notify"`
	List       []struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Account struct {
			ID    int    `json:"id"`
			Email string `json:"email"`
		} `json:"account"`
		Disk struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"disk"`
		ActualSizeMib    int `json:"actual_size_mib"`
		EstimatedSizeMib int `json:"estimated_size_mib"`
		Cluster          struct {
			ID               int    `json:"id"`
			DatacenterType   string `json:"datacenter_type"`
			ImageStoragePath string `json:"image_storage_path"`
			Name             string `json:"name"`
		} `json:"cluster"`
		Host           int    `json:"host"`
		Schedule       int    `json:"schedule"`
		DateCreate     string `json:"date_create"`
		AvailableUntil string `json:"available_until"`
		Comment        string `json:"comment"`
		BackupLocation struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"backup_location"`
	} `json:"list"`
}

type Backup struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Account        string `json:"account"`
	ForAll         bool   `json:"for_all"`
	Type           string `json:"type"`
	Comment        string `json:"comment"`
	State          string `json:"state"`
	ExpandPart     string `json:"expand_part"`
	BackupLocation struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"backup_location"`
}

type ChangeBackup struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

type Task struct {
	ID   int `json:"id"`
	Task int `json:"task"`
}
type RelocateTask struct {
	ID           int `json:"id"`
	Task         int `json:"task"`
	RelocateTask int `json:"relocate_task"`
}

type Node struct {
	Node int `json:"node"`
}

type MoveBackup struct {
	Source      int `json:"source"`
	Destination int `json:"destination"`
}

type RelocateBackup struct {
	Destination int `json:"destination"`
}

type State struct {
	State string `json:"state"`
}

type DiskBackup struct {
	Name            string `json:"name"`
	Comment         string `json:"comment"`
	BackupLocations []int  `json:"backup_locations"`
	Schedule        int    `json:"schedule"`
}

type VMBackupParams struct {
	Name            string `json:"name"`
	Comment         string `json:"comment"`
	BackupLocations []int  `json:"backup_locations"`
	Schedule        int    `json:"schedule"`
}

type BackupByHostIDResponse struct {
	LastNotify int `json:"last_notify"`
	List       []struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		Account        string `json:"account"`
		ForAll         bool   `json:"for_all"`
		Type           string `json:"type"`
		Comment        string `json:"comment"`
		State          string `json:"state"`
		ExpandPart     string `json:"expand_part"`
		BackupLocation struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"backup_location"`
	} `json:"list"`
}

type BackupLocationResponse struct {
	LastNotify int `json:"last_notify"`
	List       []struct {
		ID               int    `json:"id"`
		Name             string `json:"name"`
		Comment          string `json:"comment"`
		State            string `json:"state"`
		QuotaMib         int    `json:"quota_mib"`
		Type             string `json:"type"`
		ConnectionParams struct {
		} `json:"connection_params"`
		Clusters []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"clusters"`
	} `json:"list"`
}

type DeletedBackupResponse struct {
	Deleted []struct {
		ID int `json:"id"`
	} `json:"deleted"`
}

type BackupLocationParams struct {
	Name             string `json:"name"`
	Comment          string `json:"comment"`
	QuotaMib         int    `json:"quota_mib"`
	Type             string `json:"type"`
	ConnectionParams struct {
	} `json:"connection_params"`
	Clusters            []int `json:"clusters"`
	Schedules           []int `json:"schedules"`
	SkipConnectionCheck bool  `json:"skip_connection_check"`
}

// Cluster

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

// Disk

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

// Host

type HostsResponse struct {
	LastNotify int `json:"last_notify"`
	List       []struct {
		ExpandPart string `json:"expand_part"`
		ID         int    `json:"id"`
		Name       string `json:"name"`
		IP4        []struct {
			IP string `json:"ip"`
		} `json:"ip4"`
		Interfaces []struct {
			HostInterface string `json:"host_interface"`
			Mac           string `json:"mac"`
			NodeInterface int    `json:"node_interface"`
		} `json:"interfaces"`
		Node struct {
			ID     int    `json:"id"`
			Name   string `json:"name"`
			IPAddr string `json:"ip_addr"`
		} `json:"node"`
		Cluster struct {
			ID int `json:"id"`
		} `json:"cluster"`
		State   string `json:"state"`
		Domain  string `json:"domain"`
		Account struct {
			ID    int    `json:"id"`
			Email string `json:"email"`
		} `json:"account"`
		Comment string `json:"comment"`
		Disk    struct {
			ID         int `json:"id"`
			DiskMib    int `json:"disk_mib"`
			DiskMibNew int `json:"disk_mib_new"`
		} `json:"disk"`
		DiskCount                 int      `json:"disk_count"`
		CPUNumber                 int      `json:"cpu_number"`
		CPUNumberNew              int      `json:"cpu_number_new"`
		RAMMib                    int      `json:"ram_mib"`
		RAMMibNew                 int      `json:"ram_mib_new"`
		NetBandwidthMbitps        int      `json:"net_bandwidth_mbitps"`
		NetBandwidthMbitpsChanged bool     `json:"net_bandwidth_mbitps_changed"`
		IPAutomation              string   `json:"ip_automation"`
		NetIsSynced               bool     `json:"net_is_synced"`
		Tags                      []string `json:"tags"`
		OsName                    string   `json:"os_name"`
		OsGroup                   string   `json:"os_group"`
		Uptime                    int      `json:"uptime"`
		RescueMode                bool     `json:"rescue_mode"`
		IsoMounted                bool     `json:"iso_mounted"`
		CPUMode                   string   `json:"cpu_mode"`
		Nesting                   bool     `json:"nesting"`
		CPUCustomModel            string   `json:"cpu_custom_model"`
		CPUWeight                 int      `json:"cpu_weight"`
		IoWeight                  int      `json:"io_weight"`
		IoReadMbitps              int      `json:"io_read_mbitps"`
		IoWriteMbitps             int      `json:"io_write_mbitps"`
		IoReadIops                int      `json:"io_read_iops"`
		IoWriteIops               int      `json:"io_write_iops"`
		NetInMbitps               int      `json:"net_in_mbitps"`
		NetOutMbitps              int      `json:"net_out_mbitps"`
		NetWeight                 int      `json:"net_weight"`
		AntiSpoofing              bool     `json:"anti_spoofing"`
		Disabled                  bool     `json:"disabled"`
		TCPConnectionsIn          int      `json:"tcp_connections_in"`
		TCPConnectionsOut         int      `json:"tcp_connections_out"`
		ProcessNumber             int      `json:"process_number"`
		Vxlan                     struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Tag  int    `json:"tag"`
		} `json:"vxlan"`
		FirewallRules []struct {
			Action    string   `json:"action"`
			Direction string   `json:"direction"`
			Protocols []string `json:"protocols"`
			Portstart int      `json:"portstart"`
			Portend   int      `json:"portend"`
		} `json:"firewall_rules"`
		HasNonameIface  bool `json:"has_noname_iface"`
		VM5Restrictions struct {
			NetIfaceCount      bool `json:"net_iface_count"`
			NatOrExtra         bool `json:"nat_or_extra"`
			Ipv6               bool `json:"ipv6"`
			UnsupportedStorage bool `json:"unsupported_storage"`
			Iso                bool `json:"iso"`
			Snapshot           bool `json:"snapshot"`
		} `json:"vm5_restrictions"`
	} `json:"list"`
}

type NewHostRequest struct {
	Name       string `json:"name"`
	Cluster    int    `json:"cluster"`
	Node       int    `json:"node"`
	Storage    int    `json:"storage"`
	Account    int    `json:"account"`
	Domain     string `json:"domain"`
	Preset     int    `json:"preset"`
	Os         int    `json:"os"`
	Image      int    `json:"image"`
	ExpandPart string `json:"expand_part"`
	RecipeList []struct {
		Recipe       int `json:"recipe"`
		RecipeParams []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"recipe_params"`
	} `json:"recipe_list"`
	Recipe       int `json:"recipe"`
	RecipeParams []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"recipe_params"`
	IgnoreRecipeFilters bool   `json:"ignore_recipe_filters"`
	Password            string `json:"password"`
	RAMMib              int    `json:"ram_mib"`
	HddMib              int    `json:"hdd_mib"`
	Disks               []struct {
		Name       string `json:"name"`
		SizeMib    int    `json:"size_mib"`
		BootOrder  int    `json:"boot_order"`
		ExpandPart string `json:"expand_part"`
		Storage    int    `json:"storage"`
		Tags       []int  `json:"tags"`
	} `json:"disks"`
	CPUNumber          int `json:"cpu_number"`
	NetBandwidthMbitps int `json:"net_bandwidth_mbitps"`
	IPAddr             struct {
		Name              string `json:"name"`
		IPPool            int    `json:"ip_pool"`
		IPNetwork         int    `json:"ip_network"`
		WithoutAllocation bool   `json:"without_allocation"`
	} `json:"ip_addr"`
	Ipv6Enabled bool   `json:"ipv6_enabled"`
	Ipv6Pool    []int  `json:"ipv6_pool"`
	Ipv6Prefix  int    `json:"ipv6_prefix"`
	Ipv4Pool    []int  `json:"ipv4_pool"`
	Ipv4Number  int    `json:"ipv4_number"`
	Comment     string `json:"comment"`
	Interfaces  []struct {
	} `json:"interfaces"`
	CustomInterfaces []struct {
	} `json:"custom_interfaces"`
	CPUMode           string `json:"cpu_mode"`
	Nesting           bool   `json:"nesting"`
	CPUCustomModel    string `json:"cpu_custom_model"`
	CPUWeight         int    `json:"cpu_weight"`
	IoWeight          int    `json:"io_weight"`
	IoReadMbitps      int    `json:"io_read_mbitps"`
	IoWriteMbitps     int    `json:"io_write_mbitps"`
	IoReadIops        int    `json:"io_read_iops"`
	IoWriteIops       int    `json:"io_write_iops"`
	NetInMbitps       int    `json:"net_in_mbitps"`
	NetOutMbitps      int    `json:"net_out_mbitps"`
	NetWeight         int    `json:"net_weight"`
	AntiSpoofing      bool   `json:"anti_spoofing"`
	TCPConnectionsIn  int    `json:"tcp_connections_in"`
	TCPConnectionsOut int    `json:"tcp_connections_out"`
	ProcessNumber     int    `json:"process_number"`
	FirewallRules     []struct {
		Action    string   `json:"action"`
		Direction string   `json:"direction"`
		Protocols []string `json:"protocols"`
		Portstart int      `json:"portstart"`
		Portend   int      `json:"portend"`
	} `json:"firewall_rules"`
	SendEmailMode string `json:"send_email_mode"`
	Vxlan         []struct {
		ID         int `json:"id"`
		Ipv4Number int `json:"ipv4_number"`
		Ipnet      int `json:"ipnet"`
	} `json:"vxlan"`
}

type RecipeTask struct {
	ID             int   `json:"id"`
	Task           int   `json:"task"`
	RecipeTaskList []int `json:"recipe_task_list"`
	RecipeTask     int   `json:"recipe_task"`
}

type HostResponse struct {
	ID           int    `json:"id"`
	InternalName string `json:"internal_name"`
	Node         int    `json:"node"`
	Name         string `json:"name"`
	State        string `json:"state"`
	CPUNumber    int    `json:"cpu_number"`
	RAMMib       int    `json:"ram_mib"`
	ExpandPart   string `json:"expand_part"`
	NetIsSynced  bool   `json:"net_is_synced"`
	IPAutomation string `json:"ip_automation"`
}

type HostUpdateRequest struct {
	Values struct {
		Password         string `json:"password"`
		VncPassword      string `json:"vnc_password"`
		State            string `json:"state"`
		StateUpdatedDate string `json:"state_updated_date"`
		GuestAgent       bool   `json:"guest_agent"`
		VncPort          int    `json:"vnc_port"`
		Property         struct {
		} `json:"property"`
		StartDate int `json:"start_date"`
	} `json:"values"`
	Comment           string `json:"comment"`
	Name              string `json:"name"`
	HaRestoreOnFail   bool   `json:"ha_restore_on_fail"`
	HaRestorePriority int    `json:"ha_restore_priority"`
	ExpandPart        string `json:"expand_part"`
}

type Account struct {
	Account int `json:"account"`
}

type ImageSize struct {
	ImageGib int `json:"image_gib"`
}

type HostHistoryResponse struct {
	List []struct {
		Name       string `json:"name"`
		DateCreate string `json:"date_create"`
		State      string `json:"state"`
		Params     struct {
		} `json:"params"`
		User string `json:"user"`
	} `json:"list"`
}

type IFace struct {
	LastNotify int `json:"last_notify"`
	List       []struct {
		ID     int      `json:"id"`
		Name   string   `json:"name"`
		Bridge string   `json:"bridge"`
		IP     []string `json:"ip"`
		Mac    string   `json:"mac"`
		Model  string   `json:"model"`
		Vxlan  int      `json:"vxlan"`
	} `json:"list"`
}

type IFaceParams struct {
	Name          string `json:"name"`
	Mac           string `json:"mac"`
	Bridge        string `json:"bridge"`
	IsMainNetwork bool   `json:"is_main_network"`
	Ippool        int    `json:"ippool"`
	IPCount       int    `json:"ip_count"`
}

type IFaceModel struct {
	Model string `json:"model"`
}

type AddIPToHostRequest struct {
	IPAddr struct {
		Name              string `json:"name"`
		IPPool            int    `json:"ip_pool"`
		IPNetwork         int    `json:"ip_network"`
		WithoutAllocation bool   `json:"without_allocation"`
	} `json:"ip_addr"`
	Ipv6Pool    []int `json:"ipv6_pool"`
	Ipv6Prefix  int   `json:"ipv6_prefix"`
	Ipv6Enabled bool  `json:"ipv6_enabled"`
	Ipv4Pool    []int `json:"ipv4_pool"`
	Ipv4Number  int   `json:"ipv4_number"`
	Interfaces  []struct {
	} `json:"interfaces"`
	Vxlan []struct {
		ID         int `json:"id"`
		Ipv4Number int `json:"ipv4_number"`
		Ipnet      int `json:"ipnet"`
	} `json:"vxlan"`
}

type IPAutomationType struct {
	IPAutomationType string `json:"ip_automation_type"`
}

type IPV4Info struct {
	LastNotify int `json:"last_notify"`
	List       []struct {
		ID      int    `json:"id"`
		IPAddr  string `json:"ip_addr"`
		Domain  string `json:"domain"`
		Gateway string `json:"gateway"`
		Mask    string `json:"mask"`
		State   string `json:"state"`
		Family  int    `json:"family"`
		Ippool  int    `json:"ippool"`
		Network int    `json:"network"`
		Host    struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			Interface int    `json:"interface"`
		} `json:"host"`
		ClusterInterface int `json:"cluster_interface"`
		Vxlan            struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Tag  int    `json:"tag"`
		} `json:"vxlan"`
	} `json:"list"`
	Size int `json:"size"`
}

type IPV6Info struct {
	LastNotify     int  `json:"last_notify"`
	Ipv6Enabled    bool `json:"ipv6_enabled"`
	Ipv6PtrEnabled bool `json:"ipv6_ptr_enabled"`
}

type ISOTags struct {
	URL  string   `json:"url"`
	Tags []string `json:"tags"`
}

type Socket struct {
	Socket string `json:"socket"`
}

type Metadata struct {
	Metadata struct {
	} `json:"metadata"`
}

type HostMigrateResponse struct {
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

type HostMigrateRequest struct {
	Plain bool `json:"plain"`
	Node  int  `json:"node"`
	Disks []struct {
		ID         int `json:"id"`
		DstStorage int `json:"dst_storage"`
	} `json:"disks"`
}

type PTRUpdateRequest struct {
	Name   string `json:"name"`
	Domain string `json:"domain"`
}

type PTRResponse struct {
	List []struct {
		Family  int    `json:"family"`
		Gateway string `json:"gateway"`
		IP      []struct {
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Domain string `json:"domain"`
			Status string `json:"status"`
		} `json:"ip"`
		IPCount int    `json:"ip_count"`
		Network string `json:"network"`
		Prefix  int    `json:"prefix"`
	} `json:"list"`
}

type Domain struct {
	Domain string `json:"domain"`
}

type HostReinstallRequest struct {
	Os         int `json:"os"`
	Image      int `json:"image"`
	RecipeList []struct {
		Recipe       int `json:"recipe"`
		RecipeParams []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"recipe_params"`
	} `json:"recipe_list"`
	Recipe       int `json:"recipe"`
	RecipeParams []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"recipe_params"`
	IgnoreRecipeFilters bool   `json:"ignore_recipe_filters"`
	SendEmailMode       string `json:"send_email_mode"`
	Password            string `json:"password"`
	Disk                int    `json:"disk"`
}

type RescueMode struct {
	RescueMode bool `json:"rescue_mode"`
}

type HostResourceUpdateRequest struct {
	CPUNumber          int    `json:"cpu_number"`
	RAMMib             int    `json:"ram_mib"`
	NetBandwidthMbitps int    `json:"net_bandwidth_mbitps"`
	CPUMode            string `json:"cpu_mode"`
	Nesting            bool   `json:"nesting"`
	CPUCustomModel     string `json:"cpu_custom_model"`
	CPUWeight          int    `json:"cpu_weight"`
	IoWeight           int    `json:"io_weight"`
	IoReadMbitps       int    `json:"io_read_mbitps"`
	IoWriteMbitps      int    `json:"io_write_mbitps"`
	IoReadIops         int    `json:"io_read_iops"`
	IoWriteIops        int    `json:"io_write_iops"`
	NetInMbitps        int    `json:"net_in_mbitps"`
	NetOutMbitps       int    `json:"net_out_mbitps"`
	NetWeight          int    `json:"net_weight"`
	AntiSpoofing       bool   `json:"anti_spoofing"`
	TCPConnectionsIn   int    `json:"tcp_connections_in"`
	TCPConnectionsOut  int    `json:"tcp_connections_out"`
	ProcessNumber      int    `json:"process_number"`
	FirewallRules      []struct {
		Action    string   `json:"action"`
		Direction string   `json:"direction"`
		Protocols []string `json:"protocols"`
		Portstart int      `json:"portstart"`
		Portend   int      `json:"portend"`
	} `json:"firewall_rules"`
	Defer struct {
		Action string `json:"action"`
	} `json:"defer"`
}

type HostBackup struct {
	Backup int `json:"backup"`
}

type Recipe struct {
	Recipe       int    `json:"recipe"`
	Body         string `json:"body"`
	RecipeParams []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"recipe_params"`
	SendEmail           bool `json:"send_email"`
	IgnoreRecipeFilters bool `json:"ignore_recipe_filters"`
	Recipients          []struct {
		Lang  string `json:"lang"`
		Email string `json:"email"`
	} `json:"recipients"`
}

type VNCPortUpdateRequest struct {
	List []struct {
		Domain  string `json:"domain"`
		VncPort int    `json:"vnc_port"`
		VncPass string `json:"vnc_pass"`
	} `json:"list"`
}

type IP struct {
	LastNotify int `json:"last_notify"`
	List       []struct {
		ID      int    `json:"id"`
		IPAddr  string `json:"ip_addr"`
		Domain  string `json:"domain"`
		Gateway string `json:"gateway"`
		Mask    string `json:"mask"`
		State   string `json:"state"`
		Family  int    `json:"family"`
		Ippool  int    `json:"ippool"`
		Network int    `json:"network"`
		Host    struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			Interface int    `json:"interface"`
		} `json:"host"`
		ClusterInterface int `json:"cluster_interface"`
	} `json:"list"`
	Size int `json:"size"`
}

type UpdateIPResponse struct {
	Values struct {
		Domain  string `json:"domain"`
		Gateway string `json:"gateway"`
		IPAddr  string `json:"ip_addr"`
		Netmask string `json:"netmask"`
	} `json:"values"`
}

type ScheduleListResponse struct {
	SheduleList []struct {
		Name           string `json:"name"`
		Handler        string `json:"handler"`
		Service        string `json:"service"`
		Method         string `json:"method"`
		InstanceID     int    `json:"instance_id"`
		URLQueryParams struct {
		} `json:"url_query_params"`
		PostParams struct {
		} `json:"post_params"`
		CronExpression string `json:"cron_expression"`
	} `json:"shedule_list"`
}
