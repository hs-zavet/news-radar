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

// checks if the UpdateAuthorData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAuthorData{}

// UpdateAuthorData struct for UpdateAuthorData
type UpdateAuthorData struct {
	Id string `json:"id"`
	Type string `json:"type"`
	Attributes UpdateAuthorDataAttributes `json:"attributes"`
}

type _UpdateAuthorData UpdateAuthorData

// NewUpdateAuthorData instantiates a new UpdateAuthorData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAuthorData(id string, type_ string, attributes UpdateAuthorDataAttributes) *UpdateAuthorData {
	this := UpdateAuthorData{}
	this.Id = id
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewUpdateAuthorDataWithDefaults instantiates a new UpdateAuthorData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAuthorDataWithDefaults() *UpdateAuthorData {
	this := UpdateAuthorData{}
	return &this
}

// GetId returns the Id field value
func (o *UpdateAuthorData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateAuthorData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateAuthorData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *UpdateAuthorData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UpdateAuthorData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UpdateAuthorData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *UpdateAuthorData) GetAttributes() UpdateAuthorDataAttributes {
	if o == nil {
		var ret UpdateAuthorDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *UpdateAuthorData) GetAttributesOk() (*UpdateAuthorDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *UpdateAuthorData) SetAttributes(v UpdateAuthorDataAttributes) {
	o.Attributes = v
}

func (o UpdateAuthorData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAuthorData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *UpdateAuthorData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varUpdateAuthorData := _UpdateAuthorData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateAuthorData)

	if err != nil {
		return err
	}

	*o = UpdateAuthorData(varUpdateAuthorData)

	return err
}

type NullableUpdateAuthorData struct {
	value *UpdateAuthorData
	isSet bool
}

func (v NullableUpdateAuthorData) Get() *UpdateAuthorData {
	return v.value
}

func (v *NullableUpdateAuthorData) Set(val *UpdateAuthorData) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAuthorData) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAuthorData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAuthorData(val *UpdateAuthorData) *NullableUpdateAuthorData {
	return &NullableUpdateAuthorData{value: val, isSet: true}
}

func (v NullableUpdateAuthorData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAuthorData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


