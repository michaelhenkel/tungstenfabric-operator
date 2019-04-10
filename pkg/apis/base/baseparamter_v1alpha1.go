package base

type BaseParameter struct {
	Size int32 `json:"size"`
	Image string `json:"image"`
	HostNetwork bool `json:"hostnetwork"`
	ImagePullPolicy string `json:"imagepullpolicy"`
}

