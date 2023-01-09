# AccountUpdaterSummaryData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **int64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 10 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Unique ID assigned to the account updater request. | [optional] [default to null]
**MerchantId** | **int64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 16 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Unique ID assigned to a Merchant. | [optional] [default to null]
**Status** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 16 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Status of the request. The status can be one of the following               &lt;ul&gt;              &lt;li&gt;INPROGRESS - Account updater request has been validated, the request is ready to be sent to the Account updater engine. &lt;/li&gt;              &lt;li&gt;QUEUED - Account updater request is sent to the Account updater engine. The request can no longer be canceled. &lt;/li&gt;              &lt;li&gt;COMPLETED - Account updater request has completed.&lt;/li&gt;              &lt;li&gt;CANCELED - Account updater request has been canceled.&lt;/li&gt;              &lt;/ul&gt; | [optional] [default to null]
**CntRequest** | **int32** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 10 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Number of cards in this request. | [optional] [default to null]
**CntResponse** | **int32** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 10 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Number of cards in this request that were reviewed by the account updater service. (This should match the cnt_request.). The cnt_response field will be zero if the status field is not COMPLETED. | [optional] [default to null]
**CntUpdate** | **int32** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 10 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Number of cards in this request that were updated by the account updater service. The cnt_update field will be zero if the status field is not COMPLETED or if there are no updates by the account updater service. | [optional] [default to null]
**RequestDate** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 10 AN, in  YYYY-MM-DD format&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Account updater request date. Date when a harvest request was sent to the account updater service. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

