# qualpay{{classname}}

All URIs are relative to */platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAusJsonRequest**](AccountUpdaterApi.md#AddAusJsonRequest) | **Post** /aus/add | Submit Account Updater Request
[**BrowseAusRequestsDetail**](AccountUpdaterApi.md#BrowseAusRequestsDetail) | **Get** /aus/request/detail | Get Detailed List of all Account Updater Results
[**BrowseAusRequestsSummary**](AccountUpdaterApi.md#BrowseAusRequestsSummary) | **Get** /aus/request/summary | Get summary of all Account Updater Results
[**CancelAusRequest**](AccountUpdaterApi.md#CancelAusRequest) | **Post** /aus/request/cancel/{request_id} | Cancel Account Updater Request
[**CreateAusRequest**](AccountUpdaterApi.md#CreateAusRequest) | **Post** /aus/request/create | Create an Account Updater Request using Card IDs
[**GetAusDetail**](AccountUpdaterApi.md#GetAusDetail) | **Get** /aus/request/detail/{request_id} | Get detailed list of an Account Updater Request
[**GetAusResponse**](AccountUpdaterApi.md#GetAusResponse) | **Get** /aus/{requestId} | Get Account Updater Results (Full Card Number Requests)
[**GetAusSummary**](AccountUpdaterApi.md#GetAusSummary) | **Get** /aus/request/summary/{request_id} | Get summary of an Account Updater Request
[**PutOnAusHold**](AccountUpdaterApi.md#PutOnAusHold) | **Post** /aus/hold | Put on Account Updater Harvest Hold
[**ReleaseFromAusHold**](AccountUpdaterApi.md#ReleaseFromAusHold) | **Post** /aus/release | Release from Account Updater Harvest Hold

# **AddAusJsonRequest**
> AusRequestResponse AddAusJsonRequest(ctx, body)
Submit Account Updater Request

Submit an Account Updater request using full card number and expiration date.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AusRequest**](AusRequest.md)| aus request | 

### Return type

[**AusRequestResponse**](AusRequestResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BrowseAusRequestsDetail**
> AusListDetailResponse BrowseAusRequestsDetail(ctx, optional)
Get Detailed List of all Account Updater Results

Browses a paginated list of all cards sent to the Account Updater Engine for harvesting and the responses if received. You can use filters to limit the results returned. Optional query parameters determine size and sort order of the returned array. Use response_code to filter out only the cards that have been updated. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountUpdaterApiBrowseAusRequestsDetailOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountUpdaterApiBrowseAusRequestsDetailOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| The number of records in the result. | [default to 10]
 **orderOn** | **optional.String**| The field on which the results will be sorted on. Refer to the response model for available fields. | [default to request_date]
 **orderBy** | **optional.String**| Ascending or Descending Sort order of the result. Possible values are: asc (Ascending sort order), desc (Descending sort order) | [default to desc]
 **page** | **optional.Int32**| Zero-based page number, use this to choose a page when there are more results than the count parameter. | [default to 0]
 **filter** | [**optional.Interface of []string**](string.md)| Results can be filtered by custom filter criteria. Refer to [Filter](/api/reference#filters) to use the filter parameter. | 

### Return type

[**AusListDetailResponse**](AusListDetailResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BrowseAusRequestsSummary**
> AusListSummaryResponse BrowseAusRequestsSummary(ctx, optional)
Get summary of all Account Updater Results

Browses a paginated summary of Account Updater results. You can use filters to limit the results returned. Optional query parameters determine size and sort order of the returned array.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountUpdaterApiBrowseAusRequestsSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountUpdaterApiBrowseAusRequestsSummaryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| The number of records in the result. | [default to 10]
 **orderOn** | **optional.String**| The field on which the results will be sorted on. Refer to the response model for available fields. | [default to request_date]
 **orderBy** | **optional.String**| Ascending or Descending Sort order of the result. Possible values are: asc (Ascending sort order), desc (Descending sort order) | [default to desc]
 **page** | **optional.Int32**| Zero-based page number, use this to choose a page when there are more results than the count parameter. | [default to 0]
 **filter** | [**optional.Interface of []string**](string.md)| Results can be filtered by custom filter criteria. Refer to [Filter](/api/reference#filters) to use the filter parameter. | 

### Return type

[**AusListSummaryResponse**](AusListSummaryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelAusRequest**
> AusCancelResponse CancelAusRequest(ctx, requestId)
Cancel Account Updater Request

Cancel a previously submitted Account Updater request. Only Account Updater requests in INPROGRESS state can be canceled. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **int64**|  | 

### Return type

[**AusCancelResponse**](AusCancelResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAusRequest**
> CreateAusResponse CreateAusRequest(ctx, body)
Create an Account Updater Request using Card IDs

Submits an Account Updater request to harvest specific card IDs stored in our card vault. A unique request_id is genereated by the API, use the request_id to query the status of your request. You can subscribe to the account_updater_request_created webhook event to be notified when the request creation is complete. Alternatively you can also query the 'Get Summary of an Account Updater Request' API using the request_id to check the status of your request at any time. Upto a maximum of 100 card IDs can be harvested in a single request. You can submit multiple Account Updater requests in a single day. It typically takes 3-4 business days to receive a response from the issuers on card updates. You can subscribe to the account_updater_request_completed webhook event to be notified when the response is received for a request. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateAccountUpdaterRequest**](CreateAccountUpdaterRequest.md)| Account Updater Request | 

### Return type

[**CreateAusResponse**](CreateAusResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAusDetail**
> AusDetailResponse GetAusDetail(ctx, requestId, optional)
Get detailed list of an Account Updater Request

Browses a paginated list of all cards in the specified request id and the responses received from the Account Updater engine. You can use filters to limit the results returned. Optional query parameters determine size and sort order of the returned array. Use response_code to filter out only the cards that have an update. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **int64**|  | 
 **optional** | ***AccountUpdaterApiGetAusDetailOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountUpdaterApiGetAusDetailOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| The number of records in the result. | [default to 10]
 **orderOn** | **optional.String**| The field on which the results will be sorted on. Refer to the response model for available fields. | [default to request_date]
 **orderBy** | **optional.String**| Ascending or Descending Sort order of the result. Possible values are: asc (Ascending sort order), desc (Descending sort order) | [default to desc]
 **page** | **optional.Int32**| Zero-based page number, use this to choose a page when there are more results than the count parameter. | [default to 0]
 **filter** | [**optional.Interface of []string**](string.md)| Results can be filtered by custom filter criteria. Refer to [Filter](/api/reference#filters) to use the filter parameter. | 

### Return type

[**AusDetailResponse**](AusDetailResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAusResponse**
> AusUpdatedResponse GetAusResponse(ctx, requestId)
Get Account Updater Results (Full Card Number Requests)

Get the results of an Account Updater request using the requestId returned in the response of the \"Submit Account Updater\" request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **int64**| Request ID | 

### Return type

[**AusUpdatedResponse**](AusUpdatedResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAusSummary**
> AusSummaryResponse GetAusSummary(ctx, requestId)
Get summary of an Account Updater Request

Get the summary of Account Updater request for a specific requestId. The requestId is returned in the response of the \"Create Account Updater\" request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **int64**|  | 

### Return type

[**AusSummaryResponse**](AusSummaryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutOnAusHold**
> QpApiResponse PutOnAusHold(ctx, body)
Put on Account Updater Harvest Hold

Put a card ID or all payment cards belonging to a customer ID on Account Updater harvest hold. Held cards are not included in your Accout Updater harvest. Account updater service cannot process ACH payments or non U.S. issued cards, hence this request is applicable to only U.S. issued credit cards.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AusHoldRequest**](AusHoldRequest.md)| AusOnHold | 

### Return type

[**QpApiResponse**](QPApiResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReleaseFromAusHold**
> QpApiResponse ReleaseFromAusHold(ctx, body)
Release from Account Updater Harvest Hold

Release a card ID or all payment methods belonging to a customer ID from Account Updater harvest hold. Account updater service cannot process ACH payments or non U.S. issued cards, hence this request is applicable to only U.S. issued credit cards.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AusReleaseRequest**](AusReleaseRequest.md)| AusRelease | 

### Return type

[**QpApiResponse**](QPApiResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

