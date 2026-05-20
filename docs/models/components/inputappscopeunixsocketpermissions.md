# InputAppscopeUNIXSocketPermissions

Permissions to set for socket e.g., 777. If empty, falls back to the runtime user's default permissions.


## Supported Types

### 

```go
inputAppscopeUNIXSocketPermissions := components.CreateInputAppscopeUNIXSocketPermissionsStr(string{/* values here */})
```

### 

```go
inputAppscopeUNIXSocketPermissions := components.CreateInputAppscopeUNIXSocketPermissionsNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch inputAppscopeUNIXSocketPermissions.Type {
	case components.InputAppscopeUNIXSocketPermissionsTypeStr:
		// inputAppscopeUNIXSocketPermissions.Str is populated
	case components.InputAppscopeUNIXSocketPermissionsTypeNumber:
		// inputAppscopeUNIXSocketPermissions.Number is populated
}
```
