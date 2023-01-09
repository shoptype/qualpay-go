# PgApiEmailReceiptRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | **int64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 12 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Unique identifier on our system. | [default to null]
**DeveloperId** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 32 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Use to indicate which company developed the integration or the name of the payment solution that is connected to us. Suggested usage is softwareABCv1.0 or companyXYZv2.0.  | [optional] [default to null]
**EmailAddress** | **[]string** |  AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;An array of email addresses to which the transaction receipt should be sent to.  | [default to null]
**LogoUrl** | **string** |  AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;A link to the logo image that will be included in the transaction receipt.  | [optional] [default to null]
**VendorId** | **int64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 12 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Identifies the vendor to which this email receipt request applies. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

