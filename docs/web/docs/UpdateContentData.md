# UpdateContentData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | article id | 
**Type** | **string** |  | 
**Attributes** | [**UpdateContentDataAttributes**](UpdateContentDataAttributes.md) |  | 

## Methods

### NewUpdateContentData

`func NewUpdateContentData(id string, type_ string, attributes UpdateContentDataAttributes, ) *UpdateContentData`

NewUpdateContentData instantiates a new UpdateContentData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateContentDataWithDefaults

`func NewUpdateContentDataWithDefaults() *UpdateContentData`

NewUpdateContentDataWithDefaults instantiates a new UpdateContentData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateContentData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateContentData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateContentData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *UpdateContentData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateContentData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateContentData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *UpdateContentData) GetAttributes() UpdateContentDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UpdateContentData) GetAttributesOk() (*UpdateContentDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UpdateContentData) SetAttributes(v UpdateContentDataAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


