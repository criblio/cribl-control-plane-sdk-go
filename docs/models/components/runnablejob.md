# RunnableJob


## Supported Types

### RunnableJobCollection

```go
runnableJob := components.CreateRunnableJobRunnableJobCollection(components.RunnableJobCollection{/* values here */})
```

### RunnableJobExecutor

```go
runnableJob := components.CreateRunnableJobRunnableJobExecutor(components.RunnableJobExecutor{/* values here */})
```

### RunnableJobScheduledSearch

```go
runnableJob := components.CreateRunnableJobRunnableJobScheduledSearch(components.RunnableJobScheduledSearch{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch runnableJob.Type {
	case components.RunnableJobTypeRunnableJobCollection:
		// runnableJob.RunnableJobCollection is populated
	case components.RunnableJobTypeRunnableJobExecutor:
		// runnableJob.RunnableJobExecutor is populated
	case components.RunnableJobTypeRunnableJobScheduledSearch:
		// runnableJob.RunnableJobScheduledSearch is populated
}
```
