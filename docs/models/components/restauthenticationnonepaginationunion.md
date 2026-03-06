# RestAuthenticationNonePaginationUnion


## Supported Types

### RestAuthenticationNoneRestPaginationTypeNone

```go
restAuthenticationNonePaginationUnion := components.CreateRestAuthenticationNonePaginationUnionNone(components.RestAuthenticationNoneRestPaginationTypeNone{/* values here */})
```

### RestAuthenticationNoneRestPaginationTypeResponseBody

```go
restAuthenticationNonePaginationUnion := components.CreateRestAuthenticationNonePaginationUnionResponseBody(components.RestAuthenticationNoneRestPaginationTypeResponseBody{/* values here */})
```

### RestAuthenticationNoneRestPaginationTypeResponseHeader

```go
restAuthenticationNonePaginationUnion := components.CreateRestAuthenticationNonePaginationUnionResponseHeader(components.RestAuthenticationNoneRestPaginationTypeResponseHeader{/* values here */})
```

### RestAuthenticationNoneRestPaginationTypeResponseHeaderLink

```go
restAuthenticationNonePaginationUnion := components.CreateRestAuthenticationNonePaginationUnionResponseHeaderLink(components.RestAuthenticationNoneRestPaginationTypeResponseHeaderLink{/* values here */})
```

### RestAuthenticationNoneRestPaginationTypeRequestOffset

```go
restAuthenticationNonePaginationUnion := components.CreateRestAuthenticationNonePaginationUnionRequestOffset(components.RestAuthenticationNoneRestPaginationTypeRequestOffset{/* values here */})
```

### RestAuthenticationNoneRestPaginationTypeRequestPage

```go
restAuthenticationNonePaginationUnion := components.CreateRestAuthenticationNonePaginationUnionRequestPage(components.RestAuthenticationNoneRestPaginationTypeRequestPage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationNonePaginationUnion.Type {
	case components.RestAuthenticationNonePaginationUnionTypeNone:
		// restAuthenticationNonePaginationUnion.RestAuthenticationNoneRestPaginationTypeNone is populated
	case components.RestAuthenticationNonePaginationUnionTypeResponseBody:
		// restAuthenticationNonePaginationUnion.RestAuthenticationNoneRestPaginationTypeResponseBody is populated
	case components.RestAuthenticationNonePaginationUnionTypeResponseHeader:
		// restAuthenticationNonePaginationUnion.RestAuthenticationNoneRestPaginationTypeResponseHeader is populated
	case components.RestAuthenticationNonePaginationUnionTypeResponseHeaderLink:
		// restAuthenticationNonePaginationUnion.RestAuthenticationNoneRestPaginationTypeResponseHeaderLink is populated
	case components.RestAuthenticationNonePaginationUnionTypeRequestOffset:
		// restAuthenticationNonePaginationUnion.RestAuthenticationNoneRestPaginationTypeRequestOffset is populated
	case components.RestAuthenticationNonePaginationUnionTypeRequestPage:
		// restAuthenticationNonePaginationUnion.RestAuthenticationNoneRestPaginationTypeRequestPage is populated
	default:
		// Unknown type - use restAuthenticationNonePaginationUnion.GetUnknownRaw() for raw JSON
}
```
