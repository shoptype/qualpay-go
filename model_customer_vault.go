/*
 * APIs
 *
 * This document describes the Platform and PG API.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package qualpay

type CustomerVault struct {
	// <strong>Format: </strong>Variable length, up to 32 AN<br><strong>Description: </strong>Unique ID established by merchant to identify a customer. The customer ID can be used to add subscriptions using the Recurring Billing API or make payments using the Payment Gateway API. Once established, this ID cannot be updated. This field is case sensitive. Only letters and numbers are allowed in a Customer ID.
	CustomerId string `json:"customer_id,omitempty"`
	// For INTERNAL USE ONLY.
	RecId int64 `json:"rec_id,omitempty"`
	// <strong>Format: </strong>Variable length, up to 16 N<br><strong>Description: </strong>The Merchant ID to which this customer belongs.
	NodeId int64 `json:"node_id,omitempty"`
	// <strong>Format: </strong>Variable length, up to 32 AN<br><strong>Description: </strong>Customer's first name.
	CustomerFirstName string `json:"customer_first_name,omitempty"`
	// <strong>Format: </strong>Variable length, up to 32 AN<br><strong>Description: </strong>Customer's last name.
	CustomerLastName string `json:"customer_last_name,omitempty"`
	// <strong>Format: </strong>Variable length, up to 64 AN<br><strong>Description: </strong>Customer's business name, if applicable.
	CustomerFirmName string `json:"customer_firm_name,omitempty"`
	// <strong>Format: </strong>Variable length, up to 16 AN<br><strong>Description: </strong>Customer's phone number.
	CustomerPhone string `json:"customer_phone,omitempty"`
	// <strong>Format: </strong>Variable length, up to 64 AN<br><strong>Description: </strong>Customer's email address.
	CustomerEmail string `json:"customer_email,omitempty"`
	// <strong>Format: </strong>Variable length, up to 128 AN<br><strong>Description: </strong>Merchant provided reference value that will be stored with the customer vault data.
	ReferenceId string `json:"reference_id,omitempty"`
	// <strong>Format: </strong>Variable length, up to 128 AN<br><strong>Description: </strong>Comment
	Comments string `json:"comments,omitempty"`
	// <strong>Format: </strong>Variable length, up to 32 AN<br><strong>Description: </strong>Use to indicate which company developed the integration, or the name of the solution or system that is connected to us. Suggested usage is softwareABCv1.0 or companyXYZv2.0.
	DeveloperId string `json:"developer_id,omitempty"`
	// Default Card Number for this customer.
	PrimaryCardNumber string `json:"primary_card_number,omitempty"`
	// <br><strong>Description: </strong>An array of shipping addresses.
	ShippingAddresses []ShippingAddress `json:"shipping_addresses,omitempty"`
	// <br><strong>Description: </strong>An array of billing cards.
	BillingCards []BillingCard `json:"billing_cards,omitempty"`
}
