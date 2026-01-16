# FunctionConfSchemaSamplingRule


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `Filter`                                                                           | *string*                                                                           | :heavy_check_mark:                                                                 | JavaScript filter expression matching events to be sampled. Use true to match all. |
| `Rate`                                                                             | *int64*                                                                            | :heavy_check_mark:                                                                 | Sampling rate; picks one out of N matching events                                  |