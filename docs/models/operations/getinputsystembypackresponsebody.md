# GetInputSystemByPackResponseBody

List of Source objects.


## Supported Types

### CountedInputResponse

```go
getInputSystemByPackResponseBody := operations.CreateGetInputSystemByPackResponseBodyCountedInputResponse(components.CountedInputResponse{/* values here */})
```

### PaginatedInputResponse

```go
getInputSystemByPackResponseBody := operations.CreateGetInputSystemByPackResponseBodyPaginatedInputResponse(components.PaginatedInputResponse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch getInputSystemByPackResponseBody.Type {
	case operations.GetInputSystemByPackResponseBodyTypeCountedInputResponse:
		// getInputSystemByPackResponseBody.CountedInputResponse is populated
	case operations.GetInputSystemByPackResponseBodyTypePaginatedInputResponse:
		// getInputSystemByPackResponseBody.PaginatedInputResponse is populated
}
```
