/*
 * APIs
 *
 * This document describes the Platform and PG API.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type WebhookEvent struct {
	// An array of events that will trigger the POST request. Refer to <a href=\"/developer/api/reference#guides/webhooks-schema\" target=\"_blank\">Webhook Request Schema</a> for a list of available events.
	Events []string `json:"events"`
}
