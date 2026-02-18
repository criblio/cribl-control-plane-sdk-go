# RestCollectMethodGetPaginationUnion


## Supported Types

### RestCollectMethodGetRestPaginationTypeNone

```go
restCollectMethodGetPaginationUnion := components.CreateRestCollectMethodGetPaginationUnionNone(components.RestCollectMethodGetRestPaginationTypeNone{/* values here */})
```

### RestCollectMethodGetRestPaginationTypeResponseBody

```go
restCollectMethodGetPaginationUnion := components.CreateRestCollectMethodGetPaginationUnionResponseBody(components.RestCollectMethodGetRestPaginationTypeResponseBody{/* values here */})
```

### RestCollectMethodGetRestPaginationTypeResponseHeader

```go
restCollectMethodGetPaginationUnion := components.CreateRestCollectMethodGetPaginationUnionResponseHeader(components.RestCollectMethodGetRestPaginationTypeResponseHeader{/* values here */})
```

### RestCollectMethodGetRestPaginationTypeResponseHeaderLink

```go
restCollectMethodGetPaginationUnion := components.CreateRestCollectMethodGetPaginationUnionResponseHeaderLink(components.RestCollectMethodGetRestPaginationTypeResponseHeaderLink{/* values here */})
```

### RestCollectMethodGetRestPaginationTypeRequestOffset

```go
restCollectMethodGetPaginationUnion := components.CreateRestCollectMethodGetPaginationUnionRequestOffset(components.RestCollectMethodGetRestPaginationTypeRequestOffset{/* values here */})
```

### RestCollectMethodGetRestPaginationTypeRequestPage

```go
restCollectMethodGetPaginationUnion := components.CreateRestCollectMethodGetPaginationUnionRequestPage(components.RestCollectMethodGetRestPaginationTypeRequestPage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restCollectMethodGetPaginationUnion.Type {
	case components.RestCollectMethodGetPaginationUnionTypeNone:
		// restCollectMethodGetPaginationUnion.RestCollectMethodGetRestPaginationTypeNone is populated
	case components.RestCollectMethodGetPaginationUnionTypeResponseBody:
		// restCollectMethodGetPaginationUnion.RestCollectMethodGetRestPaginationTypeResponseBody is populated
	case components.RestCollectMethodGetPaginationUnionTypeResponseHeader:
		// restCollectMethodGetPaginationUnion.RestCollectMethodGetRestPaginationTypeResponseHeader is populated
	case components.RestCollectMethodGetPaginationUnionTypeResponseHeaderLink:
		// restCollectMethodGetPaginationUnion.RestCollectMethodGetRestPaginationTypeResponseHeaderLink is populated
	case components.RestCollectMethodGetPaginationUnionTypeRequestOffset:
		// restCollectMethodGetPaginationUnion.RestCollectMethodGetRestPaginationTypeRequestOffset is populated
	case components.RestCollectMethodGetPaginationUnionTypeRequestPage:
		// restCollectMethodGetPaginationUnion.RestCollectMethodGetRestPaginationTypeRequestPage is populated
	default:
		// Unknown type - use restCollectMethodGetPaginationUnion.GetUnknownRaw() for raw JSON
}
```
