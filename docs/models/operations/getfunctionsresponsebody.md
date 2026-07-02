# GetFunctionsResponseBody

List of Function objects.


## Supported Types

### CountedFunctionResponse

```go
getFunctionsResponseBody := operations.CreateGetFunctionsResponseBodyCountedFunctionResponse(components.CountedFunctionResponse{/* values here */})
```

### PaginatedFunctionResponse

```go
getFunctionsResponseBody := operations.CreateGetFunctionsResponseBodyPaginatedFunctionResponse(components.PaginatedFunctionResponse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch getFunctionsResponseBody.Type {
	case operations.GetFunctionsResponseBodyTypeCountedFunctionResponse:
		// getFunctionsResponseBody.CountedFunctionResponse is populated
	case operations.GetFunctionsResponseBodyTypePaginatedFunctionResponse:
		// getFunctionsResponseBody.PaginatedFunctionResponse is populated
}
```
