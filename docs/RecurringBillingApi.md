# qualpay{{classname}}

All URIs are relative to */platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPlan**](RecurringBillingApi.md#AddPlan) | **Post** /plan | Add a Recurring Plan
[**AddSubscription**](RecurringBillingApi.md#AddSubscription) | **Post** /subscription | Add a Subscription
[**ArchivePlan**](RecurringBillingApi.md#ArchivePlan) | **Post** /plan/{plan_code}/archive | Archive a Recurring Plan
[**BrowsePlans**](RecurringBillingApi.md#BrowsePlans) | **Get** /plan | Get all Recurring Plans
[**BrowseSubscriptions**](RecurringBillingApi.md#BrowseSubscriptions) | **Get** /subscription | Get all Subscriptions
[**CancelSubscription**](RecurringBillingApi.md#CancelSubscription) | **Post** /subscription/{subscription_id}/cancel | Cancel a Subscription
[**DeletePlan**](RecurringBillingApi.md#DeletePlan) | **Delete** /plan/{plan_id}/delete | Delete a Recurring Plan
[**GetAllSubscriptionTransactions**](RecurringBillingApi.md#GetAllSubscriptionTransactions) | **Get** /subscription/transactions | Get all subscription transactions
[**GetPlan**](RecurringBillingApi.md#GetPlan) | **Get** /plan/{plan_code} | Find Recurring Plan by Plan Code
[**GetSubscription**](RecurringBillingApi.md#GetSubscription) | **Get** /subscription/{subscription_id} | Get Subscription by Subscription ID
[**GetSubscriptionTransactions**](RecurringBillingApi.md#GetSubscriptionTransactions) | **Get** /subscription/transactions/{subscription_id} | Get transactions by Subscription ID
[**PauseSubscription**](RecurringBillingApi.md#PauseSubscription) | **Post** /subscription/{subscription_id}/pause | Pause a Subscription
[**ResumeSubscription**](RecurringBillingApi.md#ResumeSubscription) | **Post** /subscription/{subscription_id}/resume | Resume a Subscription
[**SubscriptionLabelsAdd**](RecurringBillingApi.md#SubscriptionLabelsAdd) | **Post** /subscription/{subscription_id}/labels/add | Add a custom field to a subscription
[**SubscriptionLabelsBrowse**](RecurringBillingApi.md#SubscriptionLabelsBrowse) | **Get** /subscription/{subscription_id}/labels/browse | Get the custom fields for a subscription
[**SubscriptionLabelsDelete**](RecurringBillingApi.md#SubscriptionLabelsDelete) | **Post** /subscription/{subscription_id}/labels/delete | Removes a custom field
[**SubscriptionLabelsEdit**](RecurringBillingApi.md#SubscriptionLabelsEdit) | **Post** /subscription/{subscription_id}/labels/edit | Edits custom field
[**UpdatePlan**](RecurringBillingApi.md#UpdatePlan) | **Put** /plan/{plan_code} | Update a Recurring Plan
[**UpdateSubscription**](RecurringBillingApi.md#UpdateSubscription) | **Put** /subscription/{subscription_id} | Update a Subscription

# **AddPlan**
> RecurringPlanResponse AddPlan(ctx, body)
Add a Recurring Plan

Adds a new Recurring Plan. Save the generated unique plan_id, which is required to delete a plan. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddRecurringPlanRequest**](AddRecurringPlanRequest.md)| Recurring Plan Object that needs to be added | 

### Return type

[**RecurringPlanResponse**](RecurringPlanResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddSubscription**
> SubscriptionResponse AddSubscription(ctx, optional)
Add a Subscription

Creates a new subscription on the specified start date. Returns the subscription_id; save this id to interact with this subscription using the API. When a subscription is added, with a one-time fee, a payment gateway sale request is  made immediately to bill the customer the one-time fee. Check the response in the return model to check the status of the payment gateway request. Note that the subscription remains active even if the payment gateway request for the one-time fee fails. An “off plan” subscription, a subscription without a plan, can be created by excluding the plan_code from your request and sending applicable fields

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RecurringBillingApiAddSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecurringBillingApiAddSubscriptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AddSubscriptionRequest**](AddSubscriptionRequest.md)| Subscription Request | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArchivePlan**
> RecurringPlanResponse ArchivePlan(ctx, planCode, optional)
Archive a Recurring Plan

Archives a Plan. Only active plans can be archived. Note that if there are subscribers to this plan, then all subscriptions belonging to this plan will continue to be active. No updates can be made to an archived plan. New subscribers cannot be added to a archived plan. You can always pull up information on an archived plan from the system. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planCode** | **string**| Plan Code that will be archived | 
 **optional** | ***RecurringBillingApiArchivePlanOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecurringBillingApiArchivePlanOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ArchiveRecurringPlanRequest**](ArchiveRecurringPlanRequest.md)| Plan Name | 

### Return type

[**RecurringPlanResponse**](RecurringPlanResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BrowsePlans**
> RecurringPlanListResponse BrowsePlans(ctx, optional)
Get all Recurring Plans

Gets a paginated list of recurring plans. Optional query parameters determines, size and sort order of returned array

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RecurringBillingApiBrowsePlansOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecurringBillingApiBrowsePlansOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| The number of records in the result. | [default to 10]
 **orderOn** | **optional.String**| The field on which the results will be sorted on. Refer to the response model for available fields. | [default to plan_code]
 **orderBy** | **optional.String**| Ascending or Descending Sort order of the result. Possible values are: asc (Ascending sort order), desc (Descending sort order) | [default to asc]
 **page** | **optional.Int32**| Zero-based page number, use this to choose a page when there are more results than the count parameter. | [default to 0]
 **filter** | [**optional.Interface of []string**](string.md)| Results can be filtered by custom filter criteria. Refer to [Filter](/api/reference#filters) to use the filter parameter. | 
 **merchantId** | **optional.Int64**| Unique ID assigned to a merchant. | [default to 0]

### Return type

[**RecurringPlanListResponse**](RecurringPlanListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BrowseSubscriptions**
> SubscriptionListResponse BrowseSubscriptions(ctx, optional)
Get all Subscriptions

Gets an array of subscription objects. Optional query parameters determines, size and sort order of returned array

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RecurringBillingApiBrowseSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecurringBillingApiBrowseSubscriptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| The number of records in the result. | [default to 10]
 **orderOn** | **optional.String**| The field on which the results will be sorted on. Refer to the response model for available fields. | [default to date_next]
 **orderBy** | **optional.String**| Ascending or Descending Sort order of the result. Possible values are: asc (Ascending sort order), desc (Descending sort order) | [default to desc]
 **page** | **optional.Int32**| Zero-based page number, use this to choose a page when there are more results than the count parameter. | [default to 0]
 **filter** | [**optional.Interface of []string**](string.md)| Results can be filtered by custom filter criteria. Refer to [Filter](/api/reference#filters) to use the filter parameter. | 
 **merchantId** | **optional.Int64**| Unique ID assigned to a merchant. | [default to 0]

### Return type

[**SubscriptionListResponse**](SubscriptionListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelSubscription**
> SubscriptionResponse CancelSubscription(ctx, body, subscriptionId)
Cancel a Subscription

Cancels a subscription. Only active, suspended or paused subscriptions can be cancelled. A cancelled subscription cannot be resumed. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CancelSubscriptionRequest**](CancelSubscriptionRequest.md)| Customer ID | 
  **subscriptionId** | **int64**| Subscription ID | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePlan**
> RecurringPlanResponse DeletePlan(ctx, planId, optional)
Delete a Recurring Plan

Deletes a Plan. Any plan, active or not can be deleted. If there are subscribers to the plan, then all subscriptions related to this plan will be cancelled. A deleted plan cannot be updated, neither can new subscrbers be added to a deleted plan. Even if a plan is deleted, you can query the system to get information about the deleted plan. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int64**| Plan ID that will flagged as deleted | 
 **optional** | ***RecurringBillingApiDeletePlanOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecurringBillingApiDeletePlanOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **merchantId** | **optional.Int64**| Unique ID assigned to a merchant. | [default to 0]

### Return type

[**RecurringPlanResponse**](RecurringPlanResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllSubscriptionTransactions**
> TransactionListResponse GetAllSubscriptionTransactions(ctx, optional)
Get all subscription transactions

Gets all subscription transactions. Optional Parameters will help filter and restrict the result. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RecurringBillingApiGetAllSubscriptionTransactionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecurringBillingApiGetAllSubscriptionTransactionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| The number of records in the result. | [default to 10]
 **orderOn** | **optional.String**| The field on which the results will be sorted on. Refer to the response model for available fields. | [default to tran_time]
 **orderBy** | **optional.String**| Ascending or Descending Sort order of the result. Possible values are: asc (Ascending sort order), desc (Descending sort order) | [default to desc]
 **page** | **optional.Int32**| Zero-based page number, use this to choose a page when there are more results than the count parameter. | [default to 0]
 **filter** | [**optional.Interface of []string**](string.md)| Results can be filtered by custom filter criteria. Refer to [Filter](/api/reference#filters) to use the filter parameter. | 
 **merchantId** | **optional.Int64**| Unique ID assigned to a merchant. | [default to 0]

### Return type

[**TransactionListResponse**](TransactionListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPlan**
> RecurringPlanListResponse GetPlan(ctx, planCode, optional)
Find Recurring Plan by Plan Code

Returns a list of recurring plans for the plan_code. Search result includes active, archived and deleted plans. Optional query parameters determines, size and sort order of returned array

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planCode** | **string**| Plan Code | 
 **optional** | ***RecurringBillingApiGetPlanOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecurringBillingApiGetPlanOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| The number of records in the result. | [default to 10]
 **orderOn** | **optional.String**| The field on which the results will be sorted on. Refer to the response model for available fields. | [default to plan_code]
 **orderBy** | **optional.String**| Ascending or Descending Sort order of the result. Possible values are: asc (Ascending sort order), desc (Descending sort order) | [default to asc]
 **page** | **optional.Int32**| Zero-based page number, use this to choose a page when there are more results than the count parameter. | [default to 0]
 **filter** | [**optional.Interface of []string**](string.md)| Results can be filtered by custom filter criteria. Refer to [Filter](/api/reference#filters) to use the filter parameter. | 
 **merchantId** | **optional.Int64**| Unique ID assigned to a merchant. | [default to 0]

### Return type

[**RecurringPlanListResponse**](RecurringPlanListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubscription**
> SubscriptionResponse GetSubscription(ctx, subscriptionId, optional)
Get Subscription by Subscription ID

Gets details of a subscription.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionId** | **int64**| Subscription ID | 
 **optional** | ***RecurringBillingApiGetSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecurringBillingApiGetSubscriptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **merchantId** | **optional.Int64**| Unique ID assigned to a merchant. | [default to 0]

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubscriptionTransactions**
> TransactionListResponse GetSubscriptionTransactions(ctx, subscriptionId, optional)
Get transactions by Subscription ID

Gets all transactions for a subscription. Optional parameters will help filter and restrict the result. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionId** | **int64**| Subscription ID | 
 **optional** | ***RecurringBillingApiGetSubscriptionTransactionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecurringBillingApiGetSubscriptionTransactionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| The number of records in the result. | [default to 10]
 **orderOn** | **optional.String**| The field on which the results will be sorted on. Refer to the response model for available fields. | [default to tran_time]
 **orderBy** | **optional.String**| Ascending or Descending Sort order of the result. Possible values are: asc (Ascending sort order), desc (Descending sort order) | [default to desc]
 **page** | **optional.Int32**| Zero-based page number, use this to choose a page when there are more results than the count parameter. | [default to 0]
 **filter** | [**optional.Interface of []string**](string.md)| Results can be filtered by custom filter criteria. Refer to [Filter](/api/reference#filters) to use the filter parameter. | 
 **merchantId** | **optional.Int64**| Unique ID assigned to a merchant. | [default to 0]

### Return type

[**TransactionListResponse**](TransactionListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PauseSubscription**
> SubscriptionResponse PauseSubscription(ctx, body, subscriptionId)
Pause a Subscription

Pauses an active subscription. Recurring payments will be skipped when a subscription is paused. Only active subscriptions can be paused. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PauseSubscriptionRequest**](PauseSubscriptionRequest.md)| Customer ID | 
  **subscriptionId** | **int64**|  | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResumeSubscription**
> SubscriptionResponse ResumeSubscription(ctx, body, subscriptionId)
Resume a Subscription

Resumes a suspended or paused subscription. When a suspended subscription is resumed, the subscription engine will initiate all the missed subscription transactions. When a paused subscription is resumed, all missed payments are skipped. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ResumeSubscriptionRequest**](ResumeSubscriptionRequest.md)| Customer ID | 
  **subscriptionId** | **int64**|  | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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
 **optional** | ***RecurringBillingApiSubscriptionLabelsBrowseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecurringBillingApiSubscriptionLabelsBrowseOpts struct
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

# **UpdatePlan**
> RecurringPlanResponse UpdatePlan(ctx, body, planCode)
Update a Recurring Plan

Update an active recurring plan. Only the fields sent in the request body will be updated. Only an active plan can be updated. If there are subscribers to this plan, then this plan will be archived and a copy of the plan with a new plan_id will be generated. All updates will be made on the new plan. Save the new plan_id to manage a plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateRecurringPlanRequest**](UpdateRecurringPlanRequest.md)| Recurring Plan Object. Send only the fields that require an update. Read only fields will be ignored | 
  **planCode** | **string**| Plan Code of the plan that will be updated | 

### Return type

[**RecurringPlanResponse**](RecurringPlanResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSubscription**
> SubscriptionResponse UpdateSubscription(ctx, body, subscriptionId)
Update a Subscription

Updates the start date of an existing subscription. Only subscriptions that has not yet started can be updated. Only start date can be updated. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateSubscriptionRequest**](UpdateSubscriptionRequest.md)| Subscription Request | 
  **subscriptionId** | **int64**|  | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

