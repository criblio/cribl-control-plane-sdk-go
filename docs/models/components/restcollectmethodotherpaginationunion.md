# RestCollectMethodOtherPaginationUnion


## Supported Types

### RestCollectMethodOtherRestPaginationTypeNone

```go
restCollectMethodOtherPaginationUnion := components.CreateRestCollectMethodOtherPaginationUnionNone(components.RestCollectMethodOtherRestPaginationTypeNone{/* values here */})
```

### RestCollectMethodOtherRestPaginationTypeResponseBody

```go
restCollectMethodOtherPaginationUnion := components.CreateRestCollectMethodOtherPaginationUnionResponseBody(components.RestCollectMethodOtherRestPaginationTypeResponseBody{/* values here */})
```

### RestCollectMethodOtherRestPaginationTypeResponseHeader

```go
restCollectMethodOtherPaginationUnion := components.CreateRestCollectMethodOtherPaginationUnionResponseHeader(components.RestCollectMethodOtherRestPaginationTypeResponseHeader{/* values here */})
```

### RestCollectMethodOtherRestPaginationTypeResponseHeaderLink

```go
restCollectMethodOtherPaginationUnion := components.CreateRestCollectMethodOtherPaginationUnionResponseHeaderLink(components.RestCollectMethodOtherRestPaginationTypeResponseHeaderLink{/* values here */})
```

### RestCollectMethodOtherRestPaginationTypeRequestOffset

```go
restCollectMethodOtherPaginationUnion := components.CreateRestCollectMethodOtherPaginationUnionRequestOffset(components.RestCollectMethodOtherRestPaginationTypeRequestOffset{/* values here */})
```

### RestCollectMethodOtherRestPaginationTypeRequestPage

```go
restCollectMethodOtherPaginationUnion := components.CreateRestCollectMethodOtherPaginationUnionRequestPage(components.RestCollectMethodOtherRestPaginationTypeRequestPage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restCollectMethodOtherPaginationUnion.Type {
	case components.RestCollectMethodOtherPaginationUnionTypeNone:
		// restCollectMethodOtherPaginationUnion.RestCollectMethodOtherRestPaginationTypeNone is populated
	case components.RestCollectMethodOtherPaginationUnionTypeResponseBody:
		// restCollectMethodOtherPaginationUnion.RestCollectMethodOtherRestPaginationTypeResponseBody is populated
	case components.RestCollectMethodOtherPaginationUnionTypeResponseHeader:
		// restCollectMethodOtherPaginationUnion.RestCollectMethodOtherRestPaginationTypeResponseHeader is populated
	case components.RestCollectMethodOtherPaginationUnionTypeResponseHeaderLink:
		// restCollectMethodOtherPaginationUnion.RestCollectMethodOtherRestPaginationTypeResponseHeaderLink is populated
	case components.RestCollectMethodOtherPaginationUnionTypeRequestOffset:
		// restCollectMethodOtherPaginationUnion.RestCollectMethodOtherRestPaginationTypeRequestOffset is populated
	case components.RestCollectMethodOtherPaginationUnionTypeRequestPage:
		// restCollectMethodOtherPaginationUnion.RestCollectMethodOtherRestPaginationTypeRequestPage is populated
	default:
		// Unknown type - use restCollectMethodOtherPaginationUnion.GetUnknownRaw() for raw JSON
}
```
