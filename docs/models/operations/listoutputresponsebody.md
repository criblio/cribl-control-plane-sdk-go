# ListOutputResponseBody

List of Destination objects.


## Supported Types

### CountedOutputResponse

```go
listOutputResponseBody := operations.CreateListOutputResponseBodyCountedOutputResponse(components.CountedOutputResponse{/* values here */})
```

### PaginatedOutputResponse

```go
listOutputResponseBody := operations.CreateListOutputResponseBodyPaginatedOutputResponse(components.PaginatedOutputResponse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch listOutputResponseBody.Type {
	case operations.ListOutputResponseBodyTypeCountedOutputResponse:
		// listOutputResponseBody.CountedOutputResponse is populated
	case operations.ListOutputResponseBodyTypePaginatedOutputResponse:
		// listOutputResponseBody.PaginatedOutputResponse is populated
}
```
