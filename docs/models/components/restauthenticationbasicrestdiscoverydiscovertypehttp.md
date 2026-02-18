# RestAuthenticationBasicRestDiscoveryDiscoverTypeHTTP


## Supported Types

### RestAuthenticationBasicRestDiscoveryDiscoverTypeHTTPDiscoverMethodGet

```go
restAuthenticationBasicRestDiscoveryDiscoverTypeHTTP := components.CreateRestAuthenticationBasicRestDiscoveryDiscoverTypeHTTPGet(components.RestAuthenticationBasicRestDiscoveryDiscoverTypeHTTPDiscoverMethodGet{/* values here */})
```

### RestAuthenticationBasicRestDiscoveryDiscoverTypeHTTPDiscoverMethodPost

```go
restAuthenticationBasicRestDiscoveryDiscoverTypeHTTP := components.CreateRestAuthenticationBasicRestDiscoveryDiscoverTypeHTTPPost(components.RestAuthenticationBasicRestDiscoveryDiscoverTypeHTTPDiscoverMethodPost{/* values here */})
```

### RestAuthenticationBasicRestDiscoveryDiscoverTypeHTTPDiscoverMethodPostWithBody

```go
restAuthenticationBasicRestDiscoveryDiscoverTypeHTTP := components.CreateRestAuthenticationBasicRestDiscoveryDiscoverTypeHTTPPostWithBody(components.RestAuthenticationBasicRestDiscoveryDiscoverTypeHTTPDiscoverMethodPostWithBody{/* values here */})
```

### RestAuthenticationBasicRestDiscoveryDiscoverTypeHTTPDiscoverMethodOther

```go
restAuthenticationBasicRestDiscoveryDiscoverTypeHTTP := components.CreateRestAuthenticationBasicRestDiscoveryDiscoverTypeHTTPOther(components.RestAuthenticationBasicRestDiscoveryDiscoverTypeHTTPDiscoverMethodOther{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationBasicRestDiscoveryDiscoverTypeHTTP.Type {
	case components.RestAuthenticationBasicRestDiscoveryDiscoverTypeHTTPTypeGet:
		// restAuthenticationBasicRestDiscoveryDiscoverTypeHTTP.RestAuthenticationBasicRestDiscoveryDiscoverTypeHTTPDiscoverMethodGet is populated
	case components.RestAuthenticationBasicRestDiscoveryDiscoverTypeHTTPTypePost:
		// restAuthenticationBasicRestDiscoveryDiscoverTypeHTTP.RestAuthenticationBasicRestDiscoveryDiscoverTypeHTTPDiscoverMethodPost is populated
	case components.RestAuthenticationBasicRestDiscoveryDiscoverTypeHTTPTypePostWithBody:
		// restAuthenticationBasicRestDiscoveryDiscoverTypeHTTP.RestAuthenticationBasicRestDiscoveryDiscoverTypeHTTPDiscoverMethodPostWithBody is populated
	case components.RestAuthenticationBasicRestDiscoveryDiscoverTypeHTTPTypeOther:
		// restAuthenticationBasicRestDiscoveryDiscoverTypeHTTP.RestAuthenticationBasicRestDiscoveryDiscoverTypeHTTPDiscoverMethodOther is populated
	default:
		// Unknown type - use restAuthenticationBasicRestDiscoveryDiscoverTypeHTTP.GetUnknownRaw() for raw JSON
}
```
