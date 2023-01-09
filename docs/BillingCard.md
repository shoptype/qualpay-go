# BillingCard

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardNumber** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 16 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The masked card number.  | [optional] [default to null]
**ExpDate** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 4 N, MMYY format&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Expiration date of card number.  | [optional] [default to null]
**CardId** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 32 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Card ID received from a tokenization request. | [optional] [default to null]
**BillingFirstName** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 32 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Billing first name.  | [optional] [default to null]
**BillingLastName** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 32 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Billing last name.  | [optional] [default to null]
**BillingFirmName** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 64 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Business name on billing card, if applicable.  | [optional] [default to null]
**BillingEmail** | **[]string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 128 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Array of billing email address. | [optional] [default to null]
**BillingAddr1** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 128 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Billing street address. This address will also be used for AVS verification if AVS verification is enabled.  | [optional] [default to null]
**BillingAddr2** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 128 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Billing Address line item 2. | [optional] [default to null]
**BillingCity** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 64 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Billing city. | [optional] [default to null]
**BillingState** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 3 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Billing state. | [optional] [default to null]
**BillingZip** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 10 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Billing zip. | [optional] [default to null]
**BillingZip4** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 4 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Billing zip+4 code if applicable. | [optional] [default to null]
**BillingCountry** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 128 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Billing Country. | [optional] [default to null]
**BillingCountryCode** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 3 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;ISO numeric country code for the billing address. Refer to &lt;a href&#x3D;\&quot;/developer/api/reference#country-codes\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Country Codes&lt;/a&gt; for a list of country codes. | [optional] [default to null]
**Primary** | **bool** | &lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;True if this is the default card. The default card will be used in recurring billing payments.&lt;br&gt;&lt;strong&gt;Default: &lt;/strong&gt;false | [optional] [default to null]
**CardType** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 2 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Card Type. The card type is derived from the card number. Refer to &lt;a href&#x3D;\&quot;/developer/api/reference#card-types\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Card Types&lt;/a&gt; for a list of card types.  | [optional] [default to null]
**VerifiedDate** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Verified Date. The date the card was last verified successfully. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

