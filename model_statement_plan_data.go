/*
 * APIs
 *
 * This document describes the Platform and PG API.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package qualpay

type StatementPlanData struct {
	// For INTERNAL USE ONLY.
	RecIdLinked int64 `json:"rec_id_linked,omitempty"`
	// <strong>Format: </strong>Variable length, up to 16 N<br><strong>Description: </strong>Unique ID assigned to a merchant.
	MerchantId int64 `json:"merchant_id,omitempty"`
	// Billing month of the statement in YYYY-MM-DD
	BillingMonth string `json:"billing_month,omitempty"`
	// Defines the payment instrument type of Visa, MasterCard, American Express, or Discover.
	PlanCode string `json:"plan_code,omitempty"`
	// The total funded sales transactions count.
	CntPurch int64 `json:"cnt_purch,omitempty"`
	// The total funded sales dollar amount.
	AmtPurch float64 `json:"amt_purch,omitempty"`
	// The total funded credit transactions count.
	CntReturn int64 `json:"cnt_return,omitempty"`
	// The total funded credit dollar amount.
	AmtReturn float64 `json:"amt_return,omitempty"`
	// The net dollar amount of transactions.
	AmtNet float64 `json:"amt_net,omitempty"`
	// Average ticket amount.
	AmtAvgTicket float64 `json:"amt_avg_ticket,omitempty"`
	// The per item rate for the payment type.
	PerItem float64 `json:"per_item,omitempty"`
	// The discount rate for the payment type is applied to the Sales Amount.
	Rate float64 `json:"rate,omitempty"`
	// The Amount of discount due
	DiscountDue float64 `json:"discount_due,omitempty"`
}
