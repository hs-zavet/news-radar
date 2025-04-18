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

// checks if the CreateHashtagData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateHashtagData{}

// CreateHashtagData struct for CreateHashtagData
type CreateHashtagData struct {
	Type string `json:"type"`
	Attributes CreateHashtagDataAttributes `json:"attributes"`
}

type _CreateHashtagData CreateHashtagData

// NewCreateHashtagData instantiates a new CreateHashtagData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateHashtagData(type_ string, attributes CreateHashtagDataAttributes) *CreateHashtagData {
	this := CreateHashtagData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewCreateHashtagDataWithDefaults instantiates a new CreateHashtagData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateHashtagDataWithDefaults() *CreateHashtagData {
	this := CreateHashtagData{}
	return &this
}

// GetType returns the Type field value
func (o *CreateHashtagData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateHashtagData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateHashtagData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *CreateHashtagData) GetAttributes() CreateHashtagDataAttributes {
	if o == nil {
		var ret CreateHashtagDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CreateHashtagData) GetAttributesOk() (*CreateHashtagDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CreateHashtagData) SetAttributes(v CreateHashtagDataAttributes) {
	o.Attributes = v
}

func (o CreateHashtagData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateHashtagData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *CreateHashtagData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"attributes",
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

	varCreateHashtagData := _CreateHashtagData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateHashtagData)

	if err != nil {
		return err
	}

	*o = CreateHashtagData(varCreateHashtagData)

	return err
}

type NullableCreateHashtagData struct {
	value *CreateHashtagData
	isSet bool
}

func (v NullableCreateHashtagData) Get() *CreateHashtagData {
	return v.value
}

func (v *NullableCreateHashtagData) Set(val *CreateHashtagData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateHashtagData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateHashtagData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateHashtagData(val *CreateHashtagData) *NullableCreateHashtagData {
	return &NullableCreateHashtagData{value: val, isSet: true}
}

func (v NullableCreateHashtagData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateHashtagData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


