# EmbeddedKey

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransientKey** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 32 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;A single use token used for loading embedded fields. The key will be invalidated   when a card is successfully verified. If unused, the token will expire in 30 minutes.  | [optional] [default to null]
**MerchantId** | **int64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 16 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Unique ID assigned to a merchant. | [optional] [default to null]
**DbTimestamp** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;This field contains the transient key creation timestamp.  | [optional] [default to null]
**ExpiryTime** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;This field contains the timestamp when the transient key will expire. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

