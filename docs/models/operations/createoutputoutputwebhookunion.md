# CreateOutputOutputWebhookUnion


## Supported Types

### CreateOutputOutputWebhookWebhook1

```go
createOutputOutputWebhookUnion := operations.CreateCreateOutputOutputWebhookUnionCreateOutputOutputWebhookWebhook1(operations.CreateOutputOutputWebhookWebhook1{/* values here */})
```

### CreateOutputOutputWebhookWebhook2

```go
createOutputOutputWebhookUnion := operations.CreateCreateOutputOutputWebhookUnionCreateOutputOutputWebhookWebhook2(operations.CreateOutputOutputWebhookWebhook2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createOutputOutputWebhookUnion.Type {
	case operations.CreateOutputOutputWebhookUnionTypeCreateOutputOutputWebhookWebhook1:
		// createOutputOutputWebhookUnion.CreateOutputOutputWebhookWebhook1 is populated
	case operations.CreateOutputOutputWebhookUnionTypeCreateOutputOutputWebhookWebhook2:
		// createOutputOutputWebhookUnion.CreateOutputOutputWebhookWebhook2 is populated
}
```
