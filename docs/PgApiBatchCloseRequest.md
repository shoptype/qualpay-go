# PgApiBatchCloseRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | **int64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 12 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Unique identifier on our system. | [default to null]
**DeveloperId** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 32 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Use to indicate which company developed the integration or the name of the payment solution that is connected to us. Suggested usage is softwareABCv1.0 or companyXYZv2.0.  | [optional] [default to null]
**EchoFields** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;This field contains a JSON array of field data that will be echoed back in the response message. | [optional] [default to null]
**LocId** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 4 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;When a merchant has more than one location using the same currency, this value is used to specify the specific location for this request. | [optional] [default to null]
**ProfileId** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 20 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Explicitly identifies which Payment Gateway profile should be used for the request. | [optional] [default to null]
**ReportData** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;This field contains a JSON array of field data that will be included with the transaction data reported in Merchant Manager. | [optional] [default to null]
**SessionId** | **string** | INTERNAL USE ONLY. | [optional] [default to null]
**TranCurrency** | **int32** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 3 N&lt;br&gt;&lt;strong&gt;Default: &lt;/strong&gt;840&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;ISO numeric currency code for the transaction. Refer to &lt;a href&#x3D;\&quot;/developer/api/reference#country-codes\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Country Codes&lt;/a&gt; for a list of currency codes. | [optional] [default to null]
**UserId** | **int64** | INTERNAL USE ONLY. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

