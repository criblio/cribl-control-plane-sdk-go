# RestAuthenticationGoogleOauthSecretRestPaginationTypeResponseBodyResponseAttributes

Names of attributes within the response that contain next-page information


## Supported Types

### 

```go
restAuthenticationGoogleOauthSecretRestPaginationTypeResponseBodyResponseAttributes := components.CreateRestAuthenticationGoogleOauthSecretRestPaginationTypeResponseBodyResponseAttributesArrayOfAny([]any{/* values here */})
```

### 

```go
restAuthenticationGoogleOauthSecretRestPaginationTypeResponseBodyResponseAttributes := components.CreateRestAuthenticationGoogleOauthSecretRestPaginationTypeResponseBodyResponseAttributesStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationGoogleOauthSecretRestPaginationTypeResponseBodyResponseAttributes.Type {
	case components.RestAuthenticationGoogleOauthSecretRestPaginationTypeResponseBodyResponseAttributesTypeArrayOfAny:
		// restAuthenticationGoogleOauthSecretRestPaginationTypeResponseBodyResponseAttributes.ArrayOfAny is populated
	case components.RestAuthenticationGoogleOauthSecretRestPaginationTypeResponseBodyResponseAttributesTypeStr:
		// restAuthenticationGoogleOauthSecretRestPaginationTypeResponseBodyResponseAttributes.Str is populated
}
```
