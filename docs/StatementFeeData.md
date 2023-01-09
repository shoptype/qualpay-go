# StatementFeeData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecIdLinked** | **int64** | For INTERNAL USE ONLY. | [optional] [default to null]
**MerchantId** | **int64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 16 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Unique ID assigned to a merchant. | [optional] [default to null]
**BillingMonth** | **string** | Billing month of the statement in YYYY-MM-DD. | [optional] [default to null]
**FeeCategory** | **string** | The category defines the type of fees. | [optional] [default to null]
**Count** | **int64** | The count of items applicable to the category. | [optional] [default to null]
**AmtFees** | **float64** | The dollar amount of items applicable to the category. | [optional] [default to null]
**PerItem** | **float64** | The percentage rate applied to the category item. | [optional] [default to null]
**Rate** | **float64** | The per item rate applied to the category item. | [optional] [default to null]
**DiscountDue** | **float64** |  | [default to null]
**Description** | **string** | The description of a category item. | [optional] [default to null]
**FeesPaid** | **float64** | Any fees paid during the month. | [optional] [default to null]
**FeesTotal** | **float64** | The total amount of fees due for the category item. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

