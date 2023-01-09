# qualpay{{classname}}

All URIs are relative to */platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCheckout**](CheckoutApi.md#AddCheckout) | **Post** /checkout | Create a Checkout Link
[**GetDetails**](CheckoutApi.md#GetDetails) | **Get** /checkout/{checkoutId} | Lookup Checkout Payment

# **AddCheckout**
> CheckoutLinkResponse AddCheckout(ctx, body)
Create a Checkout Link

Create a Checkout link to make payments, or, add or edit payment methods stored in Customer Vault. The type of link generated, a payment link or vault link, is based on the checkout profile used in the request. Checkout Profiles are created in Merchant Manager. This operation generates a unique checkout_id. For payment links, the checkout_id can be used to query the status of the payment using the Lookup Checkout Payment API. <br><br><strong>For a payment:</strong>               Use a Dynamic Checkout profile_id in the request body.<br><strong>To add a new payment method:</strong> Use a Customer Vault profile_id and include the customer_id in the request. New customers are automatically added to the Customer Vault. Include customer_first_name and customer_last_name or customer_firm_name if you wish to create a new customer, and to reduce the amount of information entered by your customer. <bt><strong>To update payment details:</strong>   Use a Customer Vault Checkout profile_id and include the customer_id and card_id in the request. The card_id must be associated with the customer_id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CheckoutRequest**](CheckoutRequest.md)| Checkout Resource | 

### Return type

[**CheckoutLinkResponse**](CheckoutLinkResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDetails**
> CheckoutResponse GetDetails(ctx, checkoutId)
Lookup Checkout Payment

Queries the status of a payment made through Hosted Checkout. The operation returns a checkout object which contains information about the transaction.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **checkoutId** | **string**| Checkout ID | 

### Return type

[**CheckoutResponse**](CheckoutResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

