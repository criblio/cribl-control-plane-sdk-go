# DeleteOutputPqByIDResponse


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `HTTPMeta`                                                            | [components.HTTPMetadata](../../models/components/httpmetadata.md)    | :heavy_check_mark:                                                    | N/A                                                                   |
| `CountedString`                                                       | [*components.CountedString](../../models/components/countedstring.md) | :heavy_minus_sign:                                                    | The job ID for the background job that clears the persistent queue.   |