# RestAuthenticationGoogleOauthSecretRestPaginationTypeResponseHeaderResponseAttributes

Names of attributes within the response that contain next-page information


## Supported Types

### 

```go
restAuthenticationGoogleOauthSecretRestPaginationTypeResponseHeaderResponseAttributes := components.CreateRestAuthenticationGoogleOauthSecretRestPaginationTypeResponseHeaderResponseAttributesArrayOfAny([]any{/* values here */})
```

### 

```go
restAuthenticationGoogleOauthSecretRestPaginationTypeResponseHeaderResponseAttributes := components.CreateRestAuthenticationGoogleOauthSecretRestPaginationTypeResponseHeaderResponseAttributesStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationGoogleOauthSecretRestPaginationTypeResponseHeaderResponseAttributes.Type {
	case components.RestAuthenticationGoogleOauthSecretRestPaginationTypeResponseHeaderResponseAttributesTypeArrayOfAny:
		// restAuthenticationGoogleOauthSecretRestPaginationTypeResponseHeaderResponseAttributes.ArrayOfAny is populated
	case components.RestAuthenticationGoogleOauthSecretRestPaginationTypeResponseHeaderResponseAttributesTypeStr:
		// restAuthenticationGoogleOauthSecretRestPaginationTypeResponseHeaderResponseAttributes.Str is populated
}
```
