# RestAuthenticationLoginRestDiscoveryDiscoverTypeHTTP


## Supported Types

### RestAuthenticationLoginRestDiscoveryDiscoverTypeHTTPDiscoverMethodGet

```go
restAuthenticationLoginRestDiscoveryDiscoverTypeHTTP := components.CreateRestAuthenticationLoginRestDiscoveryDiscoverTypeHTTPGet(components.RestAuthenticationLoginRestDiscoveryDiscoverTypeHTTPDiscoverMethodGet{/* values here */})
```

### RestAuthenticationLoginRestDiscoveryDiscoverTypeHTTPDiscoverMethodPost

```go
restAuthenticationLoginRestDiscoveryDiscoverTypeHTTP := components.CreateRestAuthenticationLoginRestDiscoveryDiscoverTypeHTTPPost(components.RestAuthenticationLoginRestDiscoveryDiscoverTypeHTTPDiscoverMethodPost{/* values here */})
```

### RestAuthenticationLoginRestDiscoveryDiscoverTypeHTTPDiscoverMethodPostWithBody

```go
restAuthenticationLoginRestDiscoveryDiscoverTypeHTTP := components.CreateRestAuthenticationLoginRestDiscoveryDiscoverTypeHTTPPostWithBody(components.RestAuthenticationLoginRestDiscoveryDiscoverTypeHTTPDiscoverMethodPostWithBody{/* values here */})
```

### RestAuthenticationLoginRestDiscoveryDiscoverTypeHTTPDiscoverMethodOther

```go
restAuthenticationLoginRestDiscoveryDiscoverTypeHTTP := components.CreateRestAuthenticationLoginRestDiscoveryDiscoverTypeHTTPOther(components.RestAuthenticationLoginRestDiscoveryDiscoverTypeHTTPDiscoverMethodOther{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationLoginRestDiscoveryDiscoverTypeHTTP.Type {
	case components.RestAuthenticationLoginRestDiscoveryDiscoverTypeHTTPTypeGet:
		// restAuthenticationLoginRestDiscoveryDiscoverTypeHTTP.RestAuthenticationLoginRestDiscoveryDiscoverTypeHTTPDiscoverMethodGet is populated
	case components.RestAuthenticationLoginRestDiscoveryDiscoverTypeHTTPTypePost:
		// restAuthenticationLoginRestDiscoveryDiscoverTypeHTTP.RestAuthenticationLoginRestDiscoveryDiscoverTypeHTTPDiscoverMethodPost is populated
	case components.RestAuthenticationLoginRestDiscoveryDiscoverTypeHTTPTypePostWithBody:
		// restAuthenticationLoginRestDiscoveryDiscoverTypeHTTP.RestAuthenticationLoginRestDiscoveryDiscoverTypeHTTPDiscoverMethodPostWithBody is populated
	case components.RestAuthenticationLoginRestDiscoveryDiscoverTypeHTTPTypeOther:
		// restAuthenticationLoginRestDiscoveryDiscoverTypeHTTP.RestAuthenticationLoginRestDiscoveryDiscoverTypeHTTPDiscoverMethodOther is populated
	default:
		// Unknown type - use restAuthenticationLoginRestDiscoveryDiscoverTypeHTTP.GetUnknownRaw() for raw JSON
}
```
