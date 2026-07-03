# DistributedSummary

Summary of the deployment for the specified Cribl product (Stream or Edge).


## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `Groups`                                                                                   | [components.DistributedSummaryGroups](../../models/components/distributedsummarygroups.md) | :heavy_check_mark:                                                                         | Resource counts for Worker Groups or Edge Fleets in the deployment summary.                |
| `Workers`                                                                                  | [*components.Workers](../../models/components/workers.md)                                  | :heavy_minus_sign:                                                                         | Worker or Edge Node counts and health statistics in the deployment summary.                |