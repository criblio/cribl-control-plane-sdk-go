# RestAuthenticationBasicSecretPagination


## Supported Types

### RestAuthenticationBasicSecretRestPaginationTypeNone

```go
restAuthenticationBasicSecretPagination := components.CreateRestAuthenticationBasicSecretPaginationNone(components.RestAuthenticationBasicSecretRestPaginationTypeNone{/* values here */})
```

### RestAuthenticationBasicSecretRestPaginationTypeResponseBody

```go
restAuthenticationBasicSecretPagination := components.CreateRestAuthenticationBasicSecretPaginationResponseBody(components.RestAuthenticationBasicSecretRestPaginationTypeResponseBody{/* values here */})
```

### RestAuthenticationBasicSecretRestPaginationTypeResponseHeader

```go
restAuthenticationBasicSecretPagination := components.CreateRestAuthenticationBasicSecretPaginationResponseHeader(components.RestAuthenticationBasicSecretRestPaginationTypeResponseHeader{/* values here */})
```

### RestAuthenticationBasicSecretRestPaginationTypeResponseHeaderLink

```go
restAuthenticationBasicSecretPagination := components.CreateRestAuthenticationBasicSecretPaginationResponseHeaderLink(components.RestAuthenticationBasicSecretRestPaginationTypeResponseHeaderLink{/* values here */})
```

### RestAuthenticationBasicSecretRestPaginationTypeRequestOffset

```go
restAuthenticationBasicSecretPagination := components.CreateRestAuthenticationBasicSecretPaginationRequestOffset(components.RestAuthenticationBasicSecretRestPaginationTypeRequestOffset{/* values here */})
```

### RestAuthenticationBasicSecretRestPaginationTypeRequestPage

```go
restAuthenticationBasicSecretPagination := components.CreateRestAuthenticationBasicSecretPaginationRequestPage(components.RestAuthenticationBasicSecretRestPaginationTypeRequestPage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationBasicSecretPagination.Type {
	case components.RestAuthenticationBasicSecretPaginationTypeNone:
		// restAuthenticationBasicSecretPagination.RestAuthenticationBasicSecretRestPaginationTypeNone is populated
	case components.RestAuthenticationBasicSecretPaginationTypeResponseBody:
		// restAuthenticationBasicSecretPagination.RestAuthenticationBasicSecretRestPaginationTypeResponseBody is populated
	case components.RestAuthenticationBasicSecretPaginationTypeResponseHeader:
		// restAuthenticationBasicSecretPagination.RestAuthenticationBasicSecretRestPaginationTypeResponseHeader is populated
	case components.RestAuthenticationBasicSecretPaginationTypeResponseHeaderLink:
		// restAuthenticationBasicSecretPagination.RestAuthenticationBasicSecretRestPaginationTypeResponseHeaderLink is populated
	case components.RestAuthenticationBasicSecretPaginationTypeRequestOffset:
		// restAuthenticationBasicSecretPagination.RestAuthenticationBasicSecretRestPaginationTypeRequestOffset is populated
	case components.RestAuthenticationBasicSecretPaginationTypeRequestPage:
		// restAuthenticationBasicSecretPagination.RestAuthenticationBasicSecretRestPaginationTypeRequestPage is populated
	default:
		// Unknown type - use restAuthenticationBasicSecretPagination.GetUnknownRaw() for raw JSON
}
```
