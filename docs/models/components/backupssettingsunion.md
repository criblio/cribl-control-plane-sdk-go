# BackupsSettingsUnion


## Supported Types

### BackupsSettings1

```go
backupsSettingsUnion := components.CreateBackupsSettingsUnionBackupsSettings1(components.BackupsSettings1{/* values here */})
```

### BackupsSettings2

```go
backupsSettingsUnion := components.CreateBackupsSettingsUnionBackupsSettings2(components.BackupsSettings2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch backupsSettingsUnion.Type {
	case components.BackupsSettingsUnionTypeBackupsSettings1:
		// backupsSettingsUnion.BackupsSettings1 is populated
	case components.BackupsSettingsUnionTypeBackupsSettings2:
		// backupsSettingsUnion.BackupsSettings2 is populated
}
```
