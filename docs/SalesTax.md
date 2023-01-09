# SalesTax

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 10 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Sales Tax type. Possible values are  &lt;br&gt;&lt;strong&gt;RATE&lt;strong&gt; The value field contains the tax rate. The system will automatically calculate the tax amount based on the rate.  &lt;br&gt;&lt;strong&gt;AMOUNT&lt;strong&gt; The value field contains the tax amount. | [optional] [default to null]
**Value** | **float64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 7,5 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Sales Tax value.If type is RATE, set this field to the sales tax rate. The system will automatically calculate the  sales tax amount (amt_tax) based on the rate. &lt;br&gt;If type is AMOUNT, set this field to the sales tax amount. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

