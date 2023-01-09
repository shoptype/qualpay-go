/*
 * APIs
 *
 * This document describes the Platform and PG API.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type AddShippingAddressRequest struct {
	// <strong>Format: </strong>Variable length, up to 32 AN<br><strong>Description: </strong>Shipping first name.
	ShippingFirstName string `json:"shipping_first_name"`
	// <strong>Format: </strong>Variable length, up to 32 AN<br><strong>Description: </strong>Shipping last name.
	ShippingLastName string `json:"shipping_last_name"`
	// <strong>Format: </strong>Variable length, up to 64 AN<br><strong>Description: </strong>Business name of the shipping address, if applicable. 
	ShippingFirmName string `json:"shipping_firm_name,omitempty"`
	// <strong>Format: </strong>Variable length, up to 128 AN<br><strong>Description: </strong>Shipping address line item 1.
	ShippingAddr1 string `json:"shipping_addr1,omitempty"`
	// <strong>Format: </strong>Variable length, up to 128 AN<br><strong>Description: </strong>Shipping address line item 2.
	ShippingAddr2 string `json:"shipping_addr2,omitempty"`
	// <strong>Format: </strong>Variable length, up to 64 AN<br><strong>Description: </strong>Shipping city.
	ShippingCity string `json:"shipping_city,omitempty"`
	// <strong>Format: </strong>Fixed length, 3 AN<br><strong>Description: </strong>Shipping state.
	ShippingState string `json:"shipping_state,omitempty"`
	// <strong>Format: </strong>Variable length, up to 128 AN<br><strong>Description: </strong>Shipping country.
	ShippingCountry string `json:"shipping_country,omitempty"`
	// <strong>Format: </strong>Fixed length, 3 AN<br><strong>Description: </strong>ISO numeric country code for the shipping address. Refer to <a href=\"/developer/api/reference#country-codes\" target=\"_blank\">Country Codes</a> for a list of country codes.
	ShippingCountryCode string `json:"shipping_country_code,omitempty"`
	// <strong>Format: </strong>Fixed length, 10 AN<br><strong>Description: </strong>Shipping zip.
	ShippingZip string `json:"shipping_zip,omitempty"`
	// <strong>Format: </strong>Fixed length, 4 N<br><strong>Description: </strong>Shipping zip+4 code if applicable.
	ShippingZip4 string `json:"shipping_zip4,omitempty"`
	// <br><strong>Description: </strong>Set this field to true if this should be the default address. <br><strong>Default: </strong>false
	Primary bool `json:"primary,omitempty"`
	// <strong>Format: </strong>Variable length, up to 16 AN<br><strong>Description: </strong>Identifies the merchant to whom this request applies. Optional field, applicable only if the request is sent on behalf of another merchant.<br><strong>Conditional Requirement: </strong>Required if this request is on behalf of another merchant.
	MerchantId int64 `json:"merchant_id,omitempty"`
}
