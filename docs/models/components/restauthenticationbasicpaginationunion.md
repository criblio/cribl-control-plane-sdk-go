# RestAuthenticationBasicPaginationUnion


## Supported Types

### RestAuthenticationBasicRestPaginationTypeNone

```go
restAuthenticationBasicPaginationUnion := components.CreateRestAuthenticationBasicPaginationUnionNone(components.RestAuthenticationBasicRestPaginationTypeNone{/* values here */})
```

### RestAuthenticationBasicRestPaginationTypeResponseBody

```go
restAuthenticationBasicPaginationUnion := components.CreateRestAuthenticationBasicPaginationUnionResponseBody(components.RestAuthenticationBasicRestPaginationTypeResponseBody{/* values here */})
```

### RestAuthenticationBasicRestPaginationTypeResponseHeader

```go
restAuthenticationBasicPaginationUnion := components.CreateRestAuthenticationBasicPaginationUnionResponseHeader(components.RestAuthenticationBasicRestPaginationTypeResponseHeader{/* values here */})
```

### RestAuthenticationBasicRestPaginationTypeResponseHeaderLink

```go
restAuthenticationBasicPaginationUnion := components.CreateRestAuthenticationBasicPaginationUnionResponseHeaderLink(components.RestAuthenticationBasicRestPaginationTypeResponseHeaderLink{/* values here */})
```

### RestAuthenticationBasicRestPaginationTypeRequestOffset

```go
restAuthenticationBasicPaginationUnion := components.CreateRestAuthenticationBasicPaginationUnionRequestOffset(components.RestAuthenticationBasicRestPaginationTypeRequestOffset{/* values here */})
```

### RestAuthenticationBasicRestPaginationTypeRequestPage

```go
restAuthenticationBasicPaginationUnion := components.CreateRestAuthenticationBasicPaginationUnionRequestPage(components.RestAuthenticationBasicRestPaginationTypeRequestPage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationBasicPaginationUnion.Type {
	case components.RestAuthenticationBasicPaginationUnionTypeNone:
		// restAuthenticationBasicPaginationUnion.RestAuthenticationBasicRestPaginationTypeNone is populated
	case components.RestAuthenticationBasicPaginationUnionTypeResponseBody:
		// restAuthenticationBasicPaginationUnion.RestAuthenticationBasicRestPaginationTypeResponseBody is populated
	case components.RestAuthenticationBasicPaginationUnionTypeResponseHeader:
		// restAuthenticationBasicPaginationUnion.RestAuthenticationBasicRestPaginationTypeResponseHeader is populated
	case components.RestAuthenticationBasicPaginationUnionTypeResponseHeaderLink:
		// restAuthenticationBasicPaginationUnion.RestAuthenticationBasicRestPaginationTypeResponseHeaderLink is populated
	case components.RestAuthenticationBasicPaginationUnionTypeRequestOffset:
		// restAuthenticationBasicPaginationUnion.RestAuthenticationBasicRestPaginationTypeRequestOffset is populated
	case components.RestAuthenticationBasicPaginationUnionTypeRequestPage:
		// restAuthenticationBasicPaginationUnion.RestAuthenticationBasicRestPaginationTypeRequestPage is populated
	default:
		// Unknown type - use restAuthenticationBasicPaginationUnion.GetUnknownRaw() for raw JSON
}
```
