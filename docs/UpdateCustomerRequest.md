# UpdateCustomerRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerFirstName** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 32 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Customer&#x27;s first name.&lt;br&gt;&lt;strong&gt;Conditional Requirement: &lt;/strong&gt;Either customer name or customer firm name is required. | [optional] [default to null]
**CustomerLastName** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 32 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Customer&#x27;s last name.&lt;br&gt;&lt;strong&gt;Conditional Requirement: &lt;/strong&gt;Either customer name or customer firm name is required. | [optional] [default to null]
**CustomerFirmName** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 64 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Customer&#x27;s business name, if applicable. &lt;br&gt;&lt;strong&gt;Conditional Requirement: &lt;/strong&gt;Either customer name or customer firm name is required. | [optional] [default to null]
**CustomerPhone** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 16 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Customer&#x27;s phone number. | [optional] [default to null]
**CustomerEmail** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 64 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Customer&#x27;s email address.  | [optional] [default to null]
**ReferenceId** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 128 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Merchant provided reference value that will be stored with the customer vault data. | [optional] [default to null]
**Comments** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 128 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Comment. | [optional] [default to null]
**ShippingAddresses** | [**[]AddShippingAddressRequest**](AddShippingAddressRequest.md) | &lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;An array of shipping addresses.  | [optional] [default to null]
**BillingCards** | [**[]AddBillingCardRequest**](AddBillingCardRequest.md) | &lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;An array of billing cards.  | [optional] [default to null]
**MerchantId** | **int64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 16 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Identifies the merchant to whom this request applies. Optional field, applicable only if the request is sent on behalf of another merchant.&lt;br&gt;&lt;strong&gt;Conditional Requirement: &lt;/strong&gt;Required if this request is on behalf of another merchant. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

