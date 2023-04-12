---
title: Redis
description: Integrating Redis Metrics
keywords:
  - redis
  - otel
  - opentelemetry
  - collector
  - metrics
---

Before proceeding, ensure that you have [built][build] Aperture Agent with the
`redisreceiver` extension enabled, so that [redisreceiver][receiver] is
available.

You can configure [Custom metrics][custom-metrics] for Redis using the following
configuration in the [Aperture Agent's config][agent-config]:

```yaml
otel:
  custom_metrics:
    redis:
      per_agent_group: true
      pipeline:
        processors:
          - batch
        receivers:
          - redis
      processors:
        batch:
          send_batch_size: 10
          timeout: 10s
      receivers:
        redis: [redisreceiver configuration here]
```

[build]: /reference/aperturectl/build/agent/agent.md
[receiver]:
  https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/redisreceiver
[custom-metrics]: /reference/configuration/agent.md#custom-metrics-config
[agent-config]: /reference/configuration/agent.md#agent-o-t-e-l-config