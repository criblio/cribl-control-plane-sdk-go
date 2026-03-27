# ShardIteratorStart

Location at which to start reading a shard for the first time

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.ShardIteratorStartTrimHorizon

// Open enum: custom values can be created with a direct type cast
custom := components.ShardIteratorStart("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `ShardIteratorStartTrimHorizon` | TRIM_HORIZON                    |
| `ShardIteratorStartLatest`      | LATEST                          |