---
title: InfluxDB
description: Integrating InfluxDB Metrics
keywords:
  - influxdb
  - otel
  - opentelemetry
  - collector
  - metrics
---

:::info

See also [influxdbreceiver docs][receiver] in `opentelemetry-collector-contrib`
repository.

:::

:::note

The `influxdbreceiver` extension is available in the default agent image. If
you're [building][build] your own Aperture Agent, add
`integrations/otel/influxdbreceiver` to the `bundled_extensions` list to make
[the receiver][receiver] available.

:::

You can configure the [OpenTelemetry Collector][opentelemetry-collector] for
InfluxDB as part of [Policy resources][policy-resources] while [applying the
policy][applying-policy]:

```yaml
policy:
  resources:
    infra_meters:
      influxdb:
        agent_group: default
        per_agent_group: true
        receivers:
          influxdb: [influxdbreceiver configuration here]
```

[build]: /reference/aperturectl/build/agent/agent.md
[receiver]:
  https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/influxdbreceiver
[opentelemetry-collector]: /reference/configuration/spec.md#telemetry-collector
[applying-policy]: /use-cases/use-cases.md
[policy-resources]: /reference/configuration/spec.md#resources
