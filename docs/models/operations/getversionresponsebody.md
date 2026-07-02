# GetVersionResponseBody

List of GitLogResult objects.


## Supported Types

### CountedGitLogResult

```go
getVersionResponseBody := operations.CreateGetVersionResponseBodyCountedGitLogResult(components.CountedGitLogResult{/* values here */})
```

### PaginatedGitLogResult

```go
getVersionResponseBody := operations.CreateGetVersionResponseBodyPaginatedGitLogResult(components.PaginatedGitLogResult{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch getVersionResponseBody.Type {
	case operations.GetVersionResponseBodyTypeCountedGitLogResult:
		// getVersionResponseBody.CountedGitLogResult is populated
	case operations.GetVersionResponseBodyTypePaginatedGitLogResult:
		// getVersionResponseBody.PaginatedGitLogResult is populated
}
```
