# CollectorType

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.CollectorTypeAzureBlob

// Open enum: custom values can be created with a direct type cast
custom := components.CollectorType("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `CollectorTypeAzureBlob`          | azure_blob                        |
| `CollectorTypeCriblLake`          | cribl_lake                        |
| `CollectorTypeDatabase`           | database                          |
| `CollectorTypeFilesystem`         | filesystem                        |
| `CollectorTypeGoogleCloudStorage` | google_cloud_storage              |
| `CollectorTypeHealthCheck`        | health_check                      |
| `CollectorTypeRest`               | rest                              |
| `CollectorTypeS3`                 | s3                                |
| `CollectorTypeScript`             | script                            |
| `CollectorTypeSplunk`             | splunk                            |