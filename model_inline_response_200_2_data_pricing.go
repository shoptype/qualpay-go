/*
 * APIs
 *
 * This document describes the Platform and PG API.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package qualpay

type InlineResponse2002DataPricing struct {
	// Unique ID assigned by Qualpay to this pricing plan.
	PricingId int64        `json:"pricing_id,omitempty"`
	Plan      *interface{} `json:"plan,omitempty"`
}
