# StorageClassOptions

Storage class to select for uploaded objects

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.StorageClassOptionsStandard

// Open enum: custom values can be created with a direct type cast
custom := components.StorageClassOptions("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `StorageClassOptionsStandard`           | STANDARD                                |
| `StorageClassOptionsReducedRedundancy`  | REDUCED_REDUNDANCY                      |
| `StorageClassOptionsStandardIa`         | STANDARD_IA                             |
| `StorageClassOptionsOnezoneIa`          | ONEZONE_IA                              |
| `StorageClassOptionsIntelligentTiering` | INTELLIGENT_TIERING                     |
| `StorageClassOptionsGlacier`            | GLACIER                                 |
| `StorageClassOptionsGlacierIr`          | GLACIER_IR                              |
| `StorageClassOptionsDeepArchive`        | DEEP_ARCHIVE                            |