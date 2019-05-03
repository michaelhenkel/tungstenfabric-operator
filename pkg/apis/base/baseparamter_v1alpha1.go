package base

type BaseParameter struct {
	Size string`json:"size"`
	Image string `json:"image"`
	HostNetwork string `json:"hostnetwork"`
	ImagePullPolicy string `json:"imagepullpolicy"`
}

