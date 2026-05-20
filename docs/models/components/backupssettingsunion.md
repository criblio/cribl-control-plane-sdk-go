# BackupsSettingsUnion


## Supported Types

### BackupsSettings

```go
backupsSettingsUnion := components.CreateBackupsSettingsUnionBackupsSettings(components.BackupsSettings{/* values here */})
```

### EmptyObject

```go
backupsSettingsUnion := components.CreateBackupsSettingsUnionEmptyObject(components.EmptyObject{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch backupsSettingsUnion.Type {
	case components.BackupsSettingsUnionTypeBackupsSettings:
		// backupsSettingsUnion.BackupsSettings is populated
	case components.BackupsSettingsUnionTypeEmptyObject:
		// backupsSettingsUnion.EmptyObject is populated
}
```
