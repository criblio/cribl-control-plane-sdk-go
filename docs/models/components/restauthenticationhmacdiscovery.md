# RestAuthenticationHmacDiscovery


## Supported Types

### RestAuthenticationHmacRestDiscoveryDiscoverTypeHTTP

```go
restAuthenticationHmacDiscovery := components.CreateRestAuthenticationHmacDiscoveryHTTP(components.RestAuthenticationHmacRestDiscoveryDiscoverTypeHTTP{/* values here */})
```

### RestAuthenticationHmacRestDiscoveryDiscoverTypeJSON

```go
restAuthenticationHmacDiscovery := components.CreateRestAuthenticationHmacDiscoveryJSON(components.RestAuthenticationHmacRestDiscoveryDiscoverTypeJSON{/* values here */})
```

### RestAuthenticationHmacRestDiscoveryDiscoverTypeList

```go
restAuthenticationHmacDiscovery := components.CreateRestAuthenticationHmacDiscoveryList(components.RestAuthenticationHmacRestDiscoveryDiscoverTypeList{/* values here */})
```

### RestAuthenticationHmacRestDiscoveryDiscoverTypeNone

```go
restAuthenticationHmacDiscovery := components.CreateRestAuthenticationHmacDiscoveryNone(components.RestAuthenticationHmacRestDiscoveryDiscoverTypeNone{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationHmacDiscovery.Type {
	case components.RestAuthenticationHmacDiscoveryTypeHTTP:
		// restAuthenticationHmacDiscovery.RestAuthenticationHmacRestDiscoveryDiscoverTypeHTTP is populated
	case components.RestAuthenticationHmacDiscoveryTypeJSON:
		// restAuthenticationHmacDiscovery.RestAuthenticationHmacRestDiscoveryDiscoverTypeJSON is populated
	case components.RestAuthenticationHmacDiscoveryTypeList:
		// restAuthenticationHmacDiscovery.RestAuthenticationHmacRestDiscoveryDiscoverTypeList is populated
	case components.RestAuthenticationHmacDiscoveryTypeNone:
		// restAuthenticationHmacDiscovery.RestAuthenticationHmacRestDiscoveryDiscoverTypeNone is populated
	default:
		// Unknown type - use restAuthenticationHmacDiscovery.GetUnknownRaw() for raw JSON
}
```
