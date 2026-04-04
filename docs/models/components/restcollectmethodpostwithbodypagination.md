# RestCollectMethodPostWithBodyPagination


## Supported Types

### RestCollectMethodPostWithBodyRestPaginationTypeNone

```go
restCollectMethodPostWithBodyPagination := components.CreateRestCollectMethodPostWithBodyPaginationNone(components.RestCollectMethodPostWithBodyRestPaginationTypeNone{/* values here */})
```

### RestCollectMethodPostWithBodyRestPaginationTypeResponseBody

```go
restCollectMethodPostWithBodyPagination := components.CreateRestCollectMethodPostWithBodyPaginationResponseBody(components.RestCollectMethodPostWithBodyRestPaginationTypeResponseBody{/* values here */})
```

### RestCollectMethodPostWithBodyRestPaginationTypeResponseHeader

```go
restCollectMethodPostWithBodyPagination := components.CreateRestCollectMethodPostWithBodyPaginationResponseHeader(components.RestCollectMethodPostWithBodyRestPaginationTypeResponseHeader{/* values here */})
```

### RestCollectMethodPostWithBodyRestPaginationTypeResponseHeaderLink

```go
restCollectMethodPostWithBodyPagination := components.CreateRestCollectMethodPostWithBodyPaginationResponseHeaderLink(components.RestCollectMethodPostWithBodyRestPaginationTypeResponseHeaderLink{/* values here */})
```

### RestCollectMethodPostWithBodyRestPaginationTypeRequestOffset

```go
restCollectMethodPostWithBodyPagination := components.CreateRestCollectMethodPostWithBodyPaginationRequestOffset(components.RestCollectMethodPostWithBodyRestPaginationTypeRequestOffset{/* values here */})
```

### RestCollectMethodPostWithBodyRestPaginationTypeRequestPage

```go
restCollectMethodPostWithBodyPagination := components.CreateRestCollectMethodPostWithBodyPaginationRequestPage(components.RestCollectMethodPostWithBodyRestPaginationTypeRequestPage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restCollectMethodPostWithBodyPagination.Type {
	case components.RestCollectMethodPostWithBodyPaginationTypeNone:
		// restCollectMethodPostWithBodyPagination.RestCollectMethodPostWithBodyRestPaginationTypeNone is populated
	case components.RestCollectMethodPostWithBodyPaginationTypeResponseBody:
		// restCollectMethodPostWithBodyPagination.RestCollectMethodPostWithBodyRestPaginationTypeResponseBody is populated
	case components.RestCollectMethodPostWithBodyPaginationTypeResponseHeader:
		// restCollectMethodPostWithBodyPagination.RestCollectMethodPostWithBodyRestPaginationTypeResponseHeader is populated
	case components.RestCollectMethodPostWithBodyPaginationTypeResponseHeaderLink:
		// restCollectMethodPostWithBodyPagination.RestCollectMethodPostWithBodyRestPaginationTypeResponseHeaderLink is populated
	case components.RestCollectMethodPostWithBodyPaginationTypeRequestOffset:
		// restCollectMethodPostWithBodyPagination.RestCollectMethodPostWithBodyRestPaginationTypeRequestOffset is populated
	case components.RestCollectMethodPostWithBodyPaginationTypeRequestPage:
		// restCollectMethodPostWithBodyPagination.RestCollectMethodPostWithBodyRestPaginationTypeRequestPage is populated
	default:
		// Unknown type - use restCollectMethodPostWithBodyPagination.GetUnknownRaw() for raw JSON
}
```
