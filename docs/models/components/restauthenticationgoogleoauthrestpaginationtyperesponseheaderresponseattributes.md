# RestAuthenticationGoogleOauthRestPaginationTypeResponseHeaderResponseAttributes

Names of attributes within the response that contain next-page information


## Supported Types

### 

```go
restAuthenticationGoogleOauthRestPaginationTypeResponseHeaderResponseAttributes := components.CreateRestAuthenticationGoogleOauthRestPaginationTypeResponseHeaderResponseAttributesArrayOfAny([]any{/* values here */})
```

### 

```go
restAuthenticationGoogleOauthRestPaginationTypeResponseHeaderResponseAttributes := components.CreateRestAuthenticationGoogleOauthRestPaginationTypeResponseHeaderResponseAttributesStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationGoogleOauthRestPaginationTypeResponseHeaderResponseAttributes.Type {
	case components.RestAuthenticationGoogleOauthRestPaginationTypeResponseHeaderResponseAttributesTypeArrayOfAny:
		// restAuthenticationGoogleOauthRestPaginationTypeResponseHeaderResponseAttributes.ArrayOfAny is populated
	case components.RestAuthenticationGoogleOauthRestPaginationTypeResponseHeaderResponseAttributesTypeStr:
		// restAuthenticationGoogleOauthRestPaginationTypeResponseHeaderResponseAttributes.Str is populated
}
```
