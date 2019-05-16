package v1alpha1

type Container struct {
	Name string `json:"name"`
	Image string `json:"image,omitempty"`
	PullPolicy string `json:"imagepullpolicy"`
	LogVolumePath string `json:"logvolumepath"`
	DataVolumePath string `json:"datavolumepath"`
	UnixSocketVolume bool `json:"unixsocketvolume"`
	HostUserBinVolume bool `json:"hostuserbinvolume"`
	EtcContrailVolume bool `json:"etccontrailvolume"`
	DevVolume bool `json:"devvolume"`
	NetworkScriptsVolume bool `json:"networkscriptsvolume"`
	HostBinVolume bool `json:"hostbinvolume"`
	UsrSrcVolume bool `json:"usrsrcvolume"`
	LibModulesVolume bool `json:"libmodulesvolume"`
	VarLibContrailVolume bool `json:"varlibcontrailvolume"`
	VarCrashesVolume bool `json:"varcrashesvolume"`
	EtcCniVolume bool `json:"etccnivolume"`
	VarLogCniVolume bool `json:"varlogcnivolume"`
	StatusVolume bool `json:"statusvolume"`
	OptBinCniVolume bool `json:"optbincnivolume"`
	Privileged bool `json:"privileged"`
	Env map[string]string `json:"env"`
	Command []string `json:"command"`
	LifeCycleScript []string `json:"lifecylescript"`
	ResourceConfigMap bool `json:"resourceconfigmap"`
}

type General struct {
	Type string `json:"type"`
	Size string `json:"size"`
	Image string `json:"image,omitempty"`
	HostNetwork string `json:"hostNetwork"`
	PullPolicy string `json:"pullPolicy"`
	AaaMode string `json:"aaaMode,omitempty"`
	AuthMode string `json:"authMode,omitempty"`
	CloudAdminRole string `json:"cloudAdminRole,omitempty"`
	GlobalReadOnlyRole string `json:"globalReadOnlyRole,omitempty"`
	CloudOrchestrator string `json:"cloudOrchestrator"`
}

type Images struct {
	Nodemanger string `json:"nodemanager,omitempty"`
	Nodeinit string `json:"nodeinit,omitempty"`
	Status string `json:"status,omitempty"`
	Api string `json:"api,omitempty"`
	Devicemanager string `json:"devicemanager,omitempty"`
	Schematransfomer string `json:"schematransformer,omitempty"`
	Servicemonitor string `json:"servicemonitor,omitempty"`
	AnalyticsApi string `json:"analyticsapi,omitempty"`
	Collector string `json:"collector,omitempty"`
	Redis string `json:"redis,omitempty"`
	Control string `json:"control,omitempty"`
	Named string `json:"named,omitempty"`
	Dns string `json:"dns,omitempty"`
	Kubemanager string `json:"kubemanager,omitempty"`
	Cassandra string `json:"cassandra,omitempty"`
	Zookeeper string `json:"zookeeper,omitempty"`
	Rabbitmq string `json:"rabbitmq,omitempty"`
	Webuiweb string `json:"webuiweb,omitempty"`
	Webuijob string `json:"webuijob,omitempty"`
}

type CassandraConfig struct {
	General *General
	ListenAddress string `json:"CASSANDRA_LISTEN_ADDRESS,omitempty"`
	Port string `json:"CASSANDRA_PORT,omitempty"`
	CqlPort string `json:"CASSANDRA_CQL_PORT,omitempty"`
	SslStoragePort string `json:"CASSANDRA_SSL_STORAGE_PORT,omitempty"`
	StoragePort string `json:"CASSANDRA_STORAGE_PORT,omitempty"`
	JmxPort string `json:"CASSANDRA_JMX_LOCAL_PORT,omitempty"`
	StartRpc string `json:"CASSANDRA_START_RPC,omitempty"`
}

type KubemanagerConfig struct {
	General *General
	KubernetesApiServer string `json:"KUBERNETES_API_SERVER,omitempty"`
	KubernetesApiSecurePort string `json:"KUBERNETES_API_SECURE_PORT,omitempty"`
	PodSubnets string `json:"KUBERNETES_POD_SUBNETS,omitempty"`
	ServiceSubnets string `json:"KUBERNETES_SERVICE_SUBNETS,omitempty"`
	ClusterName string `json:"KUBERNETES_CLUSTER_NAME,omitempty"`
	IpFabricForwarding string `json:"KUBERNETES_IP_FABRIC_FORWARDING,omitempty"`
	IpFabricSnat string `json:"KUBERNETES_IP_FABRIC_SNAT,omitempty"`
	TokenFile string `json:"K8S_TOKEN_FILE,omitempty"`
}

type ZookeeperConfig struct {
	General *General
	Port string `json:"ZOOKEEPER_PORT,omitempty"`
	Ports string `json:"ZOOKEEPER_PORTS,omitempty"`
}

type RabbitmqConfig struct {
	General *General
	Port string `json:"RABBITMQ_NODE_PORT,omitempty"`
	Cookie string `json:"RABBITMQ_ERLANG_COOKIE,omitempty"`
}

type ConfigConfig struct {
	General *General
}

type ControlConfig struct {
	General *General
}

type WebuiConfig struct {
	General *General
}

type VrouterConfig struct {
	General *General
}
