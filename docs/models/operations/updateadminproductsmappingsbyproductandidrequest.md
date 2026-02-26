# UpdateAdminProductsMappingsByProductAndIDRequest


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `Product`                                                              | [components.ProductsBase](../../models/components/productsbase.md)     | :heavy_check_mark:                                                     | Name of the Cribl product to update the Mapping Ruleset for            |
| `ID`                                                                   | *string*                                                               | :heavy_check_mark:                                                     | The <code>id</code> of the Mapping Ruleset to update.                  |
| `MappingRuleset`                                                       | [components.MappingRuleset](../../models/components/mappingruleset.md) | :heavy_check_mark:                                                     | MappingRuleset object                                                  |