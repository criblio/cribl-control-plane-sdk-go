# GoogleCloudStorageCollectorConf


## Supported Types

### GoogleCloudStorageAuthTypeAuto

```go
googleCloudStorageCollectorConf := components.CreateGoogleCloudStorageCollectorConfAuto(components.GoogleCloudStorageAuthTypeAuto{/* values here */})
```

### GoogleCloudStorageAuthTypeManual

```go
googleCloudStorageCollectorConf := components.CreateGoogleCloudStorageCollectorConfManual(components.GoogleCloudStorageAuthTypeManual{/* values here */})
```

### GoogleCloudStorageAuthTypeSecret

```go
googleCloudStorageCollectorConf := components.CreateGoogleCloudStorageCollectorConfSecret(components.GoogleCloudStorageAuthTypeSecret{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch googleCloudStorageCollectorConf.Type {
	case components.GoogleCloudStorageCollectorConfTypeAuto:
		// googleCloudStorageCollectorConf.GoogleCloudStorageAuthTypeAuto is populated
	case components.GoogleCloudStorageCollectorConfTypeManual:
		// googleCloudStorageCollectorConf.GoogleCloudStorageAuthTypeManual is populated
	case components.GoogleCloudStorageCollectorConfTypeSecret:
		// googleCloudStorageCollectorConf.GoogleCloudStorageAuthTypeSecret is populated
	default:
		// Unknown type - use googleCloudStorageCollectorConf.GetUnknownRaw() for raw JSON
}
```
