/*
 * APIs
 *
 * This document describes the Platform and PG API.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package qualpay

type UpdateOutstandingRequest struct {
	FromContact    *Contact `json:"from_contact,omitempty"`
	BillingContact *Contact `json:"billing_contact,omitempty"`
}
