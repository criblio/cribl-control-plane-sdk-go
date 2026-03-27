# RestAuthenticationHmacPaginationUnion


## Supported Types

### RestAuthenticationHmacRestPaginationTypeNone

```go
restAuthenticationHmacPaginationUnion := components.CreateRestAuthenticationHmacPaginationUnionNone(components.RestAuthenticationHmacRestPaginationTypeNone{/* values here */})
```

### RestAuthenticationHmacRestPaginationTypeResponseBody

```go
restAuthenticationHmacPaginationUnion := components.CreateRestAuthenticationHmacPaginationUnionResponseBody(components.RestAuthenticationHmacRestPaginationTypeResponseBody{/* values here */})
```

### RestAuthenticationHmacRestPaginationTypeResponseHeader

```go
restAuthenticationHmacPaginationUnion := components.CreateRestAuthenticationHmacPaginationUnionResponseHeader(components.RestAuthenticationHmacRestPaginationTypeResponseHeader{/* values here */})
```

### RestAuthenticationHmacRestPaginationTypeResponseHeaderLink

```go
restAuthenticationHmacPaginationUnion := components.CreateRestAuthenticationHmacPaginationUnionResponseHeaderLink(components.RestAuthenticationHmacRestPaginationTypeResponseHeaderLink{/* values here */})
```

### RestAuthenticationHmacRestPaginationTypeRequestOffset

```go
restAuthenticationHmacPaginationUnion := components.CreateRestAuthenticationHmacPaginationUnionRequestOffset(components.RestAuthenticationHmacRestPaginationTypeRequestOffset{/* values here */})
```

### RestAuthenticationHmacRestPaginationTypeRequestPage

```go
restAuthenticationHmacPaginationUnion := components.CreateRestAuthenticationHmacPaginationUnionRequestPage(components.RestAuthenticationHmacRestPaginationTypeRequestPage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationHmacPaginationUnion.Type {
	case components.RestAuthenticationHmacPaginationUnionTypeNone:
		// restAuthenticationHmacPaginationUnion.RestAuthenticationHmacRestPaginationTypeNone is populated
	case components.RestAuthenticationHmacPaginationUnionTypeResponseBody:
		// restAuthenticationHmacPaginationUnion.RestAuthenticationHmacRestPaginationTypeResponseBody is populated
	case components.RestAuthenticationHmacPaginationUnionTypeResponseHeader:
		// restAuthenticationHmacPaginationUnion.RestAuthenticationHmacRestPaginationTypeResponseHeader is populated
	case components.RestAuthenticationHmacPaginationUnionTypeResponseHeaderLink:
		// restAuthenticationHmacPaginationUnion.RestAuthenticationHmacRestPaginationTypeResponseHeaderLink is populated
	case components.RestAuthenticationHmacPaginationUnionTypeRequestOffset:
		// restAuthenticationHmacPaginationUnion.RestAuthenticationHmacRestPaginationTypeRequestOffset is populated
	case components.RestAuthenticationHmacPaginationUnionTypeRequestPage:
		// restAuthenticationHmacPaginationUnion.RestAuthenticationHmacRestPaginationTypeRequestPage is populated
	default:
		// Unknown type - use restAuthenticationHmacPaginationUnion.GetUnknownRaw() for raw JSON
}
```
