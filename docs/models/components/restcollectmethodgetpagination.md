# RestCollectMethodGetPagination


## Supported Types

### RestCollectMethodGetRestPaginationTypeNone

```go
restCollectMethodGetPagination := components.CreateRestCollectMethodGetPaginationNone(components.RestCollectMethodGetRestPaginationTypeNone{/* values here */})
```

### RestCollectMethodGetRestPaginationTypeResponseBody

```go
restCollectMethodGetPagination := components.CreateRestCollectMethodGetPaginationResponseBody(components.RestCollectMethodGetRestPaginationTypeResponseBody{/* values here */})
```

### RestCollectMethodGetRestPaginationTypeResponseHeader

```go
restCollectMethodGetPagination := components.CreateRestCollectMethodGetPaginationResponseHeader(components.RestCollectMethodGetRestPaginationTypeResponseHeader{/* values here */})
```

### RestCollectMethodGetRestPaginationTypeResponseHeaderLink

```go
restCollectMethodGetPagination := components.CreateRestCollectMethodGetPaginationResponseHeaderLink(components.RestCollectMethodGetRestPaginationTypeResponseHeaderLink{/* values here */})
```

### RestCollectMethodGetRestPaginationTypeRequestOffset

```go
restCollectMethodGetPagination := components.CreateRestCollectMethodGetPaginationRequestOffset(components.RestCollectMethodGetRestPaginationTypeRequestOffset{/* values here */})
```

### RestCollectMethodGetRestPaginationTypeRequestPage

```go
restCollectMethodGetPagination := components.CreateRestCollectMethodGetPaginationRequestPage(components.RestCollectMethodGetRestPaginationTypeRequestPage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restCollectMethodGetPagination.Type {
	case components.RestCollectMethodGetPaginationTypeNone:
		// restCollectMethodGetPagination.RestCollectMethodGetRestPaginationTypeNone is populated
	case components.RestCollectMethodGetPaginationTypeResponseBody:
		// restCollectMethodGetPagination.RestCollectMethodGetRestPaginationTypeResponseBody is populated
	case components.RestCollectMethodGetPaginationTypeResponseHeader:
		// restCollectMethodGetPagination.RestCollectMethodGetRestPaginationTypeResponseHeader is populated
	case components.RestCollectMethodGetPaginationTypeResponseHeaderLink:
		// restCollectMethodGetPagination.RestCollectMethodGetRestPaginationTypeResponseHeaderLink is populated
	case components.RestCollectMethodGetPaginationTypeRequestOffset:
		// restCollectMethodGetPagination.RestCollectMethodGetRestPaginationTypeRequestOffset is populated
	case components.RestCollectMethodGetPaginationTypeRequestPage:
		// restCollectMethodGetPagination.RestCollectMethodGetRestPaginationTypeRequestPage is populated
	default:
		// Unknown type - use restCollectMethodGetPagination.GetUnknownRaw() for raw JSON
}
```
