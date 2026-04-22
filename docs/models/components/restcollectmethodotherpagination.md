# RestCollectMethodOtherPagination


## Supported Types

### RestCollectMethodOtherRestPaginationTypeNone

```go
restCollectMethodOtherPagination := components.CreateRestCollectMethodOtherPaginationNone(components.RestCollectMethodOtherRestPaginationTypeNone{/* values here */})
```

### RestCollectMethodOtherRestPaginationTypeResponseBody

```go
restCollectMethodOtherPagination := components.CreateRestCollectMethodOtherPaginationResponseBody(components.RestCollectMethodOtherRestPaginationTypeResponseBody{/* values here */})
```

### RestCollectMethodOtherRestPaginationTypeResponseHeader

```go
restCollectMethodOtherPagination := components.CreateRestCollectMethodOtherPaginationResponseHeader(components.RestCollectMethodOtherRestPaginationTypeResponseHeader{/* values here */})
```

### RestCollectMethodOtherRestPaginationTypeResponseHeaderLink

```go
restCollectMethodOtherPagination := components.CreateRestCollectMethodOtherPaginationResponseHeaderLink(components.RestCollectMethodOtherRestPaginationTypeResponseHeaderLink{/* values here */})
```

### RestCollectMethodOtherRestPaginationTypeRequestOffset

```go
restCollectMethodOtherPagination := components.CreateRestCollectMethodOtherPaginationRequestOffset(components.RestCollectMethodOtherRestPaginationTypeRequestOffset{/* values here */})
```

### RestCollectMethodOtherRestPaginationTypeRequestPage

```go
restCollectMethodOtherPagination := components.CreateRestCollectMethodOtherPaginationRequestPage(components.RestCollectMethodOtherRestPaginationTypeRequestPage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restCollectMethodOtherPagination.Type {
	case components.RestCollectMethodOtherPaginationTypeNone:
		// restCollectMethodOtherPagination.RestCollectMethodOtherRestPaginationTypeNone is populated
	case components.RestCollectMethodOtherPaginationTypeResponseBody:
		// restCollectMethodOtherPagination.RestCollectMethodOtherRestPaginationTypeResponseBody is populated
	case components.RestCollectMethodOtherPaginationTypeResponseHeader:
		// restCollectMethodOtherPagination.RestCollectMethodOtherRestPaginationTypeResponseHeader is populated
	case components.RestCollectMethodOtherPaginationTypeResponseHeaderLink:
		// restCollectMethodOtherPagination.RestCollectMethodOtherRestPaginationTypeResponseHeaderLink is populated
	case components.RestCollectMethodOtherPaginationTypeRequestOffset:
		// restCollectMethodOtherPagination.RestCollectMethodOtherRestPaginationTypeRequestOffset is populated
	case components.RestCollectMethodOtherPaginationTypeRequestPage:
		// restCollectMethodOtherPagination.RestCollectMethodOtherRestPaginationTypeRequestPage is populated
	default:
		// Unknown type - use restCollectMethodOtherPagination.GetUnknownRaw() for raw JSON
}
```
