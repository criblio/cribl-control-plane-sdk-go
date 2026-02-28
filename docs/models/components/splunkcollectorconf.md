# SplunkCollectorConf


## Supported Types

### SplunkAuthenticationNone

```go
splunkCollectorConf := components.CreateSplunkCollectorConfNone(components.SplunkAuthenticationNone{/* values here */})
```

### SplunkAuthenticationBasic

```go
splunkCollectorConf := components.CreateSplunkCollectorConfBasic(components.SplunkAuthenticationBasic{/* values here */})
```

### SplunkAuthenticationBasicSecret

```go
splunkCollectorConf := components.CreateSplunkCollectorConfBasicSecret(components.SplunkAuthenticationBasicSecret{/* values here */})
```

### SplunkAuthenticationToken

```go
splunkCollectorConf := components.CreateSplunkCollectorConfToken(components.SplunkAuthenticationToken{/* values here */})
```

### SplunkAuthenticationTokenSecret

```go
splunkCollectorConf := components.CreateSplunkCollectorConfTokenSecret(components.SplunkAuthenticationTokenSecret{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch splunkCollectorConf.Type {
	case components.SplunkCollectorConfTypeNone:
		// splunkCollectorConf.SplunkAuthenticationNone is populated
	case components.SplunkCollectorConfTypeBasic:
		// splunkCollectorConf.SplunkAuthenticationBasic is populated
	case components.SplunkCollectorConfTypeBasicSecret:
		// splunkCollectorConf.SplunkAuthenticationBasicSecret is populated
	case components.SplunkCollectorConfTypeToken:
		// splunkCollectorConf.SplunkAuthenticationToken is populated
	case components.SplunkCollectorConfTypeTokenSecret:
		// splunkCollectorConf.SplunkAuthenticationTokenSecret is populated
	default:
		// Unknown type - use splunkCollectorConf.GetUnknownRaw() for raw JSON
}
```
