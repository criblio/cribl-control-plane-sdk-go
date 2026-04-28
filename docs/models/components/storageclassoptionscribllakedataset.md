# StorageClassOptionsCriblLakeDataset

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.StorageClassOptionsCriblLakeDatasetStandard

// Open enum: custom values can be created with a direct type cast
custom := components.StorageClassOptionsCriblLakeDataset("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `StorageClassOptionsCriblLakeDatasetStandard`           | STANDARD                                                |
| `StorageClassOptionsCriblLakeDatasetStandardIa`         | STANDARD_IA                                             |
| `StorageClassOptionsCriblLakeDatasetOnezoneIa`          | ONEZONE_IA                                              |
| `StorageClassOptionsCriblLakeDatasetIntelligentTiering` | INTELLIGENT_TIERING                                     |
| `StorageClassOptionsCriblLakeDatasetGlacier`            | GLACIER                                                 |
| `StorageClassOptionsCriblLakeDatasetGlacierIr`          | GLACIER_IR                                              |
| `StorageClassOptionsCriblLakeDatasetDeepArchive`        | DEEP_ARCHIVE                                            |