// Code generated by gen.py. DO NOT EDIT.

package expconf

import (
	"github.com/docker/docker/api/types"
	"github.com/santhosh-tekuri/jsonschema/v2"
	k8sV1 "k8s.io/api/core/v1"

	"github.com/determined-ai/determined/master/pkg/schemas"
)

func (e EnvironmentConfigV0) Image() EnvironmentImageMapV0 {
	if e.RawImage == nil {
		panic("You must call WithDefaults on EnvironmentConfigV0 before .RawImage")
	}
	return *e.RawImage
}

func (e *EnvironmentConfigV0) SetImage(val EnvironmentImageMapV0) {
	e.RawImage = &val
}

func (e EnvironmentConfigV0) EnvironmentVariables() EnvironmentVariablesMapV0 {
	if e.RawEnvironmentVariables == nil {
		panic("You must call WithDefaults on EnvironmentConfigV0 before .RawEnvironmentVariables")
	}
	return *e.RawEnvironmentVariables
}

func (e *EnvironmentConfigV0) SetEnvironmentVariables(val EnvironmentVariablesMapV0) {
	e.RawEnvironmentVariables = &val
}

func (e EnvironmentConfigV0) Ports() map[string]int {
	return e.RawPorts
}

func (e *EnvironmentConfigV0) SetPorts(val map[string]int) {
	e.RawPorts = val
}

func (e EnvironmentConfigV0) RegistryAuth() *types.AuthConfig {
	return e.RawRegistryAuth
}

func (e *EnvironmentConfigV0) SetRegistryAuth(val *types.AuthConfig) {
	e.RawRegistryAuth = val
}

func (e EnvironmentConfigV0) ForcePullImage() bool {
	if e.RawForcePullImage == nil {
		panic("You must call WithDefaults on EnvironmentConfigV0 before .RawForcePullImage")
	}
	return *e.RawForcePullImage
}

func (e *EnvironmentConfigV0) SetForcePullImage(val bool) {
	e.RawForcePullImage = &val
}

func (e EnvironmentConfigV0) PodSpec() *k8sV1.Pod {
	return e.RawPodSpec
}

func (e *EnvironmentConfigV0) SetPodSpec(val *k8sV1.Pod) {
	e.RawPodSpec = val
}

func (e EnvironmentConfigV0) AddCapabilities() []string {
	return e.RawAddCapabilities
}

func (e *EnvironmentConfigV0) SetAddCapabilities(val []string) {
	e.RawAddCapabilities = val
}

func (e EnvironmentConfigV0) DropCapabilities() []string {
	return e.RawDropCapabilities
}

func (e *EnvironmentConfigV0) SetDropCapabilities(val []string) {
	e.RawDropCapabilities = val
}

func (e EnvironmentConfigV0) WithDefaults() EnvironmentConfigV0 {
	return schemas.WithDefaults(e).(EnvironmentConfigV0)
}

func (e EnvironmentConfigV0) Merge(other EnvironmentConfigV0) EnvironmentConfigV0 {
	return schemas.Merge(e, other).(EnvironmentConfigV0)
}

func (e EnvironmentConfigV0) ParsedSchema() interface{} {
	return schemas.ParsedEnvironmentConfigV0()
}

func (e EnvironmentConfigV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/environment.json")
}

func (e EnvironmentConfigV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/environment.json")
}
