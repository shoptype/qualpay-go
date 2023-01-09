/*
 * APIs
 *
 * This document describes the Platform and PG API.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type EmbeddedKey struct {
	// <strong>Format: </strong>Fixed length, 32 AN<br><strong>Description: </strong>A single use token used for loading embedded fields. The key will be invalidated   when a card is successfully verified. If unused, the token will expire in 30 minutes. 
	TransientKey string `json:"transient_key,omitempty"`
	// <strong>Format: </strong>Variable length, up to 16 AN<br><strong>Description: </strong>Unique ID assigned to a merchant.
	MerchantId int64 `json:"merchant_id,omitempty"`
	// <strong>Format: </strong>Variable length AN<br><strong>Description: </strong>This field contains the transient key creation timestamp. 
	DbTimestamp string `json:"db_timestamp,omitempty"`
	// <strong>Format: </strong>Variable length AN<br><strong>Description: </strong>This field contains the timestamp when the transient key will expire.
	ExpiryTime string `json:"expiry_time,omitempty"`
}
