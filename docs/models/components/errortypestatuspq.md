# ErrorTypeStatusPq

Error information for the persistent queue, if applicable.


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `Details`                                                         | map[string]*any*                                                  | :heavy_minus_sign:                                                | Additional details for the persistent queue error.                |
| `Message`                                                         | *string*                                                          | :heavy_check_mark:                                                | Human-readable message that describes the persistent queue error. |