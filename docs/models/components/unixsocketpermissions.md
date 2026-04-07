# UNIXSocketPermissions

Permissions to set for socket e.g., 777. If empty, falls back to the runtime user's default permissions.


## Supported Types

### 

```go
unixSocketPermissions := components.CreateUNIXSocketPermissionsStr(string{/* values here */})
```

### 

```go
unixSocketPermissions := components.CreateUNIXSocketPermissionsNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch unixSocketPermissions.Type {
	case components.UNIXSocketPermissionsTypeStr:
		// unixSocketPermissions.Str is populated
	case components.UNIXSocketPermissionsTypeNumber:
		// unixSocketPermissions.Number is populated
}
```
