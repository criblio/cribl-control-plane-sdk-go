# StorageClassOptionsCriblLakeDataset

Storage class used for objects written to the Dataset.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.StorageClassOptionsCriblLakeDatasetDeepArchive

// Open enum: custom values can be created with a direct type cast
custom := components.StorageClassOptionsCriblLakeDataset("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `StorageClassOptionsCriblLakeDatasetDeepArchive`        | DEEP_ARCHIVE                                            |
| `StorageClassOptionsCriblLakeDatasetGlacier`            | GLACIER                                                 |
| `StorageClassOptionsCriblLakeDatasetGlacierIr`          | GLACIER_IR                                              |
| `StorageClassOptionsCriblLakeDatasetIntelligentTiering` | INTELLIGENT_TIERING                                     |
| `StorageClassOptionsCriblLakeDatasetOnezoneIa`          | ONEZONE_IA                                              |
| `StorageClassOptionsCriblLakeDatasetStandard`           | STANDARD                                                |
| `StorageClassOptionsCriblLakeDatasetStandardIa`         | STANDARD_IA                                             |