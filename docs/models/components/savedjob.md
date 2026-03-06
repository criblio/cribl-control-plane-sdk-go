# SavedJob


## Supported Types

### SavedJobCollection

```go
savedJob := components.CreateSavedJobSavedJobCollection(components.SavedJobCollection{/* values here */})
```

### SavedJobExecutor

```go
savedJob := components.CreateSavedJobSavedJobExecutor(components.SavedJobExecutor{/* values here */})
```

### SavedJobScheduledSearch

```go
savedJob := components.CreateSavedJobSavedJobScheduledSearch(components.SavedJobScheduledSearch{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch savedJob.Type {
	case components.SavedJobTypeSavedJobCollection:
		// savedJob.SavedJobCollection is populated
	case components.SavedJobTypeSavedJobExecutor:
		// savedJob.SavedJobExecutor is populated
	case components.SavedJobTypeSavedJobScheduledSearch:
		// savedJob.SavedJobScheduledSearch is populated
}
```
