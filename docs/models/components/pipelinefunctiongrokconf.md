# PipelineFunctionGrokConf


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `Pattern`                                                                    | *string*                                                                     | :heavy_check_mark:                                                           | Grok pattern to extract fields. Syntax supported: %{PATTERN_NAME:FIELD_NAME} |
| `PatternList`                                                                | [][components.PatternList](../../models/components/patternlist.md)           | :heavy_minus_sign:                                                           | N/A                                                                          |
| `Source`                                                                     | **string*                                                                    | :heavy_minus_sign:                                                           | Field on which to perform Grok extractions                                   |