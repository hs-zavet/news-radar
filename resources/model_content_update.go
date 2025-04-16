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

// checks if the ContentUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentUpdate{}

// ContentUpdate struct for ContentUpdate
type ContentUpdate struct {
	Data ContentUpdateData `json:"data"`
}

type _ContentUpdate ContentUpdate

// NewContentUpdate instantiates a new ContentUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentUpdate(data ContentUpdateData) *ContentUpdate {
	this := ContentUpdate{}
	this.Data = data
	return &this
}

// NewContentUpdateWithDefaults instantiates a new ContentUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentUpdateWithDefaults() *ContentUpdate {
	this := ContentUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *ContentUpdate) GetData() ContentUpdateData {
	if o == nil {
		var ret ContentUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ContentUpdate) GetDataOk() (*ContentUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ContentUpdate) SetData(v ContentUpdateData) {
	o.Data = v
}

func (o ContentUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ContentUpdate) UnmarshalJSON(data []byte) (err error) {
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

	varContentUpdate := _ContentUpdate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContentUpdate)

	if err != nil {
		return err
	}

	*o = ContentUpdate(varContentUpdate)

	return err
}

type NullableContentUpdate struct {
	value *ContentUpdate
	isSet bool
}

func (v NullableContentUpdate) Get() *ContentUpdate {
	return v.value
}

func (v *NullableContentUpdate) Set(val *ContentUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableContentUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableContentUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentUpdate(val *ContentUpdate) *NullableContentUpdate {
	return &NullableContentUpdate{value: val, isSet: true}
}

func (v NullableContentUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


