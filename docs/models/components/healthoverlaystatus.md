# HealthOverlayStatus


## Supported Types

### ActiveHealthOverlayStatus

```go
healthOverlayStatus := components.CreateHealthOverlayStatusActiveHealthOverlayStatus(components.ActiveHealthOverlayStatus{/* values here */})
```

### NoActiveHealthOverlayStatus

```go
healthOverlayStatus := components.CreateHealthOverlayStatusNoActiveHealthOverlayStatus(components.NoActiveHealthOverlayStatus{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch healthOverlayStatus.Type {
	case components.HealthOverlayStatusTypeActiveHealthOverlayStatus:
		// healthOverlayStatus.ActiveHealthOverlayStatus is populated
	case components.HealthOverlayStatusTypeNoActiveHealthOverlayStatus:
		// healthOverlayStatus.NoActiveHealthOverlayStatus is populated
}
```
