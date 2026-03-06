# ErrorTypeStatus

Error information, if applicable.


## Fields

| Field                                                                                                     | Type                                                                                                      | Required                                                                                                  | Description                                                                                               |
| --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- |
| `Details`                                                                                                 | map[string]*any*                                                                                          | :heavy_minus_sign:                                                                                        | Additional error details, which may include stack traces, request information, and other diagnostic data. |
| `Message`                                                                                                 | *string*                                                                                                  | :heavy_check_mark:                                                                                        | Human-readable message that describes the error.                                                          |