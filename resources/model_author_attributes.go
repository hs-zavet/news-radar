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

// checks if the AuthorAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorAttributes{}

// AuthorAttributes struct for AuthorAttributes
type AuthorAttributes struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
	Avatar string `json:"avatar"`
	Email *string `json:"email,omitempty"`
	Telegram *string `json:"telegram,omitempty"`
	Twitter *string `json:"twitter,omitempty"`
}

type _AuthorAttributes AuthorAttributes

// NewAuthorAttributes instantiates a new AuthorAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorAttributes(name string, desc string, avatar string) *AuthorAttributes {
	this := AuthorAttributes{}
	this.Name = name
	this.Desc = desc
	this.Avatar = avatar
	return &this
}

// NewAuthorAttributesWithDefaults instantiates a new AuthorAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorAttributesWithDefaults() *AuthorAttributes {
	this := AuthorAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *AuthorAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuthorAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuthorAttributes) SetName(v string) {
	o.Name = v
}

// GetDesc returns the Desc field value
func (o *AuthorAttributes) GetDesc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Desc
}

// GetDescOk returns a tuple with the Desc field value
// and a boolean to check if the value has been set.
func (o *AuthorAttributes) GetDescOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Desc, true
}

// SetDesc sets field value
func (o *AuthorAttributes) SetDesc(v string) {
	o.Desc = v
}

// GetAvatar returns the Avatar field value
func (o *AuthorAttributes) GetAvatar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value
// and a boolean to check if the value has been set.
func (o *AuthorAttributes) GetAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Avatar, true
}

// SetAvatar sets field value
func (o *AuthorAttributes) SetAvatar(v string) {
	o.Avatar = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *AuthorAttributes) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorAttributes) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *AuthorAttributes) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *AuthorAttributes) SetEmail(v string) {
	o.Email = &v
}

// GetTelegram returns the Telegram field value if set, zero value otherwise.
func (o *AuthorAttributes) GetTelegram() string {
	if o == nil || IsNil(o.Telegram) {
		var ret string
		return ret
	}
	return *o.Telegram
}

// GetTelegramOk returns a tuple with the Telegram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorAttributes) GetTelegramOk() (*string, bool) {
	if o == nil || IsNil(o.Telegram) {
		return nil, false
	}
	return o.Telegram, true
}

// HasTelegram returns a boolean if a field has been set.
func (o *AuthorAttributes) HasTelegram() bool {
	if o != nil && !IsNil(o.Telegram) {
		return true
	}

	return false
}

// SetTelegram gets a reference to the given string and assigns it to the Telegram field.
func (o *AuthorAttributes) SetTelegram(v string) {
	o.Telegram = &v
}

// GetTwitter returns the Twitter field value if set, zero value otherwise.
func (o *AuthorAttributes) GetTwitter() string {
	if o == nil || IsNil(o.Twitter) {
		var ret string
		return ret
	}
	return *o.Twitter
}

// GetTwitterOk returns a tuple with the Twitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorAttributes) GetTwitterOk() (*string, bool) {
	if o == nil || IsNil(o.Twitter) {
		return nil, false
	}
	return o.Twitter, true
}

// HasTwitter returns a boolean if a field has been set.
func (o *AuthorAttributes) HasTwitter() bool {
	if o != nil && !IsNil(o.Twitter) {
		return true
	}

	return false
}

// SetTwitter gets a reference to the given string and assigns it to the Twitter field.
func (o *AuthorAttributes) SetTwitter(v string) {
	o.Twitter = &v
}

func (o AuthorAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["desc"] = o.Desc
	toSerialize["avatar"] = o.Avatar
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Telegram) {
		toSerialize["telegram"] = o.Telegram
	}
	if !IsNil(o.Twitter) {
		toSerialize["twitter"] = o.Twitter
	}
	return toSerialize, nil
}

func (o *AuthorAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"desc",
		"avatar",
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

	varAuthorAttributes := _AuthorAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAuthorAttributes)

	if err != nil {
		return err
	}

	*o = AuthorAttributes(varAuthorAttributes)

	return err
}

type NullableAuthorAttributes struct {
	value *AuthorAttributes
	isSet bool
}

func (v NullableAuthorAttributes) Get() *AuthorAttributes {
	return v.value
}

func (v *NullableAuthorAttributes) Set(val *AuthorAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorAttributes(val *AuthorAttributes) *NullableAuthorAttributes {
	return &NullableAuthorAttributes{value: val, isSet: true}
}

func (v NullableAuthorAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


