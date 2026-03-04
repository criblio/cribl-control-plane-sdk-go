# NotificationTargetConfigUnion


## Supported Types

### NotificationTargetConfig

```go
notificationTargetConfigUnion := components.CreateNotificationTargetConfigUnionNotificationTargetConfig(components.NotificationTargetConfig{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch notificationTargetConfigUnion.Type {
	case components.NotificationTargetConfigUnionTypeNotificationTargetConfig:
		// notificationTargetConfigUnion.NotificationTargetConfig is populated
}
```
