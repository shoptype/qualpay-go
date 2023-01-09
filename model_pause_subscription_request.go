/*
 * APIs
 *
 * This document describes the Platform and PG API.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type PauseSubscriptionRequest struct {
	// <strong>Format: </strong>Variable length, up to 32 AN<br><strong>Description: </strong>Customer id of the subscriber. 
	CustomerId string `json:"customer_id"`
	// <strong>Format: </strong>Variable length, up to 16 AN<br><strong>Description: </strong>Identifies the merchant to whom this request applies. Optional field, applicable only if the request is sent on behalf of another merchant.<br><strong>Conditional Requirement: </strong>Required if this request is on behalf of another merchant.
	MerchantId int64 `json:"merchant_id,omitempty"`
}
