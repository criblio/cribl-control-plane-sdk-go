# SignatureVersionOptions2

Signature version to use for signing Kinesis stream requests

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.SignatureVersionOptions2V2

// Open enum: custom values can be created with a direct type cast
custom := components.SignatureVersionOptions2("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `SignatureVersionOptions2V2` | v2                           |
| `SignatureVersionOptions2V4` | v4                           |