# GetProductsSummaryWorkersByProductRequest


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `Product`                                                                  | [components.ProductsBase](../../models/components/productsbase.md)         | :heavy_check_mark:                                                         | Name of the Cribl product to get the count of Worker or Edge Nodes for.    |
| `FilterExp`                                                                | **string*                                                                  | :heavy_minus_sign:                                                         | Filter expression to evaluate against Nodes for inclusion in the response. |