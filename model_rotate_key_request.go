/*
 * APIs
 *
 * This document describes the Platform and PG API.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type RotateKeyRequest struct {
	// The number of minutes the previous key will be valid. 
	Minutes int32 `json:"minutes,omitempty"`
}