# SavedJobResponse


## Supported Types

### SavedJobResponseCollection

```go
savedJobResponse := components.CreateSavedJobResponseSavedJobResponseCollection(components.SavedJobResponseCollection{/* values here */})
```

### SavedJobResponseExecutor

```go
savedJobResponse := components.CreateSavedJobResponseSavedJobResponseExecutor(components.SavedJobResponseExecutor{/* values here */})
```

### SavedJobResponseScheduledSearch

```go
savedJobResponse := components.CreateSavedJobResponseSavedJobResponseScheduledSearch(components.SavedJobResponseScheduledSearch{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch savedJobResponse.Type {
	case components.SavedJobResponseTypeSavedJobResponseCollection:
		// savedJobResponse.SavedJobResponseCollection is populated
	case components.SavedJobResponseTypeSavedJobResponseExecutor:
		// savedJobResponse.SavedJobResponseExecutor is populated
	case components.SavedJobResponseTypeSavedJobResponseScheduledSearch:
		// savedJobResponse.SavedJobResponseScheduledSearch is populated
}
```
