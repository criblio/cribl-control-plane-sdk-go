# ListInputResponseBody

List of Source objects.


## Supported Types

### CountedInputResponse

```go
listInputResponseBody := operations.CreateListInputResponseBodyCountedInputResponse(components.CountedInputResponse{/* values here */})
```

### PaginatedInputResponse

```go
listInputResponseBody := operations.CreateListInputResponseBodyPaginatedInputResponse(components.PaginatedInputResponse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch listInputResponseBody.Type {
	case operations.ListInputResponseBodyTypeCountedInputResponse:
		// listInputResponseBody.CountedInputResponse is populated
	case operations.ListInputResponseBodyTypePaginatedInputResponse:
		// listInputResponseBody.PaginatedInputResponse is populated
}
```
