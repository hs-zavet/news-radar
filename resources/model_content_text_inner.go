/*
REST API

REST API

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
)

// checks if the ContentTextInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentTextInner{}

// ContentTextInner struct for ContentTextInner
type ContentTextInner struct {
	Text *string `json:"text,omitempty"`
	Marks []string `json:"marks,omitempty"`
	Color *string `json:"color,omitempty"`
	Link *string `json:"link,omitempty"`
}

// NewContentTextInner instantiates a new ContentTextInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentTextInner() *ContentTextInner {
	this := ContentTextInner{}
	return &this
}

// NewContentTextInnerWithDefaults instantiates a new ContentTextInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentTextInnerWithDefaults() *ContentTextInner {
	this := ContentTextInner{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *ContentTextInner) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentTextInner) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *ContentTextInner) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *ContentTextInner) SetText(v string) {
	o.Text = &v
}

// GetMarks returns the Marks field value if set, zero value otherwise.
func (o *ContentTextInner) GetMarks() []string {
	if o == nil || IsNil(o.Marks) {
		var ret []string
		return ret
	}
	return o.Marks
}

// GetMarksOk returns a tuple with the Marks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentTextInner) GetMarksOk() ([]string, bool) {
	if o == nil || IsNil(o.Marks) {
		return nil, false
	}
	return o.Marks, true
}

// HasMarks returns a boolean if a field has been set.
func (o *ContentTextInner) HasMarks() bool {
	if o != nil && !IsNil(o.Marks) {
		return true
	}

	return false
}

// SetMarks gets a reference to the given []string and assigns it to the Marks field.
func (o *ContentTextInner) SetMarks(v []string) {
	o.Marks = v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *ContentTextInner) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentTextInner) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *ContentTextInner) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *ContentTextInner) SetColor(v string) {
	o.Color = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *ContentTextInner) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentTextInner) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *ContentTextInner) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *ContentTextInner) SetLink(v string) {
	o.Link = &v
}

func (o ContentTextInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentTextInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Marks) {
		toSerialize["marks"] = o.Marks
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	return toSerialize, nil
}

type NullableContentTextInner struct {
	value *ContentTextInner
	isSet bool
}

func (v NullableContentTextInner) Get() *ContentTextInner {
	return v.value
}

func (v *NullableContentTextInner) Set(val *ContentTextInner) {
	v.value = val
	v.isSet = true
}

func (v NullableContentTextInner) IsSet() bool {
	return v.isSet
}

func (v *NullableContentTextInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentTextInner(val *ContentTextInner) *NullableContentTextInner {
	return &NullableContentTextInner{value: val, isSet: true}
}

func (v NullableContentTextInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentTextInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


