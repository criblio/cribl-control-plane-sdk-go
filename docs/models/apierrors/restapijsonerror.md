# RestAPIJSONError

JSON body returned for many REST failures that use RESTEndpoint.sendError (and similar handlers).


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `Details`                                                               | map[string]`any`                                                        | :heavy_minus_sign:                                                      | Optional structured details about the error (e.g. validation failures). |
| `Message`                                                               | `string`                                                                | :heavy_check_mark:                                                      | Human-readable message or serialized validation details for the error.  |
| `Status`                                                                | `string`                                                                | :heavy_check_mark:                                                      | Always <code>error</code> for API error responses.                      |
| `HTTPMeta`                                                              | [components.HTTPMetadata](../../models/components/httpmetadata.md)      | :heavy_check_mark:                                                      | N/A                                                                     |