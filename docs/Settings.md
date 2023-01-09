# Settings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | **int64** | The application ID for this merchant account.  Only returned if an application exists. | [optional] [default to null]
**MerchantId** | **int64** | The merchant ID this merchant account. | [optional] [default to null]
**DbaName** | **string** | The DBA name for the merchant account. | [optional] [default to null]
**PaymentsAccepted** | [**[]Payment**](Payment.md) | An array of the card types and currency accepted by the merchant account. | [optional] [default to null]
**PaymentProfiles** | [**[]PaymentProfile**](PaymentProfile.md) | An array payment profiles available for the merchant account. | [optional] [default to null]
**UserLogin** | **string** | The user login created for the merchant account.  Only returned if a new user was requested during account creation. | [optional] [default to null]
**ResetUrl** | **string** | The URL to establish a new password for the created user.  Only returned if a new user was requested during account creation. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

