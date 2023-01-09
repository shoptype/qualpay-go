# TransactionData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TranDate** | **string** | The date the transaction was captured by the merchant. | [optional] [default to null]
**ReferenceNumber** | **string** | The bank reference number of the deposit. | [optional] [default to null]
**TranStatus** | **string** | Transaction Status&lt;ul&gt;&lt;li&gt;A - Transaction is approved&lt;/li&gt;&lt;li&gt;H - Transaction Held&lt;/li&gt;&lt;li&gt;C - Transaction is captured&lt;/li&gt;&lt;li&gt;V - Transaction is voided by Merchant&lt;/li&gt;&lt;li&gt;v - Transaction is voided by System&lt;/li&gt;&lt;li&gt;K - Transaction is cancelled&lt;/li&gt;&lt;li&gt;D - Transaction is declined by issuer&lt;/li&gt;&lt;li&gt;F - Transaction failures other than Issuer Declines&lt;/li&gt;&lt;li&gt;S - Transaction Settled&lt;/li&gt;&lt;li&gt;P - Deposit Sent&lt;/li&gt;&lt;li&gt;N - Transaction Settled, and will be funded directly by issuer&lt;/li&gt;&lt;li&gt;R - Transaction Rejected&lt;/li&gt;&lt;/ul&gt; | [optional] [default to null]
**PurchaseId** | **string** | Purchase ID of the transaction | [optional] [default to null]
**AmtTran** | **float64** | Transaction Amount | [optional] [default to null]
**TranCurrency** | **string** | Numeric Currency Code. Refer to &lt;a href&#x3D;\&quot;/developer/api/reference#country-codes\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Country Codes&lt;/a&gt; for a list of currency codes. | [optional] [default to null]
**BatchDate** | **string** | The date the batch was settled to us. | [optional] [default to null]
**CardNumber** | **string** | The truncated card number of the dispute. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

