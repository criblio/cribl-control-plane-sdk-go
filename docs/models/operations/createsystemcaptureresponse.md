# CreateSystemCaptureResponse


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `HTTPMeta`                                                         | [components.HTTPMetadata](../../models/components/httpmetadata.md) | :heavy_check_mark:                                                 | N/A                                                                |
| `CapturedEvent`                                                    | **jsonl.JsonLStream[map[string]any]*                               | :heavy_minus_sign:                                                 | CapturedEvent object                                               |