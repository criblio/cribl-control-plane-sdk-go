# RestAuthenticationNoneRestDiscoveryDiscoverTypeHTTP


## Supported Types

### RestAuthenticationNoneRestDiscoveryDiscoverTypeHTTPDiscoverMethodGet

```go
restAuthenticationNoneRestDiscoveryDiscoverTypeHTTP := components.CreateRestAuthenticationNoneRestDiscoveryDiscoverTypeHTTPGet(components.RestAuthenticationNoneRestDiscoveryDiscoverTypeHTTPDiscoverMethodGet{/* values here */})
```

### RestAuthenticationNoneRestDiscoveryDiscoverTypeHTTPDiscoverMethodPost

```go
restAuthenticationNoneRestDiscoveryDiscoverTypeHTTP := components.CreateRestAuthenticationNoneRestDiscoveryDiscoverTypeHTTPPost(components.RestAuthenticationNoneRestDiscoveryDiscoverTypeHTTPDiscoverMethodPost{/* values here */})
```

### RestAuthenticationNoneRestDiscoveryDiscoverTypeHTTPDiscoverMethodPostWithBody

```go
restAuthenticationNoneRestDiscoveryDiscoverTypeHTTP := components.CreateRestAuthenticationNoneRestDiscoveryDiscoverTypeHTTPPostWithBody(components.RestAuthenticationNoneRestDiscoveryDiscoverTypeHTTPDiscoverMethodPostWithBody{/* values here */})
```

### RestAuthenticationNoneRestDiscoveryDiscoverTypeHTTPDiscoverMethodOther

```go
restAuthenticationNoneRestDiscoveryDiscoverTypeHTTP := components.CreateRestAuthenticationNoneRestDiscoveryDiscoverTypeHTTPOther(components.RestAuthenticationNoneRestDiscoveryDiscoverTypeHTTPDiscoverMethodOther{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationNoneRestDiscoveryDiscoverTypeHTTP.Type {
	case components.RestAuthenticationNoneRestDiscoveryDiscoverTypeHTTPTypeGet:
		// restAuthenticationNoneRestDiscoveryDiscoverTypeHTTP.RestAuthenticationNoneRestDiscoveryDiscoverTypeHTTPDiscoverMethodGet is populated
	case components.RestAuthenticationNoneRestDiscoveryDiscoverTypeHTTPTypePost:
		// restAuthenticationNoneRestDiscoveryDiscoverTypeHTTP.RestAuthenticationNoneRestDiscoveryDiscoverTypeHTTPDiscoverMethodPost is populated
	case components.RestAuthenticationNoneRestDiscoveryDiscoverTypeHTTPTypePostWithBody:
		// restAuthenticationNoneRestDiscoveryDiscoverTypeHTTP.RestAuthenticationNoneRestDiscoveryDiscoverTypeHTTPDiscoverMethodPostWithBody is populated
	case components.RestAuthenticationNoneRestDiscoveryDiscoverTypeHTTPTypeOther:
		// restAuthenticationNoneRestDiscoveryDiscoverTypeHTTP.RestAuthenticationNoneRestDiscoveryDiscoverTypeHTTPDiscoverMethodOther is populated
	default:
		// Unknown type - use restAuthenticationNoneRestDiscoveryDiscoverTypeHTTP.GetUnknownRaw() for raw JSON
}
```
