# qualpay{{classname}}

All URIs are relative to */platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptDispute**](DisputeManagementApi.md#AcceptDispute) | **Get** /dispute/{recId}/accept | Accept Dispute
[**CreateDispute**](DisputeManagementApi.md#CreateDispute) | **Get** /dispute/{merchantId}/create | Create Dispute Data
[**GetCreditedTran**](DisputeManagementApi.md#GetCreditedTran) | **Get** /dispute/{recId}/creditedtrans | Get Credited Transaction Detail
[**GetDisputeResponse**](DisputeManagementApi.md#GetDisputeResponse) | **Get** /dispute/{recId}/response | Get Submitted Dispute Response
[**GetDisputes**](DisputeManagementApi.md#GetDisputes) | **Get** /dispute/browse | Get Disputes
[**GetNonDisputedTran**](DisputeManagementApi.md#GetNonDisputedTran) | **Get** /dispute/{recId}/nondisputedtrans | Get Non Disputed Transaction 
[**ResetDispute**](DisputeManagementApi.md#ResetDispute) | **Get** /dispute/{recId}/reset | Reset Dispute Data
[**SubmitDisputesResponse**](DisputeManagementApi.md#SubmitDisputesResponse) | **Post** /dispute/{recId}/respond | Submit Dispute Response

# **AcceptDispute**
> QpApiResponse AcceptDispute(ctx, recId)
Accept Dispute

Accept first time Chargeback and Pre-Arbitration Withdraw from Arbitration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recId** | **int64**| Control Number | 

### Return type

[**QpApiResponse**](QPApiResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDispute**
> QpApiListResponse CreateDispute(ctx, merchantId, optional)
Create Dispute Data

For testing purposes, create a dispute for a specific reason code or for all the reason codes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **merchantId** | **int64**| Merchant ID | 
 **optional** | ***DisputeManagementApiCreateDisputeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DisputeManagementApiCreateDisputeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reasonCode** | **optional.String**| Reason Code | 

### Return type

[**QpApiListResponse**](QPApiListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCreditedTran**
> CorrespondingTransactionResponse GetCreditedTran(ctx, recId)
Get Credited Transaction Detail

Request the credited transactions associated with dispute.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recId** | **int64**| Control Number | 

### Return type

[**CorrespondingTransactionResponse**](CorrespondingTransactionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDisputeResponse**
> DisputeResponse GetDisputeResponse(ctx, recId)
Get Submitted Dispute Response

Get previously submitted dispute response.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recId** | **int64**| Control Number | 

### Return type

[**DisputeResponse**](DisputeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDisputes**
> DisputeDetail GetDisputes(ctx, optional)
Get Disputes

Request all disputes and their detail associated with a vendor or a node.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DisputeManagementApiGetDisputesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DisputeManagementApiGetDisputesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| The number of records in the result. | [default to 10]
 **orderOn** | **optional.String**| The field on which the results will be sorted on. Refer to the response model for available fields. | [default to date_payment]
 **orderBy** | **optional.String**| Ascending or Descending Sort order of the result. Possible values are: asc (Ascending sort order), desc (Descending sort order) | [default to desc]
 **page** | **optional.Int32**| Zero-based page number, use this to choose a page when there are more results than the count parameter. | [default to 0]
 **filter** | [**optional.Interface of []string**](string.md)| Results can be filtered by custom filter criteria. Refer to [Filter](/api/reference#filters) to use the filter parameter. | 

### Return type

[**DisputeDetail**](DisputeDetail.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNonDisputedTran**
> CorrespondingTransactionResponse GetNonDisputedTran(ctx, recId)
Get Non Disputed Transaction 

Request the non-disputed transactions associated with with dispute.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recId** | **int64**| Control Number | 

### Return type

[**CorrespondingTransactionResponse**](CorrespondingTransactionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetDispute**
> QpApiResponse ResetDispute(ctx, recId)
Reset Dispute Data

For non production enviornment, request to reset a dispute case to new status.For Production, request to reset dispute to new status till a dispute is in Qualpay Review (Status=Q).Reset If need to submit different response or want to cancel a rebuttal submitted in error

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recId** | **int64**| Control Number | 

### Return type

[**QpApiResponse**](QPApiResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubmitDisputesResponse**
> QpApiResponse SubmitDisputesResponse(ctx, file, disputeResponse, recId)
Submit Dispute Response

Submit dispute responses with supporting documentation. Response options are dynamic and are based on the reason code.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **file** | ***os.File*****os.File**|  | 
  **disputeResponse** | **string**|  | 
  **recId** | **int64**| Control Number | 

### Return type

[**QpApiResponse**](QPApiResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

