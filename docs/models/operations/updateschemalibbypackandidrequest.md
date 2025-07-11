# UpdateSchemaLibByPackAndIDRequest


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ID`                                                                   | *string*                                                               | :heavy_check_mark:                                                     | Unique ID to PATCH for pack                                            |
| `Pack`                                                                 | *string*                                                               | :heavy_check_mark:                                                     | pack ID to PATCH                                                       |
| `SchemaLibEntry`                                                       | [components.SchemaLibEntry](../../models/components/schemalibentry.md) | :heavy_check_mark:                                                     | Schema object to be updated                                            |