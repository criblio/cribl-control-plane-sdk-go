# GetProductsSummaryByProductRequest


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Product`                                                          | [components.ProductsBase](../../models/components/productsbase.md) | :heavy_check_mark:                                                 | Name of the Cribl product to get the summary for.                  |
| `Offset`                                                           | `*int64`                                                           | :heavy_minus_sign:                                                 | Pagination offset                                                  |
| `Limit`                                                            | `*int64`                                                           | :heavy_minus_sign:                                                 | Maximum number of items to return                                  |