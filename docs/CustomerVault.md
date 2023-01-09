# CustomerVault

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 32 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Unique ID established by merchant to identify a customer. The customer ID can be used to add subscriptions using the Recurring Billing API or make payments using the Payment Gateway API. Once established, this ID cannot be updated. This field is case sensitive. Only letters and numbers are allowed in a Customer ID. | [optional] [default to null]
**RecId** | **int64** | For INTERNAL USE ONLY. | [optional] [default to null]
**NodeId** | **int64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 16 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;The Merchant ID to which this customer belongs. | [optional] [default to null]
**CustomerFirstName** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 32 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Customer&#x27;s first name. | [optional] [default to null]
**CustomerLastName** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 32 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Customer&#x27;s last name. | [optional] [default to null]
**CustomerFirmName** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 64 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Customer&#x27;s business name, if applicable.  | [optional] [default to null]
**CustomerPhone** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 16 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Customer&#x27;s phone number. | [optional] [default to null]
**CustomerEmail** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 64 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Customer&#x27;s email address.  | [optional] [default to null]
**ReferenceId** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 128 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Merchant provided reference value that will be stored with the customer vault data. | [optional] [default to null]
**Comments** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 128 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Comment | [optional] [default to null]
**DeveloperId** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 32 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Use to indicate which company developed the integration, or the name of the solution or system that is connected to us. Suggested usage is softwareABCv1.0 or companyXYZv2.0. | [optional] [default to null]
**PrimaryCardNumber** | **string** | Default Card Number for this customer. | [optional] [default to null]
**ShippingAddresses** | [**[]ShippingAddress**](ShippingAddress.md) | &lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;An array of shipping addresses.  | [optional] [default to null]
**BillingCards** | [**[]BillingCard**](BillingCard.md) | &lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;An array of billing cards.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

