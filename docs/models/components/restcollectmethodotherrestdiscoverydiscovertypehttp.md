# RestCollectMethodOtherRestDiscoveryDiscoverTypeHTTP


## Supported Types

### RestCollectMethodOtherRestDiscoveryDiscoverTypeHTTPDiscoverMethodGet

```go
restCollectMethodOtherRestDiscoveryDiscoverTypeHTTP := components.CreateRestCollectMethodOtherRestDiscoveryDiscoverTypeHTTPGet(components.RestCollectMethodOtherRestDiscoveryDiscoverTypeHTTPDiscoverMethodGet{/* values here */})
```

### RestCollectMethodOtherRestDiscoveryDiscoverTypeHTTPDiscoverMethodPost

```go
restCollectMethodOtherRestDiscoveryDiscoverTypeHTTP := components.CreateRestCollectMethodOtherRestDiscoveryDiscoverTypeHTTPPost(components.RestCollectMethodOtherRestDiscoveryDiscoverTypeHTTPDiscoverMethodPost{/* values here */})
```

### RestCollectMethodOtherRestDiscoveryDiscoverTypeHTTPDiscoverMethodPostWithBody

```go
restCollectMethodOtherRestDiscoveryDiscoverTypeHTTP := components.CreateRestCollectMethodOtherRestDiscoveryDiscoverTypeHTTPPostWithBody(components.RestCollectMethodOtherRestDiscoveryDiscoverTypeHTTPDiscoverMethodPostWithBody{/* values here */})
```

### RestCollectMethodOtherRestDiscoveryDiscoverTypeHTTPDiscoverMethodOther

```go
restCollectMethodOtherRestDiscoveryDiscoverTypeHTTP := components.CreateRestCollectMethodOtherRestDiscoveryDiscoverTypeHTTPOther(components.RestCollectMethodOtherRestDiscoveryDiscoverTypeHTTPDiscoverMethodOther{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restCollectMethodOtherRestDiscoveryDiscoverTypeHTTP.Type {
	case components.RestCollectMethodOtherRestDiscoveryDiscoverTypeHTTPTypeGet:
		// restCollectMethodOtherRestDiscoveryDiscoverTypeHTTP.RestCollectMethodOtherRestDiscoveryDiscoverTypeHTTPDiscoverMethodGet is populated
	case components.RestCollectMethodOtherRestDiscoveryDiscoverTypeHTTPTypePost:
		// restCollectMethodOtherRestDiscoveryDiscoverTypeHTTP.RestCollectMethodOtherRestDiscoveryDiscoverTypeHTTPDiscoverMethodPost is populated
	case components.RestCollectMethodOtherRestDiscoveryDiscoverTypeHTTPTypePostWithBody:
		// restCollectMethodOtherRestDiscoveryDiscoverTypeHTTP.RestCollectMethodOtherRestDiscoveryDiscoverTypeHTTPDiscoverMethodPostWithBody is populated
	case components.RestCollectMethodOtherRestDiscoveryDiscoverTypeHTTPTypeOther:
		// restCollectMethodOtherRestDiscoveryDiscoverTypeHTTP.RestCollectMethodOtherRestDiscoveryDiscoverTypeHTTPDiscoverMethodOther is populated
	default:
		// Unknown type - use restCollectMethodOtherRestDiscoveryDiscoverTypeHTTP.GetUnknownRaw() for raw JSON
}
```
