# RestAuthenticationBasicDiscovery


## Supported Types

### RestAuthenticationBasicRestDiscoveryDiscoverTypeHTTP

```go
restAuthenticationBasicDiscovery := components.CreateRestAuthenticationBasicDiscoveryHTTP(components.RestAuthenticationBasicRestDiscoveryDiscoverTypeHTTP{/* values here */})
```

### RestAuthenticationBasicRestDiscoveryDiscoverTypeJSON

```go
restAuthenticationBasicDiscovery := components.CreateRestAuthenticationBasicDiscoveryJSON(components.RestAuthenticationBasicRestDiscoveryDiscoverTypeJSON{/* values here */})
```

### RestAuthenticationBasicRestDiscoveryDiscoverTypeList

```go
restAuthenticationBasicDiscovery := components.CreateRestAuthenticationBasicDiscoveryList(components.RestAuthenticationBasicRestDiscoveryDiscoverTypeList{/* values here */})
```

### RestAuthenticationBasicRestDiscoveryDiscoverTypeNone

```go
restAuthenticationBasicDiscovery := components.CreateRestAuthenticationBasicDiscoveryNone(components.RestAuthenticationBasicRestDiscoveryDiscoverTypeNone{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationBasicDiscovery.Type {
	case components.RestAuthenticationBasicDiscoveryTypeHTTP:
		// restAuthenticationBasicDiscovery.RestAuthenticationBasicRestDiscoveryDiscoverTypeHTTP is populated
	case components.RestAuthenticationBasicDiscoveryTypeJSON:
		// restAuthenticationBasicDiscovery.RestAuthenticationBasicRestDiscoveryDiscoverTypeJSON is populated
	case components.RestAuthenticationBasicDiscoveryTypeList:
		// restAuthenticationBasicDiscovery.RestAuthenticationBasicRestDiscoveryDiscoverTypeList is populated
	case components.RestAuthenticationBasicDiscoveryTypeNone:
		// restAuthenticationBasicDiscovery.RestAuthenticationBasicRestDiscoveryDiscoverTypeNone is populated
	default:
		// Unknown type - use restAuthenticationBasicDiscovery.GetUnknownRaw() for raw JSON
}
```
