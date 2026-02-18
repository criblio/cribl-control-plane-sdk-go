# RestAuthenticationBasicSecretDiscovery


## Supported Types

### RestAuthenticationBasicSecretRestDiscoveryDiscoverTypeHTTP

```go
restAuthenticationBasicSecretDiscovery := components.CreateRestAuthenticationBasicSecretDiscoveryHTTP(components.RestAuthenticationBasicSecretRestDiscoveryDiscoverTypeHTTP{/* values here */})
```

### RestAuthenticationBasicSecretRestDiscoveryDiscoverTypeJSON

```go
restAuthenticationBasicSecretDiscovery := components.CreateRestAuthenticationBasicSecretDiscoveryJSON(components.RestAuthenticationBasicSecretRestDiscoveryDiscoverTypeJSON{/* values here */})
```

### RestAuthenticationBasicSecretRestDiscoveryDiscoverTypeList

```go
restAuthenticationBasicSecretDiscovery := components.CreateRestAuthenticationBasicSecretDiscoveryList(components.RestAuthenticationBasicSecretRestDiscoveryDiscoverTypeList{/* values here */})
```

### RestAuthenticationBasicSecretRestDiscoveryDiscoverTypeNone

```go
restAuthenticationBasicSecretDiscovery := components.CreateRestAuthenticationBasicSecretDiscoveryNone(components.RestAuthenticationBasicSecretRestDiscoveryDiscoverTypeNone{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationBasicSecretDiscovery.Type {
	case components.RestAuthenticationBasicSecretDiscoveryTypeHTTP:
		// restAuthenticationBasicSecretDiscovery.RestAuthenticationBasicSecretRestDiscoveryDiscoverTypeHTTP is populated
	case components.RestAuthenticationBasicSecretDiscoveryTypeJSON:
		// restAuthenticationBasicSecretDiscovery.RestAuthenticationBasicSecretRestDiscoveryDiscoverTypeJSON is populated
	case components.RestAuthenticationBasicSecretDiscoveryTypeList:
		// restAuthenticationBasicSecretDiscovery.RestAuthenticationBasicSecretRestDiscoveryDiscoverTypeList is populated
	case components.RestAuthenticationBasicSecretDiscoveryTypeNone:
		// restAuthenticationBasicSecretDiscovery.RestAuthenticationBasicSecretRestDiscoveryDiscoverTypeNone is populated
	default:
		// Unknown type - use restAuthenticationBasicSecretDiscovery.GetUnknownRaw() for raw JSON
}
```
