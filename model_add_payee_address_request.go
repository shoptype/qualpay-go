/*
 * APIs
 *
 * This document describes the Platform and PG API.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type AddPayeeAddressRequest struct {
	// <strong>Format: </strong>Variable length<br><strong>Description:</strong>Line1 postal address.
	Address1 string `json:"address1,omitempty"`
	// <strong>Format: </strong>Variable length<br><strong>Description:</strong>Line2 postal address.
	Address2 string `json:"address2,omitempty"`
	// <strong>Format: </strong>Variable length<br><strong>Description:</strong>City.
	City string `json:"city,omitempty"`
	// <strong>Format: </strong>Variable length<br><strong>Description:</strong>State.
	State string `json:"state,omitempty"`
	// <strong>Format: </strong>Variable length, up to 16 AN<br><strong>Description: </strong>Postal zip code.
	Zip string `json:"zip,omitempty"`
	// <strong>Format: </strong>Fixed length 3N<br><strong>Description: </strong>Numeric Country Code.
	CountryCode string `json:"country_code,omitempty"`
	// <strong>Format: </strong>Variable length<br><strong>Description: </strong>Country name.
	Country string `json:"country,omitempty"`
}