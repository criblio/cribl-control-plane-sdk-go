# RestCollectMethodPostRestPaginationTypeResponseHeaderResponseAttributes

Names of attributes within the response that contain next-page information


## Supported Types

### 

```go
restCollectMethodPostRestPaginationTypeResponseHeaderResponseAttributes := components.CreateRestCollectMethodPostRestPaginationTypeResponseHeaderResponseAttributesArrayOfAny([]any{/* values here */})
```

### 

```go
restCollectMethodPostRestPaginationTypeResponseHeaderResponseAttributes := components.CreateRestCollectMethodPostRestPaginationTypeResponseHeaderResponseAttributesStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restCollectMethodPostRestPaginationTypeResponseHeaderResponseAttributes.Type {
	case components.RestCollectMethodPostRestPaginationTypeResponseHeaderResponseAttributesTypeArrayOfAny:
		// restCollectMethodPostRestPaginationTypeResponseHeaderResponseAttributes.ArrayOfAny is populated
	case components.RestCollectMethodPostRestPaginationTypeResponseHeaderResponseAttributesTypeStr:
		// restCollectMethodPostRestPaginationTypeResponseHeaderResponseAttributes.Str is populated
}
```
