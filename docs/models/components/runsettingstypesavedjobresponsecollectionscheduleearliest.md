# RunSettingsTypeSavedJobResponseCollectionScheduleEarliest

Earliest time to collect data for the selected timezone


## Supported Types

### 

```go
runSettingsTypeSavedJobResponseCollectionScheduleEarliest := components.CreateRunSettingsTypeSavedJobResponseCollectionScheduleEarliestNumber(float64{/* values here */})
```

### 

```go
runSettingsTypeSavedJobResponseCollectionScheduleEarliest := components.CreateRunSettingsTypeSavedJobResponseCollectionScheduleEarliestStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch runSettingsTypeSavedJobResponseCollectionScheduleEarliest.Type {
	case components.RunSettingsTypeSavedJobResponseCollectionScheduleEarliestTypeNumber:
		// runSettingsTypeSavedJobResponseCollectionScheduleEarliest.Number is populated
	case components.RunSettingsTypeSavedJobResponseCollectionScheduleEarliestTypeStr:
		// runSettingsTypeSavedJobResponseCollectionScheduleEarliest.Str is populated
}
```
