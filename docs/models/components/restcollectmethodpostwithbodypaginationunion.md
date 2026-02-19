# RestCollectMethodPostWithBodyPaginationUnion


## Supported Types

### RestCollectMethodPostWithBodyRestPaginationTypeNone

```go
restCollectMethodPostWithBodyPaginationUnion := components.CreateRestCollectMethodPostWithBodyPaginationUnionNone(components.RestCollectMethodPostWithBodyRestPaginationTypeNone{/* values here */})
```

### RestCollectMethodPostWithBodyRestPaginationTypeResponseBody

```go
restCollectMethodPostWithBodyPaginationUnion := components.CreateRestCollectMethodPostWithBodyPaginationUnionResponseBody(components.RestCollectMethodPostWithBodyRestPaginationTypeResponseBody{/* values here */})
```

### RestCollectMethodPostWithBodyRestPaginationTypeResponseHeader

```go
restCollectMethodPostWithBodyPaginationUnion := components.CreateRestCollectMethodPostWithBodyPaginationUnionResponseHeader(components.RestCollectMethodPostWithBodyRestPaginationTypeResponseHeader{/* values here */})
```

### RestCollectMethodPostWithBodyRestPaginationTypeResponseHeaderLink

```go
restCollectMethodPostWithBodyPaginationUnion := components.CreateRestCollectMethodPostWithBodyPaginationUnionResponseHeaderLink(components.RestCollectMethodPostWithBodyRestPaginationTypeResponseHeaderLink{/* values here */})
```

### RestCollectMethodPostWithBodyRestPaginationTypeRequestOffset

```go
restCollectMethodPostWithBodyPaginationUnion := components.CreateRestCollectMethodPostWithBodyPaginationUnionRequestOffset(components.RestCollectMethodPostWithBodyRestPaginationTypeRequestOffset{/* values here */})
```

### RestCollectMethodPostWithBodyRestPaginationTypeRequestPage

```go
restCollectMethodPostWithBodyPaginationUnion := components.CreateRestCollectMethodPostWithBodyPaginationUnionRequestPage(components.RestCollectMethodPostWithBodyRestPaginationTypeRequestPage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restCollectMethodPostWithBodyPaginationUnion.Type {
	case components.RestCollectMethodPostWithBodyPaginationUnionTypeNone:
		// restCollectMethodPostWithBodyPaginationUnion.RestCollectMethodPostWithBodyRestPaginationTypeNone is populated
	case components.RestCollectMethodPostWithBodyPaginationUnionTypeResponseBody:
		// restCollectMethodPostWithBodyPaginationUnion.RestCollectMethodPostWithBodyRestPaginationTypeResponseBody is populated
	case components.RestCollectMethodPostWithBodyPaginationUnionTypeResponseHeader:
		// restCollectMethodPostWithBodyPaginationUnion.RestCollectMethodPostWithBodyRestPaginationTypeResponseHeader is populated
	case components.RestCollectMethodPostWithBodyPaginationUnionTypeResponseHeaderLink:
		// restCollectMethodPostWithBodyPaginationUnion.RestCollectMethodPostWithBodyRestPaginationTypeResponseHeaderLink is populated
	case components.RestCollectMethodPostWithBodyPaginationUnionTypeRequestOffset:
		// restCollectMethodPostWithBodyPaginationUnion.RestCollectMethodPostWithBodyRestPaginationTypeRequestOffset is populated
	case components.RestCollectMethodPostWithBodyPaginationUnionTypeRequestPage:
		// restCollectMethodPostWithBodyPaginationUnion.RestCollectMethodPostWithBodyRestPaginationTypeRequestPage is populated
	default:
		// Unknown type - use restCollectMethodPostWithBodyPaginationUnion.GetUnknownRaw() for raw JSON
}
```
