# RestAuthenticationHmacPagination


## Supported Types

### RestAuthenticationHmacRestPaginationTypeNone

```go
restAuthenticationHmacPagination := components.CreateRestAuthenticationHmacPaginationNone(components.RestAuthenticationHmacRestPaginationTypeNone{/* values here */})
```

### RestAuthenticationHmacRestPaginationTypeResponseBody

```go
restAuthenticationHmacPagination := components.CreateRestAuthenticationHmacPaginationResponseBody(components.RestAuthenticationHmacRestPaginationTypeResponseBody{/* values here */})
```

### RestAuthenticationHmacRestPaginationTypeResponseHeader

```go
restAuthenticationHmacPagination := components.CreateRestAuthenticationHmacPaginationResponseHeader(components.RestAuthenticationHmacRestPaginationTypeResponseHeader{/* values here */})
```

### RestAuthenticationHmacRestPaginationTypeResponseHeaderLink

```go
restAuthenticationHmacPagination := components.CreateRestAuthenticationHmacPaginationResponseHeaderLink(components.RestAuthenticationHmacRestPaginationTypeResponseHeaderLink{/* values here */})
```

### RestAuthenticationHmacRestPaginationTypeRequestOffset

```go
restAuthenticationHmacPagination := components.CreateRestAuthenticationHmacPaginationRequestOffset(components.RestAuthenticationHmacRestPaginationTypeRequestOffset{/* values here */})
```

### RestAuthenticationHmacRestPaginationTypeRequestPage

```go
restAuthenticationHmacPagination := components.CreateRestAuthenticationHmacPaginationRequestPage(components.RestAuthenticationHmacRestPaginationTypeRequestPage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationHmacPagination.Type {
	case components.RestAuthenticationHmacPaginationTypeNone:
		// restAuthenticationHmacPagination.RestAuthenticationHmacRestPaginationTypeNone is populated
	case components.RestAuthenticationHmacPaginationTypeResponseBody:
		// restAuthenticationHmacPagination.RestAuthenticationHmacRestPaginationTypeResponseBody is populated
	case components.RestAuthenticationHmacPaginationTypeResponseHeader:
		// restAuthenticationHmacPagination.RestAuthenticationHmacRestPaginationTypeResponseHeader is populated
	case components.RestAuthenticationHmacPaginationTypeResponseHeaderLink:
		// restAuthenticationHmacPagination.RestAuthenticationHmacRestPaginationTypeResponseHeaderLink is populated
	case components.RestAuthenticationHmacPaginationTypeRequestOffset:
		// restAuthenticationHmacPagination.RestAuthenticationHmacRestPaginationTypeRequestOffset is populated
	case components.RestAuthenticationHmacPaginationTypeRequestPage:
		// restAuthenticationHmacPagination.RestAuthenticationHmacRestPaginationTypeRequestPage is populated
	default:
		// Unknown type - use restAuthenticationHmacPagination.GetUnknownRaw() for raw JSON
}
```
