# PgApiTransactionRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | **int64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 12 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Unique identifier on our system. | [default to null]
**AmtConvenienceFee** | **float64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 8,2 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Amount of convenience fee. A convenience fee is a fee charged to your customer for the \&quot;convenience\&quot; of being able to pay using an alternative payment channel outside your merchant&#x27;s customary payment channel. Must be a flat/fixed fee amount per transaction. This field tracks the convenience fee amount for display purposes, but the amount of the fee must be included in amt_tran. | [optional] [default to null]
**AmtFbo** | **float64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 12,2 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Total amount of transaction to be transferred to the \&quot;for benefit of\&quot; (FBO) account. | [optional] [default to null]
**AmtTax** | **float64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 12,2 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Amount of sales tax included in the total transaction amount. This field tracks the tax amount for display and interchange purposes, but the amount of the tax must be included in amt_tran.&lt;br&gt;&lt;strong&gt;Conditional Requirement: &lt;/strong&gt;Required for Level 2 and Level 3 interchange qualification. | [optional] [default to null]
**AmtTran** | **float64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 12,2 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Total amount of transaction including sales tax (amt_tax), convenience fee (amt_convenience_fee), and/or surcharge (amt_tran_fee) if applicable. | [optional] [default to null]
**AmtTranFee** | **float64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 8,2 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Amount of transaction surcharge fee. A surcharge is a fee added to the cost of a purchase for the \&quot;privilege\&quot; of using a credit card instead of another form of payment, and can be a percentage of the transaction amount or fixed amount of up to 4% of amt_tran. This field tracks the surcharge amount of the transaction for display purposes, but the amount of the fee must be included in amt_tran. | [optional] [default to null]
**AuthCode** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 6 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;This field should contain the 6-character authorization code that was received during a voice or Automated Response Unit(ARU) authorization for force request type. This is field is applicable to only force request type.&lt;br&gt;&lt;strong&gt;Conditional Requirement: &lt;/strong&gt;This field is required in force request type. | [optional] [default to null]
**AvsAddress** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 20 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Street address of the cardholder. If present, it will be included in the authorization request sent to the issuing bank. | [optional] [default to null]
**AvsZip** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 9 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Zip code of the cardholder. If present, it will be included in the authorization request sent to the issuing bank.&lt;br&gt;&lt;strong&gt;Conditional Requirement: &lt;/strong&gt;This field is required if avs_address is present. | [optional] [default to null]
**CardId** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 32 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Card ID received from a tokenization request. The card_id may be used in place of a card number or card swipe.&lt;br&gt;&lt;strong&gt;Conditional Requirement: &lt;/strong&gt;Refer to &lt;a href&#x3D;\&quot;/developer/api/reference#card-source-conditional-requirements\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Card or Bank Account Data Sources and Conditional Requirements&lt;/a&gt; | [optional] [default to null]
**CardNumber** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 19 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Cardholder&#x27;s card number.&lt;br&gt;&lt;strong&gt;Conditional Requirement: &lt;/strong&gt;Refer to &lt;a href&#x3D;\&quot;/developer/api/reference#card-source-conditional-requirements\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Card or Bank Account Data Sources and Conditional Requirements&lt;/a&gt; | [optional] [default to null]
**CardSwipe** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 79 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Contains either track 1 or track 2 magnetic stripe data. If the magnetic stripe reader provides both track 1 and track 2 data in a single read, it is the responsibility of the implementer to send data for only one of the two tracks.&lt;br&gt;&lt;strong&gt;Conditional Requirement: &lt;/strong&gt;Refer to &lt;a href&#x3D;\&quot;/developer/api/reference#card-source-conditional-requirements\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Card or Bank Account Data Sources and Conditional Requirements&lt;/a&gt; | [optional] [default to null]
**CardholderName** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 64 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;When provided in a tokenize request, the cardholder name will be stored in the Card Vault along with the cardholder card number and expiration date. | [optional] [default to null]
**Cavv3ds** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 28 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Base 64 encoded CAVV returned from the merchant’s third-party 3-D Secure Merchant Plug-in (MPI). Use for Visa 3D Secure transactions. | [optional] [default to null]
**ClientIp** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Client IP address. | [optional] [default to null]
**Customer** | [***Customer**](Customer.md) |  | [optional] [default to null]
**CustomerCode** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 17 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Reference code supplied by the cardholder to the merchant. | [optional] [default to null]
**CustomerEmail** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;&lt;strong&gt;[Deprecated use email_address]&lt;/strong&gt; Comma-separated list of e-mail addresses to which a receipt should be sent. | [optional] [default to null]
**CustomerId** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 32 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Customer ID value established by the merchant. The customer_id may be used in place of a card number in requests requiring cardholder account data. When used with a card_id or card_number or card_swipe, the request will be tied to the customer_id in our reporting. &lt;br&gt;&lt;strong&gt;Conditional Requirement: &lt;/strong&gt;Refer to &lt;a href&#x3D;\&quot;/developer/api/reference#card-source-conditional-requirements\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Card or Bank Account Data Sources and Conditional Requirements&lt;/a&gt; | [optional] [default to null]
**Cvv2** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 4 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;CVV2 or CID value from the signature panel on the back of the cardholder&#x27;s card. If present during a request that requires authorization, the value will be sent to the issuer for validation. | [optional] [default to null]
**DbaName** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 21 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;When the merchant has been authorized to send dynamic DBA information, this field will contain the DBA name used by Qulapay in the authorization and clearing messages.&lt;br&gt;Note: the payment gateway will automatically add a prefix plus an asterisk (*) to the dba_name value. For example, if the prefix is ABC and the dba_name value is SHOE CO, the DBA name will show as \&quot;ABC*SHOE CO\&quot; on the cardholder&#x27;s credit card statement. | [optional] [default to null]
**DeveloperId** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 32 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Use to indicate which company developed the integration or the name of the payment solution that is connected to our service. Suggested usage is softwareABCv1.0 or companyXYZv2.0.  | [optional] [default to null]
**DuplicateSeconds** | **int64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 5 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Duplicate transaction window in seconds. We will reject any transactions after a successful transaction within the duplicate_seconds window with a duplicate Account Number and optionally Purchase ID or, and, Merchant Reference Number. This value overrides any value set for a merchant on Merchant Manager. | [optional] [default to null]
**EchoFields** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;This field contains a JSON array of field data that will be echoed back in the response message. | [optional] [default to null]
**EmailAddress** | **[]string** |  AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;An array of email addresses to which the transaction receipt should be sent to.  | [optional] [default to null]
**EmailReceipt** | **bool** | &lt;br&gt;&lt;strong&gt;Default: &lt;/strong&gt;false&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;When this field is provided and set to true, a customer_email must also be provided. When these two fields are provided, a transaction receipt will be sent via e-mail to the address(es) provided in the customer_email field. | [optional] [default to null]
**EmvTranId** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 36 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Base64 encoded MasterCard UCAF Transaction ID returned from the merchant’s third-party 3D Secure Merchant Plug-in (MPI). Use for MasterCard 3-D Secure transactions. | [optional] [default to null]
**ExpDate** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 4 N, MMYY format&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Expiration date of cardholder card number.  When card_id or customer_id is present in the request this field may also be present; if it is not, then the expiration date from the Card Vault will be used.&lt;br&gt;&lt;strong&gt;Conditional Requirement: &lt;/strong&gt;Refer to &lt;a href&#x3D;\&quot;/developer/api/reference#card-source-conditional-requirements\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Card or Bank Account Data Sources and Conditional Requirements&lt;/a&gt; | [optional] [default to null]
**FboId** | **int64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 16 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;For Benefit Of (FBO) merchant account identifier on our system. Contact customer support to obtain your FBO information. | [optional] [default to null]
**LineItems** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;JSON array of JSON objects. Each object represents a single line item detail element related to the transaction. Each detail element has required subfields: &lt;br&gt;quantity (7N)&lt;br&gt; description (26AN)&lt;br&gt; unit_of_measure (12AN)&lt;br&gt; product_code (12AN) - cannot be all zeroes&lt;br&gt; debit_credit_ind (1 AN)&lt;br&gt; unit_cost (12,2N)&lt;br&gt; Optional subfields: &lt;br&gt;type_of_supply (2AN) - visa only&lt;br&gt;commodity_code - visa only(12AN)&lt;br&gt;&lt;strong&gt;Conditional Requirement: &lt;/strong&gt; This field is required for Level 3 interchange qualification. | [optional] [default to null]
**LocId** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 4 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;When a merchant has more than one location using the same currency, this value is used to specify the specific location for this request. | [optional] [default to null]
**McUcafData** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 32 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Base64 encoded MasterCard UCAF Field Data returned from the merchant’s third-party 3D Secure Merchant Plug-in (MPI). Use for MasterCard 3-D Secure transactions. | [optional] [default to null]
**McUcafInd** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 1 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;MasterCard UCAF Collection Indicator returned from the merchant’s third-party 3-D Secure Merchant Plug-in (MPI). Use for MasterCard 3-D Secure transactions. | [optional] [default to null]
**MerchRefNum** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 128 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Merchant provided reference value that will be stored with the transaction data and included with transaction data in reports within Merchant Manager. This value will also be attached to any lifecycle transactions (e.g. retrieval requests and chargebacks) that may occur. | [optional] [default to null]
**MotoEcommInd** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 1 N&lt;br&gt;&lt;strong&gt;Default: &lt;/strong&gt;7&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Indicates type of MOTO transaction: &lt;ul&gt;&lt;li&gt;0 &#x3D; Card Present (not MOTO/e-Commerce)&lt;/li&gt;&lt;li&gt; 1 &#x3D; One Time MOTO transaction&lt;/li&gt;&lt;li&gt;2 &#x3D; Recurring &lt;/li&gt;&lt;li&gt;3 &#x3D; Installment &lt;/li&gt;&lt;li&gt;5 &#x3D; Full 3D-Secure transaction&lt;/li&gt;&lt;li&gt;6 &#x3D; Merchant 3D-Secure transaction&lt;/li&gt;&lt;li&gt;7 &#x3D; e-Commerce Channel Encrypted (SSL)&lt;/li&gt;&lt;/ul&gt; | [optional] [default to null]
**PartialAuth** | **bool** | &lt;br&gt;&lt;strong&gt;Default: &lt;/strong&gt;false&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;This field must be present and set to a value of &#x27;true&#x27; in order for the request to allow for approval of a partial amount. This would be used to allow a merchant to accept a partial payment from pre-paid or debit cards. When only part of the requested amount is available, the response code will be 010 and the amt_tran field in the response will contain the amount that was approved. A second sale request  on a different card is needed  to capture the remaining amount. Applicable to auth and sale request types. | [optional] [default to null]
**PayloadGooglePay** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Google Pay payload | [optional] [default to null]
**PayloadApplePay** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Apple Pay payload | [optional] [default to null]
**PgId** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 32 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;PG ID of previously authorized transaction. This field is required when sending a capture, refund, or void request. | [optional] [default to null]
**ProfileId** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 20 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Explicitly identifies which Payment Gateway profile should be used for the request. | [optional] [default to null]
**PurchaseId** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 25 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Purchase Identifier (also referred to as the invoice number generated by the merchant).&lt;br&gt;&lt;strong&gt;Conditional Requirement: &lt;/strong&gt; This field is required for Level 2 and Level 3 interchange qualification. | [optional] [default to null]
**ReportData** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;This field contains a JSON array of field data that will be included with the transaction data reported in Merchant Manager. | [optional] [default to null]
**SessionId** | **string** | INTERNAL USE ONLY. | [optional] [default to null]
**SubscriptionId** | **int64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 10 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Identifies the recurring subscription that applies to this transaction. | [optional] [default to null]
**Tokenize** | **bool** | &lt;br&gt;&lt;strong&gt;Default: &lt;/strong&gt;false&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;In an authorization, credit, force, sale, or verify request the merchant can set tokenize to \&quot;true\&quot; and the payment gateway will store the cardholder data in the Card Vault and provide a card_id in the response. If the card_number or card_id in the request is already in the Card Vault, this flag instructs the payment gateway to update the associated data (e.g. avs_address, avs_zip, exp_date) if present.&lt;br&gt;&lt;strong&gt;Conditional Requirement: &lt;/strong&gt;Refer to &lt;a href&#x3D;\&quot;/developer/api/reference#card-source-conditional-requirements\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Card or Bank Account Data Sources and Conditional Requirements&lt;/a&gt; | [optional] [default to null]
**TranCurrency** | **int32** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 3 N&lt;br&gt;&lt;strong&gt;Default: &lt;/strong&gt;840&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;ISO numeric currency code for the transaction. Refer to &lt;a href&#x3D;\&quot;/developer/api/reference#country-codes\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Country Codes&lt;/a&gt; for a list of currency codes. | [optional] [default to null]
**UserId** | **int64** | INTERNAL USE ONLY. | [optional] [default to null]
**Xid3ds** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 28 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Base64 encoded transaction ID (XID) returned from the merchant’s third-party 3D Secure Merchant Plug-in (MPI). Use for Visa 3-D Secure transactions. | [optional] [default to null]
**DbaSuffix** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 9 AN&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;For use by merchants using negative option marketing.  This field must be used in the first transaction at the conclusion of the free or reduced trial. This suffix will be appended to the end of your DBA and the result will appear on the cardholder statement. (If your DBA and suffix contain more that 25 characters, your DBA will be truncated.) Possible values are: &lt;ul&gt;&lt;li&gt;END DSCNT&lt;/li&gt;&lt;li&gt;END OFFER&lt;/li&gt;&lt;li&gt;END PROMO&lt;/li&gt;&lt;li&gt;END TRIAL&lt;/li&gt;&lt;/ul&gt; | [optional] [default to null]
**DdaNumber** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 17 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Owner&#x27;s account number at the bank. Applicable for ACH payments.&lt;br&gt;&lt;strong&gt;Conditional Requirement: &lt;/strong&gt;Refer to &lt;a href&#x3D;\&quot;/developer/api/reference#card-source-conditional-requirements\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Card or Bank Account Data Sources and Conditional Requirements&lt;/a&gt; | [optional] [default to null]
**TrNumber** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 9 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Bank transit/routing number. Applicable for ACH payments.&lt;br&gt;&lt;strong&gt;Conditional Requirement: &lt;/strong&gt;Refer to &lt;a href&#x3D;\&quot;/developer/api/reference#card-source-conditional-requirements\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Card or Bank Account Data Sources and Conditional Requirements&lt;/a&gt; | [optional] [default to null]
**TypeId** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 1 AN&lt;br&gt;&lt;strong&gt;Default: &lt;/strong&gt;C&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Bank Account Type. Applicable for ACH payments. Possible values are: &lt;ul&gt;&lt;li&gt;C &#x3D; Personal checking account&lt;/li&gt;&lt;li&gt;S &#x3D; Personal savings account&lt;/li&gt;&lt;li&gt;K &#x3D; Business checking account&lt;/li&gt;&lt;li&gt;V &#x3D; Business savings account&lt;/li&gt;&lt;/ul&gt; | [optional] [default to null]
**VendorId** | **int64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 12 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Identifies the vendor to which this capture request applies. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
