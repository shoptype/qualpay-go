/*
 * APIs
 *
 * This document describes the Platform and PG API.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type QpApiResponse struct {
	// <strong>Format: </strong>Fixed length, 1 N<br><strong>Description: </strong>Response code from API. 0 indicates success. Refer to <a href=\"/developer/api/reference#api-response-codes\" target=\"_blank\">Platform API Response Codes</a> for entire list of return codes.
	Code int32 `json:"code,omitempty"`
	// <strong>Format: </strong>Variable length AN<br><strong>Description: </strong>A short description of the API response code.
	Message string `json:"message,omitempty"`
	Data *QpApiData `json:"data,omitempty"`
}
