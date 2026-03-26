# RestAuthenticationLoginSecret


## Supported Types

### RestAuthenticationLoginSecretGetAuthTokenFromHeaderFalse

```go
restAuthenticationLoginSecret := components.CreateRestAuthenticationLoginSecretRestAuthenticationLoginSecretGetAuthTokenFromHeaderFalse(components.RestAuthenticationLoginSecretGetAuthTokenFromHeaderFalse{/* values here */})
```

### RestAuthenticationLoginSecretGetAuthTokenFromHeaderTrue

```go
restAuthenticationLoginSecret := components.CreateRestAuthenticationLoginSecretRestAuthenticationLoginSecretGetAuthTokenFromHeaderTrue(components.RestAuthenticationLoginSecretGetAuthTokenFromHeaderTrue{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationLoginSecret.Type {
	case components.RestAuthenticationLoginSecretTypeRestAuthenticationLoginSecretGetAuthTokenFromHeaderFalse:
		// restAuthenticationLoginSecret.RestAuthenticationLoginSecretGetAuthTokenFromHeaderFalse is populated
	case components.RestAuthenticationLoginSecretTypeRestAuthenticationLoginSecretGetAuthTokenFromHeaderTrue:
		// restAuthenticationLoginSecret.RestAuthenticationLoginSecretGetAuthTokenFromHeaderTrue is populated
}
```
