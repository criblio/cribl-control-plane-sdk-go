# InputResponseUNIXSocketPermissions

Permissions to set for socket e.g., 777. If empty, falls back to the runtime user's default permissions.


## Supported Types

### 

```go
inputResponseUNIXSocketPermissions := components.CreateInputResponseUNIXSocketPermissionsStr(string{/* values here */})
```

### 

```go
inputResponseUNIXSocketPermissions := components.CreateInputResponseUNIXSocketPermissionsNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch inputResponseUNIXSocketPermissions.Type {
	case components.InputResponseUNIXSocketPermissionsTypeStr:
		// inputResponseUNIXSocketPermissions.Str is populated
	case components.InputResponseUNIXSocketPermissionsTypeNumber:
		// inputResponseUNIXSocketPermissions.Number is populated
}
```
