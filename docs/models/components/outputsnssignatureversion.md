# OutputSnsSignatureVersion

Signature version to use for signing SNS requests

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputSnsSignatureVersionV2

// Open enum: custom values can be created with a direct type cast
custom := components.OutputSnsSignatureVersion("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `OutputSnsSignatureVersionV2` | v2                            |
| `OutputSnsSignatureVersionV4` | v4                            |