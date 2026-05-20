# RunnableJobCollectionEarliest

Earliest time to collect data for the selected timezone


## Supported Types

### 

```go
runnableJobCollectionEarliest := components.CreateRunnableJobCollectionEarliestNumber(float64{/* values here */})
```

### 

```go
runnableJobCollectionEarliest := components.CreateRunnableJobCollectionEarliestStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch runnableJobCollectionEarliest.Type {
	case components.RunnableJobCollectionEarliestTypeNumber:
		// runnableJobCollectionEarliest.Number is populated
	case components.RunnableJobCollectionEarliestTypeStr:
		// runnableJobCollectionEarliest.Str is populated
}
```
