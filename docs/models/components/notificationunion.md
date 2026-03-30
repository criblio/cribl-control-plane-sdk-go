# NotificationUnion


## Supported Types

### Notification1

```go
notificationUnion := components.CreateNotificationUnionNotification1(components.Notification1{/* values here */})
```

### Notification2

```go
notificationUnion := components.CreateNotificationUnionNotification2(components.Notification2{/* values here */})
```

### Notification3

```go
notificationUnion := components.CreateNotificationUnionNotification3(components.Notification3{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch notificationUnion.Type {
	case components.NotificationUnionTypeNotification1:
		// notificationUnion.Notification1 is populated
	case components.NotificationUnionTypeNotification2:
		// notificationUnion.Notification2 is populated
	case components.NotificationUnionTypeNotification3:
		// notificationUnion.Notification3 is populated
}
```
