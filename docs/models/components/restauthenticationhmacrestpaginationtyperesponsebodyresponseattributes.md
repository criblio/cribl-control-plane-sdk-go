# RestAuthenticationHmacRestPaginationTypeResponseBodyResponseAttributes

Names of attributes within the response that contain next-page information


## Supported Types

### 

```go
restAuthenticationHmacRestPaginationTypeResponseBodyResponseAttributes := components.CreateRestAuthenticationHmacRestPaginationTypeResponseBodyResponseAttributesArrayOfAny([]any{/* values here */})
```

### 

```go
restAuthenticationHmacRestPaginationTypeResponseBodyResponseAttributes := components.CreateRestAuthenticationHmacRestPaginationTypeResponseBodyResponseAttributesStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationHmacRestPaginationTypeResponseBodyResponseAttributes.Type {
	case components.RestAuthenticationHmacRestPaginationTypeResponseBodyResponseAttributesTypeArrayOfAny:
		// restAuthenticationHmacRestPaginationTypeResponseBodyResponseAttributes.ArrayOfAny is populated
	case components.RestAuthenticationHmacRestPaginationTypeResponseBodyResponseAttributesTypeStr:
		// restAuthenticationHmacRestPaginationTypeResponseBodyResponseAttributes.Str is populated
}
```
