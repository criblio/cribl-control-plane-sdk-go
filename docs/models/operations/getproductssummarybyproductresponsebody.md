# GetProductsSummaryByProductResponseBody

List of DistributedSummary objects.


## Supported Types

### CountedDistributedSummary

```go
getProductsSummaryByProductResponseBody := operations.CreateGetProductsSummaryByProductResponseBodyCountedDistributedSummary(components.CountedDistributedSummary{/* values here */})
```

### PaginatedDistributedSummary

```go
getProductsSummaryByProductResponseBody := operations.CreateGetProductsSummaryByProductResponseBodyPaginatedDistributedSummary(components.PaginatedDistributedSummary{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch getProductsSummaryByProductResponseBody.Type {
	case operations.GetProductsSummaryByProductResponseBodyTypeCountedDistributedSummary:
		// getProductsSummaryByProductResponseBody.CountedDistributedSummary is populated
	case operations.GetProductsSummaryByProductResponseBodyTypePaginatedDistributedSummary:
		// getProductsSummaryByProductResponseBody.PaginatedDistributedSummary is populated
}
```
