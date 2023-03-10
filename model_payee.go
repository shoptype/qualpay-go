/*
 * APIs
 *
 * This document describes the Platform and PG API.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package qualpay

type Payee struct {
	// <strong>Format: </strong>Variable length, up to 32 AN<br><strong>Description: </strong>A unique ID generated to identify a payee.
	PayeeId string `json:"payee_id,omitempty"`
	// <strong>Format: </strong>Variable length, up to 64AN <br><strong>Description:</strong>Payee name.
	PayeeName string `json:"payee_name,omitempty"`
	// <strong>Format: </strong>Variable length, up to 32 N<br><strong>Description: </strong> Payee phone number. Prepend with + and country code.
	PayeePhone string `json:"payee_phone,omitempty"`
	// <strong>Description:</strong>Payee email addresses
	PayeeEmail   []string                `json:"payee_email,omitempty"`
	PayeeAddress *AddPayeeAddressRequest `json:"payee_address,omitempty"`
	PayeeAccount *PayeeAccountResponse   `json:"payee_account,omitempty"`
}
