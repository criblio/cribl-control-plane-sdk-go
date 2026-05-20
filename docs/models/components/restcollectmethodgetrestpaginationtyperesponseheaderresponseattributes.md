# RestCollectMethodGetRestPaginationTypeResponseHeaderResponseAttributes

Names of attributes within the response that contain next-page information


## Supported Types

### 

```go
restCollectMethodGetRestPaginationTypeResponseHeaderResponseAttributes := components.CreateRestCollectMethodGetRestPaginationTypeResponseHeaderResponseAttributesArrayOfAny([]any{/* values here */})
```

### 

```go
restCollectMethodGetRestPaginationTypeResponseHeaderResponseAttributes := components.CreateRestCollectMethodGetRestPaginationTypeResponseHeaderResponseAttributesStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restCollectMethodGetRestPaginationTypeResponseHeaderResponseAttributes.Type {
	case components.RestCollectMethodGetRestPaginationTypeResponseHeaderResponseAttributesTypeArrayOfAny:
		// restCollectMethodGetRestPaginationTypeResponseHeaderResponseAttributes.ArrayOfAny is populated
	case components.RestCollectMethodGetRestPaginationTypeResponseHeaderResponseAttributesTypeStr:
		// restCollectMethodGetRestPaginationTypeResponseHeaderResponseAttributes.Str is populated
}
```
