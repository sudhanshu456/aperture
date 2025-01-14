local utils = import '../utils/policy_utils.libsonnet';
local timeSeriesPanel = import '../utils/time_series_panel.libsonnet';

function(cfg) {
  local stringFilters = utils.dictToPrometheusFilter(cfg.dashboard.extra_filters { policy_name: cfg.policy.policy_name }),

  local quotaScheduler = timeSeriesPanel('Quota Checks',
                                         cfg.dashboard.datasource.name,
                                         'sum by(decision_type) (rate(workload_requests_total{%(filters)s}[$__rate_interval]))',
                                         stringFilters,
                                         'Decisions',
                                         'reqps'),

  panel: quotaScheduler.panel,
}
