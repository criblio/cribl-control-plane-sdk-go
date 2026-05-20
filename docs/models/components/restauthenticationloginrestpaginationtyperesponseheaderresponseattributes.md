# RestAuthenticationLoginRestPaginationTypeResponseHeaderResponseAttributes

Names of attributes within the response that contain next-page information


## Supported Types

### 

```go
restAuthenticationLoginRestPaginationTypeResponseHeaderResponseAttributes := components.CreateRestAuthenticationLoginRestPaginationTypeResponseHeaderResponseAttributesArrayOfAny([]any{/* values here */})
```

### 

```go
restAuthenticationLoginRestPaginationTypeResponseHeaderResponseAttributes := components.CreateRestAuthenticationLoginRestPaginationTypeResponseHeaderResponseAttributesStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationLoginRestPaginationTypeResponseHeaderResponseAttributes.Type {
	case components.RestAuthenticationLoginRestPaginationTypeResponseHeaderResponseAttributesTypeArrayOfAny:
		// restAuthenticationLoginRestPaginationTypeResponseHeaderResponseAttributes.ArrayOfAny is populated
	case components.RestAuthenticationLoginRestPaginationTypeResponseHeaderResponseAttributesTypeStr:
		// restAuthenticationLoginRestPaginationTypeResponseHeaderResponseAttributes.Str is populated
}
```
