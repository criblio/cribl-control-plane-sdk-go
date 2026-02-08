# PreviewResponse


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `Count`                                                                | *float64*                                                              | :heavy_check_mark:                                                     | Total count of captured events.                                        |
| `Items`                                                                | [][components.CapturedEvent](../../models/components/capturedevent.md) | :heavy_check_mark:                                                     | Array of events captured during Pipeline preview.                      |
| `Message`                                                              | **string*                                                              | :heavy_minus_sign:                                                     | N/A                                                                    |
| `Profile`                                                              | **string*                                                              | :heavy_minus_sign:                                                     | N/A                                                                    |
| `Stats`                                                                | [*components.CaptureStats](../../models/components/capturestats.md)    | :heavy_minus_sign:                                                     | N/A                                                                    |
| `Stderr`                                                               | *string*                                                               | :heavy_check_mark:                                                     | N/A                                                                    |