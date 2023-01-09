/*
 * APIs
 *
 * This document describes the Platform and PG API.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package qualpay

type PgApiTransactionRequest struct {
	// <strong>Format: </strong>Variable length, up to 12 N<br><strong>Description: </strong>Unique identifier on our system.
	MerchantId int64 `json:"merchant_id"`
	// <strong>Format: </strong>Variable length, up to 8,2 N<br><strong>Description: </strong>Amount of convenience fee. A convenience fee is a fee charged to your customer for the \"convenience\" of being able to pay using an alternative payment channel outside your merchant's customary payment channel. Must be a flat/fixed fee amount per transaction. This field tracks the convenience fee amount for display purposes, but the amount of the fee must be included in amt_tran.
	AmtConvenienceFee float64 `json:"amt_convenience_fee,omitempty"`
	// <strong>Format: </strong>Variable length, up to 12,2 N<br><strong>Description: </strong>Total amount of transaction to be transferred to the \"for benefit of\" (FBO) account.
	AmtFbo float64 `json:"amt_fbo,omitempty"`
	// <strong>Format: </strong>Variable length, up to 12,2 N<br><strong>Description: </strong>Amount of sales tax included in the total transaction amount. This field tracks the tax amount for display and interchange purposes, but the amount of the tax must be included in amt_tran.<br><strong>Conditional Requirement: </strong>Required for Level 2 and Level 3 interchange qualification.
	AmtTax float64 `json:"amt_tax,omitempty"`
	// <strong>Format: </strong>Variable length, up to 12,2 N<br><strong>Description: </strong>Total amount of transaction including sales tax (amt_tax), convenience fee (amt_convenience_fee), and/or surcharge (amt_tran_fee) if applicable.
	AmtTran float64 `json:"amt_tran,omitempty"`
	// <strong>Format: </strong>Variable length, up to 8,2 N<br><strong>Description: </strong>Amount of transaction surcharge fee. A surcharge is a fee added to the cost of a purchase for the \"privilege\" of using a credit card instead of another form of payment, and can be a percentage of the transaction amount or fixed amount of up to 4% of amt_tran. This field tracks the surcharge amount of the transaction for display purposes, but the amount of the fee must be included in amt_tran.
	AmtTranFee float64 `json:"amt_tran_fee,omitempty"`
	// <strong>Format: </strong>Fixed length, 6 AN<br><strong>Description: </strong>This field should contain the 6-character authorization code that was received during a voice or Automated Response Unit(ARU) authorization for force request type. This is field is applicable to only force request type.<br><strong>Conditional Requirement: </strong>This field is required in force request type.
	AuthCode string `json:"auth_code,omitempty"`
	// <strong>Format: </strong>Variable length, up to 20 AN<br><strong>Description: </strong>Street address of the cardholder. If present, it will be included in the authorization request sent to the issuing bank.
	AvsAddress string `json:"avs_address,omitempty"`
	// <strong>Format: </strong>Variable length, up to 9 AN<br><strong>Description: </strong>Zip code of the cardholder. If present, it will be included in the authorization request sent to the issuing bank.<br><strong>Conditional Requirement: </strong>This field is required if avs_address is present.
	AvsZip string `json:"avs_zip,omitempty"`
	// <strong>Format: </strong>Fixed length, 32 AN<br><strong>Description: </strong>Card ID received from a tokenization request. The card_id may be used in place of a card number or card swipe.<br><strong>Conditional Requirement: </strong>Refer to <a href=\"/developer/api/reference#card-source-conditional-requirements\" target=\"_blank\">Card or Bank Account Data Sources and Conditional Requirements</a>
	CardId string `json:"card_id,omitempty"`
	// <strong>Format: </strong>Variable length, up to 19 N<br><strong>Description: </strong>Cardholder's card number.<br><strong>Conditional Requirement: </strong>Refer to <a href=\"/developer/api/reference#card-source-conditional-requirements\" target=\"_blank\">Card or Bank Account Data Sources and Conditional Requirements</a>
	CardNumber string `json:"card_number,omitempty"`
	// <strong>Format: </strong>Variable length, up to 79 AN<br><strong>Description: </strong>Contains either track 1 or track 2 magnetic stripe data. If the magnetic stripe reader provides both track 1 and track 2 data in a single read, it is the responsibility of the implementer to send data for only one of the two tracks.<br><strong>Conditional Requirement: </strong>Refer to <a href=\"/developer/api/reference#card-source-conditional-requirements\" target=\"_blank\">Card or Bank Account Data Sources and Conditional Requirements</a>
	CardSwipe string `json:"card_swipe,omitempty"`
	// <strong>Format: </strong>Variable length, up to 64 AN<br><strong>Description: </strong>When provided in a tokenize request, the cardholder name will be stored in the Card Vault along with the cardholder card number and expiration date.
	CardholderName string `json:"cardholder_name,omitempty"`
	// <strong>Format: </strong>Fixed length, 28 AN<br><strong>Description: </strong>Base 64 encoded CAVV returned from the merchant’s third-party 3-D Secure Merchant Plug-in (MPI). Use for Visa 3D Secure transactions.
	Cavv3ds string `json:"cavv_3ds,omitempty"`
	// <strong>Format: </strong>Variable length AN<br><strong>Description: </strong>Client IP address.
	ClientIp string    `json:"client_ip,omitempty"`
	Customer *Customer `json:"customer,omitempty"`
	// <strong>Format: </strong>Variable length, up to 17 AN<br><strong>Description: </strong>Reference code supplied by the cardholder to the merchant.
	CustomerCode string `json:"customer_code,omitempty"`
	// <strong>Format: </strong>Variable length<br><strong>Description: </strong><strong>[Deprecated use email_address]</strong> Comma-separated list of e-mail addresses to which a receipt should be sent.
	CustomerEmail string `json:"customer_email,omitempty"`
	// <strong>Format: </strong>Variable length, up to 32 AN<br><strong>Description: </strong>Customer ID value established by the merchant. The customer_id may be used in place of a card number in requests requiring cardholder account data. When used with a card_id or card_number or card_swipe, the request will be tied to the customer_id in our reporting. <br><strong>Conditional Requirement: </strong>Refer to <a href=\"/developer/api/reference#card-source-conditional-requirements\" target=\"_blank\">Card or Bank Account Data Sources and Conditional Requirements</a>
	CustomerId string `json:"customer_id,omitempty"`
	// <strong>Format: </strong>Variable length, up to 4 N<br><strong>Description: </strong>CVV2 or CID value from the signature panel on the back of the cardholder's card. If present during a request that requires authorization, the value will be sent to the issuer for validation.
	Cvv2 string `json:"cvv2,omitempty"`
	// <strong>Format: </strong>Variable length, up to 21 AN<br><strong>Description: </strong>When the merchant has been authorized to send dynamic DBA information, this field will contain the DBA name used by Qulapay in the authorization and clearing messages.<br>Note: the payment gateway will automatically add a prefix plus an asterisk (*) to the dba_name value. For example, if the prefix is ABC and the dba_name value is SHOE CO, the DBA name will show as \"ABC*SHOE CO\" on the cardholder's credit card statement.
	DbaName string `json:"dba_name,omitempty"`
	// <strong>Format: </strong>Variable length, up to 32 AN<br><strong>Description: </strong>Use to indicate which company developed the integration or the name of the payment solution that is connected to our service. Suggested usage is softwareABCv1.0 or companyXYZv2.0.
	DeveloperId string `json:"developer_id,omitempty"`
	// <strong>Format: </strong>Variable length, up to 5 N<br><strong>Description: </strong>Duplicate transaction window in seconds. We will reject any transactions after a successful transaction within the duplicate_seconds window with a duplicate Account Number and optionally Purchase ID or, and, Merchant Reference Number. This value overrides any value set for a merchant on Merchant Manager.
	DuplicateSeconds int64 `json:"duplicate_seconds,omitempty"`
	// <strong>Format: </strong>Variable length<br><strong>Description: </strong>This field contains a JSON array of field data that will be echoed back in the response message.
	EchoFields string `json:"echo_fields,omitempty"`
	//  AN<br><strong>Description: </strong>An array of email addresses to which the transaction receipt should be sent to.
	EmailAddress []string `json:"email_address,omitempty"`
	// <br><strong>Default: </strong>false<br><strong>Description: </strong>When this field is provided and set to true, a customer_email must also be provided. When these two fields are provided, a transaction receipt will be sent via e-mail to the address(es) provided in the customer_email field.
	EmailReceipt bool `json:"email_receipt,omitempty"`
	// <strong>Format: </strong>Variable length, up to 36 AN<br><strong>Description: </strong>Base64 encoded MasterCard UCAF Transaction ID returned from the merchant’s third-party 3D Secure Merchant Plug-in (MPI). Use for MasterCard 3-D Secure transactions.
	EmvTranId string `json:"emv_tran_id,omitempty"`
	// <strong>Format: </strong>Fixed length, 4 N, MMYY format<br><strong>Description: </strong>Expiration date of cardholder card number.  When card_id or customer_id is present in the request this field may also be present; if it is not, then the expiration date from the Card Vault will be used.<br><strong>Conditional Requirement: </strong>Refer to <a href=\"/developer/api/reference#card-source-conditional-requirements\" target=\"_blank\">Card or Bank Account Data Sources and Conditional Requirements</a>
	ExpDate string `json:"exp_date,omitempty"`
	// <strong>Format: </strong>Variable length, up to 16 N<br><strong>Description: </strong>For Benefit Of (FBO) merchant account identifier on our system. Contact customer support to obtain your FBO information.
	FboId int64 `json:"fbo_id,omitempty"`
	// <strong>Format: </strong>Variable length<br><strong>Description: </strong>JSON array of JSON objects. Each object represents a single line item detail element related to the transaction. Each detail element has required subfields: <br>quantity (7N)<br> description (26AN)<br> unit_of_measure (12AN)<br> product_code (12AN) - cannot be all zeroes<br> debit_credit_ind (1 AN)<br> unit_cost (12,2N)<br> Optional subfields: <br>type_of_supply (2AN) - visa only<br>commodity_code - visa only(12AN)<br><strong>Conditional Requirement: </strong> This field is required for Level 3 interchange qualification.
	LineItems string `json:"line_items,omitempty"`
	// <strong>Format: </strong>Variable length, up to 4 N<br><strong>Description: </strong>When a merchant has more than one location using the same currency, this value is used to specify the specific location for this request.
	LocId string `json:"loc_id,omitempty"`
	// <strong>Format: </strong>Variable length, up to 32 AN<br><strong>Description: </strong>Base64 encoded MasterCard UCAF Field Data returned from the merchant’s third-party 3D Secure Merchant Plug-in (MPI). Use for MasterCard 3-D Secure transactions.
	McUcafData string `json:"mc_ucaf_data,omitempty"`
	// <strong>Format: </strong>Fixed length, 1 AN<br><strong>Description: </strong>MasterCard UCAF Collection Indicator returned from the merchant’s third-party 3-D Secure Merchant Plug-in (MPI). Use for MasterCard 3-D Secure transactions.
	McUcafInd string `json:"mc_ucaf_ind,omitempty"`
	// <strong>Format: </strong>Variable length, up to 128 AN<br><strong>Description: </strong>Merchant provided reference value that will be stored with the transaction data and included with transaction data in reports within Merchant Manager. This value will also be attached to any lifecycle transactions (e.g. retrieval requests and chargebacks) that may occur.
	MerchRefNum string `json:"merch_ref_num,omitempty"`
	// <strong>Format: </strong>Fixed length, 1 N<br><strong>Default: </strong>7<br><strong>Description: </strong>Indicates type of MOTO transaction: <ul><li>0 = Card Present (not MOTO/e-Commerce)</li><li> 1 = One Time MOTO transaction</li><li>2 = Recurring </li><li>3 = Installment </li><li>5 = Full 3D-Secure transaction</li><li>6 = Merchant 3D-Secure transaction</li><li>7 = e-Commerce Channel Encrypted (SSL)</li></ul>
	MotoEcommInd string `json:"moto_ecomm_ind,omitempty"`
	// <br><strong>Default: </strong>false<br><strong>Description: </strong>This field must be present and set to a value of 'true' in order for the request to allow for approval of a partial amount. This would be used to allow a merchant to accept a partial payment from pre-paid or debit cards. When only part of the requested amount is available, the response code will be 010 and the amt_tran field in the response will contain the amount that was approved. A second sale request  on a different card is needed  to capture the remaining amount. Applicable to auth and sale request types.
	PartialAuth bool `json:"partial_auth,omitempty"`
	// <strong>Format: </strong>Variable length<br><strong>Description: </strong>Google Pay payload
	PayloadGooglePay string `json:"payload_google_pay,omitempty"`
	// <strong>Format: </strong>Variable length<br><strong>Description: </strong>Apple Pay payload
	PayloadApplePay string `json:"payload_apple_pay,omitempty"`
	// <strong>Format: </strong>Fixed length, 32 AN<br><strong>Description: </strong>PG ID of previously authorized transaction. This field is required when sending a capture, refund, or void request.
	PgId string `json:"pg_id,omitempty"`
	// <strong>Format: </strong>Fixed length, 20 N<br><strong>Description: </strong>Explicitly identifies which Payment Gateway profile should be used for the request.
	ProfileId string `json:"profile_id,omitempty"`
	// <strong>Format: </strong>Variable length, up to 25 AN<br><strong>Description: </strong>Purchase Identifier (also referred to as the invoice number generated by the merchant).<br><strong>Conditional Requirement: </strong> This field is required for Level 2 and Level 3 interchange qualification.
	PurchaseId string `json:"purchase_id,omitempty"`
	// <strong>Format: </strong>Variable length<br><strong>Description: </strong>This field contains a JSON array of field data that will be included with the transaction data reported in Merchant Manager.
	ReportData string `json:"report_data,omitempty"`
	// INTERNAL USE ONLY.
	SessionId string `json:"session_id,omitempty"`
	// <strong>Format: </strong>Variable length, up to 10 N<br><strong>Description: </strong>Identifies the recurring subscription that applies to this transaction.
	SubscriptionId int64 `json:"subscription_id,omitempty"`
	// <br><strong>Default: </strong>false<br><strong>Description: </strong>In an authorization, credit, force, sale, or verify request the merchant can set tokenize to \"true\" and the payment gateway will store the cardholder data in the Card Vault and provide a card_id in the response. If the card_number or card_id in the request is already in the Card Vault, this flag instructs the payment gateway to update the associated data (e.g. avs_address, avs_zip, exp_date) if present.<br><strong>Conditional Requirement: </strong>Refer to <a href=\"/developer/api/reference#card-source-conditional-requirements\" target=\"_blank\">Card or Bank Account Data Sources and Conditional Requirements</a>
	Tokenize bool `json:"tokenize,omitempty"`
	// <strong>Format: </strong>Fixed length, 3 N<br><strong>Default: </strong>840<br><strong>Description: </strong>ISO numeric currency code for the transaction. Refer to <a href=\"/developer/api/reference#country-codes\" target=\"_blank\">Country Codes</a> for a list of currency codes.
	TranCurrency int32 `json:"tran_currency,omitempty"`
	// INTERNAL USE ONLY.
	UserId int64 `json:"user_id,omitempty"`
	// <strong>Format: </strong>Fixed length, 28 AN<br><strong>Description: </strong>Base64 encoded transaction ID (XID) returned from the merchant’s third-party 3D Secure Merchant Plug-in (MPI). Use for Visa 3-D Secure transactions.
	Xid3ds string `json:"xid_3ds,omitempty"`
	// <strong>Format: </strong>Fixed length, 9 AN<br><strong>Description: </strong>For use by merchants using negative option marketing.  This field must be used in the first transaction at the conclusion of the free or reduced trial. This suffix will be appended to the end of your DBA and the result will appear on the cardholder statement. (If your DBA and suffix contain more that 25 characters, your DBA will be truncated.) Possible values are: <ul><li>END DSCNT</li><li>END OFFER</li><li>END PROMO</li><li>END TRIAL</li></ul>
	DbaSuffix string `json:"dba_suffix,omitempty"`
	// <strong>Format: </strong>Variable length, up to 17 N<br><strong>Description: </strong>Owner's account number at the bank. Applicable for ACH payments.<br><strong>Conditional Requirement: </strong>Refer to <a href=\"/developer/api/reference#card-source-conditional-requirements\" target=\"_blank\">Card or Bank Account Data Sources and Conditional Requirements</a>
	DdaNumber string `json:"dda_number,omitempty"`
	// <strong>Format: </strong>Fixed length, 9 N<br><strong>Description: </strong>Bank transit/routing number. Applicable for ACH payments.<br><strong>Conditional Requirement: </strong>Refer to <a href=\"/developer/api/reference#card-source-conditional-requirements\" target=\"_blank\">Card or Bank Account Data Sources and Conditional Requirements</a>
	TrNumber string `json:"tr_number,omitempty"`
	// <strong>Format: </strong>Fixed length, 1 AN<br><strong>Default: </strong>C<br><strong>Description: </strong>Bank Account Type. Applicable for ACH payments. Possible values are: <ul><li>C = Personal checking account</li><li>S = Personal savings account</li><li>K = Business checking account</li><li>V = Business savings account</li></ul>
	TypeId string `json:"type_id,omitempty"`
	// <strong>Format: </strong>Variable length, up to 12 N<br><strong>Description: </strong>Identifies the vendor to which this capture request applies.
	VendorId int64 `json:"vendor_id,omitempty"`
}
