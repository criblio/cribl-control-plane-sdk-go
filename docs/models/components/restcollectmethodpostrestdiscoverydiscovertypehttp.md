# RestCollectMethodPostRestDiscoveryDiscoverTypeHTTP


## Supported Types

### RestCollectMethodPostRestDiscoveryDiscoverTypeHTTPDiscoverMethodGet

```go
restCollectMethodPostRestDiscoveryDiscoverTypeHTTP := components.CreateRestCollectMethodPostRestDiscoveryDiscoverTypeHTTPGet(components.RestCollectMethodPostRestDiscoveryDiscoverTypeHTTPDiscoverMethodGet{/* values here */})
```

### RestCollectMethodPostRestDiscoveryDiscoverTypeHTTPDiscoverMethodPost

```go
restCollectMethodPostRestDiscoveryDiscoverTypeHTTP := components.CreateRestCollectMethodPostRestDiscoveryDiscoverTypeHTTPPost(components.RestCollectMethodPostRestDiscoveryDiscoverTypeHTTPDiscoverMethodPost{/* values here */})
```

### RestCollectMethodPostRestDiscoveryDiscoverTypeHTTPDiscoverMethodPostWithBody

```go
restCollectMethodPostRestDiscoveryDiscoverTypeHTTP := components.CreateRestCollectMethodPostRestDiscoveryDiscoverTypeHTTPPostWithBody(components.RestCollectMethodPostRestDiscoveryDiscoverTypeHTTPDiscoverMethodPostWithBody{/* values here */})
```

### RestCollectMethodPostRestDiscoveryDiscoverTypeHTTPDiscoverMethodOther

```go
restCollectMethodPostRestDiscoveryDiscoverTypeHTTP := components.CreateRestCollectMethodPostRestDiscoveryDiscoverTypeHTTPOther(components.RestCollectMethodPostRestDiscoveryDiscoverTypeHTTPDiscoverMethodOther{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restCollectMethodPostRestDiscoveryDiscoverTypeHTTP.Type {
	case components.RestCollectMethodPostRestDiscoveryDiscoverTypeHTTPTypeGet:
		// restCollectMethodPostRestDiscoveryDiscoverTypeHTTP.RestCollectMethodPostRestDiscoveryDiscoverTypeHTTPDiscoverMethodGet is populated
	case components.RestCollectMethodPostRestDiscoveryDiscoverTypeHTTPTypePost:
		// restCollectMethodPostRestDiscoveryDiscoverTypeHTTP.RestCollectMethodPostRestDiscoveryDiscoverTypeHTTPDiscoverMethodPost is populated
	case components.RestCollectMethodPostRestDiscoveryDiscoverTypeHTTPTypePostWithBody:
		// restCollectMethodPostRestDiscoveryDiscoverTypeHTTP.RestCollectMethodPostRestDiscoveryDiscoverTypeHTTPDiscoverMethodPostWithBody is populated
	case components.RestCollectMethodPostRestDiscoveryDiscoverTypeHTTPTypeOther:
		// restCollectMethodPostRestDiscoveryDiscoverTypeHTTP.RestCollectMethodPostRestDiscoveryDiscoverTypeHTTPDiscoverMethodOther is populated
	default:
		// Unknown type - use restCollectMethodPostRestDiscoveryDiscoverTypeHTTP.GetUnknownRaw() for raw JSON
}
```
