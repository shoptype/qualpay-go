/*
 * APIs
 *
 * This document describes the Platform and PG API.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package qualpay

type PgApiBatchCloseRequest struct {
	// <strong>Format: </strong>Variable length, up to 12 N<br><strong>Description: </strong>Unique identifier on our system.
	MerchantId int64 `json:"merchant_id"`
	// <strong>Format: </strong>Variable length, up to 32 AN<br><strong>Description: </strong>Use to indicate which company developed the integration or the name of the payment solution that is connected to us. Suggested usage is softwareABCv1.0 or companyXYZv2.0.
	DeveloperId string `json:"developer_id,omitempty"`
	// <strong>Format: </strong>Variable length<br><strong>Description: </strong>This field contains a JSON array of field data that will be echoed back in the response message.
	EchoFields string `json:"echo_fields,omitempty"`
	// <strong>Format: </strong>Variable length, up to 4 N<br><strong>Description: </strong>When a merchant has more than one location using the same currency, this value is used to specify the specific location for this request.
	LocId string `json:"loc_id,omitempty"`
	// <strong>Format: </strong>Fixed length, 20 N<br><strong>Description: </strong>Explicitly identifies which Payment Gateway profile should be used for the request.
	ProfileId string `json:"profile_id,omitempty"`
	// <strong>Format: </strong>Variable length<br><strong>Description: </strong>This field contains a JSON array of field data that will be included with the transaction data reported in Merchant Manager.
	ReportData string `json:"report_data,omitempty"`
	// INTERNAL USE ONLY.
	SessionId string `json:"session_id,omitempty"`
	// <strong>Format: </strong>Variable length, up to 3 N<br><strong>Default: </strong>840<br><strong>Description: </strong>ISO numeric currency code for the transaction. Refer to <a href=\"/developer/api/reference#country-codes\" target=\"_blank\">Country Codes</a> for a list of currency codes.
	TranCurrency int32 `json:"tran_currency,omitempty"`
	// INTERNAL USE ONLY.
	UserId int64 `json:"user_id,omitempty"`
}
