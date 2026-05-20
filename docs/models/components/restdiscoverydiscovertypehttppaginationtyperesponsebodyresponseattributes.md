# RestDiscoveryDiscoverTypeHTTPPaginationTypeResponseBodyResponseAttributes

Names of attributes within the response that contain next-page information


## Supported Types

### 

```go
restDiscoveryDiscoverTypeHTTPPaginationTypeResponseBodyResponseAttributes := components.CreateRestDiscoveryDiscoverTypeHTTPPaginationTypeResponseBodyResponseAttributesArrayOfAny([]any{/* values here */})
```

### 

```go
restDiscoveryDiscoverTypeHTTPPaginationTypeResponseBodyResponseAttributes := components.CreateRestDiscoveryDiscoverTypeHTTPPaginationTypeResponseBodyResponseAttributesStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restDiscoveryDiscoverTypeHTTPPaginationTypeResponseBodyResponseAttributes.Type {
	case components.RestDiscoveryDiscoverTypeHTTPPaginationTypeResponseBodyResponseAttributesTypeArrayOfAny:
		// restDiscoveryDiscoverTypeHTTPPaginationTypeResponseBodyResponseAttributes.ArrayOfAny is populated
	case components.RestDiscoveryDiscoverTypeHTTPPaginationTypeResponseBodyResponseAttributesTypeStr:
		// restDiscoveryDiscoverTypeHTTPPaginationTypeResponseBodyResponseAttributes.Str is populated
}
```
