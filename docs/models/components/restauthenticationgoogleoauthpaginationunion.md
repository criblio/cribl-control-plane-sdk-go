# RestAuthenticationGoogleOauthPaginationUnion


## Supported Types

### RestAuthenticationGoogleOauthRestPaginationTypeNone

```go
restAuthenticationGoogleOauthPaginationUnion := components.CreateRestAuthenticationGoogleOauthPaginationUnionNone(components.RestAuthenticationGoogleOauthRestPaginationTypeNone{/* values here */})
```

### RestAuthenticationGoogleOauthRestPaginationTypeResponseBody

```go
restAuthenticationGoogleOauthPaginationUnion := components.CreateRestAuthenticationGoogleOauthPaginationUnionResponseBody(components.RestAuthenticationGoogleOauthRestPaginationTypeResponseBody{/* values here */})
```

### RestAuthenticationGoogleOauthRestPaginationTypeResponseHeader

```go
restAuthenticationGoogleOauthPaginationUnion := components.CreateRestAuthenticationGoogleOauthPaginationUnionResponseHeader(components.RestAuthenticationGoogleOauthRestPaginationTypeResponseHeader{/* values here */})
```

### RestAuthenticationGoogleOauthRestPaginationTypeResponseHeaderLink

```go
restAuthenticationGoogleOauthPaginationUnion := components.CreateRestAuthenticationGoogleOauthPaginationUnionResponseHeaderLink(components.RestAuthenticationGoogleOauthRestPaginationTypeResponseHeaderLink{/* values here */})
```

### RestAuthenticationGoogleOauthRestPaginationTypeRequestOffset

```go
restAuthenticationGoogleOauthPaginationUnion := components.CreateRestAuthenticationGoogleOauthPaginationUnionRequestOffset(components.RestAuthenticationGoogleOauthRestPaginationTypeRequestOffset{/* values here */})
```

### RestAuthenticationGoogleOauthRestPaginationTypeRequestPage

```go
restAuthenticationGoogleOauthPaginationUnion := components.CreateRestAuthenticationGoogleOauthPaginationUnionRequestPage(components.RestAuthenticationGoogleOauthRestPaginationTypeRequestPage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationGoogleOauthPaginationUnion.Type {
	case components.RestAuthenticationGoogleOauthPaginationUnionTypeNone:
		// restAuthenticationGoogleOauthPaginationUnion.RestAuthenticationGoogleOauthRestPaginationTypeNone is populated
	case components.RestAuthenticationGoogleOauthPaginationUnionTypeResponseBody:
		// restAuthenticationGoogleOauthPaginationUnion.RestAuthenticationGoogleOauthRestPaginationTypeResponseBody is populated
	case components.RestAuthenticationGoogleOauthPaginationUnionTypeResponseHeader:
		// restAuthenticationGoogleOauthPaginationUnion.RestAuthenticationGoogleOauthRestPaginationTypeResponseHeader is populated
	case components.RestAuthenticationGoogleOauthPaginationUnionTypeResponseHeaderLink:
		// restAuthenticationGoogleOauthPaginationUnion.RestAuthenticationGoogleOauthRestPaginationTypeResponseHeaderLink is populated
	case components.RestAuthenticationGoogleOauthPaginationUnionTypeRequestOffset:
		// restAuthenticationGoogleOauthPaginationUnion.RestAuthenticationGoogleOauthRestPaginationTypeRequestOffset is populated
	case components.RestAuthenticationGoogleOauthPaginationUnionTypeRequestPage:
		// restAuthenticationGoogleOauthPaginationUnion.RestAuthenticationGoogleOauthRestPaginationTypeRequestPage is populated
	default:
		// Unknown type - use restAuthenticationGoogleOauthPaginationUnion.GetUnknownRaw() for raw JSON
}
```
