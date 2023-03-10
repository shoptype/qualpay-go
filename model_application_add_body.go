/*
 * APIs
 *
 * This document describes the Platform and PG API.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package qualpay

type ApplicationAddBody struct {
	// The sales channel to start this application in. The channel ID is provided to you by Qualpay.
	ChannelId int64 `json:"channel_id"`
	// The sales rep and contact for this application in. This is the qualpay ID for the sales rep user. The service at sales/browse can assist in finding the user_id.
	SalesRep int64 `json:"sales_rep"`
	// The point at which to board the merchant in the Qualpay hierarchy. A default node is configured for you, however you may choose a different node by using this field.
	TerminationNode int64 `json:"termination_node,omitempty"`
}
