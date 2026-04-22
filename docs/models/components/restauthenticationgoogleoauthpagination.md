# RestAuthenticationGoogleOauthPagination


## Supported Types

### RestAuthenticationGoogleOauthRestPaginationTypeNone

```go
restAuthenticationGoogleOauthPagination := components.CreateRestAuthenticationGoogleOauthPaginationNone(components.RestAuthenticationGoogleOauthRestPaginationTypeNone{/* values here */})
```

### RestAuthenticationGoogleOauthRestPaginationTypeResponseBody

```go
restAuthenticationGoogleOauthPagination := components.CreateRestAuthenticationGoogleOauthPaginationResponseBody(components.RestAuthenticationGoogleOauthRestPaginationTypeResponseBody{/* values here */})
```

### RestAuthenticationGoogleOauthRestPaginationTypeResponseHeader

```go
restAuthenticationGoogleOauthPagination := components.CreateRestAuthenticationGoogleOauthPaginationResponseHeader(components.RestAuthenticationGoogleOauthRestPaginationTypeResponseHeader{/* values here */})
```

### RestAuthenticationGoogleOauthRestPaginationTypeResponseHeaderLink

```go
restAuthenticationGoogleOauthPagination := components.CreateRestAuthenticationGoogleOauthPaginationResponseHeaderLink(components.RestAuthenticationGoogleOauthRestPaginationTypeResponseHeaderLink{/* values here */})
```

### RestAuthenticationGoogleOauthRestPaginationTypeRequestOffset

```go
restAuthenticationGoogleOauthPagination := components.CreateRestAuthenticationGoogleOauthPaginationRequestOffset(components.RestAuthenticationGoogleOauthRestPaginationTypeRequestOffset{/* values here */})
```

### RestAuthenticationGoogleOauthRestPaginationTypeRequestPage

```go
restAuthenticationGoogleOauthPagination := components.CreateRestAuthenticationGoogleOauthPaginationRequestPage(components.RestAuthenticationGoogleOauthRestPaginationTypeRequestPage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationGoogleOauthPagination.Type {
	case components.RestAuthenticationGoogleOauthPaginationTypeNone:
		// restAuthenticationGoogleOauthPagination.RestAuthenticationGoogleOauthRestPaginationTypeNone is populated
	case components.RestAuthenticationGoogleOauthPaginationTypeResponseBody:
		// restAuthenticationGoogleOauthPagination.RestAuthenticationGoogleOauthRestPaginationTypeResponseBody is populated
	case components.RestAuthenticationGoogleOauthPaginationTypeResponseHeader:
		// restAuthenticationGoogleOauthPagination.RestAuthenticationGoogleOauthRestPaginationTypeResponseHeader is populated
	case components.RestAuthenticationGoogleOauthPaginationTypeResponseHeaderLink:
		// restAuthenticationGoogleOauthPagination.RestAuthenticationGoogleOauthRestPaginationTypeResponseHeaderLink is populated
	case components.RestAuthenticationGoogleOauthPaginationTypeRequestOffset:
		// restAuthenticationGoogleOauthPagination.RestAuthenticationGoogleOauthRestPaginationTypeRequestOffset is populated
	case components.RestAuthenticationGoogleOauthPaginationTypeRequestPage:
		// restAuthenticationGoogleOauthPagination.RestAuthenticationGoogleOauthRestPaginationTypeRequestPage is populated
	default:
		// Unknown type - use restAuthenticationGoogleOauthPagination.GetUnknownRaw() for raw JSON
}
```
