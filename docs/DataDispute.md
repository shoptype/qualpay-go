# DataDispute

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecId** | **int64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 18 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Unique ID assigned to this dispute. | [optional] [default to null]
**DataType** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 1 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The type of dispute. For all types, please see &lt;a href&#x3D;\&quot;/developer/api/reference#data_type\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Dispute Types&lt;/a&gt;. | [optional] [default to null]
**MerchantId** | **int64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 16 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Unique ID assigned to a merchant. | [optional] [default to null]
**IncomingDate** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 10 AN, in YYYY-MM-DD format&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The date we received the dispute from the card issuer. | [optional] [default to null]
**DbaName** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 25 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The doing business as name of the merchant. | [optional] [default to null]
**BatchId** | **int64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 10 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Unique ID assigned to this batch. | [optional] [default to null]
**CardNumber** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 16 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The truncated card number of the dispute. | [optional] [default to null]
**CardType** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 2 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The card brand of this dispute. | [optional] [default to null]
**TranDate** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 10 AN, in YYYY-MM-DD format&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The date the initial transaction occured. | [optional] [default to null]
**AmtTran** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 12,2 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The amount of the original settled transaction. | [optional] [default to null]
**AmtDispute** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 12,2 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The amount of the dispute: equal to, or less than the amt_tran (when in USD). | [optional] [default to null]
**TranCurrency** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 3 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The ISO 4217 numeric currency code of the dispute. | [optional] [default to null]
**PurchaseId** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 25 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;A merchant supplied tracking number, generally an invoice or purchase number. This number may be visible to the cardholder, depending on card issuer. | [optional] [default to null]
**MerchRefNum** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 128 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;A merchant supplied tracking number which is stored by us, and does not pass to the card issuer. | [optional] [default to null]
**PgId** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 32 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;If the transaction originated through Virtual Terminal, or Payment Gateway, this is the tracking ID returned in the gateway response. | [optional] [default to null]
**AuthCode** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 6 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The authorization code provided by the card issuer when the card was approved. | [optional] [default to null]
**ReasonCode** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 8 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The card association reason why the dispute was issued. | [optional] [default to null]
**RecIdLinked** | **string** | TODO | [optional] [default to null]
**AuthAvsResult** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 1 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The AVS (address validation service) match code of the original transaction. | [optional] [default to null]
**AuthCvv2Result** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 1 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The CVV2 match code of the original transaction. | [optional] [default to null]
**CardNumberOriginal** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 16 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The truncated card number of the original captured transactions. In some rare cases a dispute may be issued on a different card than the original. This can occur when, for example, the card was was re-issued between the transaction date, and the dispute date. | [optional] [default to null]
**DateStatusChange** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 10 AN, in YYYY-MM-DD format&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Records dispute status date. The dispute status of a dispute record will change over time and as a result of merchant actions. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

