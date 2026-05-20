# RestAuthenticationBasicRestPaginationTypeResponseHeaderResponseAttributes

Names of attributes within the response that contain next-page information


## Supported Types

### 

```go
restAuthenticationBasicRestPaginationTypeResponseHeaderResponseAttributes := components.CreateRestAuthenticationBasicRestPaginationTypeResponseHeaderResponseAttributesArrayOfAny([]any{/* values here */})
```

### 

```go
restAuthenticationBasicRestPaginationTypeResponseHeaderResponseAttributes := components.CreateRestAuthenticationBasicRestPaginationTypeResponseHeaderResponseAttributesStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationBasicRestPaginationTypeResponseHeaderResponseAttributes.Type {
	case components.RestAuthenticationBasicRestPaginationTypeResponseHeaderResponseAttributesTypeArrayOfAny:
		// restAuthenticationBasicRestPaginationTypeResponseHeaderResponseAttributes.ArrayOfAny is populated
	case components.RestAuthenticationBasicRestPaginationTypeResponseHeaderResponseAttributesTypeStr:
		// restAuthenticationBasicRestPaginationTypeResponseHeaderResponseAttributes.Str is populated
}
```
