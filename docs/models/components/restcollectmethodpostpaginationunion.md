# RestCollectMethodPostPaginationUnion


## Supported Types

### RestCollectMethodPostRestPaginationTypeNone

```go
restCollectMethodPostPaginationUnion := components.CreateRestCollectMethodPostPaginationUnionNone(components.RestCollectMethodPostRestPaginationTypeNone{/* values here */})
```

### RestCollectMethodPostRestPaginationTypeResponseBody

```go
restCollectMethodPostPaginationUnion := components.CreateRestCollectMethodPostPaginationUnionResponseBody(components.RestCollectMethodPostRestPaginationTypeResponseBody{/* values here */})
```

### RestCollectMethodPostRestPaginationTypeResponseHeader

```go
restCollectMethodPostPaginationUnion := components.CreateRestCollectMethodPostPaginationUnionResponseHeader(components.RestCollectMethodPostRestPaginationTypeResponseHeader{/* values here */})
```

### RestCollectMethodPostRestPaginationTypeResponseHeaderLink

```go
restCollectMethodPostPaginationUnion := components.CreateRestCollectMethodPostPaginationUnionResponseHeaderLink(components.RestCollectMethodPostRestPaginationTypeResponseHeaderLink{/* values here */})
```

### RestCollectMethodPostRestPaginationTypeRequestOffset

```go
restCollectMethodPostPaginationUnion := components.CreateRestCollectMethodPostPaginationUnionRequestOffset(components.RestCollectMethodPostRestPaginationTypeRequestOffset{/* values here */})
```

### RestCollectMethodPostRestPaginationTypeRequestPage

```go
restCollectMethodPostPaginationUnion := components.CreateRestCollectMethodPostPaginationUnionRequestPage(components.RestCollectMethodPostRestPaginationTypeRequestPage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restCollectMethodPostPaginationUnion.Type {
	case components.RestCollectMethodPostPaginationUnionTypeNone:
		// restCollectMethodPostPaginationUnion.RestCollectMethodPostRestPaginationTypeNone is populated
	case components.RestCollectMethodPostPaginationUnionTypeResponseBody:
		// restCollectMethodPostPaginationUnion.RestCollectMethodPostRestPaginationTypeResponseBody is populated
	case components.RestCollectMethodPostPaginationUnionTypeResponseHeader:
		// restCollectMethodPostPaginationUnion.RestCollectMethodPostRestPaginationTypeResponseHeader is populated
	case components.RestCollectMethodPostPaginationUnionTypeResponseHeaderLink:
		// restCollectMethodPostPaginationUnion.RestCollectMethodPostRestPaginationTypeResponseHeaderLink is populated
	case components.RestCollectMethodPostPaginationUnionTypeRequestOffset:
		// restCollectMethodPostPaginationUnion.RestCollectMethodPostRestPaginationTypeRequestOffset is populated
	case components.RestCollectMethodPostPaginationUnionTypeRequestPage:
		// restCollectMethodPostPaginationUnion.RestCollectMethodPostRestPaginationTypeRequestPage is populated
	default:
		// Unknown type - use restCollectMethodPostPaginationUnion.GetUnknownRaw() for raw JSON
}
```
