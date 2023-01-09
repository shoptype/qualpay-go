# Checkout

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckoutId** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 32 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;A unique identifier that identifies a checkout. | [optional] [default to null]
**CheckoutLink** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 128 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;A redirect link that will direct a customer to the payment page. This link is unique for each checkout. | [optional] [default to null]
**DbTimestamp** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The checkout resource creation timestamp.  | [optional] [default to null]
**MerchantId** | **int64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 16 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Unique ID assigned to a merchant. | [optional] [default to null]
**ExpiryTime** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The expiration timestamp for the checkout_link. Once expired, the link cannot be used to make a payment.   | [optional] [default to null]
**AmtTran** | **float64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 12,2 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The total amount of the transaction including sales tax (if applicable). | [default to null]
**AmtBalance** | **float64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 12,2 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Amount pending. Valid only for partial payments. If amount is paid in full, the amt_balance will be 0. | [optional] [default to null]
**TranCurrency** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 3 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Numeric currency code of the transaction. Refer to &lt;a href&#x3D;\&quot;/developer/api/reference#country-codes\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Country Codes&lt;/a&gt; for a list of currency codes. If the profile_id field is provided, this field will be overridden by the profile’s tran_currency. | [optional] [default to null]
**PurchaseId** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 25 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The purchase identifier (also referred to as the invoice number) generated by the merchant. | [optional] [default to null]
**ProfileId** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 20 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The unique profile ID that will be used during a payment gateway request.  | [optional] [default to null]
**MerchRefNum** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 128 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Merchant provided reference value that will be stored with the transaction data and will be included with the transaction data reported in the Merchant Manager. This value will also be attached to any lifecycle transactions (e.g. retrieval requests and chargebacks) that may occur. | [optional] [default to null]
**CustomerFirstName** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 32 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The customer&#x27;s first name.  | [optional] [default to null]
**CustomerLastName** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 32 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The customer&#x27;s last name. | [optional] [default to null]
**CustomerEmail** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 64 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The customer&#x27;s email address. | [optional] [default to null]
**CustomerPhone** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 16 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The customer&#x27;s phone number with country code. | [optional] [default to null]
**BillingAddr1** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 64 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The billing address of the cardholder. | [optional] [default to null]
**BillingState** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 2 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The billing state of the cardholder. | [optional] [default to null]
**BillingCity** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 64 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The billing city of the cardholder. | [optional] [default to null]
**BillingZip** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 10 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The billing zip code of the cardholder. | [optional] [default to null]
**CustomerId** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 32 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The customer id associated with this transaction. | [optional] [default to null]
**Preferences** | [***CheckoutPreferences**](CheckoutPreferences.md) |  | [optional] [default to null]
**Transactions** | [**[]GatewayResponse**](GatewayResponse.md) | An array of responses from payment gateway. Each element of the array represents a transaction for this checkout. In addition, if subscribed to a checkout webhook, the transaction object is also posted to the Webhook URL once a checkout payment is made. | [optional] [default to null]
**Subscription** | [***CheckoutSubscription**](CheckoutSubscription.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
