# RestCollectorConf


## Supported Types

### RestAuthenticationNone

```go
restCollectorConf := components.CreateRestCollectorConfNone(components.RestAuthenticationNone{/* values here */})
```

### RestAuthenticationBasic

```go
restCollectorConf := components.CreateRestCollectorConfBasic(components.RestAuthenticationBasic{/* values here */})
```

### RestAuthenticationBasicSecret

```go
restCollectorConf := components.CreateRestCollectorConfBasicSecret(components.RestAuthenticationBasicSecret{/* values here */})
```

### RestAuthenticationLogin

```go
restCollectorConf := components.CreateRestCollectorConfLogin(components.RestAuthenticationLogin{/* values here */})
```

### RestAuthenticationLoginSecret

```go
restCollectorConf := components.CreateRestCollectorConfLoginSecret(components.RestAuthenticationLoginSecret{/* values here */})
```

### RestAuthenticationOauth

```go
restCollectorConf := components.CreateRestCollectorConfOauth(components.RestAuthenticationOauth{/* values here */})
```

### RestAuthenticationOauthSecret

```go
restCollectorConf := components.CreateRestCollectorConfOauthSecret(components.RestAuthenticationOauthSecret{/* values here */})
```

### RestAuthenticationGoogleOauth

```go
restCollectorConf := components.CreateRestCollectorConfGoogleOauth(components.RestAuthenticationGoogleOauth{/* values here */})
```

### RestAuthenticationGoogleOauthSecret

```go
restCollectorConf := components.CreateRestCollectorConfGoogleOauthSecret(components.RestAuthenticationGoogleOauthSecret{/* values here */})
```

### RestAuthenticationHmac

```go
restCollectorConf := components.CreateRestCollectorConfHmac(components.RestAuthenticationHmac{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restCollectorConf.Type {
	case components.RestCollectorConfTypeNone:
		// restCollectorConf.RestAuthenticationNone is populated
	case components.RestCollectorConfTypeBasic:
		// restCollectorConf.RestAuthenticationBasic is populated
	case components.RestCollectorConfTypeBasicSecret:
		// restCollectorConf.RestAuthenticationBasicSecret is populated
	case components.RestCollectorConfTypeLogin:
		// restCollectorConf.RestAuthenticationLogin is populated
	case components.RestCollectorConfTypeLoginSecret:
		// restCollectorConf.RestAuthenticationLoginSecret is populated
	case components.RestCollectorConfTypeOauth:
		// restCollectorConf.RestAuthenticationOauth is populated
	case components.RestCollectorConfTypeOauthSecret:
		// restCollectorConf.RestAuthenticationOauthSecret is populated
	case components.RestCollectorConfTypeGoogleOauth:
		// restCollectorConf.RestAuthenticationGoogleOauth is populated
	case components.RestCollectorConfTypeGoogleOauthSecret:
		// restCollectorConf.RestAuthenticationGoogleOauthSecret is populated
	case components.RestCollectorConfTypeHmac:
		// restCollectorConf.RestAuthenticationHmac is populated
	default:
		// Unknown type - use restCollectorConf.GetUnknownRaw() for raw JSON
}
```
