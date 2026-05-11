# InputResponseCheckpointStore

The backing store used to persist consumer checkpoints. Select "None" to disable checkpointing (consumers will restart from the configured start position).

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputResponseCheckpointStoreNone

// Open enum: custom values can be created with a direct type cast
custom := components.InputResponseCheckpointStore("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `InputResponseCheckpointStoreNone`      | none                                    |
| `InputResponseCheckpointStoreAzureBlob` | azureBlob                               |