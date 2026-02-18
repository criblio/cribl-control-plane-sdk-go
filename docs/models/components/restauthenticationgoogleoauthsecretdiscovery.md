# RestAuthenticationGoogleOauthSecretDiscovery


## Supported Types

### RestAuthenticationGoogleOauthSecretRestDiscoveryDiscoverTypeHTTP

```go
restAuthenticationGoogleOauthSecretDiscovery := components.CreateRestAuthenticationGoogleOauthSecretDiscoveryHTTP(components.RestAuthenticationGoogleOauthSecretRestDiscoveryDiscoverTypeHTTP{/* values here */})
```

### RestAuthenticationGoogleOauthSecretRestDiscoveryDiscoverTypeJSON

```go
restAuthenticationGoogleOauthSecretDiscovery := components.CreateRestAuthenticationGoogleOauthSecretDiscoveryJSON(components.RestAuthenticationGoogleOauthSecretRestDiscoveryDiscoverTypeJSON{/* values here */})
```

### RestAuthenticationGoogleOauthSecretRestDiscoveryDiscoverTypeList

```go
restAuthenticationGoogleOauthSecretDiscovery := components.CreateRestAuthenticationGoogleOauthSecretDiscoveryList(components.RestAuthenticationGoogleOauthSecretRestDiscoveryDiscoverTypeList{/* values here */})
```

### RestAuthenticationGoogleOauthSecretRestDiscoveryDiscoverTypeNone

```go
restAuthenticationGoogleOauthSecretDiscovery := components.CreateRestAuthenticationGoogleOauthSecretDiscoveryNone(components.RestAuthenticationGoogleOauthSecretRestDiscoveryDiscoverTypeNone{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationGoogleOauthSecretDiscovery.Type {
	case components.RestAuthenticationGoogleOauthSecretDiscoveryTypeHTTP:
		// restAuthenticationGoogleOauthSecretDiscovery.RestAuthenticationGoogleOauthSecretRestDiscoveryDiscoverTypeHTTP is populated
	case components.RestAuthenticationGoogleOauthSecretDiscoveryTypeJSON:
		// restAuthenticationGoogleOauthSecretDiscovery.RestAuthenticationGoogleOauthSecretRestDiscoveryDiscoverTypeJSON is populated
	case components.RestAuthenticationGoogleOauthSecretDiscoveryTypeList:
		// restAuthenticationGoogleOauthSecretDiscovery.RestAuthenticationGoogleOauthSecretRestDiscoveryDiscoverTypeList is populated
	case components.RestAuthenticationGoogleOauthSecretDiscoveryTypeNone:
		// restAuthenticationGoogleOauthSecretDiscovery.RestAuthenticationGoogleOauthSecretRestDiscoveryDiscoverTypeNone is populated
	default:
		// Unknown type - use restAuthenticationGoogleOauthSecretDiscovery.GetUnknownRaw() for raw JSON
}
```
