package kubernetes

const (
	SelfResourcesPrefix        = "tesseract-"
	FrontPodName               = SelfResourcesPrefix + "front"
	FrontServiceName           = FrontPodName
	HubPodName                 = SelfResourcesPrefix + "hub"
	HubServiceName             = HubPodName
	K8sAllNamespaces           = ""
	MinKubernetesServerVersion = "1.16.0"
)