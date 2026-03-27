# RestAuthenticationHmacRestDiscoveryDiscoverTypeHTTP


## Supported Types

### RestAuthenticationHmacRestDiscoveryDiscoverTypeHTTPDiscoverMethodGet

```go
restAuthenticationHmacRestDiscoveryDiscoverTypeHTTP := components.CreateRestAuthenticationHmacRestDiscoveryDiscoverTypeHTTPGet(components.RestAuthenticationHmacRestDiscoveryDiscoverTypeHTTPDiscoverMethodGet{/* values here */})
```

### RestAuthenticationHmacRestDiscoveryDiscoverTypeHTTPDiscoverMethodPost

```go
restAuthenticationHmacRestDiscoveryDiscoverTypeHTTP := components.CreateRestAuthenticationHmacRestDiscoveryDiscoverTypeHTTPPost(components.RestAuthenticationHmacRestDiscoveryDiscoverTypeHTTPDiscoverMethodPost{/* values here */})
```

### RestAuthenticationHmacRestDiscoveryDiscoverTypeHTTPDiscoverMethodPostWithBody

```go
restAuthenticationHmacRestDiscoveryDiscoverTypeHTTP := components.CreateRestAuthenticationHmacRestDiscoveryDiscoverTypeHTTPPostWithBody(components.RestAuthenticationHmacRestDiscoveryDiscoverTypeHTTPDiscoverMethodPostWithBody{/* values here */})
```

### RestAuthenticationHmacRestDiscoveryDiscoverTypeHTTPDiscoverMethodOther

```go
restAuthenticationHmacRestDiscoveryDiscoverTypeHTTP := components.CreateRestAuthenticationHmacRestDiscoveryDiscoverTypeHTTPOther(components.RestAuthenticationHmacRestDiscoveryDiscoverTypeHTTPDiscoverMethodOther{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationHmacRestDiscoveryDiscoverTypeHTTP.Type {
	case components.RestAuthenticationHmacRestDiscoveryDiscoverTypeHTTPTypeGet:
		// restAuthenticationHmacRestDiscoveryDiscoverTypeHTTP.RestAuthenticationHmacRestDiscoveryDiscoverTypeHTTPDiscoverMethodGet is populated
	case components.RestAuthenticationHmacRestDiscoveryDiscoverTypeHTTPTypePost:
		// restAuthenticationHmacRestDiscoveryDiscoverTypeHTTP.RestAuthenticationHmacRestDiscoveryDiscoverTypeHTTPDiscoverMethodPost is populated
	case components.RestAuthenticationHmacRestDiscoveryDiscoverTypeHTTPTypePostWithBody:
		// restAuthenticationHmacRestDiscoveryDiscoverTypeHTTP.RestAuthenticationHmacRestDiscoveryDiscoverTypeHTTPDiscoverMethodPostWithBody is populated
	case components.RestAuthenticationHmacRestDiscoveryDiscoverTypeHTTPTypeOther:
		// restAuthenticationHmacRestDiscoveryDiscoverTypeHTTP.RestAuthenticationHmacRestDiscoveryDiscoverTypeHTTPDiscoverMethodOther is populated
	default:
		// Unknown type - use restAuthenticationHmacRestDiscoveryDiscoverTypeHTTP.GetUnknownRaw() for raw JSON
}
```
