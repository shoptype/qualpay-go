# StatementReserveData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecIdLinked** | **int64** | For INTERNAL USE ONLY. | [optional] [default to null]
**MerchantId** | **int64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 16 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Unique ID assigned to a merchant. | [optional] [default to null]
**BillingMonth** | **string** | Identifies the date funds are posted to the depository account for either settled transactions, net of reserve or adjustments for release of reserves. in YYYY-MM-DD. | [optional] [default to null]
**BatchDate** | [**time.Time**](time.Time.md) | Identifies the date funds are posted to the depository account for either settled transactions, net of reserve or adjustments for release of reserves. in YYYY-MM-DD. | [optional] [default to null]
**AmtReserve** | **float64** | Is the total dollar amount deducted from settled transactions prior to funding. The amount is calculated by applying an agreed upon percentage to the sales volume in a batch. | [optional] [default to null]
**AmtReleasedReserve** | **float64** | Is the total dollar amount released as an adjustment to the depository account. | [optional] [default to null]
**AmtReserveBalance** | **float64** | Total amount of reserve held as of the ACH Post Date. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

