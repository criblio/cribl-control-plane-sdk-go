# GetSearchQueryResponse


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `HTTPMeta`                                                                       | [components.HTTPMetadata](../../models/components/httpmetadata.md)               | :heavy_check_mark:                                                               | N/A                                                                              |
| `SearchJobResults`                                                               | **jsonl.JsonLStream[components.SearchJobResults]*                                | :heavy_minus_sign:                                                               | SearchResultsResults for the Search /results and /results-poll endpoints. object |