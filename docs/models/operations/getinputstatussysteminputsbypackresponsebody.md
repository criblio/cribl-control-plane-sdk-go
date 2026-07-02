# GetInputStatusSystemInputsByPackResponseBody

List of Source status objects.


## Supported Types

### CountedInputStatus

```go
getInputStatusSystemInputsByPackResponseBody := operations.CreateGetInputStatusSystemInputsByPackResponseBodyCountedInputStatus(components.CountedInputStatus{/* values here */})
```

### PaginatedInputStatus

```go
getInputStatusSystemInputsByPackResponseBody := operations.CreateGetInputStatusSystemInputsByPackResponseBodyPaginatedInputStatus(components.PaginatedInputStatus{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch getInputStatusSystemInputsByPackResponseBody.Type {
	case operations.GetInputStatusSystemInputsByPackResponseBodyTypeCountedInputStatus:
		// getInputStatusSystemInputsByPackResponseBody.CountedInputStatus is populated
	case operations.GetInputStatusSystemInputsByPackResponseBodyTypePaginatedInputStatus:
		// getInputStatusSystemInputsByPackResponseBody.PaginatedInputStatus is populated
}
```
