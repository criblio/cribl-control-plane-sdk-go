# S3CollectorConf


## Supported Types

### S3AwsAuthenticationMethodAuto

```go
s3CollectorConf := components.CreateS3CollectorConfAuto(components.S3AwsAuthenticationMethodAuto{/* values here */})
```

### S3AwsAuthenticationMethodManual

```go
s3CollectorConf := components.CreateS3CollectorConfManual(components.S3AwsAuthenticationMethodManual{/* values here */})
```

### S3AwsAuthenticationMethodSecret

```go
s3CollectorConf := components.CreateS3CollectorConfSecret(components.S3AwsAuthenticationMethodSecret{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch s3CollectorConf.Type {
	case components.S3CollectorConfTypeAuto:
		// s3CollectorConf.S3AwsAuthenticationMethodAuto is populated
	case components.S3CollectorConfTypeManual:
		// s3CollectorConf.S3AwsAuthenticationMethodManual is populated
	case components.S3CollectorConfTypeSecret:
		// s3CollectorConf.S3AwsAuthenticationMethodSecret is populated
	default:
		// Unknown type - use s3CollectorConf.GetUnknownRaw() for raw JSON
}
```
