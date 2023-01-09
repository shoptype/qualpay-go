/*
 * APIs
 *
 * This document describes the Platform and PG API.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type GatewayResponse struct {
	// <strong>Format: </strong>Variable length, up to 16 AN<br><strong>Description: </strong>a generated payment gateway ID for the transaction
	PgId string `json:"pg_id,omitempty"`
	// <strong>Format: </strong>Fixed length, 3 AN<br><strong>Description: </strong>Gateway Response Code. Refer to <a href=\"/developer/api/reference#payment-gateway-response-codes\" target=\"_blank\">Payment Gateway Response Codes</a> for possible values.
	Rcode string `json:"rcode,omitempty"`
	// <strong>Format: </strong>Variable length AN<br><strong>Description: </strong>A short description of the response code. Refer to <a href=\"/developer/api/reference#gateway-response-codes\" target=\"_blank\">Payment Gateway Response Codes</a> for response code descriptions.
	Rmsg string `json:"rmsg,omitempty"`
	// <strong>Format: </strong>Variable length, up to 16 AN<br><strong>Description: </strong>Status of the gateway transaction.
	Status string `json:"status,omitempty"`
	// <strong>Format: </strong>Variable length, up to 6 AN<br><strong>Description: </strong>Authorization code from issuer.
	AuthCode string `json:"auth_code,omitempty"`
	// <br><strong>Description: </strong>The card_id generated by the system if customer chooses to save card. Applicable only for Checkout transactions with a customer id.
	CardId string `json:"card_id,omitempty"`
	// <strong>Format: </strong>Variable length, up to 16 AN<br><strong>Description: </strong>Masked card number used in transaction. 
	CardNumber string `json:"card_number,omitempty"`
	// <strong>Format: </strong>Variable length, in YYYY-MM-DD HH:MM:ss format<br><strong>Description: </strong>Transaction time. All times are pacific. 
	TranTime string `json:"tran_time,omitempty"`
	// <strong>Format: </strong>Variable length, up to 12,2 N<br><strong>Description: </strong>Transaction amount.
	AmtTran float64 `json:"amt_tran,omitempty"`
	// <strong>Format: </strong>Variable length, up to 12,2 N<br><strong>Description: </strong>Surcharge fee amount if the transaction includes a surcharge.
	AmtTranFee float64 `json:"amt_tran_fee,omitempty"`
	// <strong>Format: </strong>Variable length, up to 25 AN<br><strong>Description: </strong>Purchase ID of the transaction.
	PurchaseId string `json:"purchase_id,omitempty"`
	// <strong>Format: </strong>Fixed length, 1 AN<br><strong>Description: </strong>If CVV is sent, then the result from the card issuer will be returned in this field. Refer to <a href=\"/developer/api/reference#cvv2-result-codes\" target=\"_blank\">Payment Result Codes for CVV2</a> for valid CVV2 result codes.
	Cvv2Result string `json:"cvv2_result,omitempty"`
	// <strong>Format: </strong>Fixed length, 1 AN<br><strong>Description: </strong>If the billing_zip (and optionally the billing_street_addr1) is sent, then the result from the card issuer will be returned in this field. Refer to <a href=\"/developer/api/reference#avs-result-codes\" target=\"_blank\">Payment Result Codes for AVS</a> for
	AvsResult string `json:"avs_result,omitempty"`
	// <strong>Format: </strong>Variable length, up to 16 N<br><strong>Description: </strong>Duration of Gateway request in ms.
	Duration int32 `json:"duration,omitempty"`
	// <strong>Format: </strong>Variable length AN<br><strong>Description: </strong>HTTP status of the payment gateway request.
	HttpStatus int32 `json:"http_status,omitempty"`
	// <strong>Format: </strong>Variable length, up to 16 N<br><strong>Description: </strong>Unique ID assigned to a merchant.
	MerchantId int64 `json:"merchant_id,omitempty"`
	// <strong>Format: </strong>Variable length, up to 20 AN<br><strong>Description: </strong>The unique profile ID used in the payment gateway requests.
	ProfileId string `json:"profile_id,omitempty"`
	// <strong>Format: </strong>Variable length, up to 128 AN<br><strong>Description: </strong>This field contains a merchant provided reference value that will be stored with the transaction data and will be included with the transaction data reported in the Merchant Manager. This value will also be attached to any lifecycle transactions (e.g. retrieval requests and chargebacks) that may occur.
	MerchRefNum string `json:"merch_ref_num,omitempty"`
	// <strong>Format: </strong>Fixed length, 3 AN<br><strong>Description: </strong>Numeric currency code of the transaction. Refer to <a href=\"/developer/api/reference#country-codes\" target=\"_blank\">Country Codes</a> for a list of currency codes.
	TranCurrency string `json:"tran_currency,omitempty"`
	// <strong>Format: </strong>Variable length, up to 25 AN<br><strong>Description: </strong>DBA name associated with the profile.
	DbaName string `json:"dba_name,omitempty"`
	// <strong>Format: </strong>Variable length, up to 32 AN<br><strong>Description: </strong>Unique ID established by merchant to identify a customer. Customer ID is included in the response if it is sent in the request. 
	CustomerId string `json:"customer_id,omitempty"`
	Customer *CheckoutCustomer `json:"customer,omitempty"`
	OptionInfo *CheckoutOptionFields `json:"option_info,omitempty"`
	// <br><strong>Description: </strong>True for recurring checkout, false for non recurring checkout.
	Recurring bool `json:"recurring,omitempty"`
}
