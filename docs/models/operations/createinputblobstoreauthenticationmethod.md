# CreateInputBlobStoreAuthenticationMethod

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputBlobStoreAuthenticationMethodSecret

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputBlobStoreAuthenticationMethod("custom_value")
```


## Values

| Name                                                         | Value                                                        |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| `CreateInputBlobStoreAuthenticationMethodSecret`             | secret                                                       |
| `CreateInputBlobStoreAuthenticationMethodClientSecret`       | clientSecret                                                 |
| `CreateInputBlobStoreAuthenticationMethodClientCert`         | clientCert                                                   |
| `CreateInputBlobStoreAuthenticationMethodClientAssertion`    | clientAssertion                                              |
| `CreateInputBlobStoreAuthenticationMethodClientAssertionRPC` | clientAssertion_rpc                                          |