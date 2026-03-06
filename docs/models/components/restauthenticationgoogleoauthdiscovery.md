# RestAuthenticationGoogleOauthDiscovery


## Supported Types

### RestAuthenticationGoogleOauthRestDiscoveryDiscoverTypeHTTP

```go
restAuthenticationGoogleOauthDiscovery := components.CreateRestAuthenticationGoogleOauthDiscoveryHTTP(components.RestAuthenticationGoogleOauthRestDiscoveryDiscoverTypeHTTP{/* values here */})
```

### RestAuthenticationGoogleOauthRestDiscoveryDiscoverTypeJSON

```go
restAuthenticationGoogleOauthDiscovery := components.CreateRestAuthenticationGoogleOauthDiscoveryJSON(components.RestAuthenticationGoogleOauthRestDiscoveryDiscoverTypeJSON{/* values here */})
```

### RestAuthenticationGoogleOauthRestDiscoveryDiscoverTypeList

```go
restAuthenticationGoogleOauthDiscovery := components.CreateRestAuthenticationGoogleOauthDiscoveryList(components.RestAuthenticationGoogleOauthRestDiscoveryDiscoverTypeList{/* values here */})
```

### RestAuthenticationGoogleOauthRestDiscoveryDiscoverTypeNone

```go
restAuthenticationGoogleOauthDiscovery := components.CreateRestAuthenticationGoogleOauthDiscoveryNone(components.RestAuthenticationGoogleOauthRestDiscoveryDiscoverTypeNone{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationGoogleOauthDiscovery.Type {
	case components.RestAuthenticationGoogleOauthDiscoveryTypeHTTP:
		// restAuthenticationGoogleOauthDiscovery.RestAuthenticationGoogleOauthRestDiscoveryDiscoverTypeHTTP is populated
	case components.RestAuthenticationGoogleOauthDiscoveryTypeJSON:
		// restAuthenticationGoogleOauthDiscovery.RestAuthenticationGoogleOauthRestDiscoveryDiscoverTypeJSON is populated
	case components.RestAuthenticationGoogleOauthDiscoveryTypeList:
		// restAuthenticationGoogleOauthDiscovery.RestAuthenticationGoogleOauthRestDiscoveryDiscoverTypeList is populated
	case components.RestAuthenticationGoogleOauthDiscoveryTypeNone:
		// restAuthenticationGoogleOauthDiscovery.RestAuthenticationGoogleOauthRestDiscoveryDiscoverTypeNone is populated
	default:
		// Unknown type - use restAuthenticationGoogleOauthDiscovery.GetUnknownRaw() for raw JSON
}
```
