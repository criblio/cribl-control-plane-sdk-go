# RestCollectMethodPostWithBodyDiscovery


## Supported Types

### RestCollectMethodPostWithBodyRestDiscoveryDiscoverTypeHTTP

```go
restCollectMethodPostWithBodyDiscovery := components.CreateRestCollectMethodPostWithBodyDiscoveryHTTP(components.RestCollectMethodPostWithBodyRestDiscoveryDiscoverTypeHTTP{/* values here */})
```

### RestCollectMethodPostWithBodyRestDiscoveryDiscoverTypeJSON

```go
restCollectMethodPostWithBodyDiscovery := components.CreateRestCollectMethodPostWithBodyDiscoveryJSON(components.RestCollectMethodPostWithBodyRestDiscoveryDiscoverTypeJSON{/* values here */})
```

### RestCollectMethodPostWithBodyRestDiscoveryDiscoverTypeList

```go
restCollectMethodPostWithBodyDiscovery := components.CreateRestCollectMethodPostWithBodyDiscoveryList(components.RestCollectMethodPostWithBodyRestDiscoveryDiscoverTypeList{/* values here */})
```

### RestCollectMethodPostWithBodyRestDiscoveryDiscoverTypeNone

```go
restCollectMethodPostWithBodyDiscovery := components.CreateRestCollectMethodPostWithBodyDiscoveryNone(components.RestCollectMethodPostWithBodyRestDiscoveryDiscoverTypeNone{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restCollectMethodPostWithBodyDiscovery.Type {
	case components.RestCollectMethodPostWithBodyDiscoveryTypeHTTP:
		// restCollectMethodPostWithBodyDiscovery.RestCollectMethodPostWithBodyRestDiscoveryDiscoverTypeHTTP is populated
	case components.RestCollectMethodPostWithBodyDiscoveryTypeJSON:
		// restCollectMethodPostWithBodyDiscovery.RestCollectMethodPostWithBodyRestDiscoveryDiscoverTypeJSON is populated
	case components.RestCollectMethodPostWithBodyDiscoveryTypeList:
		// restCollectMethodPostWithBodyDiscovery.RestCollectMethodPostWithBodyRestDiscoveryDiscoverTypeList is populated
	case components.RestCollectMethodPostWithBodyDiscoveryTypeNone:
		// restCollectMethodPostWithBodyDiscovery.RestCollectMethodPostWithBodyRestDiscoveryDiscoverTypeNone is populated
	default:
		// Unknown type - use restCollectMethodPostWithBodyDiscovery.GetUnknownRaw() for raw JSON
}
```
