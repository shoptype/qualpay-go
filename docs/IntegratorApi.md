# qualpay{{classname}}

All URIs are relative to */platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTestAccount**](IntegratorApi.md#CreateTestAccount) | **Post** /vendor/settings/account/add | Generate Test Account
[**GetMerchantSettings**](IntegratorApi.md#GetMerchantSettings) | **Get** /vendor/settings/{mid} | Merchant Settings
[**RotateApiKey**](IntegratorApi.md#RotateApiKey) | **Post** /vendor/settings/key/{mid} | Generate Security Key

# **CreateTestAccount**
> MerchantSettingsResponse CreateTestAccount(ctx, optional)
Generate Test Account

Generates a new test merchant account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IntegratorApiCreateTestAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IntegratorApiCreateTestAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of NewAccountRequest**](NewAccountRequest.md)| NewAccount | 

### Return type

[**MerchantSettingsResponse**](MerchantSettingsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMerchantSettings**
> MerchantSettingsResponse GetMerchantSettings(ctx, mid)
Merchant Settings

Request the card types and currencies supported by the requested merchant ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mid** | **int64**| Merchant ID | 

### Return type

[**MerchantSettingsResponse**](MerchantSettingsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RotateApiKey**
> RotateKeyResponse RotateApiKey(ctx, body, mid)
Generate Security Key

Generates an API key for the specified merchant account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RotateKeyRequest**](RotateKeyRequest.md)| Data | 
  **mid** | **int64**| Merchant ID | 

### Return type

[**RotateKeyResponse**](RotateKeyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

