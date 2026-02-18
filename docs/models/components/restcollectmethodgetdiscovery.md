# RestCollectMethodGetDiscovery


## Supported Types

### RestCollectMethodGetRestDiscoveryDiscoverTypeHTTP

```go
restCollectMethodGetDiscovery := components.CreateRestCollectMethodGetDiscoveryHTTP(components.RestCollectMethodGetRestDiscoveryDiscoverTypeHTTP{/* values here */})
```

### RestCollectMethodGetRestDiscoveryDiscoverTypeJSON

```go
restCollectMethodGetDiscovery := components.CreateRestCollectMethodGetDiscoveryJSON(components.RestCollectMethodGetRestDiscoveryDiscoverTypeJSON{/* values here */})
```

### RestCollectMethodGetRestDiscoveryDiscoverTypeList

```go
restCollectMethodGetDiscovery := components.CreateRestCollectMethodGetDiscoveryList(components.RestCollectMethodGetRestDiscoveryDiscoverTypeList{/* values here */})
```

### RestCollectMethodGetRestDiscoveryDiscoverTypeNone

```go
restCollectMethodGetDiscovery := components.CreateRestCollectMethodGetDiscoveryNone(components.RestCollectMethodGetRestDiscoveryDiscoverTypeNone{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restCollectMethodGetDiscovery.Type {
	case components.RestCollectMethodGetDiscoveryTypeHTTP:
		// restCollectMethodGetDiscovery.RestCollectMethodGetRestDiscoveryDiscoverTypeHTTP is populated
	case components.RestCollectMethodGetDiscoveryTypeJSON:
		// restCollectMethodGetDiscovery.RestCollectMethodGetRestDiscoveryDiscoverTypeJSON is populated
	case components.RestCollectMethodGetDiscoveryTypeList:
		// restCollectMethodGetDiscovery.RestCollectMethodGetRestDiscoveryDiscoverTypeList is populated
	case components.RestCollectMethodGetDiscoveryTypeNone:
		// restCollectMethodGetDiscovery.RestCollectMethodGetRestDiscoveryDiscoverTypeNone is populated
	default:
		// Unknown type - use restCollectMethodGetDiscovery.GetUnknownRaw() for raw JSON
}
```
