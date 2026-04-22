# RestAuthenticationBasicPagination


## Supported Types

### RestAuthenticationBasicRestPaginationTypeNone

```go
restAuthenticationBasicPagination := components.CreateRestAuthenticationBasicPaginationNone(components.RestAuthenticationBasicRestPaginationTypeNone{/* values here */})
```

### RestAuthenticationBasicRestPaginationTypeResponseBody

```go
restAuthenticationBasicPagination := components.CreateRestAuthenticationBasicPaginationResponseBody(components.RestAuthenticationBasicRestPaginationTypeResponseBody{/* values here */})
```

### RestAuthenticationBasicRestPaginationTypeResponseHeader

```go
restAuthenticationBasicPagination := components.CreateRestAuthenticationBasicPaginationResponseHeader(components.RestAuthenticationBasicRestPaginationTypeResponseHeader{/* values here */})
```

### RestAuthenticationBasicRestPaginationTypeResponseHeaderLink

```go
restAuthenticationBasicPagination := components.CreateRestAuthenticationBasicPaginationResponseHeaderLink(components.RestAuthenticationBasicRestPaginationTypeResponseHeaderLink{/* values here */})
```

### RestAuthenticationBasicRestPaginationTypeRequestOffset

```go
restAuthenticationBasicPagination := components.CreateRestAuthenticationBasicPaginationRequestOffset(components.RestAuthenticationBasicRestPaginationTypeRequestOffset{/* values here */})
```

### RestAuthenticationBasicRestPaginationTypeRequestPage

```go
restAuthenticationBasicPagination := components.CreateRestAuthenticationBasicPaginationRequestPage(components.RestAuthenticationBasicRestPaginationTypeRequestPage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationBasicPagination.Type {
	case components.RestAuthenticationBasicPaginationTypeNone:
		// restAuthenticationBasicPagination.RestAuthenticationBasicRestPaginationTypeNone is populated
	case components.RestAuthenticationBasicPaginationTypeResponseBody:
		// restAuthenticationBasicPagination.RestAuthenticationBasicRestPaginationTypeResponseBody is populated
	case components.RestAuthenticationBasicPaginationTypeResponseHeader:
		// restAuthenticationBasicPagination.RestAuthenticationBasicRestPaginationTypeResponseHeader is populated
	case components.RestAuthenticationBasicPaginationTypeResponseHeaderLink:
		// restAuthenticationBasicPagination.RestAuthenticationBasicRestPaginationTypeResponseHeaderLink is populated
	case components.RestAuthenticationBasicPaginationTypeRequestOffset:
		// restAuthenticationBasicPagination.RestAuthenticationBasicRestPaginationTypeRequestOffset is populated
	case components.RestAuthenticationBasicPaginationTypeRequestPage:
		// restAuthenticationBasicPagination.RestAuthenticationBasicRestPaginationTypeRequestPage is populated
	default:
		// Unknown type - use restAuthenticationBasicPagination.GetUnknownRaw() for raw JSON
}
```
