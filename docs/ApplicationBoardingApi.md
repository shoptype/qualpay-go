# qualpay{{classname}}

All URIs are relative to */platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddApp**](ApplicationBoardingApi.md#AddApp) | **Post** /application/add | Create Application
[**BrowseApps**](ApplicationBoardingApi.md#BrowseApps) | **Get** /application/browse | Browse applications
[**BrowseSalesReps**](ApplicationBoardingApi.md#BrowseSalesReps) | **Get** /application/sales/browse | Browse Sales Reps
[**EmailApp**](ApplicationBoardingApi.md#EmailApp) | **Post** /application/{appId}/email | Email Application
[**GetApp**](ApplicationBoardingApi.md#GetApp) | **Get** /application/{appId}/get | Get Application
[**SaveData**](ApplicationBoardingApi.md#SaveData) | **Post** /application/{appId}/data | Save Data
[**UploadFile**](ApplicationBoardingApi.md#UploadFile) | **Post** /application/{appId}/upload | Upload Document
[**ValidateApp**](ApplicationBoardingApi.md#ValidateApp) | **Get** /application/{appId}/validate | Validate Application

# **AddApp**
> InlineResponse2002 AddApp(ctx, optional)
Create Application

Creates a new merchant application, returning the new application's ID as well as the pricing options available.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplicationBoardingApiAddAppOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationBoardingApiAddAppOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApplicationAddBody**](ApplicationAddBody.md)| Application Request | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BrowseApps**
> InlineResponse200 BrowseApps(ctx, optional)
Browse applications

Retrieves an array of Applications. Optional query parameters determine the page size and sorting of the data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplicationBoardingApiBrowseAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationBoardingApiBrowseAppsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| The number of records in the result. | [default to 10]
 **orderOn** | **optional.String**| The field on which the results will be sorted on. Refer to the response model for available fields. | 
 **orderBy** | **optional.String**| Ascending or Descending Sort order of the result. Possible values are: asc (Ascending sort order), desc (Descending sort order) | 
 **page** | **optional.Int32**| Zero-based page number, use this to choose a page when there are more results than the count parameter. | [default to 0]
 **filter** | [**optional.Interface of []string**](string.md)| Results can be filtered by custom filter criteria. Refer to [Filter](/api/reference#filters) to use the filter parameter. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BrowseSalesReps**
> InlineResponse2001 BrowseSalesReps(ctx, optional)
Browse Sales Reps

Retrieves an array of Users who can be used during the creation of an application to identify the sales representitive. Optional query parameters determine the page size and sorting of the data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplicationBoardingApiBrowseSalesRepsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationBoardingApiBrowseSalesRepsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| The number of records in the result. | [default to 10]
 **orderOn** | **optional.String**| The field on which the results will be sorted on. Refer to the response model for available fields. | 
 **orderBy** | **optional.String**| Ascending or Descending Sort order of the result. Possible values are: asc (Ascending sort order), desc (Descending sort order) | 
 **page** | **optional.Int32**| Zero-based page number, use this to choose a page when there are more results than the count parameter. | [default to 0]
 **filter** | **optional.String**| Results can be filtered by custom filter criteria. Refer to [Filter](/api/reference#filters) to use the filter parameter. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EmailApp**
> QpApiResponse EmailApp(ctx, appId, optional)
Email Application

Emails an application to a new or existing owner. This request requires either   * `owner_id`, which matches an existing owner in the `ownership` field.   * Or an `owner` object consisting of at least an `email`, and both `name.first_name` and `name.last_name`. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int64**| Application ID | 
 **optional** | ***ApplicationBoardingApiEmailAppOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationBoardingApiEmailAppOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AppIdEmailBody**](AppIdEmailBody.md)| Application Request | 

### Return type

[**QpApiResponse**](QPApiResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApp**
> InlineResponse2002 GetApp(ctx, appId)
Get Application

Retrieve an application's details and available pricing.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int64**| Application ID | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SaveData**
> QpApiResponse SaveData(ctx, appId, optional)
Save Data

Saves one or more data fields to a new merchant application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int64**| Application ID | 
 **optional** | ***ApplicationBoardingApiSaveDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationBoardingApiSaveDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AppIdDataBody**](AppIdDataBody.md)| Application Request | 

### Return type

[**QpApiResponse**](QPApiResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadFile**
> QpApiResponse UploadFile(ctx, file, bucket, label, appId)
Upload Document

Uploads a document relevant to this application to the provided bucket.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **file** | ***os.File*****os.File**|  | 
  **bucket** | **string**|  | 
  **label** | **string**|  | 
  **appId** | **int64**| Application ID | 

### Return type

[**QpApiResponse**](QPApiResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateApp**
> QpApiResponse ValidateApp(ctx, appId)
Validate Application

All data fields in an application are validated, including dependencies they may have between eachother. An application can not be submitted until it passes validation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int64**| Application ID | 

### Return type

[**QpApiResponse**](QPApiResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

