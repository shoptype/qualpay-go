/*
 * APIs
 *
 * This document describes the Platform and PG API.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Settings struct {
	// The application ID for this merchant account.  Only returned if an application exists.
	AppId int64 `json:"app_id,omitempty"`
	// The merchant ID this merchant account.
	MerchantId int64 `json:"merchant_id,omitempty"`
	// The DBA name for the merchant account.
	DbaName string `json:"dba_name,omitempty"`
	// An array of the card types and currency accepted by the merchant account.
	PaymentsAccepted []Payment `json:"payments_accepted,omitempty"`
	// An array payment profiles available for the merchant account.
	PaymentProfiles []PaymentProfile `json:"payment_profiles,omitempty"`
	// The user login created for the merchant account.  Only returned if a new user was requested during account creation.
	UserLogin string `json:"user_login,omitempty"`
	// The URL to establish a new password for the created user.  Only returned if a new user was requested during account creation.
	ResetUrl string `json:"reset_url,omitempty"`
}