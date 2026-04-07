# RestAuthenticationNoneRestPaginationTypeResponseHeaderResponseAttributes

Names of attributes within the response that contain next-page information


## Supported Types

### 

```go
restAuthenticationNoneRestPaginationTypeResponseHeaderResponseAttributes := components.CreateRestAuthenticationNoneRestPaginationTypeResponseHeaderResponseAttributesArrayOfAny([]any{/* values here */})
```

### 

```go
restAuthenticationNoneRestPaginationTypeResponseHeaderResponseAttributes := components.CreateRestAuthenticationNoneRestPaginationTypeResponseHeaderResponseAttributesStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationNoneRestPaginationTypeResponseHeaderResponseAttributes.Type {
	case components.RestAuthenticationNoneRestPaginationTypeResponseHeaderResponseAttributesTypeArrayOfAny:
		// restAuthenticationNoneRestPaginationTypeResponseHeaderResponseAttributes.ArrayOfAny is populated
	case components.RestAuthenticationNoneRestPaginationTypeResponseHeaderResponseAttributesTypeStr:
		// restAuthenticationNoneRestPaginationTypeResponseHeaderResponseAttributes.Str is populated
}
```
