# RestCollectMethodOtherRestPaginationTypeResponseHeaderResponseAttributes

Names of attributes within the response that contain next-page information


## Supported Types

### 

```go
restCollectMethodOtherRestPaginationTypeResponseHeaderResponseAttributes := components.CreateRestCollectMethodOtherRestPaginationTypeResponseHeaderResponseAttributesArrayOfAny([]any{/* values here */})
```

### 

```go
restCollectMethodOtherRestPaginationTypeResponseHeaderResponseAttributes := components.CreateRestCollectMethodOtherRestPaginationTypeResponseHeaderResponseAttributesStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restCollectMethodOtherRestPaginationTypeResponseHeaderResponseAttributes.Type {
	case components.RestCollectMethodOtherRestPaginationTypeResponseHeaderResponseAttributesTypeArrayOfAny:
		// restCollectMethodOtherRestPaginationTypeResponseHeaderResponseAttributes.ArrayOfAny is populated
	case components.RestCollectMethodOtherRestPaginationTypeResponseHeaderResponseAttributesTypeStr:
		// restCollectMethodOtherRestPaginationTypeResponseHeaderResponseAttributes.Str is populated
}
```
