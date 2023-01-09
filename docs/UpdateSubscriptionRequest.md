# UpdateSubscriptionRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateStart** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 10, YYYY-MM-DD format&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Date the subscription will start. | [optional] [default to null]
**CardId** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 32 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Update the card id to be used for recurring charges. The card id can be updated any time on an active subscription. Send an empty string in this field, if you want to udpate the record to charge by the primary card.  | [optional] [default to null]
**ProfileId** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 20 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Payment Gateway Profile id that will be used when billing transactions. The new profile should belong the same currency as original.  | [optional] [default to null]
**MerchantId** | **int64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 16 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Identifies the merchant to whom this request applies. Optional field, applicable only if the request is sent on behalf of another merchant.&lt;br&gt;&lt;strong&gt;Conditional Requirement: &lt;/strong&gt;Required if this request is on behalf of another merchant. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

