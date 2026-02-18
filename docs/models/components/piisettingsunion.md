# PiiSettingsUnion


## Supported Types

### PiiSettings1

```go
piiSettingsUnion := components.CreatePiiSettingsUnionPiiSettings1(components.PiiSettings1{/* values here */})
```

### PiiSettings2

```go
piiSettingsUnion := components.CreatePiiSettingsUnionPiiSettings2(components.PiiSettings2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch piiSettingsUnion.Type {
	case components.PiiSettingsUnionTypePiiSettings1:
		// piiSettingsUnion.PiiSettings1 is populated
	case components.PiiSettingsUnionTypePiiSettings2:
		// piiSettingsUnion.PiiSettings2 is populated
}
```
