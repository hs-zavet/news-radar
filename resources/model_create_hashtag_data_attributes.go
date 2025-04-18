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

// checks if the CreateHashtagDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateHashtagDataAttributes{}

// CreateHashtagDataAttributes struct for CreateHashtagDataAttributes
type CreateHashtagDataAttributes struct {
	Tags []string `json:"tags"`
	// The ID of the article.
	ArticleID string `json:"articleID"`
}

type _CreateHashtagDataAttributes CreateHashtagDataAttributes

// NewCreateHashtagDataAttributes instantiates a new CreateHashtagDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateHashtagDataAttributes(tags []string, articleID string) *CreateHashtagDataAttributes {
	this := CreateHashtagDataAttributes{}
	this.Tags = tags
	this.ArticleID = articleID
	return &this
}

// NewCreateHashtagDataAttributesWithDefaults instantiates a new CreateHashtagDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateHashtagDataAttributesWithDefaults() *CreateHashtagDataAttributes {
	this := CreateHashtagDataAttributes{}
	return &this
}

// GetTags returns the Tags field value
func (o *CreateHashtagDataAttributes) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *CreateHashtagDataAttributes) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *CreateHashtagDataAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetArticleID returns the ArticleID field value
func (o *CreateHashtagDataAttributes) GetArticleID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ArticleID
}

// GetArticleIDOk returns a tuple with the ArticleID field value
// and a boolean to check if the value has been set.
func (o *CreateHashtagDataAttributes) GetArticleIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArticleID, true
}

// SetArticleID sets field value
func (o *CreateHashtagDataAttributes) SetArticleID(v string) {
	o.ArticleID = v
}

func (o CreateHashtagDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateHashtagDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tags"] = o.Tags
	toSerialize["articleID"] = o.ArticleID
	return toSerialize, nil
}

func (o *CreateHashtagDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tags",
		"articleID",
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

	varCreateHashtagDataAttributes := _CreateHashtagDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateHashtagDataAttributes)

	if err != nil {
		return err
	}

	*o = CreateHashtagDataAttributes(varCreateHashtagDataAttributes)

	return err
}

type NullableCreateHashtagDataAttributes struct {
	value *CreateHashtagDataAttributes
	isSet bool
}

func (v NullableCreateHashtagDataAttributes) Get() *CreateHashtagDataAttributes {
	return v.value
}

func (v *NullableCreateHashtagDataAttributes) Set(val *CreateHashtagDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateHashtagDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateHashtagDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateHashtagDataAttributes(val *CreateHashtagDataAttributes) *NullableCreateHashtagDataAttributes {
	return &NullableCreateHashtagDataAttributes{value: val, isSet: true}
}

func (v NullableCreateHashtagDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateHashtagDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


