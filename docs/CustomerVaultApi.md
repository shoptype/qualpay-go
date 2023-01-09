# qualpay{{classname}}

All URIs are relative to */platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBillingCard**](CustomerVaultApi.md#AddBillingCard) | **Post** /vault/customer/{customer_id}/billing | Add a Payment Method
[**AddCustomer**](CustomerVaultApi.md#AddCustomer) | **Post** /vault/customer | Add a Customer
[**AddShippingAddress**](CustomerVaultApi.md#AddShippingAddress) | **Post** /vault/customer/{customer_id}/shipping | Add a Shipping Address
[**BrowseCustomers**](CustomerVaultApi.md#BrowseCustomers) | **Get** /vault/customer | Get all Customers
[**DeleteBillingCard**](CustomerVaultApi.md#DeleteBillingCard) | **Put** /vault/customer/{customer_id}/billing/delete | Delete a Payment Method
[**DeleteCustomer**](CustomerVaultApi.md#DeleteCustomer) | **Delete** /vault/customer/{customer_id} | Delete a Customer
[**DeleteShippingAddress**](CustomerVaultApi.md#DeleteShippingAddress) | **Delete** /vault/customer/{customer_id}/shipping/{id} | Delete a Shipping Address
[**GetBillingCards**](CustomerVaultApi.md#GetBillingCards) | **Get** /vault/customer/{customer_id}/billing | Get Payment Methods
[**GetCustomer**](CustomerVaultApi.md#GetCustomer) | **Get** /vault/customer/{customer_id} | Get by Customer ID
[**GetShippingAddresses**](CustomerVaultApi.md#GetShippingAddresses) | **Get** /vault/customer/{customer_id}/shipping | Get Shipping Addresses
[**SetPrimaryBillingCard**](CustomerVaultApi.md#SetPrimaryBillingCard) | **Put** /vault/customer/{customer_id}/billing/primary | Set Primary Payment Method
[**SetPrimaryShippingAddress**](CustomerVaultApi.md#SetPrimaryShippingAddress) | **Put** /vault/customer/{customer_id}/shipping/primary | Set Primary Shipping Address
[**UpdateBillingCard**](CustomerVaultApi.md#UpdateBillingCard) | **Put** /vault/customer/{customer_id}/billing | Update a Payment Method
[**UpdateCustomer**](CustomerVaultApi.md#UpdateCustomer) | **Put** /vault/customer/{customer_id} | Update a Customer
[**UpdateShippingAddress**](CustomerVaultApi.md#UpdateShippingAddress) | **Put** /vault/customer/{customer_id}/shipping | Update a Shipping Address

# **AddBillingCard**
> CustomerResponse AddBillingCard(ctx, body, customerId)
Add a Payment Method

Tokenizes and adds a payment method to a customer record. The response will include the tokenized card_number in the field, card_id, which can be used to manage the card and can be used in place of the card_number in Payment Gateway API requests. There is no limit on the number of payment methods you can add to a customer. The first payment method added will be designated as the primary payment method, used for subscription payments. You can use the Set Primary Payment Method API to change the default payment method.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddBillingCardRequest**](AddBillingCardRequest.md)| Customer | 
  **customerId** | **string**| Customer ID | 

### Return type

[**CustomerResponse**](CustomerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddCustomer**
> CustomerResponse AddCustomer(ctx, body)
Add a Customer

Adds a customer to the Customer Vault. A customer_id is required to create invoice and subscription payments. Payment methods, billing addresses and shipping addresses may be included in this request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddCustomerRequest**](AddCustomerRequest.md)| Customer | 

### Return type

[**CustomerResponse**](CustomerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddShippingAddress**
> CustomerResponse AddShippingAddress(ctx, body, customerId)
Add a Shipping Address

Adds a shipping address to a customer. The response will include a unique shipping_id. The shipping_id can be used to manage the shipping_address.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddShippingAddressRequest**](AddShippingAddressRequest.md)| Shipping Address | 
  **customerId** | **string**| Shipping Address | 

### Return type

[**CustomerResponse**](CustomerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BrowseCustomers**
> CustomerListResponse BrowseCustomers(ctx, optional)
Get all Customers

Gets an array of customer vault objects. You can use filters to limit the results returned. Optional query parameters determine size and sort order of the returned array.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CustomerVaultApiBrowseCustomersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerVaultApiBrowseCustomersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| The number of records in the result. | [default to 10]
 **orderOn** | **optional.String**| The field on which the results will be sorted on. Refer to the response model for available fields. | [default to customer_id]
 **orderBy** | **optional.String**| Ascending or Descending Sort order of the result. Possible values are: asc (Ascending sort order), desc (Descending sort order) | [default to desc]
 **page** | **optional.Int32**| Zero-based page number, use this to choose a page when there are more results than the count parameter. | [default to 0]
 **filter** | [**optional.Interface of []string**](string.md)| Results can be filtered by custom filter criteria. Refer to [Filter](/api/reference#filters) to use the filter parameter. | 
 **merchantId** | **optional.Int64**| Unique ID assigned to a merchant. | [default to 0]

### Return type

[**CustomerListResponse**](CustomerListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBillingCard**
> CustomerResponse DeleteBillingCard(ctx, body, customerId)
Delete a Payment Method

Deletes a customer's specific payment method based on the card_id provided in the request. The card_id is required in the request body to delete the payment method. A payment method with active subscriptions cannot be deleted. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeleteBillingCardRequest**](DeleteBillingCardRequest.md)| Customer | 
  **customerId** | **string**| Customer ID | 

### Return type

[**CustomerResponse**](CustomerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCustomer**
> QpApiResponse DeleteCustomer(ctx, customerId, optional)
Delete a Customer

Deletes a customer from the customer vault. Deleting a customer will delete all the subscriptions associated with the customer. **This operation cannot be reversed.**

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| Customer ID | 
 **optional** | ***CustomerVaultApiDeleteCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerVaultApiDeleteCustomerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **merchantId** | **optional.Int64**| Unique ID assigned to a merchant. | [default to 0]

### Return type

[**QpApiResponse**](QPApiResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteShippingAddress**
> CustomerResponse DeleteShippingAddress(ctx, customerId, id, optional)
Delete a Shipping Address

Deletes a customer's specific shipping address based on the shipping_id provided in the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| Customer ID | 
  **id** | **int64**| Shipping ID | 
 **optional** | ***CustomerVaultApiDeleteShippingAddressOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerVaultApiDeleteShippingAddressOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **merchantId** | **optional.Int64**| Unique ID assigned to a merchant. | [default to 0]

### Return type

[**CustomerResponse**](CustomerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBillingCards**
> GetBillingResponse GetBillingCards(ctx, customerId, optional)
Get Payment Methods

Gets a list of payment methods associated with a customer_id. The response will contain an array of billing_cards.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| Customer ID | 
 **optional** | ***CustomerVaultApiGetBillingCardsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerVaultApiGetBillingCardsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **merchantId** | **optional.Int64**| Unique ID assigned to a merchant. | [default to 0]

### Return type

[**GetBillingResponse**](GetBillingResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomer**
> CustomerResponse GetCustomer(ctx, customerId, optional)
Get by Customer ID

Gets a customer vault record by customer_id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| Customer ID | 
 **optional** | ***CustomerVaultApiGetCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerVaultApiGetCustomerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **merchantId** | **optional.Int64**| Unique ID assigned to a merchant. | [default to 0]

### Return type

[**CustomerResponse**](CustomerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetShippingAddresses**
> GetShippingResponse GetShippingAddresses(ctx, customerId, optional)
Get Shipping Addresses

Gets all shipping addresses for a customer. The response will include an array of shipping addresses. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| Customer ID | 
 **optional** | ***CustomerVaultApiGetShippingAddressesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerVaultApiGetShippingAddressesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **merchantId** | **optional.Int64**| Unique ID assigned to a merchant. | [default to 0]

### Return type

[**GetShippingResponse**](GetShippingResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetPrimaryBillingCard**
> CustomerResponse SetPrimaryBillingCard(ctx, body, customerId)
Set Primary Payment Method

Sets a customer's specific payment method as the primary payment method. A primary payment method is used for subscription payments.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SetPrimaryBillingCardRequest**](SetPrimaryBillingCardRequest.md)| Customer | 
  **customerId** | **string**| Customer ID | 

### Return type

[**CustomerResponse**](CustomerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetPrimaryShippingAddress**
> CustomerResponse SetPrimaryShippingAddress(ctx, body, customerId)
Set Primary Shipping Address

Sets a customer's specific shipping address as primary.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SetPrimaryShippingAddressRequest**](SetPrimaryShippingAddressRequest.md)| Shipping Address | 
  **customerId** | **string**| Customer ID | 

### Return type

[**CustomerResponse**](CustomerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBillingCard**
> CustomerResponse UpdateBillingCard(ctx, body, customerId)
Update a Payment Method

Updates a customer's payment method by card_id. The complete billing record should be included in the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateBillingCardRequest**](UpdateBillingCardRequest.md)| Customer | 
  **customerId** | **string**| Customer ID | 

### Return type

[**CustomerResponse**](CustomerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCustomer**
> CustomerResponse UpdateCustomer(ctx, body, customerId)
Update a Customer

Updates a customer vault record. You can send the entire resource or only the fields that require an update. When updating array fields - billing_cards or shipping_addresses - the entire array should be included in the request body. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateCustomerRequest**](UpdateCustomerRequest.md)| Customer | 
  **customerId** | **string**| Customer ID | 

### Return type

[**CustomerResponse**](CustomerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateShippingAddress**
> CustomerResponse UpdateShippingAddress(ctx, body, customerId)
Update a Shipping Address

Update a shipping address for a customer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateShippingAddressRequest**](UpdateShippingAddressRequest.md)| Shipping Address | 
  **customerId** | **string**| Shipping Address | 

### Return type

[**CustomerResponse**](CustomerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

