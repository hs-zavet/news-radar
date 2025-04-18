/*
REST API

REST API

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UpdateContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateContent{}

// UpdateContent struct for UpdateContent
type UpdateContent struct {
	Data UpdateContentData `json:"data"`
}

type _UpdateContent UpdateContent

// NewUpdateContent instantiates a new UpdateContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateContent(data UpdateContentData) *UpdateContent {
	this := UpdateContent{}
	this.Data = data
	return &this
}

// NewUpdateContentWithDefaults instantiates a new UpdateContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateContentWithDefaults() *UpdateContent {
	this := UpdateContent{}
	return &this
}

// GetData returns the Data field value
func (o *UpdateContent) GetData() UpdateContentData {
	if o == nil {
		var ret UpdateContentData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UpdateContent) GetDataOk() (*UpdateContentData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UpdateContent) SetData(v UpdateContentData) {
	o.Data = v
}

func (o UpdateContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *UpdateContent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUpdateContent := _UpdateContent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateContent)

	if err != nil {
		return err
	}

	*o = UpdateContent(varUpdateContent)

	return err
}

type NullableUpdateContent struct {
	value *UpdateContent
	isSet bool
}

func (v NullableUpdateContent) Get() *UpdateContent {
	return v.value
}

func (v *NullableUpdateContent) Set(val *UpdateContent) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateContent) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateContent(val *UpdateContent) *NullableUpdateContent {
	return &NullableUpdateContent{value: val, isSet: true}
}

func (v NullableUpdateContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


