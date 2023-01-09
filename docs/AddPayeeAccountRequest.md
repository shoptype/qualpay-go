# AddPayeeAccountRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeId** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 1 AN&lt;br&gt;&lt;strong&gt;Description:&lt;/strong&gt; Bank Account Type. Possible values are: &lt;br&gt; &lt;ul&gt; &lt;li&gt;&#x60;C&#x60; - Personal checking account&lt;/li&gt; &lt;li&gt;&#x60;S&#x60; - Personal savings account&lt;/li&gt; &lt;li&gt;&#x60;K&#x60; - Business checking account&lt;/li&gt; &lt;li&gt;&#x60;V&#x60; - Business savings account&lt;/li&gt; &lt;/ul&gt; | [optional] [default to TYPE_ID.C]
**TrNumber** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Fixed length, 9 N &lt;br&gt; &lt;strong&gt;Description:&lt;/strong&gt;Bank transit/routing number. | [optional] [default to null]
**DdaNumber** | **string** | &lt;strong&gt;Format: &lt;/strong&gt;Variable length, up to 17 N&lt;br&gt;&lt;strong&gt;Description: &lt;/strong&gt;Owner&#x27;s account number at the bank. | [optional] [default to null]
**AccountholderName** | **string** | Account holder name | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

