/*
 * APIs
 *
 * This document describes the Platform and PG API.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package qualpay

type DisputeData struct {
	// Unique ID assigned to dispute.
	RecId int64 `json:"rec_id,omitempty"`
	// Unique ID assigned to a merchant.
	MerchantId int64 `json:"merchant_id,omitempty"`
	// The date the dispute was issued.
	IncomingDate string `json:"incoming_date,omitempty"`
	// The amount of the original settled transaction.
	AmtTran float64 `json:"amt_tran,omitempty"`
	// The amount of the dispute; equal to or less than the amt_tran (when in USD).
	AmtDispute float64 `json:"amt_dispute,omitempty"`
	// The card association's reason why the dispute was issued.
	ReasonCode string `json:"reason_code,omitempty"`
	// The card association's reason code desc why the dispute was issued.
	ReasonDesc string `json:"reason_desc,omitempty"`
	// The type of dispute. For all types, please refer to <a href=\"/developer/api/reference#dispute-type\" target=\"_blank\">Dispute Type</a>
	DataType string `json:"data_type,omitempty"`
	// The type of dispute workflow. For all types, please refer to <a href=\"/developer/api/reference#dispute-workflow\" target=\"_blank\">Dispute Workflow</a>
	CbrWorkflow string `json:"cbr_workflow,omitempty"`
	// Current Dispute Status.
	Status string `json:"status,omitempty"`
	// The date the dispute's status most recently changed.
	DateStatusChange string `json:"date_status_change,omitempty"`
	// The truncated card number of the dispute.
	CardNumber string `json:"card_number,omitempty"`
	// The date of the initial transaction.
	TranDate string `json:"tran_date,omitempty"`
	// The bank reference number of the deposit.
	ReferenceNumber string `json:"reference_number,omitempty"`
	// Issuer provided reference number for this transaction.
	CbRefNum string `json:"cb_ref_num,omitempty"`
	// Unique Vendor ID assigned to integrator.
	VendorId int64 `json:"vendor_id,omitempty"`
	// Purchase Identifier (also referred to as the invoice number generated by the merchant).
	PurchaseId string `json:"purchase_id,omitempty"`
}
