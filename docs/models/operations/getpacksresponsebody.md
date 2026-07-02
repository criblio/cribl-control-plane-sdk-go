# GetPacksResponseBody

List of Pack objects.


## Supported Types

### CountedPackInfo

```go
getPacksResponseBody := operations.CreateGetPacksResponseBodyCountedPackInfo(components.CountedPackInfo{/* values here */})
```

### PaginatedPackInfo

```go
getPacksResponseBody := operations.CreateGetPacksResponseBodyPaginatedPackInfo(components.PaginatedPackInfo{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch getPacksResponseBody.Type {
	case operations.GetPacksResponseBodyTypeCountedPackInfo:
		// getPacksResponseBody.CountedPackInfo is populated
	case operations.GetPacksResponseBodyTypePaginatedPackInfo:
		// getPacksResponseBody.PaginatedPackInfo is populated
}
```
