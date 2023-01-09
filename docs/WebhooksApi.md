# qualpay{{classname}}

All URIs are relative to */platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddEvent**](WebhooksApi.md#AddEvent) | **Post** /webhook/{webhook_id}/event/{event} | Add an event
[**AddWebhook**](WebhooksApi.md#AddWebhook) | **Post** /webhook | Add webhook
[**BrowseWebhook**](WebhooksApi.md#BrowseWebhook) | **Get** /webhook | Browse webhooks
[**DisableWebhook**](WebhooksApi.md#DisableWebhook) | **Put** /webhook/{webhook_id}/disable | Disable a Webhook
[**EditWebhook**](WebhooksApi.md#EditWebhook) | **Put** /webhook/{webhook_id} | Update webhook
[**EnableWebhook**](WebhooksApi.md#EnableWebhook) | **Put** /webhook/{webhook_id}/enable | Enable a Webhook
[**GetEvents**](WebhooksApi.md#GetEvents) | **Get** /webhook/{webhook_id}/event | Get events
[**GetWebhook**](WebhooksApi.md#GetWebhook) | **Get** /webhook/{webhook_id} | Get webhook
[**RemoveEvent**](WebhooksApi.md#RemoveEvent) | **Delete** /webhook/{webhook_id}/event/{event} | Delete event

# **AddEvent**
> QpApiResponse AddEvent(ctx, webhookId, event, optional)
Add an event

Add an event to a webhook. Refer to <a href=\"/developer/api/reference#guides/webhooks-schema\" target=\"_blank\">Webhook Request Schema</a> for a list of available events. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webhookId** | **int64**| Webhook ID | 
  **event** | **string**| Event | 
 **optional** | ***WebhooksApiAddEventOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhooksApiAddEventOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of QpApiData**](QpApiData.md)| Webhook | 

### Return type

[**QpApiResponse**](QPApiResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddWebhook**
> WebhookResponse AddWebhook(ctx, optional)
Add webhook

Configure a new webhook. Save the webhook_id in the response for use in future requests. Save the generated secret, the secret can be used to validate the webhook. Refer to <a href=\"/developer/api/reference#guides/webhooks\" target=\"_blank\">Webhooks</a> for secret usage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WebhooksApiAddWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhooksApiAddWebhookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Webhook**](Webhook.md)| Webhook | 

### Return type

[**WebhookResponse**](WebhookResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BrowseWebhook**
> WebhooksListResponse BrowseWebhook(ctx, optional)
Browse webhooks

Gets an array of webhook objects. Optional query parameters determines, size and sort order of returned array. Filters can be used to filter results. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WebhooksApiBrowseWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhooksApiBrowseWebhookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| The number of records in the result. | [default to 10]
 **orderOn** | **optional.String**| The field on which the results will be sorted on. Refer to the response model for available fields. | [default to webhook_id]
 **orderBy** | **optional.String**| Ascending or Descending Sort order of the result. Possible values are: asc (Ascending sort order), desc (Descending sort order) | [default to asc]
 **page** | **optional.Int32**| Zero-based page number, use this to choose a page when there are more results than the count parameter. | [default to 0]
 **filter** | [**optional.Interface of []string**](string.md)| Results can be filtered by custom filter criteria. Refer to [Filter](/api/reference#filters) to use the filter parameter. | 

### Return type

[**WebhooksListResponse**](WebhooksListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableWebhook**
> QpApiResponse DisableWebhook(ctx, webhookId, optional)
Disable a Webhook

Disable a webhook. Events will not be triggered on a disabled webhook. When disabled, all active requests for this webhook will be held. If the webhook is enabled before a request expires, we will attempt to post the request again.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webhookId** | **int64**| Webhook ID | 
 **optional** | ***WebhooksApiDisableWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhooksApiDisableWebhookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of QpApiData**](QpApiData.md)| Webhook | 

### Return type

[**QpApiResponse**](QPApiResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditWebhook**
> QpApiResponse EditWebhook(ctx, webhookId, optional)
Update webhook

Manage a webhook. Once created, the webhook node cannot be modified. Only the fields to be updated can be sent in the request.  When events array is included in the request, all events will be replaced with the events in the request. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webhookId** | **int64**| Webhook ID | 
 **optional** | ***WebhooksApiEditWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhooksApiEditWebhookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Webhook**](Webhook.md)| Webhook | 

### Return type

[**QpApiResponse**](QPApiResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableWebhook**
> QpApiResponse EnableWebhook(ctx, webhookId, optional)
Enable a Webhook

Enable a webhook. Events are triggered and requests are posted only for active webhooks. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webhookId** | **int64**| Webhook ID | 
 **optional** | ***WebhooksApiEnableWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhooksApiEnableWebhookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of QpApiData**](QpApiData.md)| Webhook | 

### Return type

[**QpApiResponse**](QPApiResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEvents**
> WebhookEventResponse GetEvents(ctx, webhookId)
Get events

Get all events for a webhook. Refer to <a href=\"/developer/api/reference#guides/webhooks-schema\" target=\"_blank\">Webhook Request Schema</a> for a list of available events. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webhookId** | **int64**| Webhook ID | 

### Return type

[**WebhookEventResponse**](WebhookEventResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebhook**
> WebhookResponse GetWebhook(ctx, webhookId)
Get webhook

Gets a webhook.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webhookId** | **int64**| Webhook ID | 

### Return type

[**WebhookResponse**](WebhookResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveEvent**
> QpApiResponse RemoveEvent(ctx, webhookId, event)
Delete event

Delete an event from a webhook. Refer to <a href=\"/developer/api/reference#guides/webhooks-schema\" target=\"_blank\">Webhook Request Schema</a> for a list of available events.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webhookId** | **int64**| Webhook ID | 
  **event** | **string**| Event | 

### Return type

[**QpApiResponse**](QPApiResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

