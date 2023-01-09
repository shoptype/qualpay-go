/*
 * APIs
 *
 * This document describes the Platform and PG API.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type PgApiPaymentRequest struct {
	// <strong>Format: </strong>Variable length, up to 12,2 N<br><strong>Description: </strong>Amount of payment.
	AmtTran float64 `json:"amt_tran"`
	// <strong>Format: </strong>Variable length, up to 16 N<br><strong>Description: </strong>Unique identifier on the payments system.
	MerchantId int64 `json:"merchant_id"`
	// <strong>Format: </strong>Variable length, up to 16 N<br><strong>Description: </strong>For Benefit Of (FBO) merchant account identifier on the payments system.
	FboId int64 `json:"fbo_id,omitempty"`
	// <strong>Format: </strong>Variable length, up to 32 AN<br><strong>Description: </strong>A unique ID generated to identify a payee.
	PayeeId string `json:"payee_id,omitempty"`
	// <strong>Format: </strong>Variable length, up to 128 AN<br><strong>Description: </strong>[Optional] Reference value that will be stored with the payment data and included in reports.
	MerchRefNum string `json:"merch_ref_num,omitempty"`
	// <strong>Format: </strong>Variable length, up to 16 N<br><strong>Description: </strong>[Optional] Identifies the vendor submitting the payment request.
	VendorId int64 `json:"vendor_id,omitempty"`
}