# FunctionConfSchemaOtlpTraces


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `DropNonTraceEvents`                                                            | **bool*                                                                         | :heavy_minus_sign:                                                              | N/A                                                                             |
| `OtlpVersion`                                                                   | [*components.OtlpVersionOptions](../../models/components/otlpversionoptions.md) | :heavy_minus_sign:                                                              | N/A                                                                             |
| `BatchOTLPTraces`                                                               | **bool*                                                                         | :heavy_minus_sign:                                                              | Batch OTLP traces by shared top-level `resource` attributes                     |