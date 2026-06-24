# FormatOptionsCriblLakeDataset

Storage format used for data persisted in the Dataset.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.FormatOptionsCriblLakeDatasetDdss

// Open enum: custom values can be created with a direct type cast
custom := components.FormatOptionsCriblLakeDataset("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `FormatOptionsCriblLakeDatasetDdss`     | ddss                                    |
| `FormatOptionsCriblLakeDatasetJSON`     | json                                    |
| `FormatOptionsCriblLakeDatasetNetskope` | netskope                                |
| `FormatOptionsCriblLakeDatasetParquet`  | parquet                                 |