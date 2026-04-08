# RestCollectMethodPostWithBodyRestPaginationTypeResponseHeaderResponseAttributes

Names of attributes within the response that contain next-page information


## Supported Types

### 

```go
restCollectMethodPostWithBodyRestPaginationTypeResponseHeaderResponseAttributes := components.CreateRestCollectMethodPostWithBodyRestPaginationTypeResponseHeaderResponseAttributesArrayOfAny([]any{/* values here */})
```

### 

```go
restCollectMethodPostWithBodyRestPaginationTypeResponseHeaderResponseAttributes := components.CreateRestCollectMethodPostWithBodyRestPaginationTypeResponseHeaderResponseAttributesStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restCollectMethodPostWithBodyRestPaginationTypeResponseHeaderResponseAttributes.Type {
	case components.RestCollectMethodPostWithBodyRestPaginationTypeResponseHeaderResponseAttributesTypeArrayOfAny:
		// restCollectMethodPostWithBodyRestPaginationTypeResponseHeaderResponseAttributes.ArrayOfAny is populated
	case components.RestCollectMethodPostWithBodyRestPaginationTypeResponseHeaderResponseAttributesTypeStr:
		// restCollectMethodPostWithBodyRestPaginationTypeResponseHeaderResponseAttributes.Str is populated
}
```
