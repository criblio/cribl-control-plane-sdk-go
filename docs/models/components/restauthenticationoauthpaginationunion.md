# RestAuthenticationOauthPaginationUnion


## Supported Types

### RestAuthenticationOauthRestPaginationTypeNone

```go
restAuthenticationOauthPaginationUnion := components.CreateRestAuthenticationOauthPaginationUnionNone(components.RestAuthenticationOauthRestPaginationTypeNone{/* values here */})
```

### RestAuthenticationOauthRestPaginationTypeResponseBody

```go
restAuthenticationOauthPaginationUnion := components.CreateRestAuthenticationOauthPaginationUnionResponseBody(components.RestAuthenticationOauthRestPaginationTypeResponseBody{/* values here */})
```

### RestAuthenticationOauthRestPaginationTypeResponseHeader

```go
restAuthenticationOauthPaginationUnion := components.CreateRestAuthenticationOauthPaginationUnionResponseHeader(components.RestAuthenticationOauthRestPaginationTypeResponseHeader{/* values here */})
```

### RestAuthenticationOauthRestPaginationTypeResponseHeaderLink

```go
restAuthenticationOauthPaginationUnion := components.CreateRestAuthenticationOauthPaginationUnionResponseHeaderLink(components.RestAuthenticationOauthRestPaginationTypeResponseHeaderLink{/* values here */})
```

### RestAuthenticationOauthRestPaginationTypeRequestOffset

```go
restAuthenticationOauthPaginationUnion := components.CreateRestAuthenticationOauthPaginationUnionRequestOffset(components.RestAuthenticationOauthRestPaginationTypeRequestOffset{/* values here */})
```

### RestAuthenticationOauthRestPaginationTypeRequestPage

```go
restAuthenticationOauthPaginationUnion := components.CreateRestAuthenticationOauthPaginationUnionRequestPage(components.RestAuthenticationOauthRestPaginationTypeRequestPage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationOauthPaginationUnion.Type {
	case components.RestAuthenticationOauthPaginationUnionTypeNone:
		// restAuthenticationOauthPaginationUnion.RestAuthenticationOauthRestPaginationTypeNone is populated
	case components.RestAuthenticationOauthPaginationUnionTypeResponseBody:
		// restAuthenticationOauthPaginationUnion.RestAuthenticationOauthRestPaginationTypeResponseBody is populated
	case components.RestAuthenticationOauthPaginationUnionTypeResponseHeader:
		// restAuthenticationOauthPaginationUnion.RestAuthenticationOauthRestPaginationTypeResponseHeader is populated
	case components.RestAuthenticationOauthPaginationUnionTypeResponseHeaderLink:
		// restAuthenticationOauthPaginationUnion.RestAuthenticationOauthRestPaginationTypeResponseHeaderLink is populated
	case components.RestAuthenticationOauthPaginationUnionTypeRequestOffset:
		// restAuthenticationOauthPaginationUnion.RestAuthenticationOauthRestPaginationTypeRequestOffset is populated
	case components.RestAuthenticationOauthPaginationUnionTypeRequestPage:
		// restAuthenticationOauthPaginationUnion.RestAuthenticationOauthRestPaginationTypeRequestPage is populated
	default:
		// Unknown type - use restAuthenticationOauthPaginationUnion.GetUnknownRaw() for raw JSON
}
```
