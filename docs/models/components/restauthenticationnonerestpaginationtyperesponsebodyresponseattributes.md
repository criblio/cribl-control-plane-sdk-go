# RestAuthenticationNoneRestPaginationTypeResponseBodyResponseAttributes

Names of attributes within the response that contain next-page information


## Supported Types

### 

```go
restAuthenticationNoneRestPaginationTypeResponseBodyResponseAttributes := components.CreateRestAuthenticationNoneRestPaginationTypeResponseBodyResponseAttributesArrayOfAny([]any{/* values here */})
```

### 

```go
restAuthenticationNoneRestPaginationTypeResponseBodyResponseAttributes := components.CreateRestAuthenticationNoneRestPaginationTypeResponseBodyResponseAttributesStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationNoneRestPaginationTypeResponseBodyResponseAttributes.Type {
	case components.RestAuthenticationNoneRestPaginationTypeResponseBodyResponseAttributesTypeArrayOfAny:
		// restAuthenticationNoneRestPaginationTypeResponseBodyResponseAttributes.ArrayOfAny is populated
	case components.RestAuthenticationNoneRestPaginationTypeResponseBodyResponseAttributesTypeStr:
		// restAuthenticationNoneRestPaginationTypeResponseBodyResponseAttributes.Str is populated
}
```
