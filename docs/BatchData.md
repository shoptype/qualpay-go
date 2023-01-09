# BatchData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchId** | **int64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 10 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Unique ID assigned to this batch. | [optional] [default to null]
**MerchantId** | **int64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 16 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Unique ID assigned to a merchant. | [optional] [default to null]
**TranCurrency** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 3 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The ISO 4217 numeric currency code of the batch. | [optional] [default to null]
**BatchNumber** | **int64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 3 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;A non-unique ID assigned by the merchant&#x27;s terminal, POS device, or gateway for this batch, in the range of 1 - 999. | [optional] [default to null]
**BatchDate** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 10 AN, in YYYY-MM-DD format&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The date the batch was settled to us. | [optional] [default to null]
**AmtTotal** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 12,2 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The total, or net amount of the batch, in the batches currency. | [optional] [default to null]
**CntTotal** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 9 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The count of transactions in the batch, including purchases and credits. | [optional] [default to null]
**AmtPurch** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 12,2 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The amount of purchases within the batch, in the batches currency. | [optional] [default to null]
**CntPurch** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 9 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The count of purchases in the batch. | [optional] [default to null]
**AmtReturn** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 12,2 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The amount of returns, or credits back to the customer, in the batch&#x27;s currency. | [optional] [default to null]
**CntReturn** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 9 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The count of returns or credits back to the customer in the batch. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

