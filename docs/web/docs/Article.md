# Article

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**ArticleData**](ArticleData.md) |  | 

## Methods

### NewArticle

`func NewArticle(data ArticleData, ) *Article`

NewArticle instantiates a new Article object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArticleWithDefaults

`func NewArticleWithDefaults() *Article`

NewArticleWithDefaults instantiates a new Article object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *Article) GetData() ArticleData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Article) GetDataOk() (*ArticleData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Article) SetData(v ArticleData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


