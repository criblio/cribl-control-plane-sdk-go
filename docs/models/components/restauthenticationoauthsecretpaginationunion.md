# RestAuthenticationOauthSecretPaginationUnion


## Supported Types

### RestAuthenticationOauthSecretRestPaginationTypeNone

```go
restAuthenticationOauthSecretPaginationUnion := components.CreateRestAuthenticationOauthSecretPaginationUnionNone(components.RestAuthenticationOauthSecretRestPaginationTypeNone{/* values here */})
```

### RestAuthenticationOauthSecretRestPaginationTypeResponseBody

```go
restAuthenticationOauthSecretPaginationUnion := components.CreateRestAuthenticationOauthSecretPaginationUnionResponseBody(components.RestAuthenticationOauthSecretRestPaginationTypeResponseBody{/* values here */})
```

### RestAuthenticationOauthSecretRestPaginationTypeResponseHeader

```go
restAuthenticationOauthSecretPaginationUnion := components.CreateRestAuthenticationOauthSecretPaginationUnionResponseHeader(components.RestAuthenticationOauthSecretRestPaginationTypeResponseHeader{/* values here */})
```

### RestAuthenticationOauthSecretRestPaginationTypeResponseHeaderLink

```go
restAuthenticationOauthSecretPaginationUnion := components.CreateRestAuthenticationOauthSecretPaginationUnionResponseHeaderLink(components.RestAuthenticationOauthSecretRestPaginationTypeResponseHeaderLink{/* values here */})
```

### RestAuthenticationOauthSecretRestPaginationTypeRequestOffset

```go
restAuthenticationOauthSecretPaginationUnion := components.CreateRestAuthenticationOauthSecretPaginationUnionRequestOffset(components.RestAuthenticationOauthSecretRestPaginationTypeRequestOffset{/* values here */})
```

### RestAuthenticationOauthSecretRestPaginationTypeRequestPage

```go
restAuthenticationOauthSecretPaginationUnion := components.CreateRestAuthenticationOauthSecretPaginationUnionRequestPage(components.RestAuthenticationOauthSecretRestPaginationTypeRequestPage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationOauthSecretPaginationUnion.Type {
	case components.RestAuthenticationOauthSecretPaginationUnionTypeNone:
		// restAuthenticationOauthSecretPaginationUnion.RestAuthenticationOauthSecretRestPaginationTypeNone is populated
	case components.RestAuthenticationOauthSecretPaginationUnionTypeResponseBody:
		// restAuthenticationOauthSecretPaginationUnion.RestAuthenticationOauthSecretRestPaginationTypeResponseBody is populated
	case components.RestAuthenticationOauthSecretPaginationUnionTypeResponseHeader:
		// restAuthenticationOauthSecretPaginationUnion.RestAuthenticationOauthSecretRestPaginationTypeResponseHeader is populated
	case components.RestAuthenticationOauthSecretPaginationUnionTypeResponseHeaderLink:
		// restAuthenticationOauthSecretPaginationUnion.RestAuthenticationOauthSecretRestPaginationTypeResponseHeaderLink is populated
	case components.RestAuthenticationOauthSecretPaginationUnionTypeRequestOffset:
		// restAuthenticationOauthSecretPaginationUnion.RestAuthenticationOauthSecretRestPaginationTypeRequestOffset is populated
	case components.RestAuthenticationOauthSecretPaginationUnionTypeRequestPage:
		// restAuthenticationOauthSecretPaginationUnion.RestAuthenticationOauthSecretRestPaginationTypeRequestPage is populated
	default:
		// Unknown type - use restAuthenticationOauthSecretPaginationUnion.GetUnknownRaw() for raw JSON
}
```
