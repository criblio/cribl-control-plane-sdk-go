# GetOutputStatusSystemOutputsByPackResponseBody

List of Destination status objects.


## Supported Types

### CountedOutputStatus

```go
getOutputStatusSystemOutputsByPackResponseBody := operations.CreateGetOutputStatusSystemOutputsByPackResponseBodyCountedOutputStatus(components.CountedOutputStatus{/* values here */})
```

### PaginatedOutputStatus

```go
getOutputStatusSystemOutputsByPackResponseBody := operations.CreateGetOutputStatusSystemOutputsByPackResponseBodyPaginatedOutputStatus(components.PaginatedOutputStatus{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch getOutputStatusSystemOutputsByPackResponseBody.Type {
	case operations.GetOutputStatusSystemOutputsByPackResponseBodyTypeCountedOutputStatus:
		// getOutputStatusSystemOutputsByPackResponseBody.CountedOutputStatus is populated
	case operations.GetOutputStatusSystemOutputsByPackResponseBodyTypePaginatedOutputStatus:
		// getOutputStatusSystemOutputsByPackResponseBody.PaginatedOutputStatus is populated
}
```
