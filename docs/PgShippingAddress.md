# PgShippingAddress

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Primary** | **bool** | &lt;br&gt;&lt;strong&gt;Default: &lt;/strong&gt;false&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Identifies the primary shipping address where there are multiple shipping addresses. | [optional] [default to null]
**ShippingAddr1** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Customer street and number, P.O. box, c/o. | [default to null]
**ShippingAddr2** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Customer apartment, suite, unit, building, floor, etc. | [optional] [default to null]
**ShippingCity** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Customer shipping city. | [default to null]
**ShippingCountry** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Customer shipping country. | [optional] [default to null]
**ShippingCountryCode** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 3 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;ISO numeric country code for the shipping address. Refer to &lt;a href&#x3D;\&quot;/developer/api/reference#country-codes\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Country Codes&lt;/a&gt; for a list of country codes. If the country code is provided, the shipping_country will be updated to match the country code.  | [optional] [default to null]
**ShippingFirmName** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 64 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Business name if applicable.  | [optional] [default to null]
**ShippingFirstName** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 32 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Customer shipping first name. | [default to null]
**ShippingLastName** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 32 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Customer shipping last name. | [default to null]
**ShippingState** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Customer shipping state (abbreviated). | [default to null]
**ShippingZip** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Customer shipping zip code. | [default to null]
**ShippingZip4** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 4 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Customer shipping zip+4 code if applicable. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

