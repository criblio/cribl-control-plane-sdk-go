# CreateInputAPIVersion

The API version to use for communicating with the server

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputAPIVersionSixDot8Dot4

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputAPIVersion("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `CreateInputAPIVersionSixDot8Dot4`   | 6.8.4                                |
| `CreateInputAPIVersionEightDot3Dot2` | 8.3.2                                |
| `CreateInputAPIVersionCustom`        | custom                               |