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

// checks if the Content type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Content{}

// Content struct for Content
type Content struct {
	Type string `json:"type"`
	Content map[string]interface{} `json:"content"`
}

type _Content Content

// NewContent instantiates a new Content object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContent(type_ string, content map[string]interface{}) *Content {
	this := Content{}
	this.Type = type_
	this.Content = content
	return &this
}

// NewContentWithDefaults instantiates a new Content object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentWithDefaults() *Content {
	this := Content{}
	return &this
}

// GetType returns the Type field value
func (o *Content) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Content) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Content) SetType(v string) {
	o.Type = v
}

// GetContent returns the Content field value
func (o *Content) GetContent() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *Content) GetContentOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Content, true
}

// SetContent sets field value
func (o *Content) SetContent(v map[string]interface{}) {
	o.Content = v
}

func (o Content) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Content) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["content"] = o.Content
	return toSerialize, nil
}

func (o *Content) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"content",
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

	varContent := _Content{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContent)

	if err != nil {
		return err
	}

	*o = Content(varContent)

	return err
}

type NullableContent struct {
	value *Content
	isSet bool
}

func (v NullableContent) Get() *Content {
	return v.value
}

func (v *NullableContent) Set(val *Content) {
	v.value = val
	v.isSet = true
}

func (v NullableContent) IsSet() bool {
	return v.isSet
}

func (v *NullableContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContent(val *Content) *NullableContent {
	return &NullableContent{value: val, isSet: true}
}

func (v NullableContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


