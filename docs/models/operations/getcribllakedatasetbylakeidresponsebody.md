# GetCriblLakeDatasetByLakeIDResponseBody

List of CriblLakeDataset objects.


## Supported Types

### CountedCriblLakeDataset

```go
getCriblLakeDatasetByLakeIDResponseBody := operations.CreateGetCriblLakeDatasetByLakeIDResponseBodyCountedCriblLakeDataset(components.CountedCriblLakeDataset{/* values here */})
```

### PaginatedCriblLakeDataset

```go
getCriblLakeDatasetByLakeIDResponseBody := operations.CreateGetCriblLakeDatasetByLakeIDResponseBodyPaginatedCriblLakeDataset(components.PaginatedCriblLakeDataset{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch getCriblLakeDatasetByLakeIDResponseBody.Type {
	case operations.GetCriblLakeDatasetByLakeIDResponseBodyTypeCountedCriblLakeDataset:
		// getCriblLakeDatasetByLakeIDResponseBody.CountedCriblLakeDataset is populated
	case operations.GetCriblLakeDatasetByLakeIDResponseBodyTypePaginatedCriblLakeDataset:
		// getCriblLakeDatasetByLakeIDResponseBody.PaginatedCriblLakeDataset is populated
}
```
