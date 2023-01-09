/*
 * APIs
 *
 * This document describes the Platform and PG API.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package qualpay

type SettledTransactionRequest struct {
	// <strong>Format: </strong>Variable length, up to 18 N<br><strong>Description: </strong>Unique ID assigned to this deposit.
	RecId int64 `json:"rec_id,omitempty"`
	// <strong>Format: </strong>Variable length, up to 16 N<br><strong>Description: </strong>Unique ID assigned to a Merchant.
	MerchantId int64 `json:"merchant_id,omitempty"`
	// <strong>Format: </strong>Variable length, up to 10 AN, in YYYY-MM-DD<br><strong>Description: </strong>The date the batch was settled to us.
	BatchDate string `json:"batch_date,omitempty"`
	// <strong>Format: </strong>Variable length, up to 25 AN<br><strong>Description: </strong>The doing business as name of the merchant.
	DbaName string `json:"dba_name,omitempty"`
	// <strong>Format: </strong>Variable length, up to 13 AN<br><strong>Description: </strong>Sent to the card issuer; either the merchant's city, or phone number.
	MerchCityPh string `json:"merch_city_ph,omitempty"`
	// <strong>Format: </strong>Variable length, up to 3 AN<br><strong>Description: </strong>Sent to the card issuer; the merchant's state.
	MerchState string `json:"merch_state,omitempty"`
	// <strong>Format: </strong>Variable length, up to 10 AN<br><strong>Description: </strong>Sent to the card issuer; the merchant's ZIP code.
	MerchZip string `json:"merch_zip,omitempty"`
	// <strong>Format: </strong>Variable length, up to 3 AN<br><strong>Description: </strong>Sent to the card issuer; the merchant's country code.
	MerchCountry string `json:"merch_country,omitempty"`
	// <strong>Format: </strong>Variable length, up to 3 N<br><strong>Description: </strong>A non-unique ID assigned by the merchant's terminal, POS device, or gateway for this batch, in the range of 1 - 999.
	BatchNumber int64 `json:"batch_number,omitempty"`
	// <strong>Format: </strong>Variable length, up to 10 N<br><strong>Description: </strong>Unique ID assigned to this batch.
	BatchId int64 `json:"batch_id,omitempty"`
	// <strong>Format: </strong>Variable length, up to 16 AN<br><strong>Description: </strong>The truncated card number of the transaction.
	CardNumber string `json:"card_number,omitempty"`
	// <strong>Format: </strong>Variable length, up to 2 AN<br><strong>Description: </strong>The card brand of this transaction.
	CardType string `json:"card_type,omitempty"`
	// <strong>Format: </strong>Variable length, up to 10 AN, in YYYY-MM-DD<br><strong>Description: </strong>The date the transaction was authorized by the merchant.
	AuthDate string `json:"auth_date,omitempty"`
	// <strong>Format: </strong>Variable length, up to 10 AN, in YYYY-MM-DD<br><strong>Description: </strong>The date the transaction was captured by the merchant.
	TranDate string `json:"tran_date,omitempty"`
	// <strong>Format: </strong>Variable length, up to 10 AN, in YYYY-MM-DD<br><strong>Description: </strong>The date we settled the transaction with the issuer.
	SettleDate string `json:"settle_date,omitempty"`
	// <strong>Format: </strong>Variable length, up to 12,2 N<br><strong>Description: </strong>The amount of the original authorization.
	AmtAuth string `json:"amt_auth,omitempty"`
	// <strong>Format: </strong>Variable length, up to 12,2 N<br><strong>Description: </strong>The amount of the settled transaction.
	AmtTran string `json:"amt_tran,omitempty"`
	// <strong>Format: </strong>Fixed length, 3 AN<br><strong>Description: </strong>The ISO 4217 numeric currency code of the dispute.
	TranCurrency string `json:"tran_currency,omitempty"`
	// <strong>Format: </strong>Variable length, up to 23 AN<br><strong>Description: </strong>The bank reference number of the deposit.
	ReferenceNumber string `json:"reference_number,omitempty"`
	// <strong>Format: </strong>Variable length, up to 25 AN<br><strong>Description: </strong>A merchant supplied tracking number, generally an invoice or purchase number. This number may be visible to the cardholder, depending on card issuer.
	PurchaseId string `json:"purchase_id,omitempty"`
	// <strong>Format: </strong>Fixed length, 6 AN<br><strong>Description: </strong>The authorization code provided by the card issuer when the card was approved.
	AuthCode string `json:"auth_code,omitempty"`
	// <strong>Format: </strong>Fixed length, 1 AN<br><strong>Description: </strong>The transaction ECI (e-commerce indicator), an indicatior of how the card was processed.
	MotoEcommInd string `json:"moto_ecomm_ind,omitempty"`
	// <strong>Format: </strong>Variable length, up to 3 AN<br><strong>Description: </strong>Indicates the way a card number was entered. For all codes, please see <a href=\"/developer/api/reference#pos_entry_mode\" target=\"_blank\">POS Entry Modes</a>.
	PosEntryMode string `json:"pos_entry_mode,omitempty"`
	// <strong>Format: </strong>Variable length, up to 12,2 N<br><strong>Description: </strong>The amount of this transaction funded to the merchant (in the funded currency).
	AmtFunded string `json:"amt_funded,omitempty"`
	// <strong>Format: </strong>Fixed length, 3 AN<br><strong>Description: </strong>The ISO 4217 numeric currency code of the transaction.
	FundedCurrency string `json:"funded_currency,omitempty"`
	// <strong>Format: </strong>Fixed length, 32 AN<br><strong>Description: </strong>32-byte unique identifier generated by the payment gateway, returned in all valid responses. Applicable only to Payment Gateway transactions. It provides a reference to the transaction and is required for some post-auth operations (e.g. capture, refund, or void).
	PgId string `json:"pg_id,omitempty"`
	// <strong>Format: </strong>Variable length, up to 128 AN<br><strong>Description: </strong>Merchant provided reference value that will be stored with the transaction data and included with transaction data in reports within Merchant Manager. This value will also be attached to any lifecycle transactions (e.g. retrieval requests and chargebacks) that may occur.
	MerchRefNum string `json:"merch_ref_num,omitempty"`
	// <strong>Format: </strong>Variable length, up to 11 N<br><strong>Description: </strong>Acquirer reference number is an 11-digit number generated by the product initiating the transaction. It is a unique number that both the acquirer and the issuer can use to identify a transaction. For chargeback adjustments, the acquirer reference number is used as the deposit reference number.
	AcqReferenceNumber string `json:"acq_reference_number,omitempty"`
	// <strong>Format: </strong>Variable length, up to 12 AN<br><strong>Description: </strong>The Authorization Retrieval Reference Number (RRN) is a unique identifier assigned by an acquirer to an authorization.
	AuthRrn string `json:"auth_rrn,omitempty"`
	// <strong>Format: </strong>Variable length, up to 15 AN<br><strong>Description: </strong>The Authorization transaction identifier is a unique identifier returned by the issuing bank for an electronically authorized transaction.
	AuthTranId string `json:"auth_tran_id,omitempty"`
}
