# GetSavedJobResponseBody

The list of Collectors in a response envelope with <code>count</code> and <code>items</code>.


## Supported Types

### CountedSavedJobResponse

```go
getSavedJobResponseBody := operations.CreateGetSavedJobResponseBodyCountedSavedJobResponse(components.CountedSavedJobResponse{/* values here */})
```

### PaginatedSavedJobResponse

```go
getSavedJobResponseBody := operations.CreateGetSavedJobResponseBodyPaginatedSavedJobResponse(components.PaginatedSavedJobResponse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch getSavedJobResponseBody.Type {
	case operations.GetSavedJobResponseBodyTypeCountedSavedJobResponse:
		// getSavedJobResponseBody.CountedSavedJobResponse is populated
	case operations.GetSavedJobResponseBodyTypePaginatedSavedJobResponse:
		// getSavedJobResponseBody.PaginatedSavedJobResponse is populated
}
```
