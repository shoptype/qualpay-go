# CheckoutPreferences

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuccessUrl** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 128 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;A URL to which the customer will be directed after a successful payment. If not defined in the API or in the checkout settings page, the customer will be directed to the default receipt page. Must be formatted as http://www.domain.com/ or https://www.domain.com/.  | [optional] [default to null]
**FailureUrl** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 128 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;A URL to which the customer will be directed if a transaction is declined. If not defined in the API or in the checkout settings page, the customer will be directed to the default receipt page. Must be formatted as http://www.domain.com/ or https://www.domain.com/.  | [optional] [default to null]
**AllowPartialPayments** | **bool** | &lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;If set to true, the customer can make changes to the transaction amount. | [optional] [default to null]
**AllowSaveCard** | **bool** | &lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Applicable only for checkouts associated with a customer. If set to true, the checkout page will display the &#x27;Save Card&#x27; checkbox that lets the customer decide if the card information can be saved. If set to false, the &#x27;Save Card&#x27; box is not displayed and the customer card information is always updated.&lt;br&gt;&lt;strong&gt;Default: &lt;/strong&gt;false | [optional] [default to null]
**AllowAchPayment** | **bool** | &lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Applicable only if your merchant account is configured to take ACH payments. If set to true, the checkout page will present ACH as one of the payment modes to the customer. &lt;br&gt;&lt;strong&gt;Default: &lt;/strong&gt;false | [optional] [default to null]
**EmailReceipt** | **bool** | &lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;If set to true and the customer email address is provided, a receipt is sent to the customer.  | [optional] [default to null]
**RequestType** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 10 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Identifies the type of request when the customer submits the payment data on the checkout page.   | [optional] [default to null]
**ExpireInSecs** | **int32** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 7 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The time period for which the checkout_link will be valid in seconds. The minimum value for this field is 900 seconds  (5 minutes),  the maximum value is 7776000 (90 days). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

