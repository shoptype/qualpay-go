/*
 * APIs
 *
 * This document describes the Platform and PG API.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type PaymentTerm struct {
	// <strong>Format: </strong>Variable length, up to 10 AN<br><strong>Description: </strong>Payment Term type. By default, this field is set to DAYS. 
	Type_ string `json:"type,omitempty"`
	// <strong>Format: </strong>Variable length, up to 25 AN<br><strong>Description: </strong>Payment Term. The invoice_due_date will be automatically updated based on this value. If type is DATE, this field should contain the invoice_due_date in 'YYYY-MM-DD' format. If type is DAYS, this field should contain the number of days from invoice date when payment is due.  
	Value string `json:"value,omitempty"`
}