---
title: CAMEL_CASE
description: Integrating CAMEL_CASE Metrics
keywords:
  - METRIC_NAME
  - otel
  - opentelemetry
  - collector
  - metrics
---

:::info

See also [RECEIVER_NAME docs][receiver] in `opentelemetry-collector-contrib`
repository.

:::

:::note

The `RECEIVER_NAME` extension is available in the default agent image. If you're
[building][build] your own Aperture Agent, add `integrations/otel/RECEIVER_NAME`
to the `bundled_extensions` list to make [the receiver][receiver] available.

:::

You can configure the [OpenTelemetry Collector][opentelemetry-collector] for
CAMEL_CASE as part of [Policy resources][policy-resources] while [applying the
policy][applying-policy]:

```yaml
policy:
  resources:
    telemetry_collectors:
      - agent_group: default
        infra_meters:
          METRIC_NAME:
            per_agent_group: true
            pipeline:
              processors:
                - batch
              receivers:
                - METRIC_NAME
            processors:
              batch:
                send_batch_size: 10
                timeout: 10s
            receivers:
              METRIC_NAME: [RECEIVER_NAME configuration here]
```

[build]: /reference/aperturectl/build/agent/agent.md
[receiver]:
  https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/RECEIVER_NAME
[opentelemetry-collector]: /reference/policies/spec.md#telemetry-collector
[applying-policy]: /applying-policies/applying-policies.md
[policy-resources]: /reference/policies/spec.md#resources