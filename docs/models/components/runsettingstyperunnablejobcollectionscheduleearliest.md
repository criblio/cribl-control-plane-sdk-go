# RunSettingsTypeRunnableJobCollectionScheduleEarliest

Earliest time to collect data for the selected timezone


## Supported Types

### 

```go
runSettingsTypeRunnableJobCollectionScheduleEarliest := components.CreateRunSettingsTypeRunnableJobCollectionScheduleEarliestNumber(float64{/* values here */})
```

### 

```go
runSettingsTypeRunnableJobCollectionScheduleEarliest := components.CreateRunSettingsTypeRunnableJobCollectionScheduleEarliestStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch runSettingsTypeRunnableJobCollectionScheduleEarliest.Type {
	case components.RunSettingsTypeRunnableJobCollectionScheduleEarliestTypeNumber:
		// runSettingsTypeRunnableJobCollectionScheduleEarliest.Number is populated
	case components.RunSettingsTypeRunnableJobCollectionScheduleEarliestTypeStr:
		// runSettingsTypeRunnableJobCollectionScheduleEarliest.Str is populated
}
```
