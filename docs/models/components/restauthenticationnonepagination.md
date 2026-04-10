# RestAuthenticationNonePagination


## Supported Types

### RestAuthenticationNoneRestPaginationTypeNone

```go
restAuthenticationNonePagination := components.CreateRestAuthenticationNonePaginationNone(components.RestAuthenticationNoneRestPaginationTypeNone{/* values here */})
```

### RestAuthenticationNoneRestPaginationTypeResponseBody

```go
restAuthenticationNonePagination := components.CreateRestAuthenticationNonePaginationResponseBody(components.RestAuthenticationNoneRestPaginationTypeResponseBody{/* values here */})
```

### RestAuthenticationNoneRestPaginationTypeResponseHeader

```go
restAuthenticationNonePagination := components.CreateRestAuthenticationNonePaginationResponseHeader(components.RestAuthenticationNoneRestPaginationTypeResponseHeader{/* values here */})
```

### RestAuthenticationNoneRestPaginationTypeResponseHeaderLink

```go
restAuthenticationNonePagination := components.CreateRestAuthenticationNonePaginationResponseHeaderLink(components.RestAuthenticationNoneRestPaginationTypeResponseHeaderLink{/* values here */})
```

### RestAuthenticationNoneRestPaginationTypeRequestOffset

```go
restAuthenticationNonePagination := components.CreateRestAuthenticationNonePaginationRequestOffset(components.RestAuthenticationNoneRestPaginationTypeRequestOffset{/* values here */})
```

### RestAuthenticationNoneRestPaginationTypeRequestPage

```go
restAuthenticationNonePagination := components.CreateRestAuthenticationNonePaginationRequestPage(components.RestAuthenticationNoneRestPaginationTypeRequestPage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationNonePagination.Type {
	case components.RestAuthenticationNonePaginationTypeNone:
		// restAuthenticationNonePagination.RestAuthenticationNoneRestPaginationTypeNone is populated
	case components.RestAuthenticationNonePaginationTypeResponseBody:
		// restAuthenticationNonePagination.RestAuthenticationNoneRestPaginationTypeResponseBody is populated
	case components.RestAuthenticationNonePaginationTypeResponseHeader:
		// restAuthenticationNonePagination.RestAuthenticationNoneRestPaginationTypeResponseHeader is populated
	case components.RestAuthenticationNonePaginationTypeResponseHeaderLink:
		// restAuthenticationNonePagination.RestAuthenticationNoneRestPaginationTypeResponseHeaderLink is populated
	case components.RestAuthenticationNonePaginationTypeRequestOffset:
		// restAuthenticationNonePagination.RestAuthenticationNoneRestPaginationTypeRequestOffset is populated
	case components.RestAuthenticationNonePaginationTypeRequestPage:
		// restAuthenticationNonePagination.RestAuthenticationNoneRestPaginationTypeRequestPage is populated
	default:
		// Unknown type - use restAuthenticationNonePagination.GetUnknownRaw() for raw JSON
}
```
