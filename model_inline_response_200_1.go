/*
 * APIs
 *
 * This document describes the Platform and PG API.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type InlineResponse2001 struct {
	// <strong>Format: </strong>Fixed length, 1 N<br><strong>Description: </strong>Response code from API. 0 indicates success. Refer to <a href=\"/developer/api/reference#api-response-codes\"target=\"_blank\">Platform API Response Codes</a> for entire list of return codes.
	Code int32 `json:"code,omitempty"`
	// <strong>Format: </strong>Variable length AN<br><strong>Description: </strong>A short description of the API response code.
	Message string `json:"message,omitempty"`
	// <strong>Format: </strong>Variable length N<br><strong>Description: </strong>Total Number of pages. 
	TotalPages int32 `json:"totalPages,omitempty"`
	// <strong>Format: </strong>Variable length, up to 3 N<br><strong>Description: </strong>Total number of records upto a maximum of 100.
	TotalRecords int32 `json:"totalRecords,omitempty"`
	// An array of the resource object.
	Data []InlineResponse2001Data `json:"data,omitempty"`
}
