/*
 * APIs
 *
 * This document describes the Platform and PG API.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type StatementDisputeData struct {
	// For INTERNAL USE ONLY.
	RecIdLinked int64 `json:"rec_id_linked,omitempty"`
	// <strong>Format: </strong>Variable length, up to 16 N<br><strong>Description: </strong>Unique ID assigned to a merchant.
	MerchantId int64 `json:"merchant_id,omitempty"`
	// Billing month of the statement in YYYY-MM-DD.
	BillingMonth string `json:"billing_month,omitempty"`
	// The date the ACH debit or credit adjustment was made to the depository account in YYYY-MM-DD format.
	BatchDate time.Time `json:"batch_date,omitempty"`
	// An 11-digit number generated by the product initiating the transaction. The number is available when the transaction is settled and is a unique number that both the acquirer and the issuer can use to identify a transaction. The reference is used as the deposit reference number.
	ReferenceNumber int64 `json:"reference_number,omitempty"`
	// SDefines the payment instrument type of Visa, MasterCard, American Express, or Discover.
	PlanType string `json:"plan_type,omitempty"`
	// The type of dispute.
	TypeId string `json:"type_id,omitempty"`
	// Defines by a Yes or No whether the dispute triggered a billable event. All billable dispute events will appear as a total line item in the Fees section in YYYY-MM-DD format.
	Billable string `json:"billable,omitempty"`
	// The dollar value of the adjustment that is a debit to the depository account.
	AmtPurch float64 `json:"amt_purch,omitempty"`
	// The dollar value of the adjustment that is a credit to the depository account.
	AmtReturn float64 `json:"amt_return,omitempty"`
}
