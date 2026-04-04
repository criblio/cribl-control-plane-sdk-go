# RestAuthenticationGoogleOauthSecretPagination


## Supported Types

### RestAuthenticationGoogleOauthSecretRestPaginationTypeNone

```go
restAuthenticationGoogleOauthSecretPagination := components.CreateRestAuthenticationGoogleOauthSecretPaginationNone(components.RestAuthenticationGoogleOauthSecretRestPaginationTypeNone{/* values here */})
```

### RestAuthenticationGoogleOauthSecretRestPaginationTypeResponseBody

```go
restAuthenticationGoogleOauthSecretPagination := components.CreateRestAuthenticationGoogleOauthSecretPaginationResponseBody(components.RestAuthenticationGoogleOauthSecretRestPaginationTypeResponseBody{/* values here */})
```

### RestAuthenticationGoogleOauthSecretRestPaginationTypeResponseHeader

```go
restAuthenticationGoogleOauthSecretPagination := components.CreateRestAuthenticationGoogleOauthSecretPaginationResponseHeader(components.RestAuthenticationGoogleOauthSecretRestPaginationTypeResponseHeader{/* values here */})
```

### RestAuthenticationGoogleOauthSecretRestPaginationTypeResponseHeaderLink

```go
restAuthenticationGoogleOauthSecretPagination := components.CreateRestAuthenticationGoogleOauthSecretPaginationResponseHeaderLink(components.RestAuthenticationGoogleOauthSecretRestPaginationTypeResponseHeaderLink{/* values here */})
```

### RestAuthenticationGoogleOauthSecretRestPaginationTypeRequestOffset

```go
restAuthenticationGoogleOauthSecretPagination := components.CreateRestAuthenticationGoogleOauthSecretPaginationRequestOffset(components.RestAuthenticationGoogleOauthSecretRestPaginationTypeRequestOffset{/* values here */})
```

### RestAuthenticationGoogleOauthSecretRestPaginationTypeRequestPage

```go
restAuthenticationGoogleOauthSecretPagination := components.CreateRestAuthenticationGoogleOauthSecretPaginationRequestPage(components.RestAuthenticationGoogleOauthSecretRestPaginationTypeRequestPage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationGoogleOauthSecretPagination.Type {
	case components.RestAuthenticationGoogleOauthSecretPaginationTypeNone:
		// restAuthenticationGoogleOauthSecretPagination.RestAuthenticationGoogleOauthSecretRestPaginationTypeNone is populated
	case components.RestAuthenticationGoogleOauthSecretPaginationTypeResponseBody:
		// restAuthenticationGoogleOauthSecretPagination.RestAuthenticationGoogleOauthSecretRestPaginationTypeResponseBody is populated
	case components.RestAuthenticationGoogleOauthSecretPaginationTypeResponseHeader:
		// restAuthenticationGoogleOauthSecretPagination.RestAuthenticationGoogleOauthSecretRestPaginationTypeResponseHeader is populated
	case components.RestAuthenticationGoogleOauthSecretPaginationTypeResponseHeaderLink:
		// restAuthenticationGoogleOauthSecretPagination.RestAuthenticationGoogleOauthSecretRestPaginationTypeResponseHeaderLink is populated
	case components.RestAuthenticationGoogleOauthSecretPaginationTypeRequestOffset:
		// restAuthenticationGoogleOauthSecretPagination.RestAuthenticationGoogleOauthSecretRestPaginationTypeRequestOffset is populated
	case components.RestAuthenticationGoogleOauthSecretPaginationTypeRequestPage:
		// restAuthenticationGoogleOauthSecretPagination.RestAuthenticationGoogleOauthSecretRestPaginationTypeRequestPage is populated
	default:
		// Unknown type - use restAuthenticationGoogleOauthSecretPagination.GetUnknownRaw() for raw JSON
}
```
