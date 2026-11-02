package krutrim

import "github.com/ola-krutrim/krutrim-go-sdk/option"

// KKSV1Service provides version 1 of the KKS API.
type KKSV1Service struct {
	Options []option.RequestOption
	Cluster KKSV1ClusterService
	Flavor  KKSV1FlavorService
	Addon   KKSV1AddonService
}

func NewKKSV1Service(opts ...option.RequestOption) KKSV1Service {
	return KKSV1Service{
		Options: opts,
		Cluster: NewKKSV1ClusterService(opts...),
		Flavor:  NewKKSV1FlavorService(opts...),
		Addon:   NewKKSV1AddonService(opts...),
	}
}
