# RestAuthenticationLoginSecretPaginationUnion


## Supported Types

### RestAuthenticationLoginSecretRestPaginationTypeNone

```go
restAuthenticationLoginSecretPaginationUnion := components.CreateRestAuthenticationLoginSecretPaginationUnionNone(components.RestAuthenticationLoginSecretRestPaginationTypeNone{/* values here */})
```

### RestAuthenticationLoginSecretRestPaginationTypeResponseBody

```go
restAuthenticationLoginSecretPaginationUnion := components.CreateRestAuthenticationLoginSecretPaginationUnionResponseBody(components.RestAuthenticationLoginSecretRestPaginationTypeResponseBody{/* values here */})
```

### RestAuthenticationLoginSecretRestPaginationTypeResponseHeader

```go
restAuthenticationLoginSecretPaginationUnion := components.CreateRestAuthenticationLoginSecretPaginationUnionResponseHeader(components.RestAuthenticationLoginSecretRestPaginationTypeResponseHeader{/* values here */})
```

### RestAuthenticationLoginSecretRestPaginationTypeResponseHeaderLink

```go
restAuthenticationLoginSecretPaginationUnion := components.CreateRestAuthenticationLoginSecretPaginationUnionResponseHeaderLink(components.RestAuthenticationLoginSecretRestPaginationTypeResponseHeaderLink{/* values here */})
```

### RestAuthenticationLoginSecretRestPaginationTypeRequestOffset

```go
restAuthenticationLoginSecretPaginationUnion := components.CreateRestAuthenticationLoginSecretPaginationUnionRequestOffset(components.RestAuthenticationLoginSecretRestPaginationTypeRequestOffset{/* values here */})
```

### RestAuthenticationLoginSecretRestPaginationTypeRequestPage

```go
restAuthenticationLoginSecretPaginationUnion := components.CreateRestAuthenticationLoginSecretPaginationUnionRequestPage(components.RestAuthenticationLoginSecretRestPaginationTypeRequestPage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationLoginSecretPaginationUnion.Type {
	case components.RestAuthenticationLoginSecretPaginationUnionTypeNone:
		// restAuthenticationLoginSecretPaginationUnion.RestAuthenticationLoginSecretRestPaginationTypeNone is populated
	case components.RestAuthenticationLoginSecretPaginationUnionTypeResponseBody:
		// restAuthenticationLoginSecretPaginationUnion.RestAuthenticationLoginSecretRestPaginationTypeResponseBody is populated
	case components.RestAuthenticationLoginSecretPaginationUnionTypeResponseHeader:
		// restAuthenticationLoginSecretPaginationUnion.RestAuthenticationLoginSecretRestPaginationTypeResponseHeader is populated
	case components.RestAuthenticationLoginSecretPaginationUnionTypeResponseHeaderLink:
		// restAuthenticationLoginSecretPaginationUnion.RestAuthenticationLoginSecretRestPaginationTypeResponseHeaderLink is populated
	case components.RestAuthenticationLoginSecretPaginationUnionTypeRequestOffset:
		// restAuthenticationLoginSecretPaginationUnion.RestAuthenticationLoginSecretRestPaginationTypeRequestOffset is populated
	case components.RestAuthenticationLoginSecretPaginationUnionTypeRequestPage:
		// restAuthenticationLoginSecretPaginationUnion.RestAuthenticationLoginSecretRestPaginationTypeRequestPage is populated
	default:
		// Unknown type - use restAuthenticationLoginSecretPaginationUnion.GetUnknownRaw() for raw JSON
}
```
