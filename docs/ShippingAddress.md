# ShippingAddress

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShippingId** | **int64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 10 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Unique id generated to identify the shipping address for a customer. | [optional] [default to null]
**ShippingFirstName** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 32 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Shipping first name. | [optional] [default to null]
**ShippingLastName** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 32 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Shipping last name. | [optional] [default to null]
**ShippingFirmName** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 64 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Shipping address business name, if applicable.  | [optional] [default to null]
**ShippingAddr1** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 128 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Shipping Address line item 1. | [optional] [default to null]
**ShippingAddr2** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 128 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Shipping Address line item 2. | [optional] [default to null]
**ShippingCity** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 64 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Shipping city. | [optional] [default to null]
**ShippingState** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 3 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Shipping state. | [optional] [default to null]
**ShippingZip** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 10 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Shipping zip. | [optional] [default to null]
**ShippingZip4** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 4 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Shipping zip+4 code, if applicable. | [optional] [default to null]
**ShippingCountry** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 128 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Shipping Country. | [optional] [default to null]
**ShippingCountryCode** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 3 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;ISO numeric country code for the shipping address. Refer to &lt;a href&#x3D;\&quot;/developer/api/reference#country-codes\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Country Codes&lt;/a&gt; for a list of country codes. | [optional] [default to null]
**Primary** | **bool** | &lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;If set to true, this is the default shipping address. &lt;br&gt;&lt;strong&gt;Default: &lt;/strong&gt;false | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

