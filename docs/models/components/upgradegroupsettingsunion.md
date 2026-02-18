# UpgradeGroupSettingsUnion


## Supported Types

### UpgradeGroupSettings1

```go
upgradeGroupSettingsUnion := components.CreateUpgradeGroupSettingsUnionUpgradeGroupSettings1(components.UpgradeGroupSettings1{/* values here */})
```

### UpgradeGroupSettings2

```go
upgradeGroupSettingsUnion := components.CreateUpgradeGroupSettingsUnionUpgradeGroupSettings2(components.UpgradeGroupSettings2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch upgradeGroupSettingsUnion.Type {
	case components.UpgradeGroupSettingsUnionTypeUpgradeGroupSettings1:
		// upgradeGroupSettingsUnion.UpgradeGroupSettings1 is populated
	case components.UpgradeGroupSettingsUnionTypeUpgradeGroupSettings2:
		// upgradeGroupSettingsUnion.UpgradeGroupSettings2 is populated
}
```
