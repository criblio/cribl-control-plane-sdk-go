# CreateInputSystemByPackBlobStoreAuthenticationMethod

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackBlobStoreAuthenticationMethodSecret

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackBlobStoreAuthenticationMethod("custom_value")
```


## Values

| Name                                                                     | Value                                                                    |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `CreateInputSystemByPackBlobStoreAuthenticationMethodSecret`             | secret                                                                   |
| `CreateInputSystemByPackBlobStoreAuthenticationMethodClientSecret`       | clientSecret                                                             |
| `CreateInputSystemByPackBlobStoreAuthenticationMethodClientCert`         | clientCert                                                               |
| `CreateInputSystemByPackBlobStoreAuthenticationMethodClientAssertion`    | clientAssertion                                                          |
| `CreateInputSystemByPackBlobStoreAuthenticationMethodClientAssertionRPC` | clientAssertion_rpc                                                      |