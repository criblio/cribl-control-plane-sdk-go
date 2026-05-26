# SniSettingsUnion


## Supported Types

### SniSettings

```go
sniSettingsUnion := components.CreateSniSettingsUnionSniSettings(components.SniSettings{/* values here */})
```

### EmptyObject

```go
sniSettingsUnion := components.CreateSniSettingsUnionEmptyObject(components.EmptyObject{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch sniSettingsUnion.Type {
	case components.SniSettingsUnionTypeSniSettings:
		// sniSettingsUnion.SniSettings is populated
	case components.SniSettingsUnionTypeEmptyObject:
		// sniSettingsUnion.EmptyObject is populated
}
```
