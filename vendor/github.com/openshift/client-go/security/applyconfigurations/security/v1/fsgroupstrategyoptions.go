// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/security/v1"
)

// FSGroupStrategyOptionsApplyConfiguration represents an declarative configuration of the FSGroupStrategyOptions type for use
// with apply.
type FSGroupStrategyOptionsApplyConfiguration struct {
	Type   *v1.FSGroupStrategyType     `json:"type,omitempty"`
	Ranges []IDRangeApplyConfiguration `json:"ranges,omitempty"`
}

// FSGroupStrategyOptionsApplyConfiguration constructs an declarative configuration of the FSGroupStrategyOptions type for use with
// apply.
func FSGroupStrategyOptions() *FSGroupStrategyOptionsApplyConfiguration {
	return &FSGroupStrategyOptionsApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *FSGroupStrategyOptionsApplyConfiguration) WithType(value v1.FSGroupStrategyType) *FSGroupStrategyOptionsApplyConfiguration {
	b.Type = &value
	return b
}

// WithRanges adds the given value to the Ranges field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Ranges field.
func (b *FSGroupStrategyOptionsApplyConfiguration) WithRanges(values ...*IDRangeApplyConfiguration) *FSGroupStrategyOptionsApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithRanges")
		}
		b.Ranges = append(b.Ranges, *values[i])
	}
	return b
}
