# OutputResponseOutputWebhookUnion


## Supported Types

### OutputResponseOutputWebhookWebhook1

```go
outputResponseOutputWebhookUnion := components.CreateOutputResponseOutputWebhookUnionOutputResponseOutputWebhookWebhook1(components.OutputResponseOutputWebhookWebhook1{/* values here */})
```

### OutputResponseOutputWebhookWebhook2

```go
outputResponseOutputWebhookUnion := components.CreateOutputResponseOutputWebhookUnionOutputResponseOutputWebhookWebhook2(components.OutputResponseOutputWebhookWebhook2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch outputResponseOutputWebhookUnion.Type {
	case components.OutputResponseOutputWebhookUnionTypeOutputResponseOutputWebhookWebhook1:
		// outputResponseOutputWebhookUnion.OutputResponseOutputWebhookWebhook1 is populated
	case components.OutputResponseOutputWebhookUnionTypeOutputResponseOutputWebhookWebhook2:
		// outputResponseOutputWebhookUnion.OutputResponseOutputWebhookWebhook2 is populated
}
```
