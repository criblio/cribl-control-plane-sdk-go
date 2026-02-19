# RestAuthenticationOauthDiscovery


## Supported Types

### RestAuthenticationOauthRestDiscoveryDiscoverTypeHTTP

```go
restAuthenticationOauthDiscovery := components.CreateRestAuthenticationOauthDiscoveryHTTP(components.RestAuthenticationOauthRestDiscoveryDiscoverTypeHTTP{/* values here */})
```

### RestAuthenticationOauthRestDiscoveryDiscoverTypeJSON

```go
restAuthenticationOauthDiscovery := components.CreateRestAuthenticationOauthDiscoveryJSON(components.RestAuthenticationOauthRestDiscoveryDiscoverTypeJSON{/* values here */})
```

### RestAuthenticationOauthRestDiscoveryDiscoverTypeList

```go
restAuthenticationOauthDiscovery := components.CreateRestAuthenticationOauthDiscoveryList(components.RestAuthenticationOauthRestDiscoveryDiscoverTypeList{/* values here */})
```

### RestAuthenticationOauthRestDiscoveryDiscoverTypeNone

```go
restAuthenticationOauthDiscovery := components.CreateRestAuthenticationOauthDiscoveryNone(components.RestAuthenticationOauthRestDiscoveryDiscoverTypeNone{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationOauthDiscovery.Type {
	case components.RestAuthenticationOauthDiscoveryTypeHTTP:
		// restAuthenticationOauthDiscovery.RestAuthenticationOauthRestDiscoveryDiscoverTypeHTTP is populated
	case components.RestAuthenticationOauthDiscoveryTypeJSON:
		// restAuthenticationOauthDiscovery.RestAuthenticationOauthRestDiscoveryDiscoverTypeJSON is populated
	case components.RestAuthenticationOauthDiscoveryTypeList:
		// restAuthenticationOauthDiscovery.RestAuthenticationOauthRestDiscoveryDiscoverTypeList is populated
	case components.RestAuthenticationOauthDiscoveryTypeNone:
		// restAuthenticationOauthDiscovery.RestAuthenticationOauthRestDiscoveryDiscoverTypeNone is populated
	default:
		// Unknown type - use restAuthenticationOauthDiscovery.GetUnknownRaw() for raw JSON
}
```
