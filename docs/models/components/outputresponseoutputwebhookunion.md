# OutputResponseOutputWebhookUnion


## Supported Types

### OutputResponseOutputWebhook1

```go
outputResponseOutputWebhookUnion := components.CreateOutputResponseOutputWebhookUnionOutputResponseOutputWebhook1(components.OutputResponseOutputWebhook1{/* values here */})
```

### OutputResponseOutputWebhook2

```go
outputResponseOutputWebhookUnion := components.CreateOutputResponseOutputWebhookUnionOutputResponseOutputWebhook2(components.OutputResponseOutputWebhook2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch outputResponseOutputWebhookUnion.Type {
	case components.OutputResponseOutputWebhookUnionTypeOutputResponseOutputWebhook1:
		// outputResponseOutputWebhookUnion.OutputResponseOutputWebhook1 is populated
	case components.OutputResponseOutputWebhookUnionTypeOutputResponseOutputWebhook2:
		// outputResponseOutputWebhookUnion.OutputResponseOutputWebhook2 is populated
}
```
