# TLSSettingsUnion


## Supported Types

### TLSSettings1

```go
tlsSettingsUnion := components.CreateTLSSettingsUnionTLSSettings1(components.TLSSettings1{/* values here */})
```

### TLSSettings2

```go
tlsSettingsUnion := components.CreateTLSSettingsUnionTLSSettings2(components.TLSSettings2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch tlsSettingsUnion.Type {
	case components.TLSSettingsUnionTypeTLSSettings1:
		// tlsSettingsUnion.TLSSettings1 is populated
	case components.TLSSettingsUnionTypeTLSSettings2:
		// tlsSettingsUnion.TLSSettings2 is populated
}
```
