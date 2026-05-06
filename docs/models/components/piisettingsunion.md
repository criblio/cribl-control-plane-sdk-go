# PiiSettingsUnion


## Supported Types

### PiiSettings

```go
piiSettingsUnion := components.CreatePiiSettingsUnionPiiSettings(components.PiiSettings{/* values here */})
```

### EmptyObject

```go
piiSettingsUnion := components.CreatePiiSettingsUnionEmptyObject(components.EmptyObject{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch piiSettingsUnion.Type {
	case components.PiiSettingsUnionTypePiiSettings:
		// piiSettingsUnion.PiiSettings is populated
	case components.PiiSettingsUnionTypeEmptyObject:
		// piiSettingsUnion.EmptyObject is populated
}
```
