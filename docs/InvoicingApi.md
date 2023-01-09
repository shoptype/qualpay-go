# qualpay{{classname}}

All URIs are relative to */platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddInvoicePayment**](InvoicingApi.md#AddInvoicePayment) | **Post** /invoice/{invoice_id}/payments | Add Payment to an Invoice
[**BrowseBouncedInvoices**](InvoicingApi.md#BrowseBouncedInvoices) | **Get** /invoice/bounced | Get Undelivered Invoices
[**BrowseInvoicePayments**](InvoicingApi.md#BrowseInvoicePayments) | **Get** /invoice/payments | Get Invoice Payments
[**BrowseInvoicePaymentsById**](InvoicingApi.md#BrowseInvoicePaymentsById) | **Get** /invoice/{invoice_id}/payments | Get invoice payments by id
[**BrowseInvoices**](InvoicingApi.md#BrowseInvoices) | **Get** /invoice | Get all Invoices
[**CancelInvoice**](InvoicingApi.md#CancelInvoice) | **Delete** /invoice/{invoice_id}/cancel | Cancel an Invoice
[**CopyInvoice**](InvoicingApi.md#CopyInvoice) | **Post** /invoice/{invoice_id}/copy | Copy an Invoice
[**CreateInvoice**](InvoicingApi.md#CreateInvoice) | **Post** /invoice | Create an invoice
[**GetInvoice**](InvoicingApi.md#GetInvoice) | **Get** /invoice/{invoice_id}/detail | Get by Invoice ID
[**RemoveInvoicePayment**](InvoicingApi.md#RemoveInvoicePayment) | **Delete** /invoice/{invoice_id}/payments/{payment_id} | Remove an Invoice Payment
[**ResendInvoice**](InvoicingApi.md#ResendInvoice) | **Post** /invoice/{invoice_id}/resend | Resend an Invoice
[**SendInvoice**](InvoicingApi.md#SendInvoice) | **Post** /invoice/{invoice_id}/send | Send an Invoice
[**UpdateDraftInvoice**](InvoicingApi.md#UpdateDraftInvoice) | **Put** /invoice/{invoice_id}/draft | Update a Draft Invoice
[**UpdateInvoicePayment**](InvoicingApi.md#UpdateInvoicePayment) | **Put** /invoice/{invoice_id}/payments/{payment_id} | Update an Invoice Payment
[**UpdateOutstandingInvoice**](InvoicingApi.md#UpdateOutstandingInvoice) | **Put** /invoice/{invoice_id}/outstanding | Update an Outstanding Invoice

# **AddInvoicePayment**
> InvoicePaymentResponse AddInvoicePayment(ctx, body, invoiceId)
Add Payment to an Invoice

Adds a payment to an invoice. A check or cash payment can be added to a saved or outstanding invoice. Credit card payments cannot be added manually to an invoice. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InvoicePaymentRequest**](InvoicePaymentRequest.md)| Invoice Payment | 
  **invoiceId** | **int64**|  | 

### Return type

[**InvoicePaymentResponse**](InvoicePaymentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BrowseBouncedInvoices**
> InvoiceBouncedResponse BrowseBouncedInvoices(ctx, optional)
Get Undelivered Invoices

Browse all undelivered invoices. Optional query parameters determines, size and sort order of returned array.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InvoicingApiBrowseBouncedInvoicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InvoicingApiBrowseBouncedInvoicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| The number of records in the result. | [default to 10]
 **orderOn** | **optional.String**| The field on which the results will be sorted on. Refer to the response model for available fields. | [default to db_timestamp]
 **orderBy** | **optional.String**| Ascending or Descending Sort order of the result. Possible values are: asc (Ascending sort order), desc (Descending sort order) | [default to desc]
 **page** | **optional.Int32**| Zero-based page number, use this to choose a page when there are more results than the count parameter. | [default to 0]
 **filter** | [**optional.Interface of []string**](string.md)| Results can be filtered by custom filter criteria. Refer to [Filter](/api/reference#filters) to use the filter parameter. | 

### Return type

[**InvoiceBouncedResponse**](InvoiceBouncedResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BrowseInvoicePayments**
> InvoicePaymentListResponse BrowseInvoicePayments(ctx, optional)
Get Invoice Payments

Browse all invoice payments. Optional query parameters determines, size and sort order of returned array.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InvoicingApiBrowseInvoicePaymentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InvoicingApiBrowseInvoicePaymentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| The number of records in the result. | [default to 10]
 **orderOn** | **optional.String**| The field on which the results will be sorted on. Refer to the response model for available fields. | [default to date_payment]
 **orderBy** | **optional.String**| Ascending or Descending Sort order of the result. Possible values are: asc (Ascending sort order), desc (Descending sort order) | [default to desc]
 **page** | **optional.Int32**| Zero-based page number, use this to choose a page when there are more results than the count parameter. | [default to 0]
 **filter** | [**optional.Interface of []string**](string.md)| Results can be filtered by custom filter criteria. Refer to [Filter](/api/reference#filters) to use the filter parameter. | 

### Return type

[**InvoicePaymentListResponse**](InvoicePaymentListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BrowseInvoicePaymentsById**
> InvoicePaymentListResponse BrowseInvoicePaymentsById(ctx, invoiceId, optional)
Get invoice payments by id

Browse all invoice payments made to an invoice. Optional query parameters determines, size and sort order of returned array.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **invoiceId** | **int64**|  | 
 **optional** | ***InvoicingApiBrowseInvoicePaymentsByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InvoicingApiBrowseInvoicePaymentsByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| The number of records in the result. | [default to 10]
 **orderOn** | **optional.String**| The field on which the results will be sorted on. Refer to the response model for available fields. | [default to date_payment]
 **orderBy** | **optional.String**| Ascending or Descending Sort order of the result. Possible values are: asc (Ascending sort order), desc (Descending sort order) | [default to desc]
 **page** | **optional.Int32**| Zero-based page number, use this to choose a page when there are more results than the count parameter. | [default to 0]
 **filter** | [**optional.Interface of []string**](string.md)| Results can be filtered by custom filter criteria. Refer to [Filter](/api/reference#filters) to use the filter parameter. | 

### Return type

[**InvoicePaymentListResponse**](InvoicePaymentListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BrowseInvoices**
> InvoiceListResponse BrowseInvoices(ctx, optional)
Get all Invoices

Gets an array of invoice objects. Optional query parameters determines, size and sort order of returned array.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InvoicingApiBrowseInvoicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InvoicingApiBrowseInvoicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| The number of records in the result. | [default to 10]
 **orderOn** | **optional.String**| The field on which the results will be sorted on. Refer to the response model for available fields. | [default to invoice_number]
 **orderBy** | **optional.String**| Ascending or Descending Sort order of the result. Possible values are: asc (Ascending sort order), desc (Descending sort order) | [default to desc]
 **page** | **optional.Int32**| Zero-based page number, use this to choose a page when there are more results than the count parameter. | [default to 0]
 **filter** | [**optional.Interface of []string**](string.md)| Results can be filtered by custom filter criteria. Refer to [Filter](/api/reference#filters) to use the filter parameter. | 

### Return type

[**InvoiceListResponse**](InvoiceListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelInvoice**
> InvoiceResponse CancelInvoice(ctx, invoiceId)
Cancel an Invoice

Cancels an invoice. A canceled invoice cannot be edited. If your customer clicks on the pay now button in the invoice e-mail an error message will be displayed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **invoiceId** | **int64**|  | 

### Return type

[**InvoiceResponse**](InvoiceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CopyInvoice**
> InvoiceResponse CopyInvoice(ctx, invoiceId, optional)
Copy an Invoice

Makes a copy of an invoice. The invoice date will be set to today's date and due date will be adjusted based on the invoice date and the payment terms. Optionally, include an invoice_number in the POST body to make a copy of the invoice with a different invoice number. Invoice payments from the original invoice will not be copied. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **invoiceId** | **int64**|  | 
 **optional** | ***InvoicingApiCopyInvoiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InvoicingApiCopyInvoiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CopyInvoiceRequest**](CopyInvoiceRequest.md)| Invoice | 

### Return type

[**InvoiceResponse**](InvoiceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateInvoice**
> InvoiceResponse CreateInvoice(ctx, body)
Create an invoice

Creates a draft invoice that you can send later using the Send Invoice API.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateInvoiceRequest**](CreateInvoiceRequest.md)| Invoice | 

### Return type

[**InvoiceResponse**](InvoiceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInvoice**
> InvoiceResponse GetInvoice(ctx, invoiceId)
Get by Invoice ID

Gets an invoice by invoice_id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **invoiceId** | **int64**| Invoice ID | 

### Return type

[**InvoiceResponse**](InvoiceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveInvoicePayment**
> QpApiResponse RemoveInvoicePayment(ctx, invoiceId, paymentId)
Remove an Invoice Payment

Removes an invoice payment. A payment can be removed on a saved or an outstanding invoice. CARD type payments cannot be removed. Payments made via a credit card cannot be removed. Payments can be deleted only from SAVED or OUTSTANDING invoices.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **invoiceId** | **int64**|  | 
  **paymentId** | **int64**|  | 

### Return type

[**QpApiResponse**](QPApiResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResendInvoice**
> QpApiResponse ResendInvoice(ctx, invoiceId, optional)
Resend an Invoice

Resends an invoice to the customer. An outstanding or paid invoice can be resent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **invoiceId** | **int64**|  | 
 **optional** | ***InvoicingApiResendInvoiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InvoicingApiResendInvoiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ResendInvoiceRequest**](ResendInvoiceRequest.md)| Email Addresses | 

### Return type

[**QpApiResponse**](QPApiResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendInvoice**
> InvoiceResponse SendInvoice(ctx, invoiceId, optional)
Send an Invoice

Sends an invoice to the customer. Sending an invoice changes the status of the invoice to outstanding. Once sent, only the from_contact and business_contact of the invoice can be updated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **invoiceId** | **int64**|  | 
 **optional** | ***InvoicingApiSendInvoiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InvoicingApiSendInvoiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SendInvoiceRequest**](SendInvoiceRequest.md)| Email Addresses | 

### Return type

[**InvoiceResponse**](InvoiceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDraftInvoice**
> InvoiceResponse UpdateDraftInvoice(ctx, body, invoiceId)
Update a Draft Invoice

Updates a draft invoice. Only the fields that need updating can be sent in the request body. If updating JSON object fields, the complete JSON should be sent in the request body. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateDraftRequest**](UpdateDraftRequest.md)| Invoice | 
  **invoiceId** | **int64**|  | 

### Return type

[**InvoiceResponse**](InvoiceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateInvoicePayment**
> InvoicePaymentResponse UpdateInvoicePayment(ctx, body, invoiceId, paymentId)
Update an Invoice Payment

Updates an invoice payment. A payment can be updated on a saved or an outstanding invoice. Payments made via credit card using the “Pay Now” button cannot be updated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InvoicePaymentRequest**](InvoicePaymentRequest.md)| Invoice Payment | 
  **invoiceId** | **int64**|  | 
  **paymentId** | **int64**|  | 

### Return type

[**InvoicePaymentResponse**](InvoicePaymentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOutstandingInvoice**
> InvoiceResponse UpdateOutstandingInvoice(ctx, body, invoiceId)
Update an Outstanding Invoice

Updates an outstanding invoice. Only the from_contact and business_contact fields can be updated on an outstanding invoice. Only the fields that need updating can be sent in the request body. If updating JSON object fields, the complete JSON should be sent in the request body. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateOutstandingRequest**](UpdateOutstandingRequest.md)| Invoice | 
  **invoiceId** | **int64**|  | 

### Return type

[**InvoiceResponse**](InvoiceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

