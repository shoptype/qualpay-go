# StatementPlanData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecIdLinked** | **int64** | For INTERNAL USE ONLY. | [optional] [default to null]
**MerchantId** | **int64** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 16 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Unique ID assigned to a merchant. | [optional] [default to null]
**BillingMonth** | **string** | Billing month of the statement in YYYY-MM-DD | [optional] [default to null]
**PlanCode** | **string** | Defines the payment instrument type of Visa, MasterCard, American Express, or Discover. | [optional] [default to null]
**CntPurch** | **int64** | The total funded sales transactions count. | [optional] [default to null]
**AmtPurch** | **float64** | The total funded sales dollar amount. | [optional] [default to null]
**CntReturn** | **int64** | The total funded credit transactions count. | [optional] [default to null]
**AmtReturn** | **float64** | The total funded credit dollar amount. | [optional] [default to null]
**AmtNet** | **float64** | The net dollar amount of transactions. | [optional] [default to null]
**AmtAvgTicket** | **float64** | Average ticket amount. | [optional] [default to null]
**PerItem** | **float64** | The per item rate for the payment type. | [optional] [default to null]
**Rate** | **float64** | The discount rate for the payment type is applied to the Sales Amount. | [optional] [default to null]
**DiscountDue** | **float64** | The Amount of discount due | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

