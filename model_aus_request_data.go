/*
 * APIs
 *
 * This document describes the Platform and PG API.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type AusRequestData struct {
	// <strong>Format: </strong>Variable length, up to 32 AN<br><strong>Description: </strong>A unique id that can be used to identify the request. The id will be included in the response and can be used to co-relate a request to a response. 
	CardId string `json:"card_id"`
	// <strong>Format: </strong>Variable length, up to 6 AN<br><strong>Description: </strong>The full card number. Only U.S. issued cards can be harvested.  
	CardNumber string `json:"card_number"`
	// <strong>Format: </strong>Fixed length, 4 N, MMYY format<br><strong>Description: </strong>Expiration date of cardholder card number. 
	ExpDate string `json:"exp_date"`
}
