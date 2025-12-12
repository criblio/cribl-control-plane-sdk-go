# DistinctConfiguration


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `GroupBy`                                                              | []*string*                                                             | :heavy_minus_sign:                                                     | Defines the properties that are concatenated to produce distinct key   |
| `MaxCombinations`                                                      | **float64*                                                             | :heavy_minus_sign:                                                     | maximum number of tracked combinations                                 |
| `MaxDepth`                                                             | **float64*                                                             | :heavy_minus_sign:                                                     | maximum number of groupBy properties                                   |
| `IsFederated`                                                          | **bool*                                                                | :heavy_minus_sign:                                                     | indicator that the operator runs on a federated executor               |
| `SuppressPreviews`                                                     | **bool*                                                                | :heavy_minus_sign:                                                     | Toggle this on to suppress generating previews of intermediate results |