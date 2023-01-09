# qualpay{{classname}}

All URIs are relative to */platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEmbeddedTransientKey**](EmbeddedFieldsApi.md#GetEmbeddedTransientKey) | **Get** /embedded | Get Transient Key

# **GetEmbeddedTransientKey**
> EmbeddedKeyResponse GetEmbeddedTransientKey(ctx, )
Get Transient Key

Gets a transient key for use with Embedded Fields. The transient key is a one time key that is used to invoke Embedded Fields. This key is valid for up to 12 hours. The key will be automatically invalidated once a card is successfully tokenized using Embedded Fields or if there are too many declines during verification.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**EmbeddedKeyResponse**](EmbeddedKeyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

