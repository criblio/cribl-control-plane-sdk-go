# RestAuthenticationOauthSecretPagination


## Supported Types

### RestAuthenticationOauthSecretRestPaginationTypeNone

```go
restAuthenticationOauthSecretPagination := components.CreateRestAuthenticationOauthSecretPaginationNone(components.RestAuthenticationOauthSecretRestPaginationTypeNone{/* values here */})
```

### RestAuthenticationOauthSecretRestPaginationTypeResponseBody

```go
restAuthenticationOauthSecretPagination := components.CreateRestAuthenticationOauthSecretPaginationResponseBody(components.RestAuthenticationOauthSecretRestPaginationTypeResponseBody{/* values here */})
```

### RestAuthenticationOauthSecretRestPaginationTypeResponseHeader

```go
restAuthenticationOauthSecretPagination := components.CreateRestAuthenticationOauthSecretPaginationResponseHeader(components.RestAuthenticationOauthSecretRestPaginationTypeResponseHeader{/* values here */})
```

### RestAuthenticationOauthSecretRestPaginationTypeResponseHeaderLink

```go
restAuthenticationOauthSecretPagination := components.CreateRestAuthenticationOauthSecretPaginationResponseHeaderLink(components.RestAuthenticationOauthSecretRestPaginationTypeResponseHeaderLink{/* values here */})
```

### RestAuthenticationOauthSecretRestPaginationTypeRequestOffset

```go
restAuthenticationOauthSecretPagination := components.CreateRestAuthenticationOauthSecretPaginationRequestOffset(components.RestAuthenticationOauthSecretRestPaginationTypeRequestOffset{/* values here */})
```

### RestAuthenticationOauthSecretRestPaginationTypeRequestPage

```go
restAuthenticationOauthSecretPagination := components.CreateRestAuthenticationOauthSecretPaginationRequestPage(components.RestAuthenticationOauthSecretRestPaginationTypeRequestPage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationOauthSecretPagination.Type {
	case components.RestAuthenticationOauthSecretPaginationTypeNone:
		// restAuthenticationOauthSecretPagination.RestAuthenticationOauthSecretRestPaginationTypeNone is populated
	case components.RestAuthenticationOauthSecretPaginationTypeResponseBody:
		// restAuthenticationOauthSecretPagination.RestAuthenticationOauthSecretRestPaginationTypeResponseBody is populated
	case components.RestAuthenticationOauthSecretPaginationTypeResponseHeader:
		// restAuthenticationOauthSecretPagination.RestAuthenticationOauthSecretRestPaginationTypeResponseHeader is populated
	case components.RestAuthenticationOauthSecretPaginationTypeResponseHeaderLink:
		// restAuthenticationOauthSecretPagination.RestAuthenticationOauthSecretRestPaginationTypeResponseHeaderLink is populated
	case components.RestAuthenticationOauthSecretPaginationTypeRequestOffset:
		// restAuthenticationOauthSecretPagination.RestAuthenticationOauthSecretRestPaginationTypeRequestOffset is populated
	case components.RestAuthenticationOauthSecretPaginationTypeRequestPage:
		// restAuthenticationOauthSecretPagination.RestAuthenticationOauthSecretRestPaginationTypeRequestPage is populated
	default:
		// Unknown type - use restAuthenticationOauthSecretPagination.GetUnknownRaw() for raw JSON
}
```
