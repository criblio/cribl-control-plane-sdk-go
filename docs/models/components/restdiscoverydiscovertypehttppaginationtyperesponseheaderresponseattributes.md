# RestDiscoveryDiscoverTypeHTTPPaginationTypeResponseHeaderResponseAttributes

Names of attributes within the response that contain next-page information


## Supported Types

### 

```go
restDiscoveryDiscoverTypeHTTPPaginationTypeResponseHeaderResponseAttributes := components.CreateRestDiscoveryDiscoverTypeHTTPPaginationTypeResponseHeaderResponseAttributesArrayOfAny([]any{/* values here */})
```

### 

```go
restDiscoveryDiscoverTypeHTTPPaginationTypeResponseHeaderResponseAttributes := components.CreateRestDiscoveryDiscoverTypeHTTPPaginationTypeResponseHeaderResponseAttributesStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restDiscoveryDiscoverTypeHTTPPaginationTypeResponseHeaderResponseAttributes.Type {
	case components.RestDiscoveryDiscoverTypeHTTPPaginationTypeResponseHeaderResponseAttributesTypeArrayOfAny:
		// restDiscoveryDiscoverTypeHTTPPaginationTypeResponseHeaderResponseAttributes.ArrayOfAny is populated
	case components.RestDiscoveryDiscoverTypeHTTPPaginationTypeResponseHeaderResponseAttributesTypeStr:
		// restDiscoveryDiscoverTypeHTTPPaginationTypeResponseHeaderResponseAttributes.Str is populated
}
```
