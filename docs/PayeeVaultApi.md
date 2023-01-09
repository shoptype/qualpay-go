# qualpay{{classname}}

All URIs are relative to */platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPayee**](PayeeVaultApi.md#AddPayee) | **Post** /vault/payee | Add a Payee
[**BrowsePayees**](PayeeVaultApi.md#BrowsePayees) | **Get** /vault/payee | Get all Payees
[**DeletePayee**](PayeeVaultApi.md#DeletePayee) | **Delete** /vault/payee/{payee_id} | Delete a Payee
[**GetPayee**](PayeeVaultApi.md#GetPayee) | **Get** /vault/payee/{payee_id} | Get by Payee ID
[**UpdatePayee**](PayeeVaultApi.md#UpdatePayee) | **Put** /vault/payee/{payee_id} | Update a Payee
[**UpdatePayeeAccount**](PayeeVaultApi.md#UpdatePayeeAccount) | **Post** /vault/payee/account/{payee_id} | Update a Payee account

# **AddPayee**
> PayeeResponse AddPayee(ctx, body)
Add a Payee

Adds a payee to the Payee Vault. A payee name is required to add a payee.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddPayeeRequest**](AddPayeeRequest.md)| Add Payee Request | 

### Return type

[**PayeeResponse**](PayeeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BrowsePayees**
> PayeeListResponse BrowsePayees(ctx, optional)
Get all Payees

Gets an array of payees from vault. You can use filters to limit the results returned. Optional query parameters determine size and sort order of the returned array.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PayeeVaultApiBrowsePayeesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PayeeVaultApiBrowsePayeesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| The number of records in the result. | [default to 10]
 **orderOn** | **optional.String**| The field on which the results will be sorted on. Refer to the response model for available fields. | [default to payee_id]
 **orderBy** | **optional.String**| Ascending or Descending Sort order of the result. Possible values are: asc (Ascending sort order), desc (Descending sort order) | [default to desc]
 **page** | **optional.Int32**| Zero-based page number, use this to choose a page when there are more results than the count parameter. | [default to 0]
 **filter** | [**optional.Interface of []string**](string.md)| Results can be filtered by custom filter criteria. Refer to [Filter](/developer/api/reference#filters) to use the filter parameter. | 
 **merchantId** | **optional.Int64**| Unique ID assigned to a merchant. Required only if the request is sent on behalf of another merchant. | [default to 0]

### Return type

[**PayeeListResponse**](PayeeListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePayee**
> QpApiResponse DeletePayee(ctx, payeeId, optional)
Delete a Payee

Deletes a payee from the Payee vault. **This operation cannot be reversed.**

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payeeId** | **string**| Payee ID | 
 **optional** | ***PayeeVaultApiDeletePayeeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PayeeVaultApiDeletePayeeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **merchantId** | **optional.Int64**| Unique ID assigned to a merchant. Required only if the request is sent on behalf of another merchant. | [default to 0]

### Return type

[**QpApiResponse**](QPApiResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPayee**
> PayeeResponse GetPayee(ctx, payeeId, optional)
Get by Payee ID

Gets a payee vault record by payee_id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payeeId** | **string**| Payee ID | 
 **optional** | ***PayeeVaultApiGetPayeeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PayeeVaultApiGetPayeeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **merchantId** | **optional.Int64**| Unique ID assigned to a merchant. Required only if the request is sent on behalf of another merchant. | 

### Return type

[**PayeeResponse**](PayeeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePayee**
> PayeeResponse UpdatePayee(ctx, body, payeeId)
Update a Payee

Updates a payee vault record. You can send only the fields that require an update in the request. Use the \"Update a Payee Account\" to update account details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdatePayeeRequest**](UpdatePayeeRequest.md)| Update Payee Request | 
  **payeeId** | **string**| Payee ID | 

### Return type

[**PayeeResponse**](PayeeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePayeeAccount**
> PayeeAccountResponse UpdatePayeeAccount(ctx, payeeId, optional)
Update a Payee account

Updates payee's account details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payeeId** | **string**| Payee ID | 
 **optional** | ***PayeeVaultApiUpdatePayeeAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PayeeVaultApiUpdatePayeeAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdatePayeeAccountRequest**](UpdatePayeeAccountRequest.md)| Update Payee Account Request | 

### Return type

[**PayeeAccountResponse**](PayeeAccountResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

