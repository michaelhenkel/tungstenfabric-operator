package v1alpha1

type BaseParameter struct {
	Type string `json:"type"`
	Size string `json:"size"`
	Image string `json:"image,omitempty"`
	HostNetwork string `json:"hostnetwork"`
	ImagePullPolicy string `json:"imagepullpolicy"`
	NodeManagerImage string `json:"nodeManagerImage,omitempty"`
	NodeInitImage string `json:"nodeInitImage,omitempty"`
	StatusImage string `json:"statusImage,omitempty"`
}

type Container struct {
	Name string `json:"name"`
	Image string `json:"image,omitempty"`
}
