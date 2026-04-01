# OutputWebhook


## Supported Types

### OutputWebhookWebhook1

```go
outputWebhook := components.CreateOutputWebhookOutputWebhookWebhook1(components.OutputWebhookWebhook1{/* values here */})
```

### OutputWebhookWebhook2

```go
outputWebhook := components.CreateOutputWebhookOutputWebhookWebhook2(components.OutputWebhookWebhook2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch outputWebhook.Type {
	case components.OutputWebhookTypeOutputWebhookWebhook1:
		// outputWebhook.OutputWebhookWebhook1 is populated
	case components.OutputWebhookTypeOutputWebhookWebhook2:
		// outputWebhook.OutputWebhookWebhook2 is populated
}
```
