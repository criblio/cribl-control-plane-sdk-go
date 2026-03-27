# RestCollectMethodGetRestDiscoveryDiscoverTypeHTTP


## Supported Types

### RestCollectMethodGetRestDiscoveryDiscoverTypeHTTPDiscoverMethodGet

```go
restCollectMethodGetRestDiscoveryDiscoverTypeHTTP := components.CreateRestCollectMethodGetRestDiscoveryDiscoverTypeHTTPGet(components.RestCollectMethodGetRestDiscoveryDiscoverTypeHTTPDiscoverMethodGet{/* values here */})
```

### RestCollectMethodGetRestDiscoveryDiscoverTypeHTTPDiscoverMethodPost

```go
restCollectMethodGetRestDiscoveryDiscoverTypeHTTP := components.CreateRestCollectMethodGetRestDiscoveryDiscoverTypeHTTPPost(components.RestCollectMethodGetRestDiscoveryDiscoverTypeHTTPDiscoverMethodPost{/* values here */})
```

### RestCollectMethodGetRestDiscoveryDiscoverTypeHTTPDiscoverMethodPostWithBody

```go
restCollectMethodGetRestDiscoveryDiscoverTypeHTTP := components.CreateRestCollectMethodGetRestDiscoveryDiscoverTypeHTTPPostWithBody(components.RestCollectMethodGetRestDiscoveryDiscoverTypeHTTPDiscoverMethodPostWithBody{/* values here */})
```

### RestCollectMethodGetRestDiscoveryDiscoverTypeHTTPDiscoverMethodOther

```go
restCollectMethodGetRestDiscoveryDiscoverTypeHTTP := components.CreateRestCollectMethodGetRestDiscoveryDiscoverTypeHTTPOther(components.RestCollectMethodGetRestDiscoveryDiscoverTypeHTTPDiscoverMethodOther{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restCollectMethodGetRestDiscoveryDiscoverTypeHTTP.Type {
	case components.RestCollectMethodGetRestDiscoveryDiscoverTypeHTTPTypeGet:
		// restCollectMethodGetRestDiscoveryDiscoverTypeHTTP.RestCollectMethodGetRestDiscoveryDiscoverTypeHTTPDiscoverMethodGet is populated
	case components.RestCollectMethodGetRestDiscoveryDiscoverTypeHTTPTypePost:
		// restCollectMethodGetRestDiscoveryDiscoverTypeHTTP.RestCollectMethodGetRestDiscoveryDiscoverTypeHTTPDiscoverMethodPost is populated
	case components.RestCollectMethodGetRestDiscoveryDiscoverTypeHTTPTypePostWithBody:
		// restCollectMethodGetRestDiscoveryDiscoverTypeHTTP.RestCollectMethodGetRestDiscoveryDiscoverTypeHTTPDiscoverMethodPostWithBody is populated
	case components.RestCollectMethodGetRestDiscoveryDiscoverTypeHTTPTypeOther:
		// restCollectMethodGetRestDiscoveryDiscoverTypeHTTP.RestCollectMethodGetRestDiscoveryDiscoverTypeHTTPDiscoverMethodOther is populated
	default:
		// Unknown type - use restCollectMethodGetRestDiscoveryDiscoverTypeHTTP.GetUnknownRaw() for raw JSON
}
```
