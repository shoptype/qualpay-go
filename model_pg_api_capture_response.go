/*
 * APIs
 *
 * This document describes the Platform and PG API.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type PgApiCaptureResponse struct {
	// <strong>Format: </strong>Fixed length, 32 AN<br><strong>Description: </strong>32-byte unique identifier generated by the payment gateway, returned in all valid responses.
	PgId string `json:"pg_id,omitempty"`
	// <strong>Format: </strong>Fixed length, 3 AN<br><strong>Description: </strong>Response code from the payment gateway. \"000\" indicates success. Refer to <a href=\"/developer/api/reference#payment-gateway-response-codes\" target=\"_blank\">Payment Gateway Response Codes</a> for entire list of Payment Gateway Response Codes.
	Rcode string `json:"rcode,omitempty"`
	// <strong>Format: </strong>Variable length<br><strong>Description: </strong>Response text from the payment gateway.
	Rmsg string `json:"rmsg,omitempty"`
	// <strong>Format: </strong>Variable length<br><strong>Description: </strong>Data that was provided in echo_fields in the request.
	EchoFields string `json:"echo_fields,omitempty"`
}
