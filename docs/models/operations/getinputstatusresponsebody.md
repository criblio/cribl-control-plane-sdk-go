# GetInputStatusResponseBody

List of Source status objects.


## Supported Types

### CountedInputStatus

```go
getInputStatusResponseBody := operations.CreateGetInputStatusResponseBodyCountedInputStatus(components.CountedInputStatus{/* values here */})
```

### PaginatedInputStatus

```go
getInputStatusResponseBody := operations.CreateGetInputStatusResponseBodyPaginatedInputStatus(components.PaginatedInputStatus{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch getInputStatusResponseBody.Type {
	case operations.GetInputStatusResponseBodyTypeCountedInputStatus:
		// getInputStatusResponseBody.CountedInputStatus is populated
	case operations.GetInputStatusResponseBodyTypePaginatedInputStatus:
		// getInputStatusResponseBody.PaginatedInputStatus is populated
}
```
