# RestAuthenticationOauthSecretDiscovery


## Supported Types

### RestAuthenticationOauthSecretRestDiscoveryDiscoverTypeHTTP

```go
restAuthenticationOauthSecretDiscovery := components.CreateRestAuthenticationOauthSecretDiscoveryHTTP(components.RestAuthenticationOauthSecretRestDiscoveryDiscoverTypeHTTP{/* values here */})
```

### RestAuthenticationOauthSecretRestDiscoveryDiscoverTypeJSON

```go
restAuthenticationOauthSecretDiscovery := components.CreateRestAuthenticationOauthSecretDiscoveryJSON(components.RestAuthenticationOauthSecretRestDiscoveryDiscoverTypeJSON{/* values here */})
```

### RestAuthenticationOauthSecretRestDiscoveryDiscoverTypeList

```go
restAuthenticationOauthSecretDiscovery := components.CreateRestAuthenticationOauthSecretDiscoveryList(components.RestAuthenticationOauthSecretRestDiscoveryDiscoverTypeList{/* values here */})
```

### RestAuthenticationOauthSecretRestDiscoveryDiscoverTypeNone

```go
restAuthenticationOauthSecretDiscovery := components.CreateRestAuthenticationOauthSecretDiscoveryNone(components.RestAuthenticationOauthSecretRestDiscoveryDiscoverTypeNone{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationOauthSecretDiscovery.Type {
	case components.RestAuthenticationOauthSecretDiscoveryTypeHTTP:
		// restAuthenticationOauthSecretDiscovery.RestAuthenticationOauthSecretRestDiscoveryDiscoverTypeHTTP is populated
	case components.RestAuthenticationOauthSecretDiscoveryTypeJSON:
		// restAuthenticationOauthSecretDiscovery.RestAuthenticationOauthSecretRestDiscoveryDiscoverTypeJSON is populated
	case components.RestAuthenticationOauthSecretDiscoveryTypeList:
		// restAuthenticationOauthSecretDiscovery.RestAuthenticationOauthSecretRestDiscoveryDiscoverTypeList is populated
	case components.RestAuthenticationOauthSecretDiscoveryTypeNone:
		// restAuthenticationOauthSecretDiscovery.RestAuthenticationOauthSecretRestDiscoveryDiscoverTypeNone is populated
	default:
		// Unknown type - use restAuthenticationOauthSecretDiscovery.GetUnknownRaw() for raw JSON
}
```
