# FunctionConfSchemaSamplingRule


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `Filter`                                                                           | **string*                                                                          | :heavy_minus_sign:                                                                 | JavaScript filter expression matching events to be sampled. Use true to match all. |
| `Rate`                                                                             | **int64*                                                                           | :heavy_minus_sign:                                                                 | Sampling rate; picks one out of N matching events                                  |