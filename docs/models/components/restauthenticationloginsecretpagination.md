# RestAuthenticationLoginSecretPagination


## Supported Types

### RestAuthenticationLoginSecretRestPaginationTypeNone

```go
restAuthenticationLoginSecretPagination := components.CreateRestAuthenticationLoginSecretPaginationNone(components.RestAuthenticationLoginSecretRestPaginationTypeNone{/* values here */})
```

### RestAuthenticationLoginSecretRestPaginationTypeResponseBody

```go
restAuthenticationLoginSecretPagination := components.CreateRestAuthenticationLoginSecretPaginationResponseBody(components.RestAuthenticationLoginSecretRestPaginationTypeResponseBody{/* values here */})
```

### RestAuthenticationLoginSecretRestPaginationTypeResponseHeader

```go
restAuthenticationLoginSecretPagination := components.CreateRestAuthenticationLoginSecretPaginationResponseHeader(components.RestAuthenticationLoginSecretRestPaginationTypeResponseHeader{/* values here */})
```

### RestAuthenticationLoginSecretRestPaginationTypeResponseHeaderLink

```go
restAuthenticationLoginSecretPagination := components.CreateRestAuthenticationLoginSecretPaginationResponseHeaderLink(components.RestAuthenticationLoginSecretRestPaginationTypeResponseHeaderLink{/* values here */})
```

### RestAuthenticationLoginSecretRestPaginationTypeRequestOffset

```go
restAuthenticationLoginSecretPagination := components.CreateRestAuthenticationLoginSecretPaginationRequestOffset(components.RestAuthenticationLoginSecretRestPaginationTypeRequestOffset{/* values here */})
```

### RestAuthenticationLoginSecretRestPaginationTypeRequestPage

```go
restAuthenticationLoginSecretPagination := components.CreateRestAuthenticationLoginSecretPaginationRequestPage(components.RestAuthenticationLoginSecretRestPaginationTypeRequestPage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationLoginSecretPagination.Type {
	case components.RestAuthenticationLoginSecretPaginationTypeNone:
		// restAuthenticationLoginSecretPagination.RestAuthenticationLoginSecretRestPaginationTypeNone is populated
	case components.RestAuthenticationLoginSecretPaginationTypeResponseBody:
		// restAuthenticationLoginSecretPagination.RestAuthenticationLoginSecretRestPaginationTypeResponseBody is populated
	case components.RestAuthenticationLoginSecretPaginationTypeResponseHeader:
		// restAuthenticationLoginSecretPagination.RestAuthenticationLoginSecretRestPaginationTypeResponseHeader is populated
	case components.RestAuthenticationLoginSecretPaginationTypeResponseHeaderLink:
		// restAuthenticationLoginSecretPagination.RestAuthenticationLoginSecretRestPaginationTypeResponseHeaderLink is populated
	case components.RestAuthenticationLoginSecretPaginationTypeRequestOffset:
		// restAuthenticationLoginSecretPagination.RestAuthenticationLoginSecretRestPaginationTypeRequestOffset is populated
	case components.RestAuthenticationLoginSecretPaginationTypeRequestPage:
		// restAuthenticationLoginSecretPagination.RestAuthenticationLoginSecretRestPaginationTypeRequestPage is populated
	default:
		// Unknown type - use restAuthenticationLoginSecretPagination.GetUnknownRaw() for raw JSON
}
```
