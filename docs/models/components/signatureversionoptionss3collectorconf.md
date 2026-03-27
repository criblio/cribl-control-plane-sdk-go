# SignatureVersionOptionsS3CollectorConf

Signature version to use for signing S3 requests

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.SignatureVersionOptionsS3CollectorConfV2

// Open enum: custom values can be created with a direct type cast
custom := components.SignatureVersionOptionsS3CollectorConf("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `SignatureVersionOptionsS3CollectorConfV2` | v2                                         |
| `SignatureVersionOptionsS3CollectorConfV4` | v4                                         |