# DeleteOutputPqByIDResponse


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `HTTPMeta`                                                                    | [components.HTTPMetadata](../../models/components/httpmetadata.md)            | :heavy_check_mark:                                                            | N/A                                                                           |
| `CountedListstring`                                                           | [*components.CountedListstring](../../models/components/countedliststring.md) | :heavy_minus_sign:                                                            | A list of job ids for the background job that clears the persistent queue     |