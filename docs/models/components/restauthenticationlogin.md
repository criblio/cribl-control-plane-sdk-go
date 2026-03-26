# RestAuthenticationLogin


## Supported Types

### RestAuthenticationLoginGetAuthTokenFromHeaderFalse

```go
restAuthenticationLogin := components.CreateRestAuthenticationLoginRestAuthenticationLoginGetAuthTokenFromHeaderFalse(components.RestAuthenticationLoginGetAuthTokenFromHeaderFalse{/* values here */})
```

### RestAuthenticationLoginGetAuthTokenFromHeaderTrue

```go
restAuthenticationLogin := components.CreateRestAuthenticationLoginRestAuthenticationLoginGetAuthTokenFromHeaderTrue(components.RestAuthenticationLoginGetAuthTokenFromHeaderTrue{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationLogin.Type {
	case components.RestAuthenticationLoginTypeRestAuthenticationLoginGetAuthTokenFromHeaderFalse:
		// restAuthenticationLogin.RestAuthenticationLoginGetAuthTokenFromHeaderFalse is populated
	case components.RestAuthenticationLoginTypeRestAuthenticationLoginGetAuthTokenFromHeaderTrue:
		// restAuthenticationLogin.RestAuthenticationLoginGetAuthTokenFromHeaderTrue is populated
}
```
