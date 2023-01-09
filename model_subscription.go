/*
 * APIs
 *
 * This document describes the Platform and PG API.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Subscription struct {
	// <strong>Format: </strong>Variable length, up to 10 N<br><strong>Description: </strong>a generated ID that identifies a subscription. Save this id to manage the subscription. 
	SubscriptionId int64 `json:"subscription_id,omitempty"`
	// <strong>Format: </strong>Variable length, up to 16 N<br><strong>Description: </strong>Unique ID assigned to a merchant.
	MerchantId int64 `json:"merchant_id,omitempty"`
	// <strong>Format: </strong>Variable length, up to 32 AN<br><strong>Description: </strong>Customer id of the subscriber.
	CustomerId string `json:"customer_id,omitempty"`
	// <strong>Format: </strong>Fixed length, 32 AN<br><strong>Description: </strong>The payment method that will be used for the recurring charge. If empty, the recurring charges will be made using the customer's primary payment method. 
	CardId string `json:"card_id,omitempty"`
	// <strong>Format: </strong>Fixed length, 1 AN<br><strong>Description: </strong> Status of the subscription. Following are possible statuses:               <ul>              <li>A - Active</li>              <li>D - Complete</li>              <li>P - Paused</li>              <li>C - Cancelled</li>              <li>S - Suspended</li>              </ul>
	Status string `json:"status,omitempty"`
	// <strong>Format: </strong>Variable length, up to 20 AN<br><strong>Description: </strong>The profile ID to be used in payment gateway requests.
	ProfileId string `json:"profile_id,omitempty"`
	// <br><strong>Description: </strong>True for on-plan subscriptions, false for off-plan subscriptions.
	SubscriptionOnPlan bool `json:"subscription_on_plan,omitempty"`
	// <strong>Format: </strong>Variable length, up to 10 N<br><strong>Description: </strong>The plan id of the recurring plan associated with this subscription. 0 if this is an off-plan subscription. 
	PlanId int64 `json:"plan_id,omitempty"`
	// <strong>Format: </strong>Variable length, up to 64 AN<br><strong>Description: </strong>A name assigned by merchant to the plan.
	PlanName string `json:"plan_name,omitempty"`
	// <strong>Format: </strong>Variable length, up to 16 AN<br><strong>Description: </strong>A code assigned by merchant to the plan.
	PlanCode string `json:"plan_code,omitempty"`
	// <strong>Format: </strong>Variable length AN<br><strong>Description: </strong>A short description of the plan. 
	PlanDesc string `json:"plan_desc,omitempty"`
	// <strong>Format: </strong>Variable length, up to 1 N<br><strong>Description: </strong>This field identifies the frequency of billing. Use one of the following codes for frequency. <ul>      <li>0 - Weekly</li>      <li>1 - Bi-Weekly</li>      <li>3 - Monthly</li>      <li>4 - Quarterly</li>      <li>5 - Bi-Annually</li>      <li>6 - Annually</li>      <li>7 - Daily</li>   </ul>
	PlanFrequency int32 `json:"plan_frequency,omitempty"`
	// <strong>Format: </strong>Variable length, up to 4 N<br><strong>Description: </strong> Number of billing cycles in the recurring transaction, -1 indicates bill until cancelled.
	PlanDuration int32 `json:"plan_duration,omitempty"`
	// <strong>Format: </strong>Variable length, up to 2 N<br><strong>Description: </strong> Applicable only for monthly frequency. Number of months in a subscription cycle.
	Interval int32 `json:"interval,omitempty"`
	// <strong>Format: </strong>Variable length, up to 32 AN<br><strong>Description: </strong>First name of the subscriber.
	CustomerFirstName string `json:"customer_first_name,omitempty"`
	// <strong>Format: </strong>Variable length, up to 32 AN<br><strong>Description: </strong>Last name of the subscriber.
	CustomerLastName string `json:"customer_last_name,omitempty"`
	// <strong>Format: </strong>Variable length, up to 10 AN, YYYY-MM-DD format<br><strong>Description: </strong>Start date of subscription. 
	DateStart string `json:"date_start,omitempty"`
	// <strong>Format: </strong>Variable length, up to 10 AN, YYYY-MM-DD format<br><strong>Description: </strong>Next billing date of subscription. This field will be empty for cancelled and completed subscriptions.
	DateNext string `json:"date_next,omitempty"`
	// <strong>Format: </strong>Variable length, up to 10 AN, YYYY-MM-DD format<br><strong>Description: </strong>Date the subscription will end. 
	DateEnd string `json:"date_end,omitempty"`
	// <strong>Format: </strong>Variable length, up to 10,2 N<br><strong>Description: </strong>One-time fee amount. This fee will be charged when a subscription is added.
	AmtSetup float64 `json:"amt_setup,omitempty"`
	// <strong>Format: </strong>Variable length, up to 10 AN, YYYY-MM-DD format<br><strong>Description: </strong>The date the customer will be billed the prorate amount. (for pro-rated subscriptions). 
	ProrateDateStart string `json:"prorate_date_start,omitempty"`
	// <strong>Format: </strong>Variable length, up to 10,2 N<br><strong>Description: </strong>The prorate amount  (for pro-rated subscriptions).
	ProrateAmt float64 `json:"prorate_amt,omitempty"`
	// <strong>Format: </strong>Variable length, up to 10 AN, YYYY-MM-DD format<br><strong>Description: </strong>The start date of the trial period.  Applicable only for subscriptions that include a trail period.
	TrialDateStart string `json:"trial_date_start,omitempty"`
	// <strong>Format: </strong>Variable length, up to 10 AN, YYYY-MM-DD format<br><strong>Description: </strong>The end date of the trial period.  Applicable only for subscriptions that include a trail period.
	TrialDateEnd string `json:"trial_date_end,omitempty"`
	// <strong>Format: </strong>Variable length, up to 10,2 N<br><strong>Description: </strong>The amount billed during the trial period.  Applicable only for subscriptions that include a trail period.
	TrialAmt float64 `json:"trial_amt,omitempty"`
	// <strong>Format: </strong>Variable length, up to 10 AN, YYYY-MM-DD format<br><strong>Description: </strong>Date regular billing cycle will start.
	RecurDateStart string `json:"recur_date_start,omitempty"`
	// <strong>Format: </strong>Variable length, up to 10 AN, YYYY-MM-DD format<br><strong>Description: </strong>Date regular billing cycle will end. 
	RecurDateEnd string `json:"recur_date_end,omitempty"`
	// <strong>Format: </strong>Variable length, up to 10,2 N<br><strong>Description: </strong>Regular billing amount.
	RecurAmt float64 `json:"recur_amt,omitempty"`
	Response *GatewayResponse `json:"response,omitempty"`
	Customer *CustomerVault `json:"customer,omitempty"`
	// <strong>Format: </strong>Fixed length, 3 AN<br><strong>Description: </strong>Numeric Currency Code. Refer to <a href=\"/developer/api/reference#country-codes\" target=\"_blank\">Country Codes</a> for possible values.<br><strong>Default: </strong>840
	TranCurrency string `json:"tran_currency,omitempty"`
	// <strong>Format: </strong>Variable length, up to 3,2 N<br><strong>Description: </strong>Percentage of transaction that will be added as a surcharge fee. A surcharge is a fee added to the cost of a purchase for the \"privilege\" of using a credit card instead of another form of payment. Surcharge can be used only if merchant is enabled to use surcharge and if the selected payment supports surcharge. 
	PctFee float64 `json:"pct_fee,omitempty"`
	AmtTranFee *Surcharge `json:"amt_tran_fee,omitempty"`
	// <strong>Format: </strong>Variable length, up to 25 AN<br><strong>Description: </strong>Purchase Identifier (also referred to as the invoice number generated by the merchant) included in the recurring transactions. If not provided, the system generates a purchase identifier based on customer data. Applicable only to one-off subscriptions.
	PurchaseId string `json:"purchase_id,omitempty"`
}
