# PaginationTypeRestDiscoveryDiscoverTypeHTTP


## Supported Types

### RestDiscoveryDiscoverTypeHTTPPaginationTypeNone

```go
paginationTypeRestDiscoveryDiscoverTypeHTTP := components.CreatePaginationTypeRestDiscoveryDiscoverTypeHTTPNone(components.RestDiscoveryDiscoverTypeHTTPPaginationTypeNone{/* values here */})
```

### RestDiscoveryDiscoverTypeHTTPPaginationTypeResponseBody

```go
paginationTypeRestDiscoveryDiscoverTypeHTTP := components.CreatePaginationTypeRestDiscoveryDiscoverTypeHTTPResponseBody(components.RestDiscoveryDiscoverTypeHTTPPaginationTypeResponseBody{/* values here */})
```

### RestDiscoveryDiscoverTypeHTTPPaginationTypeResponseHeader

```go
paginationTypeRestDiscoveryDiscoverTypeHTTP := components.CreatePaginationTypeRestDiscoveryDiscoverTypeHTTPResponseHeader(components.RestDiscoveryDiscoverTypeHTTPPaginationTypeResponseHeader{/* values here */})
```

### RestDiscoveryDiscoverTypeHTTPPaginationTypeResponseHeaderLink

```go
paginationTypeRestDiscoveryDiscoverTypeHTTP := components.CreatePaginationTypeRestDiscoveryDiscoverTypeHTTPResponseHeaderLink(components.RestDiscoveryDiscoverTypeHTTPPaginationTypeResponseHeaderLink{/* values here */})
```

### RestDiscoveryDiscoverTypeHTTPPaginationTypeRequestOffset

```go
paginationTypeRestDiscoveryDiscoverTypeHTTP := components.CreatePaginationTypeRestDiscoveryDiscoverTypeHTTPRequestOffset(components.RestDiscoveryDiscoverTypeHTTPPaginationTypeRequestOffset{/* values here */})
```

### RestDiscoveryDiscoverTypeHTTPPaginationTypeRequestPage

```go
paginationTypeRestDiscoveryDiscoverTypeHTTP := components.CreatePaginationTypeRestDiscoveryDiscoverTypeHTTPRequestPage(components.RestDiscoveryDiscoverTypeHTTPPaginationTypeRequestPage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch paginationTypeRestDiscoveryDiscoverTypeHTTP.Type {
	case components.PaginationTypeRestDiscoveryDiscoverTypeHTTPTypeNone:
		// paginationTypeRestDiscoveryDiscoverTypeHTTP.RestDiscoveryDiscoverTypeHTTPPaginationTypeNone is populated
	case components.PaginationTypeRestDiscoveryDiscoverTypeHTTPTypeResponseBody:
		// paginationTypeRestDiscoveryDiscoverTypeHTTP.RestDiscoveryDiscoverTypeHTTPPaginationTypeResponseBody is populated
	case components.PaginationTypeRestDiscoveryDiscoverTypeHTTPTypeResponseHeader:
		// paginationTypeRestDiscoveryDiscoverTypeHTTP.RestDiscoveryDiscoverTypeHTTPPaginationTypeResponseHeader is populated
	case components.PaginationTypeRestDiscoveryDiscoverTypeHTTPTypeResponseHeaderLink:
		// paginationTypeRestDiscoveryDiscoverTypeHTTP.RestDiscoveryDiscoverTypeHTTPPaginationTypeResponseHeaderLink is populated
	case components.PaginationTypeRestDiscoveryDiscoverTypeHTTPTypeRequestOffset:
		// paginationTypeRestDiscoveryDiscoverTypeHTTP.RestDiscoveryDiscoverTypeHTTPPaginationTypeRequestOffset is populated
	case components.PaginationTypeRestDiscoveryDiscoverTypeHTTPTypeRequestPage:
		// paginationTypeRestDiscoveryDiscoverTypeHTTP.RestDiscoveryDiscoverTypeHTTPPaginationTypeRequestPage is populated
	default:
		// Unknown type - use paginationTypeRestDiscoveryDiscoverTypeHTTP.GetUnknownRaw() for raw JSON
}
```
