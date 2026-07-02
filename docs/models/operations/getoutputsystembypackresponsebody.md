# GetOutputSystemByPackResponseBody

List of Destination objects.


## Supported Types

### CountedOutputResponse

```go
getOutputSystemByPackResponseBody := operations.CreateGetOutputSystemByPackResponseBodyCountedOutputResponse(components.CountedOutputResponse{/* values here */})
```

### PaginatedOutputResponse

```go
getOutputSystemByPackResponseBody := operations.CreateGetOutputSystemByPackResponseBodyPaginatedOutputResponse(components.PaginatedOutputResponse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch getOutputSystemByPackResponseBody.Type {
	case operations.GetOutputSystemByPackResponseBodyTypeCountedOutputResponse:
		// getOutputSystemByPackResponseBody.CountedOutputResponse is populated
	case operations.GetOutputSystemByPackResponseBodyTypePaginatedOutputResponse:
		// getOutputSystemByPackResponseBody.PaginatedOutputResponse is populated
}
```
