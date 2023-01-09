/*
 * APIs
 *
 * This document describes the Platform and PG API.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type PaymentProfile struct {
	// The id of the payment profile.
	ProfileId string `json:"profile_id,omitempty"`
	// Payment profile label.
	Label string `json:"label,omitempty"`
	// The currency code of the payment profile.
	Currency string `json:"currency,omitempty"`
	// The auto close time for the batch in Pacific Time.
	AutoCloseTime int32 `json:"auto_close_time,omitempty"`
}