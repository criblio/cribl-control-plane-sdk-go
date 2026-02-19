# RestCollectMethodPostDiscovery


## Supported Types

### RestCollectMethodPostRestDiscoveryDiscoverTypeHTTP

```go
restCollectMethodPostDiscovery := components.CreateRestCollectMethodPostDiscoveryHTTP(components.RestCollectMethodPostRestDiscoveryDiscoverTypeHTTP{/* values here */})
```

### RestCollectMethodPostRestDiscoveryDiscoverTypeJSON

```go
restCollectMethodPostDiscovery := components.CreateRestCollectMethodPostDiscoveryJSON(components.RestCollectMethodPostRestDiscoveryDiscoverTypeJSON{/* values here */})
```

### RestCollectMethodPostRestDiscoveryDiscoverTypeList

```go
restCollectMethodPostDiscovery := components.CreateRestCollectMethodPostDiscoveryList(components.RestCollectMethodPostRestDiscoveryDiscoverTypeList{/* values here */})
```

### RestCollectMethodPostRestDiscoveryDiscoverTypeNone

```go
restCollectMethodPostDiscovery := components.CreateRestCollectMethodPostDiscoveryNone(components.RestCollectMethodPostRestDiscoveryDiscoverTypeNone{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restCollectMethodPostDiscovery.Type {
	case components.RestCollectMethodPostDiscoveryTypeHTTP:
		// restCollectMethodPostDiscovery.RestCollectMethodPostRestDiscoveryDiscoverTypeHTTP is populated
	case components.RestCollectMethodPostDiscoveryTypeJSON:
		// restCollectMethodPostDiscovery.RestCollectMethodPostRestDiscoveryDiscoverTypeJSON is populated
	case components.RestCollectMethodPostDiscoveryTypeList:
		// restCollectMethodPostDiscovery.RestCollectMethodPostRestDiscoveryDiscoverTypeList is populated
	case components.RestCollectMethodPostDiscoveryTypeNone:
		// restCollectMethodPostDiscovery.RestCollectMethodPostRestDiscoveryDiscoverTypeNone is populated
	default:
		// Unknown type - use restCollectMethodPostDiscovery.GetUnknownRaw() for raw JSON
}
```
