# RollbackSettingsUnion


## Supported Types

### RollbackSettings1

```go
rollbackSettingsUnion := components.CreateRollbackSettingsUnionRollbackSettings1(components.RollbackSettings1{/* values here */})
```

### RollbackSettings2

```go
rollbackSettingsUnion := components.CreateRollbackSettingsUnionRollbackSettings2(components.RollbackSettings2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch rollbackSettingsUnion.Type {
	case components.RollbackSettingsUnionTypeRollbackSettings1:
		// rollbackSettingsUnion.RollbackSettings1 is populated
	case components.RollbackSettingsUnionTypeRollbackSettings2:
		// rollbackSettingsUnion.RollbackSettings2 is populated
}
```
