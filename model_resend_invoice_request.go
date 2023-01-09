/*
 * APIs
 *
 * This document describes the Platform and PG API.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ResendInvoiceRequest struct {
	// An array of email addresses to which the invoice will be sent. By default, the system will send the invoice to the customer's email address (business_contact.email_address). 
	EmailAddress []string `json:"email_address,omitempty"`
}