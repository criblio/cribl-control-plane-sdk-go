# CreateOutputSystemByPackOutputWebhookUnion


## Supported Types

### CreateOutputSystemByPackOutputWebhookWebhook1

```go
createOutputSystemByPackOutputWebhookUnion := operations.CreateCreateOutputSystemByPackOutputWebhookUnionCreateOutputSystemByPackOutputWebhookWebhook1(operations.CreateOutputSystemByPackOutputWebhookWebhook1{/* values here */})
```

### CreateOutputSystemByPackOutputWebhookWebhook2

```go
createOutputSystemByPackOutputWebhookUnion := operations.CreateCreateOutputSystemByPackOutputWebhookUnionCreateOutputSystemByPackOutputWebhookWebhook2(operations.CreateOutputSystemByPackOutputWebhookWebhook2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createOutputSystemByPackOutputWebhookUnion.Type {
	case operations.CreateOutputSystemByPackOutputWebhookUnionTypeCreateOutputSystemByPackOutputWebhookWebhook1:
		// createOutputSystemByPackOutputWebhookUnion.CreateOutputSystemByPackOutputWebhookWebhook1 is populated
	case operations.CreateOutputSystemByPackOutputWebhookUnionTypeCreateOutputSystemByPackOutputWebhookWebhook2:
		// createOutputSystemByPackOutputWebhookUnion.CreateOutputSystemByPackOutputWebhookWebhook2 is populated
}
```
