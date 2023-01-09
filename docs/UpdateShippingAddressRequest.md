# UpdateShippingAddressRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShippingId** | **int64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 10 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;ID of the shipping address that should be updated. Shipping ID should belong to the customer.  | [default to null]
**ShippingFirstName** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 32 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Shipping first name. | [default to null]
**ShippingLastName** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 32 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Shipping last name. | [default to null]
**ShippingFirmName** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 64 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Shipping address business name, if applicable.  | [optional] [default to null]
**ShippingAddr1** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 128 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Shipping address line item 1. | [optional] [default to null]
**ShippingAddr2** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 128 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Shipping address line item 2. | [optional] [default to null]
**ShippingCity** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 64 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Shipping city. | [optional] [default to null]
**ShippingState** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 3 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Shipping state. | [optional] [default to null]
**ShippingCountry** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 128 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Shipping country. | [optional] [default to null]
**ShippingCountryCode** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 3 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;ISO numeric country code for the shipping address. Refer to &lt;a href&#x3D;\&quot;/developer/api/reference#country-codes\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Country Codes&lt;/a&gt; for a list of country codes. | [optional] [default to null]
**ShippingZip** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 10 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Shipping zip. | [optional] [default to null]
**ShippingZip4** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 4 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Shipping zip+4 code if applicable. | [optional] [default to null]
**Primary** | **bool** | &lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Set this field to true if this should be the default address. &lt;br&gt;&lt;strong&gt;Default: &lt;/strong&gt;false | [optional] [default to null]
**MerchantId** | **int64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 16 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Identifies the merchant to whom this request applies. Optional field, applicable only if the request is sent on behalf of another merchant.&lt;br&gt;&lt;strong&gt;Conditional Requirement: &lt;/strong&gt;Required if this request is on behalf of another merchant. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

