# TLSSettingsUnion


## Supported Types

### TLSSettings

```go
tlsSettingsUnion := components.CreateTLSSettingsUnionTLSSettings(components.TLSSettings{/* values here */})
```

### EmptyObject

```go
tlsSettingsUnion := components.CreateTLSSettingsUnionEmptyObject(components.EmptyObject{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch tlsSettingsUnion.Type {
	case components.TLSSettingsUnionTypeTLSSettings:
		// tlsSettingsUnion.TLSSettings is populated
	case components.TLSSettingsUnionTypeEmptyObject:
		// tlsSettingsUnion.EmptyObject is populated
}
```
