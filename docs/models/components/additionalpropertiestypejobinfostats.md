# AdditionalPropertiesTypeJobInfoStats


## Supported Types

### 

```go
additionalPropertiesTypeJobInfoStats := components.CreateAdditionalPropertiesTypeJobInfoStatsNumber(float64{/* values here */})
```

### 

```go
additionalPropertiesTypeJobInfoStats := components.CreateAdditionalPropertiesTypeJobInfoStatsMapOfNumber(map[string]float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch additionalPropertiesTypeJobInfoStats.Type {
	case components.AdditionalPropertiesTypeJobInfoStatsTypeNumber:
		// additionalPropertiesTypeJobInfoStats.Number is populated
	case components.AdditionalPropertiesTypeJobInfoStatsTypeMapOfNumber:
		// additionalPropertiesTypeJobInfoStats.MapOfNumber is populated
}
```
