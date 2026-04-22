# RestAuthenticationLoginPagination


## Supported Types

### RestAuthenticationLoginRestPaginationTypeNone

```go
restAuthenticationLoginPagination := components.CreateRestAuthenticationLoginPaginationNone(components.RestAuthenticationLoginRestPaginationTypeNone{/* values here */})
```

### RestAuthenticationLoginRestPaginationTypeResponseBody

```go
restAuthenticationLoginPagination := components.CreateRestAuthenticationLoginPaginationResponseBody(components.RestAuthenticationLoginRestPaginationTypeResponseBody{/* values here */})
```

### RestAuthenticationLoginRestPaginationTypeResponseHeader

```go
restAuthenticationLoginPagination := components.CreateRestAuthenticationLoginPaginationResponseHeader(components.RestAuthenticationLoginRestPaginationTypeResponseHeader{/* values here */})
```

### RestAuthenticationLoginRestPaginationTypeResponseHeaderLink

```go
restAuthenticationLoginPagination := components.CreateRestAuthenticationLoginPaginationResponseHeaderLink(components.RestAuthenticationLoginRestPaginationTypeResponseHeaderLink{/* values here */})
```

### RestAuthenticationLoginRestPaginationTypeRequestOffset

```go
restAuthenticationLoginPagination := components.CreateRestAuthenticationLoginPaginationRequestOffset(components.RestAuthenticationLoginRestPaginationTypeRequestOffset{/* values here */})
```

### RestAuthenticationLoginRestPaginationTypeRequestPage

```go
restAuthenticationLoginPagination := components.CreateRestAuthenticationLoginPaginationRequestPage(components.RestAuthenticationLoginRestPaginationTypeRequestPage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationLoginPagination.Type {
	case components.RestAuthenticationLoginPaginationTypeNone:
		// restAuthenticationLoginPagination.RestAuthenticationLoginRestPaginationTypeNone is populated
	case components.RestAuthenticationLoginPaginationTypeResponseBody:
		// restAuthenticationLoginPagination.RestAuthenticationLoginRestPaginationTypeResponseBody is populated
	case components.RestAuthenticationLoginPaginationTypeResponseHeader:
		// restAuthenticationLoginPagination.RestAuthenticationLoginRestPaginationTypeResponseHeader is populated
	case components.RestAuthenticationLoginPaginationTypeResponseHeaderLink:
		// restAuthenticationLoginPagination.RestAuthenticationLoginRestPaginationTypeResponseHeaderLink is populated
	case components.RestAuthenticationLoginPaginationTypeRequestOffset:
		// restAuthenticationLoginPagination.RestAuthenticationLoginRestPaginationTypeRequestOffset is populated
	case components.RestAuthenticationLoginPaginationTypeRequestPage:
		// restAuthenticationLoginPagination.RestAuthenticationLoginRestPaginationTypeRequestPage is populated
	default:
		// Unknown type - use restAuthenticationLoginPagination.GetUnknownRaw() for raw JSON
}
```
