# GetProductsWorkersByProductResponseBody

List of MasterWorkerEntry objects.


## Supported Types

### CountedMasterWorkerEntry

```go
getProductsWorkersByProductResponseBody := operations.CreateGetProductsWorkersByProductResponseBodyCountedMasterWorkerEntry(components.CountedMasterWorkerEntry{/* values here */})
```

### PaginatedMasterWorkerEntry

```go
getProductsWorkersByProductResponseBody := operations.CreateGetProductsWorkersByProductResponseBodyPaginatedMasterWorkerEntry(components.PaginatedMasterWorkerEntry{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch getProductsWorkersByProductResponseBody.Type {
	case operations.GetProductsWorkersByProductResponseBodyTypeCountedMasterWorkerEntry:
		// getProductsWorkersByProductResponseBody.CountedMasterWorkerEntry is populated
	case operations.GetProductsWorkersByProductResponseBodyTypePaginatedMasterWorkerEntry:
		// getProductsWorkersByProductResponseBody.PaginatedMasterWorkerEntry is populated
}
```
