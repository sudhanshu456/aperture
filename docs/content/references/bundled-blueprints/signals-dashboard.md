---
title: Signals Dashboard
---

```mdx-code-block
import {apertureVersion} from '../../introduction.md';
```

## Blueprint Location

GitHub: <a
href={`https://github.com/fluxninja/aperture/tree/v${apertureVersion}/blueprints/lib/1.0/blueprints/signals-dashboard`}>signals-dashboard</a>

## Introduction

This blueprint provides a
[policy monitoring](/get-started/policies/monitoring.md) dashboard that
visualizes Signals flowing through the [Circuit](/concepts/policy/circuit.md).

## Configuration

<!-- Configuration Marker -->

### Common

| Parameter Name      | Parameter Type | Default      | Description         |
| ------------------- | -------------- | ------------ | ------------------- |
| `common.policyName` | `string`       | `(required)` | Name of the policy. |

### Dashboard

| Parameter Name              | Parameter Type | Default     | Description                            |
| --------------------------- | -------------- | ----------- | -------------------------------------- |
| `dashboard.refreshInterval` | `string`       | `"10s"`     | Refresh interval for dashboard panels. |
| `dashboard.timeFrom`        | `string`       | `"now-30m"` | From time of dashboard.                |
| `dashboard.timeTo`          | `string`       | `"now"`     | To time of dashboard.                  |

#### Datasource

| Parameter Name                     | Parameter Type | Default         | Description              |
| ---------------------------------- | -------------- | --------------- | ------------------------ |
| `dashboard.datasource.name`        | `string`       | `"$datasource"` | Datasource name.         |
| `dashboard.datasource.filterRegex` | `string`       | `""`            | Datasource filter regex. |