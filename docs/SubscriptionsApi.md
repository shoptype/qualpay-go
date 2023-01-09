# qualpay{{classname}}

All URIs are relative to */platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionLabelsAdd**](SubscriptionsApi.md#SubscriptionLabelsAdd) | **Post** /subscription/{subscription_id}/labels/add | Add a custom field to a subscription
[**SubscriptionLabelsBrowse**](SubscriptionsApi.md#SubscriptionLabelsBrowse) | **Get** /subscription/{subscription_id}/labels/browse | Get the custom fields for a subscription
[**SubscriptionLabelsDelete**](SubscriptionsApi.md#SubscriptionLabelsDelete) | **Post** /subscription/{subscription_id}/labels/delete | Removes a custom field
[**SubscriptionLabelsEdit**](SubscriptionsApi.md#SubscriptionLabelsEdit) | **Post** /subscription/{subscription_id}/labels/edit | Edits custom field

# **SubscriptionLabelsAdd**
> SubscriptionLabelsAddEditResponse SubscriptionLabelsAdd(ctx, body)
Add a custom field to a subscription

Adds a new custom field to a subscription.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SubscriptionLabel**](SubscriptionLabel.md)| Custom field data to be added | 

### Return type

[**SubscriptionLabelsAddEditResponse**](SubscriptionLabelsAddEditResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubscriptionLabelsBrowse**
> SubscriptionLabelsListResponse SubscriptionLabelsBrowse(ctx, subscriptionId, optional)
Get the custom fields for a subscription

Gets all of the custom fields assigned to a subscription. Optional parameters will help filter and restrict the result. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionId** | **int64**| Subscription ID | 
 **optional** | ***SubscriptionsApiSubscriptionLabelsBrowseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SubscriptionsApiSubscriptionLabelsBrowseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| The number of records in the result. | [default to 10]
 **orderOn** | **optional.String**| The field on which the results will be sorted on. Refer to the response model for available fields. | [default to tran_time]
 **orderBy** | **optional.String**| Ascending or Descending Sort order of the result. Possible values are: asc (Ascending sort order), desc (Descending sort order) | [default to desc]
 **page** | **optional.Int32**| Zero-based page number, use this to choose a page when there are more results than the count parameter. | [default to 0]
 **filter** | [**optional.Interface of []string**](string.md)| Results can be filtered by custom filter criteria. Refer to [Filter](/api/reference#filters) to use the filter parameter. | 

### Return type

[**SubscriptionLabelsListResponse**](SubscriptionLabelsListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubscriptionLabelsDelete**
> QpApiResponse SubscriptionLabelsDelete(ctx, body)
Removes a custom field

Removes a custom field assigned to a subscription.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SubscriptionLabel**](SubscriptionLabel.md)| Custom field data to be removed | 

### Return type

[**QpApiResponse**](QPApiResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubscriptionLabelsEdit**
> SubscriptionLabelsAddEditResponse SubscriptionLabelsEdit(ctx, body)
Edits custom field

Edits a custom field assigned to a subscription.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SubscriptionLabel**](SubscriptionLabel.md)| Custom field data to be edited | 

### Return type

[**SubscriptionLabelsAddEditResponse**](SubscriptionLabelsAddEditResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

