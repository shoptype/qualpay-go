/*
 * APIs
 *
 * This document describes the Platform and PG API.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type PgApiVerifyRequest struct {
	// <strong>Format: </strong>Variable length, up to 12 N<br><strong>Description: </strong>Unique identifier on our system.
	MerchantId int64 `json:"merchant_id"`
	// <strong>Format: </strong>Variable length, up to 20 AN<br><strong>Description: </strong>Street address of the cardholder. If present, it will be included in the authorization request sent to the issuing bank.
	AvsAddress string `json:"avs_address,omitempty"`
	// <strong>Format: </strong>Variable length, up to 9 N<br><strong>Description: </strong>Zip code of the cardholder. If present, it will be included in the authorization request sent to the issuing bank.<br><strong>Conditional Requirement: </strong>This field is required if avs_address is present.
	AvsZip string `json:"avs_zip,omitempty"`
	// <strong>Format: </strong>Variable length, up to 32 AN<br><strong>Description: </strong>Card ID received from a tokenization request. The card_id may be used in place of a card number or card swipe.<br><strong>Conditional Requirement: </strong>Refer to <a href=\"/developer/api/reference#card-source-conditional-requirements\" target=\"_blank\">Card or Bank Account Data Sources and Conditional Requirements</a>
	CardId string `json:"card_id,omitempty"`
	// <strong>Format: </strong>Variable length, up to 19 N<br><strong>Description: </strong>Cardholder's card number. <br><strong>Conditional Requirement: </strong>Refer to <a href=\"/developer/api/reference#card-source-conditional-requirements\" target=\"_blank\">Card or Bank Account Data Sources and Conditional Requirements</a>
	CardNumber string `json:"card_number,omitempty"`
	// <strong>Format: </strong>Variable length, up to 79 AN<br><strong>Description: </strong>Contains either track 1 or track 2 data magnetic stripe data. If the magnetic stripe reader provides both track 1 and track 2 data in a single read, it is the responsibility of the implementer to send data for only one of the two tracks.<br><strong>Conditional Requirement: </strong>Refer to <a href=\"/developer/api/reference#card-source-conditional-requirements\" target=\"_blank\">Card or Bank Account Data Sources and Conditional Requirements</a>
	CardSwipe string `json:"card_swipe,omitempty"`
	// <strong>Format: </strong>Variable length, up to 64 AN<br><strong>Description: </strong>When provided in a tokenize request, the cardholder name will be stored in the Card Vault along with the cardholder card number and expiration date.
	CardholderName string `json:"cardholder_name,omitempty"`
	// <strong>Format: </strong>Variable length AN<br><strong>Description: </strong>Client IP address.
	ClientIp string `json:"client_ip,omitempty"`
	Customer *Customer `json:"customer,omitempty"`
	// <strong>Format: </strong>Variable length, up to 17 AN<br><strong>Description: </strong>Reference code supplied by the cardholder to the merchant.
	CustomerCode string `json:"customer_code,omitempty"`
	// <strong>Format: </strong>Variable length, up to 4 N<br><strong>Description: </strong>CVV2 or CID value from the signature panel on the back of the cardholder's card. If present during a request that requires authorization, the value will be sent to the issuer for validation.
	Cvv2 string `json:"cvv2,omitempty"`
	// <strong>Format: </strong>Variable length, up to 32 AN<br><strong>Description: </strong>Use to indicate which company developed the integration or the name of the payment solution that is connected to us. Suggested usage is softwareABCv1.0 or companyXYZv2.0. 
	DeveloperId string `json:"developer_id,omitempty"`
	// <strong>Format: </strong>Variable length<br><strong>Description: </strong>This field contains a JSON array of field data that will be echoed back in the response message.
	EchoFields string `json:"echo_fields,omitempty"`
	// <strong>Format: </strong>Fixed length, 4 N, MMYY format<br><strong>Description: </strong>Expiration date of cardholder card number. Required when the field card_number is present. If card_swipe is present in the request, this field must NOT be present. When card_id or customer_id is present in the request this field may also be present; if it is not, then the expiration date from the Card Vault will be used.<br><strong>Conditional Requirement: </strong>Refer to <a href=\"/developer/api/reference#card-source-conditional-requirements\" target=\"_blank\">Card or Bank Account Data Sources and Conditional Requirements</a>
	ExpDate string `json:"exp_date,omitempty"`
	// <strong>Format: </strong>Variable length, up to 4 N<br><strong>Description: </strong>When a merchant has more than one location using the same currency, this value is used to specify the specific location for this request.
	LocId string `json:"loc_id,omitempty"`
	// <strong>Format: </strong>Variable length, up to 128 AN<br><strong>Description: </strong>Merchant provided reference value that will be stored with the transaction data and included with transaction data in reports within Merchant Manager. This value will also be attached to any lifecycle transactions (e.g. retrieval requests and chargebacks) that may occur.
	MerchRefNum string `json:"merch_ref_num,omitempty"`
	// <strong>Format: </strong>Variable length<br><strong>Description: </strong>Google Pay payload
	PayloadGooglePay string `json:"payload_google_pay,omitempty"`
	// <strong>Format: </strong>Variable length<br><strong>Description: </strong>Apple Pay payload
	PayloadApplePay string `json:"payload_apple_pay,omitempty"`
	// <strong>Format: </strong>Fixed length, 20 N<br><strong>Description: </strong>Explicitly identifies which Payment Gateway profile should be used for the request.
	ProfileId string `json:"profile_id,omitempty"`
	// <strong>Format: </strong>Variable length<br><strong>Description: </strong>This field contains a JSON array of field data that will be included with the transaction data reported in Merchant Manager.
	ReportData string `json:"report_data,omitempty"`
	// INTERNAL USE ONLY.
	SessionId string `json:"session_id,omitempty"`
	// <br><strong>Default: </strong>false<br><strong>Description: </strong>In an authorization, credit, force, sale, or verify request the merchant can set tokenize to \"true\" and the payment gateway will store the cardholder data in the Card Vault and provide a card_id in the response. If the card_number or card_id in the request is already in the Card Vault, this flag instructs the payment gateway to update the associated data (e.g. avs_address, avs_zip, exp_date) if present.
	Tokenize bool `json:"tokenize,omitempty"`
	// INTERNAL USE ONLY.
	UserId int64 `json:"user_id,omitempty"`
	// <strong>Format: </strong>Variable length, up to 17 N<br><strong>Description: </strong>Owner's account number at the bank. Applicable for ACH payments.<br><strong>Conditional Requirement: </strong>Refer to <a href=\"/developer/api/reference#card-source-conditional-requirements\" target=\"_blank\">Card or Bank Account Data Sources and Conditional Requirements</a>
	DdaNumber string `json:"dda_number,omitempty"`
	//  AN<br><strong>Description: </strong>An array of account holder email addresses.
	EmailAddress []string `json:"email_address,omitempty"`
	// <strong>Format: </strong>Fixed length, 9 N<br><strong>Description: </strong>Bank transit/routing number. Applicable for ACH payments.<br><strong>Conditional Requirement: </strong>Refer to <a href=\"/developer/api/reference#card-source-conditional-requirements\" target=\"_blank\">Card or Bank Account Data Sources and Conditional Requirements</a>
	TrNumber string `json:"tr_number,omitempty"`
	// <strong>Format: </strong>Fixed length, 1 AN<br><strong>Default: </strong>C<br><strong>Description: </strong>Bank Account Type. Applicable for ACH payments. Possible values are: <ul><li>C = Personal checking account</li><li>S = Personal savings account</li><li>K = Business checking account</li><li>V = Business savings account</li></ul>
	TypeId string `json:"type_id,omitempty"`
	// <strong>Format: </strong>Fixed length, 1 N<br><strong>Default: </strong>7<br><strong>Description: </strong>Indicates type of MOTO transaction: <ul><li>0 = Card Present (not MOTO/e-Commerce)</li><li> 1 = One Time MOTO transaction</li><li>2 = Recurring </li><li>3 = Installment </li><li>5 = Full 3D-Secure transaction</li><li>6 = Merchant 3D-Secure transaction</li><li>7 = e-Commerce Channel Encrypted (SSL)</li></ul>
	MotoEcommInd string `json:"moto_ecomm_ind,omitempty"`
}
