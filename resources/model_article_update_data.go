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

// checks if the ArticleUpdateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArticleUpdateData{}

// ArticleUpdateData struct for ArticleUpdateData
type ArticleUpdateData struct {
	Type string `json:"type"`
	Attributes ArticleUpdateDataAttributes `json:"attributes"`
}

type _ArticleUpdateData ArticleUpdateData

// NewArticleUpdateData instantiates a new ArticleUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArticleUpdateData(type_ string, attributes ArticleUpdateDataAttributes) *ArticleUpdateData {
	this := ArticleUpdateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewArticleUpdateDataWithDefaults instantiates a new ArticleUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArticleUpdateDataWithDefaults() *ArticleUpdateData {
	this := ArticleUpdateData{}
	return &this
}

// GetType returns the Type field value
func (o *ArticleUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ArticleUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ArticleUpdateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ArticleUpdateData) GetAttributes() ArticleUpdateDataAttributes {
	if o == nil {
		var ret ArticleUpdateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ArticleUpdateData) GetAttributesOk() (*ArticleUpdateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ArticleUpdateData) SetAttributes(v ArticleUpdateDataAttributes) {
	o.Attributes = v
}

func (o ArticleUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArticleUpdateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *ArticleUpdateData) UnmarshalJSON(data []byte) (err error) {
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

	varArticleUpdateData := _ArticleUpdateData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varArticleUpdateData)

	if err != nil {
		return err
	}

	*o = ArticleUpdateData(varArticleUpdateData)

	return err
}

type NullableArticleUpdateData struct {
	value *ArticleUpdateData
	isSet bool
}

func (v NullableArticleUpdateData) Get() *ArticleUpdateData {
	return v.value
}

func (v *NullableArticleUpdateData) Set(val *ArticleUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableArticleUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableArticleUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArticleUpdateData(val *ArticleUpdateData) *NullableArticleUpdateData {
	return &NullableArticleUpdateData{value: val, isSet: true}
}

func (v NullableArticleUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArticleUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


