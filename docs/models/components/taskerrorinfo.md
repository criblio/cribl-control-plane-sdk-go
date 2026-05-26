# TaskErrorInfo

Serialized error object that describes why a job entered its current <code>state</code>. Includes <code>message</code> and may include a nested <code>reason</code> for wrapped errors.


## Fields

| Field                               | Type                                | Required                            | Description                         |
| ----------------------------------- | ----------------------------------- | ----------------------------------- | ----------------------------------- |
| `Message`                           | `string`                            | :heavy_check_mark:                  | Human-readable error message.       |
| `Name`                              | `*string`                           | :heavy_minus_sign:                  | Error name, if available.           |
| `Stack`                             | `*string`                           | :heavy_minus_sign:                  | Truncated stack trace of the error. |
| `AdditionalProperties`              | map[string]`any`                    | :heavy_minus_sign:                  | N/A                                 |