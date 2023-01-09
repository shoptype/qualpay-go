# PaymentTerm

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 10 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Payment Term type. By default, this field is set to DAYS.  | [optional] [default to null]
**Value** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 25 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Payment Term. The invoice_due_date will be automatically updated based on this value. If type is DATE, this field should contain the invoice_due_date in &#x27;YYYY-MM-DD&#x27; format. If type is DAYS, this field should contain the number of days from invoice date when payment is due.   | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

