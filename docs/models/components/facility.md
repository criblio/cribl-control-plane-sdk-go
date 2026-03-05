# Facility

Default value for message facility. Will be overwritten by value of __facility if set. Defaults to user.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.FacilityZero

// Open enum: custom values can be created with a direct type cast
custom := components.Facility(999)
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `FacilityZero`      | 0                   |
| `FacilityOne`       | 1                   |
| `FacilityTwo`       | 2                   |
| `FacilityThree`     | 3                   |
| `FacilityFour`      | 4                   |
| `FacilityFive`      | 5                   |
| `FacilitySix`       | 6                   |
| `FacilitySeven`     | 7                   |
| `FacilityEight`     | 8                   |
| `FacilityNine`      | 9                   |
| `FacilityTen`       | 10                  |
| `FacilityEleven`    | 11                  |
| `FacilityTwelve`    | 12                  |
| `FacilityThirteen`  | 13                  |
| `FacilityFourteen`  | 14                  |
| `FacilityFifteen`   | 15                  |
| `FacilitySixteen`   | 16                  |
| `FacilitySeventeen` | 17                  |
| `FacilityEighteen`  | 18                  |
| `FacilityNineteen`  | 19                  |
| `FacilityTwenty`    | 20                  |
| `FacilityTwentyOne` | 21                  |