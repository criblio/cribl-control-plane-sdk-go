# GetOutputStatusResponseBody

List of Destination status objects.


## Supported Types

### CountedOutputStatus

```go
getOutputStatusResponseBody := operations.CreateGetOutputStatusResponseBodyCountedOutputStatus(components.CountedOutputStatus{/* values here */})
```

### PaginatedOutputStatus

```go
getOutputStatusResponseBody := operations.CreateGetOutputStatusResponseBodyPaginatedOutputStatus(components.PaginatedOutputStatus{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch getOutputStatusResponseBody.Type {
	case operations.GetOutputStatusResponseBodyTypeCountedOutputStatus:
		// getOutputStatusResponseBody.CountedOutputStatus is populated
	case operations.GetOutputStatusResponseBodyTypePaginatedOutputStatus:
		// getOutputStatusResponseBody.PaginatedOutputStatus is populated
}
```
