# ApplicationModel

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | **int64** | Unique ID assigned to this application. | [optional] [default to null]
**MerchantId** | **int64** | Unique ID assigned to a merchant. | [optional] [default to null]
**AppStatus** | **string** | The current status of the application. | [optional] [default to null]
**DbaName** | **string** | The &#x27;doing business as&#x27; name, as it is currently on the merchant application. | [optional] [default to null]
**DbTimestamp** | [**time.Time**](time.Time.md) | The timestamp the application was created. | [optional] [default to null]
**SubmitTimestamp** | [**time.Time**](time.Time.md) | The timestamp the application was submitted as complete. | [optional] [default to null]
**CreditTimestamp** | [**time.Time**](time.Time.md) | The timestamp the application&#x27;s credit decision was made. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

