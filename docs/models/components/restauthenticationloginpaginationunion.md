# RestAuthenticationLoginPaginationUnion


## Supported Types

### RestAuthenticationLoginRestPaginationTypeNone

```go
restAuthenticationLoginPaginationUnion := components.CreateRestAuthenticationLoginPaginationUnionNone(components.RestAuthenticationLoginRestPaginationTypeNone{/* values here */})
```

### RestAuthenticationLoginRestPaginationTypeResponseBody

```go
restAuthenticationLoginPaginationUnion := components.CreateRestAuthenticationLoginPaginationUnionResponseBody(components.RestAuthenticationLoginRestPaginationTypeResponseBody{/* values here */})
```

### RestAuthenticationLoginRestPaginationTypeResponseHeader

```go
restAuthenticationLoginPaginationUnion := components.CreateRestAuthenticationLoginPaginationUnionResponseHeader(components.RestAuthenticationLoginRestPaginationTypeResponseHeader{/* values here */})
```

### RestAuthenticationLoginRestPaginationTypeResponseHeaderLink

```go
restAuthenticationLoginPaginationUnion := components.CreateRestAuthenticationLoginPaginationUnionResponseHeaderLink(components.RestAuthenticationLoginRestPaginationTypeResponseHeaderLink{/* values here */})
```

### RestAuthenticationLoginRestPaginationTypeRequestOffset

```go
restAuthenticationLoginPaginationUnion := components.CreateRestAuthenticationLoginPaginationUnionRequestOffset(components.RestAuthenticationLoginRestPaginationTypeRequestOffset{/* values here */})
```

### RestAuthenticationLoginRestPaginationTypeRequestPage

```go
restAuthenticationLoginPaginationUnion := components.CreateRestAuthenticationLoginPaginationUnionRequestPage(components.RestAuthenticationLoginRestPaginationTypeRequestPage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationLoginPaginationUnion.Type {
	case components.RestAuthenticationLoginPaginationUnionTypeNone:
		// restAuthenticationLoginPaginationUnion.RestAuthenticationLoginRestPaginationTypeNone is populated
	case components.RestAuthenticationLoginPaginationUnionTypeResponseBody:
		// restAuthenticationLoginPaginationUnion.RestAuthenticationLoginRestPaginationTypeResponseBody is populated
	case components.RestAuthenticationLoginPaginationUnionTypeResponseHeader:
		// restAuthenticationLoginPaginationUnion.RestAuthenticationLoginRestPaginationTypeResponseHeader is populated
	case components.RestAuthenticationLoginPaginationUnionTypeResponseHeaderLink:
		// restAuthenticationLoginPaginationUnion.RestAuthenticationLoginRestPaginationTypeResponseHeaderLink is populated
	case components.RestAuthenticationLoginPaginationUnionTypeRequestOffset:
		// restAuthenticationLoginPaginationUnion.RestAuthenticationLoginRestPaginationTypeRequestOffset is populated
	case components.RestAuthenticationLoginPaginationUnionTypeRequestPage:
		// restAuthenticationLoginPaginationUnion.RestAuthenticationLoginRestPaginationTypeRequestPage is populated
	default:
		// Unknown type - use restAuthenticationLoginPaginationUnion.GetUnknownRaw() for raw JSON
}
```
