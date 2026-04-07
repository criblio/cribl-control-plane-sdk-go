# CreateInputSystemByPackUNIXSocketPermissions

Permissions to set for socket e.g., 777. If empty, falls back to the runtime user's default permissions.


## Supported Types

### 

```go
createInputSystemByPackUNIXSocketPermissions := operations.CreateCreateInputSystemByPackUNIXSocketPermissionsStr(string{/* values here */})
```

### 

```go
createInputSystemByPackUNIXSocketPermissions := operations.CreateCreateInputSystemByPackUNIXSocketPermissionsNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createInputSystemByPackUNIXSocketPermissions.Type {
	case operations.CreateInputSystemByPackUNIXSocketPermissionsTypeStr:
		// createInputSystemByPackUNIXSocketPermissions.Str is populated
	case operations.CreateInputSystemByPackUNIXSocketPermissionsTypeNumber:
		// createInputSystemByPackUNIXSocketPermissions.Number is populated
}
```
