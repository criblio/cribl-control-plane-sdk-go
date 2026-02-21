# RestAuthenticationBasicSecretPaginationUnion


## Supported Types

### RestAuthenticationBasicSecretRestPaginationTypeNone

```go
restAuthenticationBasicSecretPaginationUnion := components.CreateRestAuthenticationBasicSecretPaginationUnionNone(components.RestAuthenticationBasicSecretRestPaginationTypeNone{/* values here */})
```

### RestAuthenticationBasicSecretRestPaginationTypeResponseBody

```go
restAuthenticationBasicSecretPaginationUnion := components.CreateRestAuthenticationBasicSecretPaginationUnionResponseBody(components.RestAuthenticationBasicSecretRestPaginationTypeResponseBody{/* values here */})
```

### RestAuthenticationBasicSecretRestPaginationTypeResponseHeader

```go
restAuthenticationBasicSecretPaginationUnion := components.CreateRestAuthenticationBasicSecretPaginationUnionResponseHeader(components.RestAuthenticationBasicSecretRestPaginationTypeResponseHeader{/* values here */})
```

### RestAuthenticationBasicSecretRestPaginationTypeResponseHeaderLink

```go
restAuthenticationBasicSecretPaginationUnion := components.CreateRestAuthenticationBasicSecretPaginationUnionResponseHeaderLink(components.RestAuthenticationBasicSecretRestPaginationTypeResponseHeaderLink{/* values here */})
```

### RestAuthenticationBasicSecretRestPaginationTypeRequestOffset

```go
restAuthenticationBasicSecretPaginationUnion := components.CreateRestAuthenticationBasicSecretPaginationUnionRequestOffset(components.RestAuthenticationBasicSecretRestPaginationTypeRequestOffset{/* values here */})
```

### RestAuthenticationBasicSecretRestPaginationTypeRequestPage

```go
restAuthenticationBasicSecretPaginationUnion := components.CreateRestAuthenticationBasicSecretPaginationUnionRequestPage(components.RestAuthenticationBasicSecretRestPaginationTypeRequestPage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationBasicSecretPaginationUnion.Type {
	case components.RestAuthenticationBasicSecretPaginationUnionTypeNone:
		// restAuthenticationBasicSecretPaginationUnion.RestAuthenticationBasicSecretRestPaginationTypeNone is populated
	case components.RestAuthenticationBasicSecretPaginationUnionTypeResponseBody:
		// restAuthenticationBasicSecretPaginationUnion.RestAuthenticationBasicSecretRestPaginationTypeResponseBody is populated
	case components.RestAuthenticationBasicSecretPaginationUnionTypeResponseHeader:
		// restAuthenticationBasicSecretPaginationUnion.RestAuthenticationBasicSecretRestPaginationTypeResponseHeader is populated
	case components.RestAuthenticationBasicSecretPaginationUnionTypeResponseHeaderLink:
		// restAuthenticationBasicSecretPaginationUnion.RestAuthenticationBasicSecretRestPaginationTypeResponseHeaderLink is populated
	case components.RestAuthenticationBasicSecretPaginationUnionTypeRequestOffset:
		// restAuthenticationBasicSecretPaginationUnion.RestAuthenticationBasicSecretRestPaginationTypeRequestOffset is populated
	case components.RestAuthenticationBasicSecretPaginationUnionTypeRequestPage:
		// restAuthenticationBasicSecretPaginationUnion.RestAuthenticationBasicSecretRestPaginationTypeRequestPage is populated
	default:
		// Unknown type - use restAuthenticationBasicSecretPaginationUnion.GetUnknownRaw() for raw JSON
}
```
