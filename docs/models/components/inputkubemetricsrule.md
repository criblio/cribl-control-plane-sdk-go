# InputKubeMetricsRule


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `Filter`                                                                          | *string*                                                                          | :heavy_check_mark:                                                                | JavaScript expression applied to Kubernetes objects. Return 'true' to include it. |
| `Description`                                                                     | **string*                                                                         | :heavy_minus_sign:                                                                | Optional description of this rule's purpose                                       |