// Code generated by gen.py. DO NOT EDIT.

package expconf

import (
	"github.com/santhosh-tekuri/jsonschema/v2"

	"github.com/determined-ai/determined/master/pkg/schemas"
)

func (s SharedFSConfigV0) SaveExperimentBest() int {
	if s.RawSaveExperimentBest == nil {
		panic("You must call WithDefaults on SharedFSConfigV0 before .RawSaveExperimentBest")
	}
	return *s.RawSaveExperimentBest
}

func (s *SharedFSConfigV0) SetSaveExperimentBest(val int) {
	s.RawSaveExperimentBest = &val
}

func (s SharedFSConfigV0) SaveTrialBest() int {
	if s.RawSaveTrialBest == nil {
		panic("You must call WithDefaults on SharedFSConfigV0 before .RawSaveTrialBest")
	}
	return *s.RawSaveTrialBest
}

func (s *SharedFSConfigV0) SetSaveTrialBest(val int) {
	s.RawSaveTrialBest = &val
}

func (s SharedFSConfigV0) SaveTrialLatest() int {
	if s.RawSaveTrialLatest == nil {
		panic("You must call WithDefaults on SharedFSConfigV0 before .RawSaveTrialLatest")
	}
	return *s.RawSaveTrialLatest
}

func (s *SharedFSConfigV0) SetSaveTrialLatest(val int) {
	s.RawSaveTrialLatest = &val
}

func (s SharedFSConfigV0) HostPath() string {
	return s.RawHostPath
}

func (s *SharedFSConfigV0) SetHostPath(val string) {
	s.RawHostPath = val
}

func (s SharedFSConfigV0) ContainerPath() *string {
	return s.RawContainerPath
}

func (s *SharedFSConfigV0) SetContainerPath(val *string) {
	s.RawContainerPath = val
}

func (s SharedFSConfigV0) CheckpointPath() *string {
	return s.RawCheckpointPath
}

func (s *SharedFSConfigV0) SetCheckpointPath(val *string) {
	s.RawCheckpointPath = val
}

func (s SharedFSConfigV0) TensorboardPath() *string {
	return s.RawTensorboardPath
}

func (s *SharedFSConfigV0) SetTensorboardPath(val *string) {
	s.RawTensorboardPath = val
}

func (s SharedFSConfigV0) StoragePath() *string {
	return s.RawStoragePath
}

func (s *SharedFSConfigV0) SetStoragePath(val *string) {
	s.RawStoragePath = val
}

func (s SharedFSConfigV0) Propagation() string {
	if s.RawPropagation == nil {
		panic("You must call WithDefaults on SharedFSConfigV0 before .RawPropagation")
	}
	return *s.RawPropagation
}

func (s *SharedFSConfigV0) SetPropagation(val string) {
	s.RawPropagation = &val
}

func (s SharedFSConfigV0) WithDefaults() SharedFSConfigV0 {
	return schemas.WithDefaults(s).(SharedFSConfigV0)
}

func (s SharedFSConfigV0) Merge(other SharedFSConfigV0) SharedFSConfigV0 {
	return schemas.Merge(s, other).(SharedFSConfigV0)
}

func (s SharedFSConfigV0) ParsedSchema() interface{} {
	return schemas.ParsedSharedFSConfigV0()
}

func (s SharedFSConfigV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/shared-fs.json")
}

func (s SharedFSConfigV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/shared-fs.json")
}
