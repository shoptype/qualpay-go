# AddPayeeRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VendorId** | **int64** | Unique Vendor ID assigned to integrator. | [optional] [default to null]
**PayeeName** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 64AN &lt;br&gt;&lt;strong&gt;Description:&lt;/strong&gt;Payee name. | [default to null]
**PayeePhone** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 32 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt; Payee phone number. Prepend with + and country code. | [optional] [default to null]
**PayeeEmail** | **[]string** | &lt;strong&gt;Description:&lt;/strong&gt;Payee email addresses | [optional] [default to null]
**PayeeAddress** | [**[]AddPayeeAddressRequest**](addPayeeAddressRequest.md) |  | [optional] [default to null]
**PayeeAccount** | [**[]AddPayeeAccountRequest**](addPayeeAccountRequest.md) |  | [optional] [default to null]
**MerchantId** | **int32** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 16 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Identifies the merchant to whom this request applies. Optional field, applicable only if the request is sent on behalf of another merchant.&lt;br&gt;&lt;strong&gt;Conditional Requirement: &lt;/strong&gt;Required if this request is on behalf of another merchant. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

