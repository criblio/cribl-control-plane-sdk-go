# RollbackSettingsUnion


## Supported Types

### RollbackSettings

```go
rollbackSettingsUnion := components.CreateRollbackSettingsUnionRollbackSettings(components.RollbackSettings{/* values here */})
```

### EmptyObject

```go
rollbackSettingsUnion := components.CreateRollbackSettingsUnionEmptyObject(components.EmptyObject{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch rollbackSettingsUnion.Type {
	case components.RollbackSettingsUnionTypeRollbackSettings:
		// rollbackSettingsUnion.RollbackSettings is populated
	case components.RollbackSettingsUnionTypeEmptyObject:
		// rollbackSettingsUnion.EmptyObject is populated
}
```
