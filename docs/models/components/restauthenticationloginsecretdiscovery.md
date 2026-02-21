# RestAuthenticationLoginSecretDiscovery


## Supported Types

### RestAuthenticationLoginSecretRestDiscoveryDiscoverTypeHTTP

```go
restAuthenticationLoginSecretDiscovery := components.CreateRestAuthenticationLoginSecretDiscoveryHTTP(components.RestAuthenticationLoginSecretRestDiscoveryDiscoverTypeHTTP{/* values here */})
```

### RestAuthenticationLoginSecretRestDiscoveryDiscoverTypeJSON

```go
restAuthenticationLoginSecretDiscovery := components.CreateRestAuthenticationLoginSecretDiscoveryJSON(components.RestAuthenticationLoginSecretRestDiscoveryDiscoverTypeJSON{/* values here */})
```

### RestAuthenticationLoginSecretRestDiscoveryDiscoverTypeList

```go
restAuthenticationLoginSecretDiscovery := components.CreateRestAuthenticationLoginSecretDiscoveryList(components.RestAuthenticationLoginSecretRestDiscoveryDiscoverTypeList{/* values here */})
```

### RestAuthenticationLoginSecretRestDiscoveryDiscoverTypeNone

```go
restAuthenticationLoginSecretDiscovery := components.CreateRestAuthenticationLoginSecretDiscoveryNone(components.RestAuthenticationLoginSecretRestDiscoveryDiscoverTypeNone{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationLoginSecretDiscovery.Type {
	case components.RestAuthenticationLoginSecretDiscoveryTypeHTTP:
		// restAuthenticationLoginSecretDiscovery.RestAuthenticationLoginSecretRestDiscoveryDiscoverTypeHTTP is populated
	case components.RestAuthenticationLoginSecretDiscoveryTypeJSON:
		// restAuthenticationLoginSecretDiscovery.RestAuthenticationLoginSecretRestDiscoveryDiscoverTypeJSON is populated
	case components.RestAuthenticationLoginSecretDiscoveryTypeList:
		// restAuthenticationLoginSecretDiscovery.RestAuthenticationLoginSecretRestDiscoveryDiscoverTypeList is populated
	case components.RestAuthenticationLoginSecretDiscoveryTypeNone:
		// restAuthenticationLoginSecretDiscovery.RestAuthenticationLoginSecretRestDiscoveryDiscoverTypeNone is populated
	default:
		// Unknown type - use restAuthenticationLoginSecretDiscovery.GetUnknownRaw() for raw JSON
}
```
