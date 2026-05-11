# RunSettingsTypeRunnableJobCollectionScheduleLatest

Latest time to collect data for the selected timezone


## Supported Types

### 

```go
runSettingsTypeRunnableJobCollectionScheduleLatest := components.CreateRunSettingsTypeRunnableJobCollectionScheduleLatestNumber(float64{/* values here */})
```

### 

```go
runSettingsTypeRunnableJobCollectionScheduleLatest := components.CreateRunSettingsTypeRunnableJobCollectionScheduleLatestStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch runSettingsTypeRunnableJobCollectionScheduleLatest.Type {
	case components.RunSettingsTypeRunnableJobCollectionScheduleLatestTypeNumber:
		// runSettingsTypeRunnableJobCollectionScheduleLatest.Number is populated
	case components.RunSettingsTypeRunnableJobCollectionScheduleLatestTypeStr:
		// runSettingsTypeRunnableJobCollectionScheduleLatest.Str is populated
}
```
