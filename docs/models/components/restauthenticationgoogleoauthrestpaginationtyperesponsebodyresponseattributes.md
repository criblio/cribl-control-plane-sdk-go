# RestAuthenticationGoogleOauthRestPaginationTypeResponseBodyResponseAttributes

Names of attributes within the response that contain next-page information


## Supported Types

### 

```go
restAuthenticationGoogleOauthRestPaginationTypeResponseBodyResponseAttributes := components.CreateRestAuthenticationGoogleOauthRestPaginationTypeResponseBodyResponseAttributesArrayOfAny([]any{/* values here */})
```

### 

```go
restAuthenticationGoogleOauthRestPaginationTypeResponseBodyResponseAttributes := components.CreateRestAuthenticationGoogleOauthRestPaginationTypeResponseBodyResponseAttributesStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationGoogleOauthRestPaginationTypeResponseBodyResponseAttributes.Type {
	case components.RestAuthenticationGoogleOauthRestPaginationTypeResponseBodyResponseAttributesTypeArrayOfAny:
		// restAuthenticationGoogleOauthRestPaginationTypeResponseBodyResponseAttributes.ArrayOfAny is populated
	case components.RestAuthenticationGoogleOauthRestPaginationTypeResponseBodyResponseAttributesTypeStr:
		// restAuthenticationGoogleOauthRestPaginationTypeResponseBodyResponseAttributes.Str is populated
}
```
