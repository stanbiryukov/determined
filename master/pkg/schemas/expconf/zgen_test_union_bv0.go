// Code generated by gen.py. DO NOT EDIT.

package expconf

import (
	"github.com/santhosh-tekuri/jsonschema/v2"

	"github.com/determined-ai/determined/master/pkg/schemas"
)

func (t TestUnionBV0) Type() string {
	return t.RawType
}

func (t *TestUnionBV0) SetType(val string) {
	t.RawType = val
}

func (t TestUnionBV0) ValB() int {
	return t.RawValB
}

func (t *TestUnionBV0) SetValB(val int) {
	t.RawValB = val
}

func (t TestUnionBV0) CommonVal() string {
	if t.RawCommonVal == nil {
		panic("You must call WithDefaults on TestUnionBV0 before .RawCommonVal")
	}
	return *t.RawCommonVal
}

func (t *TestUnionBV0) SetCommonVal(val string) {
	t.RawCommonVal = &val
}

func (t TestUnionBV0) WithDefaults() TestUnionBV0 {
	return schemas.WithDefaults(t).(TestUnionBV0)
}

func (t TestUnionBV0) Merge(other TestUnionBV0) TestUnionBV0 {
	return schemas.Merge(t, other).(TestUnionBV0)
}

func (t TestUnionBV0) ParsedSchema() interface{} {
	return schemas.ParsedTestUnionBV0()
}

func (t TestUnionBV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/test-union-b.json")
}

func (t TestUnionBV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/test-union-b.json")
}
