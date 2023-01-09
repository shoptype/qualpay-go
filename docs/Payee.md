# Payee

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayeeId** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 32 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;A unique ID generated to identify a payee. | [optional] [default to null]
**PayeeName** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 64AN &lt;br&gt;&lt;strong&gt;Description:&lt;/strong&gt;Payee name. | [optional] [default to null]
**PayeePhone** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 32 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt; Payee phone number. Prepend with + and country code. | [optional] [default to null]
**PayeeEmail** | **[]string** | &lt;strong&gt;Description:&lt;/strong&gt;Payee email addresses | [optional] [default to null]
**PayeeAddress** | [***AddPayeeAddressRequest**](addPayeeAddressRequest.md) |  | [optional] [default to null]
**PayeeAccount** | [***PayeeAccountResponse**](PayeeAccountResponse.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

