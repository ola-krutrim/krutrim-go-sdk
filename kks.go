package krutrim

import "github.com/ola-krutrim/krutrim-go-sdk/option"

// KKSService provides access to the Krutrim Kubernetes Service API.
// Use NewKKSService to construct it with request options. Unlike NewClient,
// service constructors do not read options from the environment.
type KKSService struct {
	Options []option.RequestOption
	V1      KKSV1Service
}

func NewKKSService(opts ...option.RequestOption) KKSService {
	return KKSService{Options: opts, V1: NewKKSV1Service(opts...)}
}

// kksRequestOptions applies service options, the parameter's region header, and
// finally request options, so a caller can override any setting per request.
func kksRequestOptions(serviceOpts []option.RequestOption, region string, opts []option.RequestOption) []option.RequestOption {
	result := make([]option.RequestOption, 0, len(serviceOpts)+len(opts)+1)
	result = append(result, serviceOpts...)
	if region != "" {
		result = append(result, option.WithHeader("x-region", region))
	}
	return append(result, opts...)
}
