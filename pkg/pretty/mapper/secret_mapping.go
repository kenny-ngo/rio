package mapper

import (
	"github.com/rancher/rio/cli/cmd/create"
	"github.com/rancher/rio/types/apis/rio.cattle.io/v1beta1"
)

func NewSecretMapping(field string) ObjectsToSlice {
	return ObjectsToSlice{
		Field: field,
		NewObject: func() MaybeStringer {
			return &v1beta1.SecretMapping{}
		},
		ToObject: func(obj interface{}) (interface{}, error) {
			str, ok := obj.(string)
			if !ok {
				return obj, nil
			}
			objs, err := create.ParseSecrets([]string{str})
			if err != nil {
				return nil, err
			}
			return objs[0], nil
		},
	}
}
