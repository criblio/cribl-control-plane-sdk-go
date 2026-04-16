# ConfigGroupLookups


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `Context`                                                                                    | `string`                                                                                     | :heavy_check_mark:                                                                           | The Worker or Node context for the lookup deployment.                                        |
| `Lookups`                                                                                    | [][components.ConfigGroupLookupsLookup](../../models/components/configgrouplookupslookup.md) | :heavy_check_mark:                                                                           | List of lookup files deployed to this context.                                               |