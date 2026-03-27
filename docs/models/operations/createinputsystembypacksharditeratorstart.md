# CreateInputSystemByPackShardIteratorStart

Location at which to start reading a shard for the first time

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackShardIteratorStartTrimHorizon

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackShardIteratorStart("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `CreateInputSystemByPackShardIteratorStartTrimHorizon` | TRIM_HORIZON                                           |
| `CreateInputSystemByPackShardIteratorStartLatest`      | LATEST                                                 |