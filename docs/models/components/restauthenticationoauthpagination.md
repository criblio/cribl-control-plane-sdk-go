# RestAuthenticationOauthPagination


## Supported Types

### RestAuthenticationOauthRestPaginationTypeNone

```go
restAuthenticationOauthPagination := components.CreateRestAuthenticationOauthPaginationNone(components.RestAuthenticationOauthRestPaginationTypeNone{/* values here */})
```

### RestAuthenticationOauthRestPaginationTypeResponseBody

```go
restAuthenticationOauthPagination := components.CreateRestAuthenticationOauthPaginationResponseBody(components.RestAuthenticationOauthRestPaginationTypeResponseBody{/* values here */})
```

### RestAuthenticationOauthRestPaginationTypeResponseHeader

```go
restAuthenticationOauthPagination := components.CreateRestAuthenticationOauthPaginationResponseHeader(components.RestAuthenticationOauthRestPaginationTypeResponseHeader{/* values here */})
```

### RestAuthenticationOauthRestPaginationTypeResponseHeaderLink

```go
restAuthenticationOauthPagination := components.CreateRestAuthenticationOauthPaginationResponseHeaderLink(components.RestAuthenticationOauthRestPaginationTypeResponseHeaderLink{/* values here */})
```

### RestAuthenticationOauthRestPaginationTypeRequestOffset

```go
restAuthenticationOauthPagination := components.CreateRestAuthenticationOauthPaginationRequestOffset(components.RestAuthenticationOauthRestPaginationTypeRequestOffset{/* values here */})
```

### RestAuthenticationOauthRestPaginationTypeRequestPage

```go
restAuthenticationOauthPagination := components.CreateRestAuthenticationOauthPaginationRequestPage(components.RestAuthenticationOauthRestPaginationTypeRequestPage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationOauthPagination.Type {
	case components.RestAuthenticationOauthPaginationTypeNone:
		// restAuthenticationOauthPagination.RestAuthenticationOauthRestPaginationTypeNone is populated
	case components.RestAuthenticationOauthPaginationTypeResponseBody:
		// restAuthenticationOauthPagination.RestAuthenticationOauthRestPaginationTypeResponseBody is populated
	case components.RestAuthenticationOauthPaginationTypeResponseHeader:
		// restAuthenticationOauthPagination.RestAuthenticationOauthRestPaginationTypeResponseHeader is populated
	case components.RestAuthenticationOauthPaginationTypeResponseHeaderLink:
		// restAuthenticationOauthPagination.RestAuthenticationOauthRestPaginationTypeResponseHeaderLink is populated
	case components.RestAuthenticationOauthPaginationTypeRequestOffset:
		// restAuthenticationOauthPagination.RestAuthenticationOauthRestPaginationTypeRequestOffset is populated
	case components.RestAuthenticationOauthPaginationTypeRequestPage:
		// restAuthenticationOauthPagination.RestAuthenticationOauthRestPaginationTypeRequestPage is populated
	default:
		// Unknown type - use restAuthenticationOauthPagination.GetUnknownRaw() for raw JSON
}
```
