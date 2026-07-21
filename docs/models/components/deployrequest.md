# DeployRequest


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `Lookups`                                                                            | [][components.DeployRequestLookups](../../models/components/deployrequestlookups.md) | :heavy_minus_sign:                                                                   | Optional list of lookup file deployments to include with the commit deployment.      |
| `Version`                                                                            | `string`                                                                             | :heavy_check_mark:                                                                   | Commit hash to deploy to the Worker Group, Outpost Group, or Edge Fleet.             |