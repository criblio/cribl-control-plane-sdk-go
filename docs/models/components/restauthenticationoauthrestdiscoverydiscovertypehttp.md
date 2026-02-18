# RestAuthenticationOauthRestDiscoveryDiscoverTypeHTTP


## Supported Types

### RestAuthenticationOauthRestDiscoveryDiscoverTypeHTTPDiscoverMethodGet

```go
restAuthenticationOauthRestDiscoveryDiscoverTypeHTTP := components.CreateRestAuthenticationOauthRestDiscoveryDiscoverTypeHTTPGet(components.RestAuthenticationOauthRestDiscoveryDiscoverTypeHTTPDiscoverMethodGet{/* values here */})
```

### RestAuthenticationOauthRestDiscoveryDiscoverTypeHTTPDiscoverMethodPost

```go
restAuthenticationOauthRestDiscoveryDiscoverTypeHTTP := components.CreateRestAuthenticationOauthRestDiscoveryDiscoverTypeHTTPPost(components.RestAuthenticationOauthRestDiscoveryDiscoverTypeHTTPDiscoverMethodPost{/* values here */})
```

### RestAuthenticationOauthRestDiscoveryDiscoverTypeHTTPDiscoverMethodPostWithBody

```go
restAuthenticationOauthRestDiscoveryDiscoverTypeHTTP := components.CreateRestAuthenticationOauthRestDiscoveryDiscoverTypeHTTPPostWithBody(components.RestAuthenticationOauthRestDiscoveryDiscoverTypeHTTPDiscoverMethodPostWithBody{/* values here */})
```

### RestAuthenticationOauthRestDiscoveryDiscoverTypeHTTPDiscoverMethodOther

```go
restAuthenticationOauthRestDiscoveryDiscoverTypeHTTP := components.CreateRestAuthenticationOauthRestDiscoveryDiscoverTypeHTTPOther(components.RestAuthenticationOauthRestDiscoveryDiscoverTypeHTTPDiscoverMethodOther{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationOauthRestDiscoveryDiscoverTypeHTTP.Type {
	case components.RestAuthenticationOauthRestDiscoveryDiscoverTypeHTTPTypeGet:
		// restAuthenticationOauthRestDiscoveryDiscoverTypeHTTP.RestAuthenticationOauthRestDiscoveryDiscoverTypeHTTPDiscoverMethodGet is populated
	case components.RestAuthenticationOauthRestDiscoveryDiscoverTypeHTTPTypePost:
		// restAuthenticationOauthRestDiscoveryDiscoverTypeHTTP.RestAuthenticationOauthRestDiscoveryDiscoverTypeHTTPDiscoverMethodPost is populated
	case components.RestAuthenticationOauthRestDiscoveryDiscoverTypeHTTPTypePostWithBody:
		// restAuthenticationOauthRestDiscoveryDiscoverTypeHTTP.RestAuthenticationOauthRestDiscoveryDiscoverTypeHTTPDiscoverMethodPostWithBody is populated
	case components.RestAuthenticationOauthRestDiscoveryDiscoverTypeHTTPTypeOther:
		// restAuthenticationOauthRestDiscoveryDiscoverTypeHTTP.RestAuthenticationOauthRestDiscoveryDiscoverTypeHTTPDiscoverMethodOther is populated
	default:
		// Unknown type - use restAuthenticationOauthRestDiscoveryDiscoverTypeHTTP.GetUnknownRaw() for raw JSON
}
```
