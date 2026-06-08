# RestAPIJSONError

JSON body returned for many REST failures that use RESTEndpoint.sendError (and similar handlers).


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `Message`                                                              | `string`                                                               | :heavy_check_mark:                                                     | Human-readable message or serialized validation details for the error. |
| `Status`                                                               | `string`                                                               | :heavy_check_mark:                                                     | Always <code>error</code> for API error responses.                     |
| `HTTPMeta`                                                             | [components.HTTPMetadata](../../models/components/httpmetadata.md)     | :heavy_check_mark:                                                     | N/A                                                                    |