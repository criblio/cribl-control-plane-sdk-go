# PqTypeStatus

Persistent queue status information (if persistent queue is enabled).


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `Error`                                                                       | [*components.ErrorTypeStatusPq](../../models/components/errortypestatuspq.md) | :heavy_minus_sign:                                                            | Error information for the persistent queue, if applicable.                    |
| `Health`                                                                      | [components.HealthStringType](../../models/components/healthstringtype.md)    | :heavy_check_mark:                                                            | N/A                                                                           |
| `HealthCounts`                                                                | [components.HealthCountType](../../models/components/healthcounttype.md)      | :heavy_check_mark:                                                            | N/A                                                                           |
| `Timestamp`                                                                   | *float64*                                                                     | :heavy_check_mark:                                                            | Timestamp (in Unix time) when the persistent queue status was last updated.   |