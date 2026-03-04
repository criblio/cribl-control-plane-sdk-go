# GoogleCloudStorageAuthTypeManualAuthenticationMethod

Enter account credentials manually, select a secret that references your credentials, or use Google Application Default Credentials

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.GoogleCloudStorageAuthTypeManualAuthenticationMethodAuto

// Open enum: custom values can be created with a direct type cast
custom := components.GoogleCloudStorageAuthTypeManualAuthenticationMethod("custom_value")
```


## Values

| Name                                                         | Value                                                        |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| `GoogleCloudStorageAuthTypeManualAuthenticationMethodAuto`   | auto                                                         |
| `GoogleCloudStorageAuthTypeManualAuthenticationMethodManual` | manual                                                       |
| `GoogleCloudStorageAuthTypeManualAuthenticationMethodSecret` | secret                                                       |