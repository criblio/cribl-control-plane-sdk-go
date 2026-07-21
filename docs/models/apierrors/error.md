# Error


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `Status`                                                                | `string`                                                                | :heavy_check_mark:                                                      | Always "error" for API error responses.                                 |
| `Message`                                                               | `string`                                                                | :heavy_check_mark:                                                      | Human-readable message describing the error.                            |
| `Details`                                                               | `any`                                                                   | :heavy_minus_sign:                                                      | Optional structured details about the error (e.g. validation failures). |
| `HTTPMeta`                                                              | [components.HTTPMetadata](../../models/components/httpmetadata.md)      | :heavy_check_mark:                                                      | N/A                                                                     |