# Customer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingAddr1** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Customer billing address street. | [optional] [default to null]
**BillingAddr2** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Customer billing address, line 2. | [optional] [default to null]
**BillingCity** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Customer billing city. | [optional] [default to null]
**BillingCountry** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Customer billing country. | [optional] [default to null]
**BillingCountryCode** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 3 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;ISO numeric country code for the billing address. Refer to &lt;a href&#x3D;\&quot;/developer/api/reference#country-codes\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Country Codes&lt;/a&gt; for a list of country codes. | [optional] [default to null]
**BillingState** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Customer billing state (abbreviated). | [optional] [default to null]
**BillingZip** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Customer billing zip code. | [optional] [default to null]
**BillingZip4** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 4 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Customer billing zip+4 code if applicable. | [optional] [default to null]
**CustomerEmail** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 64 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Customer e-mail address. | [optional] [default to null]
**CustomerFirmName** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 64 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Customer Business name if applicable.&lt;br&gt;&lt;strong&gt;Conditional Requirement: &lt;/strong&gt;Either customer first and last name or firm name is required.  | [optional] [default to null]
**CustomerFirstName** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 32 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Customer first name.&lt;br&gt;&lt;strong&gt;Conditional Requirement: &lt;/strong&gt;Either customer first and last name or firm name is required.  | [optional] [default to null]
**CustomerLastName** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 32 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Customer last name.&lt;br&gt;&lt;strong&gt;Conditional Requirement: &lt;/strong&gt;Either customer first and last name or firm name is required.  | [optional] [default to null]
**CustomerPhone** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 16 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Customer phone number. | [optional] [default to null]
**ShippingAddresses** | [**[]PgShippingAddress**](PGShippingAddress.md) | List of shipping addresses for customer. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

