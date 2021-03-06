/*
 * API Title
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type WakeStatus struct {
	WakeDevice bool `json:"wake_device"`
}

// AssertWakeStatusRequired checks if the required fields are not zero-ed
func AssertWakeStatusRequired(obj WakeStatus) error {
	elements := map[string]interface{}{
		"wake_device": obj.WakeDevice,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseWakeStatusRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of WakeStatus (e.g. [][]WakeStatus), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseWakeStatusRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aWakeStatus, ok := obj.(WakeStatus)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertWakeStatusRequired(aWakeStatus)
	})
}
