# CreateInputUNIXSocketPermissions

Permissions to set for socket e.g., 777. If empty, falls back to the runtime user's default permissions.


## Supported Types

### 

```go
createInputUNIXSocketPermissions := operations.CreateCreateInputUNIXSocketPermissionsStr(string{/* values here */})
```

### 

```go
createInputUNIXSocketPermissions := operations.CreateCreateInputUNIXSocketPermissionsNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createInputUNIXSocketPermissions.Type {
	case operations.CreateInputUNIXSocketPermissionsTypeStr:
		// createInputUNIXSocketPermissions.Str is populated
	case operations.CreateInputUNIXSocketPermissionsTypeNumber:
		// createInputUNIXSocketPermissions.Number is populated
}
```
