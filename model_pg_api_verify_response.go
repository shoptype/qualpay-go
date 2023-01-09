/*
 * APIs
 *
 * This document describes the Platform and PG API.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type PgApiVerifyResponse struct {
	// <strong>Format: </strong>Variable length, up to 16 AN<br><strong>Description: </strong>Cardholder's card number (truncated).
	CardNumber string `json:"card_number,omitempty"`
	// <strong>Format: </strong>Fixed length 4 AN<br><strong>Description: </strong>Recurring transaction advice for MasterCard authorizations.  Refer to <a href=\"/developer/api/reference#mastercard-merchant-advice-codes\" target=\"_blank\">MasterCard Merchant Advice Codes</a> for possible values.
	MerchantAdviceCode string `json:"merchant_advice_code,omitempty"`
	// <strong>Format: </strong>Fixed length, 32 AN<br><strong>Description: </strong>32-byte unique identifier generated by the payment gateway, returned in all valid responses.
	PgId string `json:"pg_id,omitempty"`
	// <strong>Format: </strong>Fixed length, 3 AN<br><strong>Description: </strong>Response code from the payment gateway. \"000\" indicates success. Refer to <a href=\"/developer/api/reference#payment-gateway-response-codes\" target=\"_blank\">Payment Gateway Response Codes</a> for entire list of Payment Gateway Response Codes.
	Rcode string `json:"rcode,omitempty"`
	// <strong>Format: </strong>Variable length<br><strong>Description: </strong>Response text from the payment gateway.
	Rmsg string `json:"rmsg,omitempty"`
	// <strong>Format: </strong>Fixed length, 32 AN<br><strong>Description: </strong>32-byte value returned after successful tokenize request or when an authorization or sale transaction requests tokenization of the cardholder data.
	CardId string `json:"card_id,omitempty"`
	// <strong>Format: </strong>Fixed length, 6 AN<br><strong>Description: </strong>Card issuer authorization code returned on successful authorization request.
	AuthCode string `json:"auth_code,omitempty"`
	// <strong>Format: </strong>Fixed length, 1 AN<br><strong>Description: </strong>AVS result from card issuer (if avs_zip and optionally avs_address were provided in the request). Refer to <a href=\"/developer/api/reference#avs-result-codes\" target=\"_blank\">Payment Result Codes for AVS</a> for possible values.
	AuthAvsResult string `json:"auth_avs_result,omitempty"`
	// <strong>Format: </strong>Fixed length, 1 AN<br><strong>Description: </strong>CVV2 result from card issuer (if CVV2 data was sent in the request). Refer to <a href=\"/developer/api/reference#cvv2-result-codes\" target=\"_blank\">Payment Result Codes for CVV2</a> for possible values.
	AuthCvv2Result string `json:"auth_cvv2_result,omitempty"`
	// <strong>Format: </strong>Variable length<br><strong>Description: </strong>Data that was provided in echo_fields in the request.
	EchoFields string `json:"echo_fields,omitempty"`
}
