# RestAuthenticationLoginDiscovery


## Supported Types

### RestAuthenticationLoginRestDiscoveryDiscoverTypeHTTP

```go
restAuthenticationLoginDiscovery := components.CreateRestAuthenticationLoginDiscoveryHTTP(components.RestAuthenticationLoginRestDiscoveryDiscoverTypeHTTP{/* values here */})
```

### RestAuthenticationLoginRestDiscoveryDiscoverTypeJSON

```go
restAuthenticationLoginDiscovery := components.CreateRestAuthenticationLoginDiscoveryJSON(components.RestAuthenticationLoginRestDiscoveryDiscoverTypeJSON{/* values here */})
```

### RestAuthenticationLoginRestDiscoveryDiscoverTypeList

```go
restAuthenticationLoginDiscovery := components.CreateRestAuthenticationLoginDiscoveryList(components.RestAuthenticationLoginRestDiscoveryDiscoverTypeList{/* values here */})
```

### RestAuthenticationLoginRestDiscoveryDiscoverTypeNone

```go
restAuthenticationLoginDiscovery := components.CreateRestAuthenticationLoginDiscoveryNone(components.RestAuthenticationLoginRestDiscoveryDiscoverTypeNone{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationLoginDiscovery.Type {
	case components.RestAuthenticationLoginDiscoveryTypeHTTP:
		// restAuthenticationLoginDiscovery.RestAuthenticationLoginRestDiscoveryDiscoverTypeHTTP is populated
	case components.RestAuthenticationLoginDiscoveryTypeJSON:
		// restAuthenticationLoginDiscovery.RestAuthenticationLoginRestDiscoveryDiscoverTypeJSON is populated
	case components.RestAuthenticationLoginDiscoveryTypeList:
		// restAuthenticationLoginDiscovery.RestAuthenticationLoginRestDiscoveryDiscoverTypeList is populated
	case components.RestAuthenticationLoginDiscoveryTypeNone:
		// restAuthenticationLoginDiscovery.RestAuthenticationLoginRestDiscoveryDiscoverTypeNone is populated
	default:
		// Unknown type - use restAuthenticationLoginDiscovery.GetUnknownRaw() for raw JSON
}
```
