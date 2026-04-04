# RestCollectMethodPostPagination


## Supported Types

### RestCollectMethodPostRestPaginationTypeNone

```go
restCollectMethodPostPagination := components.CreateRestCollectMethodPostPaginationNone(components.RestCollectMethodPostRestPaginationTypeNone{/* values here */})
```

### RestCollectMethodPostRestPaginationTypeResponseBody

```go
restCollectMethodPostPagination := components.CreateRestCollectMethodPostPaginationResponseBody(components.RestCollectMethodPostRestPaginationTypeResponseBody{/* values here */})
```

### RestCollectMethodPostRestPaginationTypeResponseHeader

```go
restCollectMethodPostPagination := components.CreateRestCollectMethodPostPaginationResponseHeader(components.RestCollectMethodPostRestPaginationTypeResponseHeader{/* values here */})
```

### RestCollectMethodPostRestPaginationTypeResponseHeaderLink

```go
restCollectMethodPostPagination := components.CreateRestCollectMethodPostPaginationResponseHeaderLink(components.RestCollectMethodPostRestPaginationTypeResponseHeaderLink{/* values here */})
```

### RestCollectMethodPostRestPaginationTypeRequestOffset

```go
restCollectMethodPostPagination := components.CreateRestCollectMethodPostPaginationRequestOffset(components.RestCollectMethodPostRestPaginationTypeRequestOffset{/* values here */})
```

### RestCollectMethodPostRestPaginationTypeRequestPage

```go
restCollectMethodPostPagination := components.CreateRestCollectMethodPostPaginationRequestPage(components.RestCollectMethodPostRestPaginationTypeRequestPage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restCollectMethodPostPagination.Type {
	case components.RestCollectMethodPostPaginationTypeNone:
		// restCollectMethodPostPagination.RestCollectMethodPostRestPaginationTypeNone is populated
	case components.RestCollectMethodPostPaginationTypeResponseBody:
		// restCollectMethodPostPagination.RestCollectMethodPostRestPaginationTypeResponseBody is populated
	case components.RestCollectMethodPostPaginationTypeResponseHeader:
		// restCollectMethodPostPagination.RestCollectMethodPostRestPaginationTypeResponseHeader is populated
	case components.RestCollectMethodPostPaginationTypeResponseHeaderLink:
		// restCollectMethodPostPagination.RestCollectMethodPostRestPaginationTypeResponseHeaderLink is populated
	case components.RestCollectMethodPostPaginationTypeRequestOffset:
		// restCollectMethodPostPagination.RestCollectMethodPostRestPaginationTypeRequestOffset is populated
	case components.RestCollectMethodPostPaginationTypeRequestPage:
		// restCollectMethodPostPagination.RestCollectMethodPostRestPaginationTypeRequestPage is populated
	default:
		// Unknown type - use restCollectMethodPostPagination.GetUnknownRaw() for raw JSON
}
```
