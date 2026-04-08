# RunnableJobCollectionLatest

Latest time to collect data for the selected timezone


## Supported Types

### 

```go
runnableJobCollectionLatest := components.CreateRunnableJobCollectionLatestNumber(float64{/* values here */})
```

### 

```go
runnableJobCollectionLatest := components.CreateRunnableJobCollectionLatestStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch runnableJobCollectionLatest.Type {
	case components.RunnableJobCollectionLatestTypeNumber:
		// runnableJobCollectionLatest.Number is populated
	case components.RunnableJobCollectionLatestTypeStr:
		// runnableJobCollectionLatest.Str is populated
}
```
