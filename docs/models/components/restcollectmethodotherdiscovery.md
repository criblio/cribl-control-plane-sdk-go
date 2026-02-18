# RestCollectMethodOtherDiscovery


## Supported Types

### RestCollectMethodOtherRestDiscoveryDiscoverTypeHTTP

```go
restCollectMethodOtherDiscovery := components.CreateRestCollectMethodOtherDiscoveryHTTP(components.RestCollectMethodOtherRestDiscoveryDiscoverTypeHTTP{/* values here */})
```

### RestCollectMethodOtherRestDiscoveryDiscoverTypeJSON

```go
restCollectMethodOtherDiscovery := components.CreateRestCollectMethodOtherDiscoveryJSON(components.RestCollectMethodOtherRestDiscoveryDiscoverTypeJSON{/* values here */})
```

### RestCollectMethodOtherRestDiscoveryDiscoverTypeList

```go
restCollectMethodOtherDiscovery := components.CreateRestCollectMethodOtherDiscoveryList(components.RestCollectMethodOtherRestDiscoveryDiscoverTypeList{/* values here */})
```

### RestCollectMethodOtherRestDiscoveryDiscoverTypeNone

```go
restCollectMethodOtherDiscovery := components.CreateRestCollectMethodOtherDiscoveryNone(components.RestCollectMethodOtherRestDiscoveryDiscoverTypeNone{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restCollectMethodOtherDiscovery.Type {
	case components.RestCollectMethodOtherDiscoveryTypeHTTP:
		// restCollectMethodOtherDiscovery.RestCollectMethodOtherRestDiscoveryDiscoverTypeHTTP is populated
	case components.RestCollectMethodOtherDiscoveryTypeJSON:
		// restCollectMethodOtherDiscovery.RestCollectMethodOtherRestDiscoveryDiscoverTypeJSON is populated
	case components.RestCollectMethodOtherDiscoveryTypeList:
		// restCollectMethodOtherDiscovery.RestCollectMethodOtherRestDiscoveryDiscoverTypeList is populated
	case components.RestCollectMethodOtherDiscoveryTypeNone:
		// restCollectMethodOtherDiscovery.RestCollectMethodOtherRestDiscoveryDiscoverTypeNone is populated
	default:
		// Unknown type - use restCollectMethodOtherDiscovery.GetUnknownRaw() for raw JSON
}
```
