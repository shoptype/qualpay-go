# qualpay{{classname}}

All URIs are relative to */platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BrowseAusRequests**](ReportingApi.md#BrowseAusRequests) | **Get** /reporting/aus/detail | Account Updater Detail Report
[**BrowseAusSummary**](ReportingApi.md#BrowseAusSummary) | **Get** /reporting/aus/summary | Account Updater Summary Report
[**BrowseBatches**](ReportingApi.md#BrowseBatches) | **Get** /reporting/batches | Batch Report
[**BrowseDeposits**](ReportingApi.md#BrowseDeposits) | **Get** /reporting/deposits | Deposit Report
[**BrowseDisputes**](ReportingApi.md#BrowseDisputes) | **Get** /reporting/disputes | Disputes Report
[**BrowseStatementDisputeAdjustmentData**](ReportingApi.md#BrowseStatementDisputeAdjustmentData) | **Get** /reporting/statement-report/disputeAdjustment | Statement Dispute Adjustments Data
[**BrowseStatementFeeData**](ReportingApi.md#BrowseStatementFeeData) | **Get** /reporting/statement-report/fees | Statement Fee Data
[**BrowseStatementPlanData**](ReportingApi.md#BrowseStatementPlanData) | **Get** /reporting/statement-report/planType | Statement Plan Type Data
[**BrowseStatementReserveData**](ReportingApi.md#BrowseStatementReserveData) | **Get** /reporting/statement-report/reserve | Statement Reserve Data
[**BrowseStatementSettlementData**](ReportingApi.md#BrowseStatementSettlementData) | **Get** /reporting/statement-report/settlement | Statement Settlement Data
[**BrowseTrans**](ReportingApi.md#BrowseTrans) | **Get** /reporting/transactions | Settled Transaction Report
[**GetTransactionByPgId**](ReportingApi.md#GetTransactionByPgId) | **Get** /reporting/transactions/bypgid/{pg_id} | Get transaction by PG ID
[**GetTransactionRequests**](ReportingApi.md#GetTransactionRequests) | **Get** /reporting/transaction-requests | Transaction Report

# **BrowseAusRequests**
> AccountUpdaterResponse BrowseAusRequests(ctx, optional)
Account Updater Detail Report

Browses a paginated list of all cards sent to the Account Updater Engine for harvesting and the responses if received. You can use filters to limit the results returned. Optional query parameters determine size and sort order of the returned array. Use response_code to filter out only the cards that have been updated. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReportingApiBrowseAusRequestsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportingApiBrowseAusRequestsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| The number of records in the result. | [default to 10]
 **orderOn** | **optional.String**| The field on which the results will be sorted on. Refer to the response model for available fields. | 
 **orderBy** | **optional.String**| Ascending or Descending Sort order of the result. Possible values are: asc (Ascending sort order), desc (Descending sort order) | 
 **page** | **optional.Int32**| Zero-based page number, use this to choose a page when there are more results than the count parameter. | [default to 0]
 **filter** | [**optional.Interface of []string**](string.md)| Results can be filtered by custom filter criteria. Refer to [Filter](/api/reference#filters) to use the filter parameter. | 

### Return type

[**AccountUpdaterResponse**](AccountUpdaterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BrowseAusSummary**
> AccountUpdaterSummaryResponse BrowseAusSummary(ctx, optional)
Account Updater Summary Report

Browses a paginated summary of Account Updater results. You can use filters to limit the results returned. Optional query parameters determine size and sort order of the returned array.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReportingApiBrowseAusSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportingApiBrowseAusSummaryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| The number of records in the result. | [default to 10]
 **orderOn** | **optional.String**| The field on which the results will be sorted on. Refer to the response model for available fields. | 
 **orderBy** | **optional.String**| Ascending or Descending Sort order of the result. Possible values are: asc (Ascending sort order), desc (Descending sort order) | 
 **page** | **optional.Int32**| Zero-based page number, use this to choose a page when there are more results than the count parameter. | [default to 0]
 **filter** | [**optional.Interface of []string**](string.md)| Results can be filtered by custom filter criteria. Refer to [Filter](/api/reference#filters) to use the filter parameter. | 

### Return type

[**AccountUpdaterSummaryResponse**](AccountUpdaterSummaryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BrowseBatches**
> BatchResponse BrowseBatches(ctx, optional)
Batch Report

Browses a paginated list of merchant batches. These are whole batches which have been settled from the merchant's POS device, software, or gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReportingApiBrowseBatchesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportingApiBrowseBatchesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| The number of records in the result. | [default to 10]
 **orderOn** | **optional.String**| The field on which the results will be sorted on. Refer to the response model for available fields. | 
 **orderBy** | **optional.String**| Ascending or Descending Sort order of the result. Possible values are: asc (Ascending sort order), desc (Descending sort order) | 
 **page** | **optional.Int32**| Zero-based page number, use this to choose a page when there are more results than the count parameter. | [default to 0]
 **filter** | [**optional.Interface of []string**](string.md)| Results can be filtered by custom filter criteria. Refer to [Filter](/api/reference#filters) to use the filter parameter. | 

### Return type

[**BatchResponse**](BatchResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BrowseDeposits**
> DepositResponse BrowseDeposits(ctx, optional)
Deposit Report

Browses a paginated list of bank deposits and withdrawls.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReportingApiBrowseDepositsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportingApiBrowseDepositsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| The number of records in the result. | [default to 10]
 **orderOn** | **optional.String**| The field on which the results will be sorted on. Refer to the response model for available fields. | 
 **orderBy** | **optional.String**| Ascending or Descending Sort order of the result. Possible values are: asc (Ascending sort order), desc (Descending sort order) | 
 **page** | **optional.Int32**| Zero-based page number, use this to choose a page when there are more results than the count parameter. | [default to 0]
 **filter** | [**optional.Interface of []string**](string.md)| Results can be filtered by custom filter criteria. Refer to [Filter](/api/reference#filters) to use the filter parameter. | 

### Return type

[**DepositResponse**](DepositResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BrowseDisputes**
> DisputeResponse BrowseDisputes(ctx, optional)
Disputes Report

Browses a paginated list of disputes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReportingApiBrowseDisputesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportingApiBrowseDisputesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| The number of records in the result. | [default to 10]
 **orderOn** | **optional.String**| The field on which the results will be sorted on. Refer to the response model for available fields. | 
 **orderBy** | **optional.String**| Ascending or Descending Sort order of the result. Possible values are: asc (Ascending sort order), desc (Descending sort order) | 
 **page** | **optional.Int32**| Zero-based page number, use this to choose a page when there are more results than the count parameter. | [default to 0]
 **filter** | [**optional.Interface of []string**](string.md)| Results can be filtered by custom filter criteria. Refer to [Filter](/api/reference#filters) to use the filter parameter. | 

### Return type

[**DisputeResponse**](DisputeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BrowseStatementDisputeAdjustmentData**
> StatementDisputeResponse BrowseStatementDisputeAdjustmentData(ctx, optional)
Statement Dispute Adjustments Data

Browses a paginated list of Statement Disputes Adjustment Data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReportingApiBrowseStatementDisputeAdjustmentDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportingApiBrowseStatementDisputeAdjustmentDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| The number of records in the result. | [default to 10]
 **orderOn** | **optional.String**| The field on which the results will be sorted on. Refer to the response model for available fields. | [default to tran_time]
 **orderBy** | **optional.String**| Ascending or Descending Sort order of the result. Possible values are: asc (Ascending sort order), desc (Descending sort order) | [default to desc]
 **page** | **optional.Int32**| Zero-based page number, use this to choose a page when there are more results than the count parameter. | [default to 0]
 **filter** | [**optional.Interface of []string**](string.md)| Results can be filtered by custom filter criteria. Refer to [Filter](/api/reference#filters) to use the filter parameter. | 

### Return type

[**StatementDisputeResponse**](StatementDisputeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BrowseStatementFeeData**
> StatementFeeResponse BrowseStatementFeeData(ctx, optional)
Statement Fee Data

Browses a paginated list of Statement Fee Data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReportingApiBrowseStatementFeeDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportingApiBrowseStatementFeeDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| The number of records in the result. | [default to 10]
 **orderOn** | **optional.String**| The field on which the results will be sorted on. Refer to the response model for available fields. | [default to tran_time]
 **orderBy** | **optional.String**| Ascending or Descending Sort order of the result. Possible values are: asc (Ascending sort order), desc (Descending sort order) | [default to desc]
 **page** | **optional.Int32**| Zero-based page number, use this to choose a page when there are more results than the count parameter. | [default to 0]
 **filter** | [**optional.Interface of []string**](string.md)| Results can be filtered by custom filter criteria. Refer to [Filter](/api/reference#filters) to use the filter parameter. | 

### Return type

[**StatementFeeResponse**](StatementFeeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BrowseStatementPlanData**
> StatementPlanResponse BrowseStatementPlanData(ctx, optional)
Statement Plan Type Data

Browses a paginated list of Statement Plan Type Summary

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReportingApiBrowseStatementPlanDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportingApiBrowseStatementPlanDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| The number of records in the result. | [default to 10]
 **orderOn** | **optional.String**| The field on which the results will be sorted on. Refer to the response model for available fields. | [default to tran_time]
 **orderBy** | **optional.String**| Ascending or Descending Sort order of the result. Possible values are: asc (Ascending sort order), desc (Descending sort order) | [default to desc]
 **page** | **optional.Int32**| Zero-based page number, use this to choose a page when there are more results than the count parameter. | [default to 0]
 **filter** | [**optional.Interface of []string**](string.md)| Results can be filtered by custom filter criteria. Refer to [Filter](/api/reference#filters) to use the filter parameter. | 

### Return type

[**StatementPlanResponse**](StatementPlanResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BrowseStatementReserveData**
> StatementReserveResponse BrowseStatementReserveData(ctx, optional)
Statement Reserve Data

Browses a paginated list of Statement Reserve Data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReportingApiBrowseStatementReserveDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportingApiBrowseStatementReserveDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| The number of records in the result. | [default to 10]
 **orderOn** | **optional.String**| The field on which the results will be sorted on. Refer to the response model for available fields. | [default to tran_time]
 **orderBy** | **optional.String**| Ascending or Descending Sort order of the result. Possible values are: asc (Ascending sort order), desc (Descending sort order) | [default to desc]
 **page** | **optional.Int32**| Zero-based page number, use this to choose a page when there are more results than the count parameter. | [default to 0]
 **filter** | [**optional.Interface of []string**](string.md)| Results can be filtered by custom filter criteria. Refer to [Filter](/api/reference#filters) to use the filter parameter. | 

### Return type

[**StatementReserveResponse**](StatementReserveResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BrowseStatementSettlementData**
> StatementSettlementResponse BrowseStatementSettlementData(ctx, optional)
Statement Settlement Data

Browses a paginated list of Statement Settlement Data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReportingApiBrowseStatementSettlementDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportingApiBrowseStatementSettlementDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| The number of records in the result. | [default to 10]
 **orderOn** | **optional.String**| The field on which the results will be sorted on. Refer to the response model for available fields. | [default to tran_time]
 **orderBy** | **optional.String**| Ascending or Descending Sort order of the result. Possible values are: asc (Ascending sort order), desc (Descending sort order) | [default to desc]
 **page** | **optional.Int32**| Zero-based page number, use this to choose a page when there are more results than the count parameter. | [default to 0]
 **filter** | [**optional.Interface of []string**](string.md)| Results can be filtered by custom filter criteria. Refer to [Filter](/api/reference#filters) to use the filter parameter. | 

### Return type

[**StatementSettlementResponse**](StatementSettlementResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BrowseTrans**
> SettledTransactionResponse BrowseTrans(ctx, optional)
Settled Transaction Report

Browses a paginated list of settled transactions. Some additional text about deposits and stuff.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReportingApiBrowseTransOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportingApiBrowseTransOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| The number of records in the result. | [default to 10]
 **orderOn** | **optional.String**| The field on which the results will be sorted on. Refer to the response model for available fields. | 
 **orderBy** | **optional.String**| Ascending or Descending Sort order of the result. Possible values are: asc (Ascending sort order), desc (Descending sort order) | 
 **page** | **optional.Int32**| Zero-based page number, use this to choose a page when there are more results than the count parameter. | [default to 0]
 **filter** | [**optional.Interface of []string**](string.md)| Results can be filtered by custom filter criteria. Refer to [Filter](/api/reference#filters) to use the filter parameter. | 

### Return type

[**SettledTransactionResponse**](SettledTransactionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransactionByPgId**
> TransactionResponse GetTransactionByPgId(ctx, pgId)
Get transaction by PG ID

Gets a payment gateway transaction by Payment Gateway ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pgId** | **string**| PG ID | 

### Return type

[**TransactionResponse**](TransactionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransactionRequests**
> TransactionRequestResponse GetTransactionRequests(ctx, optional)
Transaction Report

Browses a paginated list of transaction requests

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReportingApiGetTransactionRequestsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportingApiGetTransactionRequestsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| The number of records in the result. | [default to 10]
 **orderOn** | **optional.String**| The field on which the results will be sorted on. Refer to the response model for available fields. | [default to tran_time]
 **orderBy** | **optional.String**| Ascending or Descending Sort order of the result. Possible values are: asc (Ascending sort order), desc (Descending sort order) | [default to desc]
 **page** | **optional.Int32**| Zero-based page number, use this to choose a page when there are more results than the count parameter. | [default to 0]
 **filter** | [**optional.Interface of []string**](string.md)| Results can be filtered by custom filter criteria. Refer to [Filter](/api/reference#filters) to use the filter parameter. | 

### Return type

[**TransactionRequestResponse**](TransactionRequestResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

