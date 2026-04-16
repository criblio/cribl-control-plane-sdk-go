# RestAuthenticationHmacRestPaginationTypeResponseHeaderResponseAttributes

Names of attributes within the response that contain next-page information


## Supported Types

### 

```go
restAuthenticationHmacRestPaginationTypeResponseHeaderResponseAttributes := components.CreateRestAuthenticationHmacRestPaginationTypeResponseHeaderResponseAttributesArrayOfAny([]any{/* values here */})
```

### 

```go
restAuthenticationHmacRestPaginationTypeResponseHeaderResponseAttributes := components.CreateRestAuthenticationHmacRestPaginationTypeResponseHeaderResponseAttributesStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationHmacRestPaginationTypeResponseHeaderResponseAttributes.Type {
	case components.RestAuthenticationHmacRestPaginationTypeResponseHeaderResponseAttributesTypeArrayOfAny:
		// restAuthenticationHmacRestPaginationTypeResponseHeaderResponseAttributes.ArrayOfAny is populated
	case components.RestAuthenticationHmacRestPaginationTypeResponseHeaderResponseAttributesTypeStr:
		// restAuthenticationHmacRestPaginationTypeResponseHeaderResponseAttributes.Str is populated
}
```
