# InvoicePaymentRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 16 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The mode of payment. &lt;br&gt;&lt;ul&gt;&lt;li&gt;&lt;strong&gt;CARD&lt;/strong&gt; Invoice checkout payment by customer using a card. A card payment cannot be manually added. &lt;/li&gt;&lt;li&gt;&lt;strong&gt;CASH&lt;/strong&gt; Cash payment.&lt;/li&gt;&lt;li&gt;&lt;strong&gt;CHECK&lt;/strong&gt; Check Payment.&lt;/li&gt;&lt;li&gt;&lt;strong&gt;OTHER&lt;/strong&gt; Other modes of payment.&lt;/li&gt;&lt;/ul&gt; | [default to null]
**DatePayment** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 10 AN, in YYYY-MM-DD format&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Date the payment was made. | [default to null]
**AmtPaid** | **float64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 10, 2 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Amount paid. The payment currency should be the same as the invoice currency.  | [default to null]
**Description** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;A short description of the payment. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

