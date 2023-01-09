# StatementSettlementResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 1 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Response code from API. 0 indicates success. Refer to &lt;a href&#x3D;\&quot;/developer/api/reference#api-response-codes\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Platform API Response Codes&lt;/a&gt; for entire list of return codes. | [optional] [default to null]
**Message** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;A short description of the API response code. | [optional] [default to null]
**TotalPages** | **int32** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Total Number of pages.  | [optional] [default to null]
**TotalRecords** | **int32** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 3 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Total number of records upto a maximum of 100. | [optional] [default to null]
**Data** | [***StatementSettlementData**](StatementSettlementData.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

