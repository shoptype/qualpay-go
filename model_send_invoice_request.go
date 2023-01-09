/*
 * APIs
 *
 * This document describes the Platform and PG API.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SendInvoiceRequest struct {
	// An array of email addresses to which the invoice will be sent.
	EmailAddress []string `json:"email_address,omitempty"`
	// An SMS of the invoice will be sent to the phone number. The numbers should be ISO Formatted. Messages can be sent only to US phone numbers. SMS feature should be enabled for your account to use SMS.
	PhoneNumber []string `json:"phone_number,omitempty"`
}
