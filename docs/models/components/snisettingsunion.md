# SniSettingsUnion


## Supported Types

### SniSettings1

```go
sniSettingsUnion := components.CreateSniSettingsUnionSniSettings1(components.SniSettings1{/* values here */})
```

### SniSettings2

```go
sniSettingsUnion := components.CreateSniSettingsUnionSniSettings2(components.SniSettings2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch sniSettingsUnion.Type {
	case components.SniSettingsUnionTypeSniSettings1:
		// sniSettingsUnion.SniSettings1 is populated
	case components.SniSettingsUnionTypeSniSettings2:
		// sniSettingsUnion.SniSettings2 is populated
}
```
