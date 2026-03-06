# AzureBlobCollectorConf


## Supported Types

### AzureBlobAuthTypeManual

```go
azureBlobCollectorConf := components.CreateAzureBlobCollectorConfManual(components.AzureBlobAuthTypeManual{/* values here */})
```

### AzureBlobAuthTypeSecret

```go
azureBlobCollectorConf := components.CreateAzureBlobCollectorConfSecret(components.AzureBlobAuthTypeSecret{/* values here */})
```

### AzureBlobAuthTypeClientSecret

```go
azureBlobCollectorConf := components.CreateAzureBlobCollectorConfClientSecret(components.AzureBlobAuthTypeClientSecret{/* values here */})
```

### AzureBlobAuthTypeClientCert

```go
azureBlobCollectorConf := components.CreateAzureBlobCollectorConfClientCert(components.AzureBlobAuthTypeClientCert{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch azureBlobCollectorConf.Type {
	case components.AzureBlobCollectorConfTypeManual:
		// azureBlobCollectorConf.AzureBlobAuthTypeManual is populated
	case components.AzureBlobCollectorConfTypeSecret:
		// azureBlobCollectorConf.AzureBlobAuthTypeSecret is populated
	case components.AzureBlobCollectorConfTypeClientSecret:
		// azureBlobCollectorConf.AzureBlobAuthTypeClientSecret is populated
	case components.AzureBlobCollectorConfTypeClientCert:
		// azureBlobCollectorConf.AzureBlobAuthTypeClientCert is populated
	default:
		// Unknown type - use azureBlobCollectorConf.GetUnknownRaw() for raw JSON
}
```
