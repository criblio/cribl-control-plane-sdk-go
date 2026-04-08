# RunSettingsTypeSavedJobResponseCollectionScheduleLatest

Latest time to collect data for the selected timezone


## Supported Types

### 

```go
runSettingsTypeSavedJobResponseCollectionScheduleLatest := components.CreateRunSettingsTypeSavedJobResponseCollectionScheduleLatestNumber(float64{/* values here */})
```

### 

```go
runSettingsTypeSavedJobResponseCollectionScheduleLatest := components.CreateRunSettingsTypeSavedJobResponseCollectionScheduleLatestStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch runSettingsTypeSavedJobResponseCollectionScheduleLatest.Type {
	case components.RunSettingsTypeSavedJobResponseCollectionScheduleLatestTypeNumber:
		// runSettingsTypeSavedJobResponseCollectionScheduleLatest.Number is populated
	case components.RunSettingsTypeSavedJobResponseCollectionScheduleLatestTypeStr:
		// runSettingsTypeSavedJobResponseCollectionScheduleLatest.Str is populated
}
```
