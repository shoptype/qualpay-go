/*
 * APIs
 *
 * This document describes the Platform and PG API.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package qualpay

type Customer struct {
	// <strong>Format: </strong>Variable length<br><strong>Description: </strong>Customer billing address street.
	BillingAddr1 string `json:"billing_addr1,omitempty"`
	// <strong>Format: </strong>Variable length<br><strong>Description: </strong>Customer billing address, line 2.
	BillingAddr2 string `json:"billing_addr2,omitempty"`
	// <strong>Format: </strong>Variable length<br><strong>Description: </strong>Customer billing city.
	BillingCity string `json:"billing_city,omitempty"`
	// <strong>Format: </strong>Variable length<br><strong>Description: </strong>Customer billing country.
	BillingCountry string `json:"billing_country,omitempty"`
	// <strong>Format: </strong>Fixed length, 3 N<br><strong>Description: </strong>ISO numeric country code for the billing address. Refer to <a href=\"/developer/api/reference#country-codes\" target=\"_blank\">Country Codes</a> for a list of country codes.
	BillingCountryCode string `json:"billing_country_code,omitempty"`
	// <strong>Format: </strong>Variable length<br><strong>Description: </strong>Customer billing state (abbreviated).
	BillingState string `json:"billing_state,omitempty"`
	// <strong>Format: </strong>Variable length<br><strong>Description: </strong>Customer billing zip code.
	BillingZip string `json:"billing_zip,omitempty"`
	// <strong>Format: </strong>Fixed length, 4 N<br><strong>Description: </strong>Customer billing zip+4 code if applicable.
	BillingZip4 string `json:"billing_zip4,omitempty"`
	// <strong>Format: </strong>Variable length, up to 64 AN<br><strong>Description: </strong>Customer e-mail address.
	CustomerEmail string `json:"customer_email,omitempty"`
	// <strong>Format: </strong>Variable length, up to 64 AN<br><strong>Description: </strong>Customer Business name if applicable.<br><strong>Conditional Requirement: </strong>Either customer first and last name or firm name is required.
	CustomerFirmName string `json:"customer_firm_name,omitempty"`
	// <strong>Format: </strong>Variable length, up to 32 AN<br><strong>Description: </strong>Customer first name.<br><strong>Conditional Requirement: </strong>Either customer first and last name or firm name is required.
	CustomerFirstName string `json:"customer_first_name,omitempty"`
	// <strong>Format: </strong>Variable length, up to 32 AN<br><strong>Description: </strong>Customer last name.<br><strong>Conditional Requirement: </strong>Either customer first and last name or firm name is required.
	CustomerLastName string `json:"customer_last_name,omitempty"`
	// <strong>Format: </strong>Variable length, up to 16 N<br><strong>Description: </strong>Customer phone number.
	CustomerPhone string `json:"customer_phone,omitempty"`
	// List of shipping addresses for customer.
	ShippingAddresses []PgShippingAddress `json:"shipping_addresses,omitempty"`
}
