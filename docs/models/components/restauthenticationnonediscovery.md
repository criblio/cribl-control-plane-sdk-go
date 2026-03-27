# RestAuthenticationNoneDiscovery


## Supported Types

### RestAuthenticationNoneRestDiscoveryDiscoverTypeHTTP

```go
restAuthenticationNoneDiscovery := components.CreateRestAuthenticationNoneDiscoveryHTTP(components.RestAuthenticationNoneRestDiscoveryDiscoverTypeHTTP{/* values here */})
```

### RestAuthenticationNoneRestDiscoveryDiscoverTypeJSON

```go
restAuthenticationNoneDiscovery := components.CreateRestAuthenticationNoneDiscoveryJSON(components.RestAuthenticationNoneRestDiscoveryDiscoverTypeJSON{/* values here */})
```

### RestAuthenticationNoneRestDiscoveryDiscoverTypeList

```go
restAuthenticationNoneDiscovery := components.CreateRestAuthenticationNoneDiscoveryList(components.RestAuthenticationNoneRestDiscoveryDiscoverTypeList{/* values here */})
```

### RestAuthenticationNoneRestDiscoveryDiscoverTypeNone

```go
restAuthenticationNoneDiscovery := components.CreateRestAuthenticationNoneDiscoveryNone(components.RestAuthenticationNoneRestDiscoveryDiscoverTypeNone{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationNoneDiscovery.Type {
	case components.RestAuthenticationNoneDiscoveryTypeHTTP:
		// restAuthenticationNoneDiscovery.RestAuthenticationNoneRestDiscoveryDiscoverTypeHTTP is populated
	case components.RestAuthenticationNoneDiscoveryTypeJSON:
		// restAuthenticationNoneDiscovery.RestAuthenticationNoneRestDiscoveryDiscoverTypeJSON is populated
	case components.RestAuthenticationNoneDiscoveryTypeList:
		// restAuthenticationNoneDiscovery.RestAuthenticationNoneRestDiscoveryDiscoverTypeList is populated
	case components.RestAuthenticationNoneDiscoveryTypeNone:
		// restAuthenticationNoneDiscovery.RestAuthenticationNoneRestDiscoveryDiscoverTypeNone is populated
	default:
		// Unknown type - use restAuthenticationNoneDiscovery.GetUnknownRaw() for raw JSON
}
```
