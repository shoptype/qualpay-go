# qualpay{{classname}}

All URIs are relative to */platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Authorization**](PaymentGatewayApi.md#Authorization) | **Post** /auth | Authorize Transaction
[**BatchClose**](PaymentGatewayApi.md#BatchClose) | **Post** /batchClose | Close Batch
[**Capture**](PaymentGatewayApi.md#Capture) | **Post** /capture/{pgIdOrig} | Capture an Authorized Transaction
[**Credit**](PaymentGatewayApi.md#Credit) | **Post** /credit | Issue Credit to Cardholder
[**Expire**](PaymentGatewayApi.md#Expire) | **Post** /expireToken | Expire Token
[**Force**](PaymentGatewayApi.md#Force) | **Post** /force | Force Transaction Approval
[**GetCardTypeInformation**](PaymentGatewayApi.md#GetCardTypeInformation) | **Post** /ardef | Get Card type Information for Visa, Mastercard, and Discover
[**InitiatePayment**](PaymentGatewayApi.md#InitiatePayment) | **Post** /payment | Initiate a payment to a payee
[**Recharge**](PaymentGatewayApi.md#Recharge) | **Post** /recharge/{pgIdOrig} | Recharge Previously Settled Transaction
[**Refund**](PaymentGatewayApi.md#Refund) | **Post** /refund/{pgIdOrig} | Refund Previously Captured Transaction
[**RequestBalance**](PaymentGatewayApi.md#RequestBalance) | **Post** /balance | Request a balance
[**Sale**](PaymentGatewayApi.md#Sale) | **Post** /sale | Sale (Auth + Capture)
[**SendReceipt**](PaymentGatewayApi.md#SendReceipt) | **Post** /emailReceipt/{pgId} | Send Transaction Receipt Email
[**Tokenize**](PaymentGatewayApi.md#Tokenize) | **Post** /tokenize | Tokenize
[**Verify**](PaymentGatewayApi.md#Verify) | **Post** /verify | Verify
[**Void**](PaymentGatewayApi.md#Void) | **Post** /void/{pgIdOrig} | Void a Previously Authorized Transaction

# **Authorization**
> PgApiTransactionResponse Authorization(ctx, body)
Authorize Transaction

Authorizes a credit card for capture at a later time. An authorized transaction will continue to be open until it expires or a capture message is received. Authorizations are automatically voided if they are not captured within 28 days, although most issuing banks will release the hold after 24 hours in retail environments or 7 days in card not present environments.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PgApiTransactionRequest**](PgApiTransactionRequest.md)| Payment Gateway Authorization Request | 

### Return type

[**PgApiTransactionResponse**](PGApiTransactionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BatchClose**
> PgApiBatchCloseResponse BatchClose(ctx, body)
Close Batch

Closes a batch. Use this request when the timing of the batch close needs to be controlled rather than relying on the once-daily automatic batch close which is 9 PM Pacific by default, and can be changed in the Qualpay Manager administration settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PgApiBatchCloseRequest**](PgApiBatchCloseRequest.md)| Payment Gateway Batch Close Request | 

### Return type

[**PgApiBatchCloseResponse**](PGApiBatchCloseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Capture**
> PgApiCaptureResponse Capture(ctx, body, pgIdOrig)
Capture an Authorized Transaction

Captures an authorized transaction for any amount up to the amount originally authorized. An authorized transaction can only be captured once.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PgApiCaptureRequest**](PgApiCaptureRequest.md)| Payment Gateway Capture Request | 
  **pgIdOrig** | **string**| pgIdOrig | 

### Return type

[**PgApiCaptureResponse**](PGApiCaptureResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Credit**
> PgApiTransactionResponse Credit(ctx, body)
Issue Credit to Cardholder

Issues an unlinked credit. Credit requests require that the cardholder data is  provided in the request. Credits are only available during the first 30 days of account opening unless you contact Qualpay support to make other arrangements. The refund request should generally be used to return money to the cardholder, as it is a reversal of a previously captured transaction. A refund request is linked to the original transaction which is helpful for reconciliation purposes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PgApiTransactionRequest**](PgApiTransactionRequest.md)| Payment Gateway Credit Request | 

### Return type

[**PgApiTransactionResponse**](PGApiTransactionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Expire**
> PgApiExpireTokenResponse Expire(ctx, body)
Expire Token

Once expired, the token (card_id) is no longer valid for use in future transactions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PgApiExpireTokenRequest**](PgApiExpireTokenRequest.md)| Payment Gateway Expire Token Request | 

### Return type

[**PgApiExpireTokenResponse**](PGApiExpireTokenResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Force**
> PgApiTransactionResponse Force(ctx, body)
Force Transaction Approval

Forces an approval, used when an online authorization request received a 'declined' reason code and you have received an authorization from a voice or automated response (ARU) system. The required fields are the same as a sale or authorization request, except that the expiration date (exp_date) is not required, and the 6-character authorization code (auth_code) is required.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PgApiTransactionRequest**](PgApiTransactionRequest.md)| Payment Gateway Force Request | 

### Return type

[**PgApiTransactionResponse**](PGApiTransactionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCardTypeInformation**
> ArdefResponse GetCardTypeInformation(ctx, body)
Get Card type Information for Visa, Mastercard, and Discover

Gets Card type information for Visa, Mastercard, and Discover. Useful if you prohibit or allow certain activity based on card type. For example, you may not want to allow a subscription to be created using a prepaid card.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ArdefRequest**](ArdefRequest.md)| Card Type Request | 

### Return type

[**ArdefResponse**](ArdefResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InitiatePayment**
> PgApiPaymentResponse InitiatePayment(ctx, body)
Initiate a payment to a payee

Initiates a payment to the specified payee.  Payee relationship will be validated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PgApiPaymentRequest**](PgApiPaymentRequest.md)| Payment Request | 

### Return type

[**PgApiPaymentResponse**](PGApiPaymentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Recharge**
> PgApiRechargeResponse Recharge(ctx, body, pgIdOrig)
Recharge Previously Settled Transaction

Creates a new sale transaction using the cardholder data from a previous successful transaction.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PgApiRechargeRequest**](PgApiRechargeRequest.md)| Payment Gateway Recharge Request | 
  **pgIdOrig** | **string**| pgIdOrig | 

### Return type

[**PgApiRechargeResponse**](PGApiRechargeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Refund**
> PgApiRefundResponse Refund(ctx, body, pgIdOrig)
Refund Previously Captured Transaction

Returns money to the cardholder from a previously captured transaction. Multiple refunds are allowed per captured transaction, provided that the sum of all refunds does not exceed the original captured transaction amount. Authorizations that have not been captured are not eligible for a refund.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PgApiRefundRequest**](PgApiRefundRequest.md)| Payment Gateway Refund Request | 
  **pgIdOrig** | **string**| pgIdOrig | 

### Return type

[**PgApiRefundResponse**](PGApiRefundResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestBalance**
> PgApiBalanceResponse RequestBalance(ctx, body)
Request a balance

Requests a balance for the specified FBO merchant account.  FBO relationship will be validated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PgApiBalanceRequest**](PgApiBalanceRequest.md)| Balance Request | 

### Return type

[**PgApiBalanceResponse**](PGApiBalanceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Sale**
> PgApiTransactionResponse Sale(ctx, body)
Sale (Auth + Capture)

Requests authorization, and, if approved, will immediately capture the transaction to be included in the next batch close. This transaction type is used in card-present environments, and also card-not-present environments where no physical goods are being shipped.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PgApiTransactionRequest**](PgApiTransactionRequest.md)| Payment Gateway Sale Request | 

### Return type

[**PgApiTransactionResponse**](PGApiTransactionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendReceipt**
> PgApiEmailReceiptResponse SendReceipt(ctx, body, pgId)
Send Transaction Receipt Email

Sends the transaction receipt to multiple email addresses. Receipts can be sent only for successful transactions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PgApiEmailReceiptRequest**](PgApiEmailReceiptRequest.md)| Payment Gateway Email Receipt Request | 
  **pgId** | **string**| pgId | 

### Return type

[**PgApiEmailReceiptResponse**](PGApiEmailReceiptResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Tokenize**
> PgApiTokenizeResponse Tokenize(ctx, body)
Tokenize

Once stored, a unique card_id is returned for use in future transactions. Optionally, tokenization can be requested in an authorization, verify, force, credit, or sale request by sending the tokenize field set to true.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PgApiTokenizeRequest**](PgApiTokenizeRequest.md)| Payment Gateway Tokenize Request | 

### Return type

[**PgApiTokenizeResponse**](PGApiTokenizeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Verify**
> PgApiVerifyResponse Verify(ctx, body)
Verify

A verify request will return success if the cardholder information was verified by the issuer. If AVS or CVV data is included in the message, then the AVS or CVV result code will be returned in the response message. This is useful if you want to determine if you have been presented with a valid card, but are not ready to authorize the card.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PgApiVerifyRequest**](PgApiVerifyRequest.md)| Payment Gateway Card Verify Request | 

### Return type

[**PgApiVerifyResponse**](PGApiVerifyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Void**
> PgApiVoidResponse Void(ctx, body, pgIdOrig)
Void a Previously Authorized Transaction

Authorizations can be voided at any time until Qualpay automatically voids them at 28 days. Captured transactions can be voided until the batch is closed. If your batch closes and you did not void the transaction in time, you may make a refund request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PgApiVoidRequest**](PgApiVoidRequest.md)| Payment Gateway Void Request | 
  **pgIdOrig** | **string**| pgIdOrig | 

### Return type

[**PgApiVoidResponse**](PGApiVoidResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

