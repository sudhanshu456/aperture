local aperture = import 'blueprints/lib/1.0/main.libsonnet';

local latencyGradientPolicy = aperture.blueprints.LatencyGradientConcurrencyLimiting.policy;

local workloadParameters = aperture.spec.v1.SchedulerWorkloadParameters;
local labelMatcher = aperture.spec.v1.LabelMatcher;
local workload = aperture.spec.v1.SchedulerWorkload;

local classifier = aperture.spec.v1.Classifier;
local fluxMeter = aperture.spec.v1.FluxMeter;
local extractor = aperture.spec.v1.Extractor;
local rule = aperture.spec.v1.Rule;
local flowSelector = aperture.spec.v1.FlowSelector;
local serviceSelector = aperture.spec.v1.ServiceSelector;
local flowMatcher = aperture.spec.v1.FlowMatcher;
local controlPoint = aperture.spec.v1.ControlPoint;
local component = aperture.spec.v1.Component;
local rateLimiter = aperture.spec.v1.RateLimiter;
local decider = aperture.spec.v1.Decider;
local switcher = aperture.spec.v1.Switcher;
local port = aperture.spec.v1.Port;

local fluxMeterSelector =
  flowSelector.new()
  + flowSelector.withServiceSelector(
    serviceSelector.new()
    + serviceSelector.withAgentGroup('default')
    + serviceSelector.withService('service3-demo-app.demoapp.svc.cluster.local')
  )
  + flowSelector.withFlowMatcher(
    flowMatcher.new()
    + flowMatcher.withControlPoint('ingress')
  );

local concurrencyLimiterFlowSelector =
  flowSelector.new()
  + flowSelector.withServiceSelector(
    serviceSelector.new()
    + serviceSelector.withAgentGroup('default')
    + serviceSelector.withService('service1-demo-app.demoapp.svc.cluster.local')
  )
  + flowSelector.withFlowMatcher(
    flowMatcher.new()
    + flowMatcher.withControlPoint('ingress')
  );

// Restrict this selector to only bot traffic
local rateLimiterSelector =
  flowSelector.new()
  + flowSelector.withServiceSelector(
    serviceSelector.new()
    + serviceSelector.withAgentGroup('default')
    + serviceSelector.withService('service1-demo-app.demoapp.svc.cluster.local')
  )
  + flowSelector.withFlowMatcher(
    flowMatcher.new()
    + flowMatcher.withControlPoint('ingress')
    + flowMatcher.withLabelMatcher(
      labelMatcher.withMatchLabels({ 'http.request.header.user_type': 'bot' })
    )
  );

local policyResource = latencyGradientPolicy({
  policyName: 'service1-demo-app',
  fluxMeter: fluxMeter.new() + fluxMeter.withFlowSelector(fluxMeterSelector),
  concurrencyLimiterFlowSelector: concurrencyLimiterFlowSelector,
  classifiers: [
    classifier.new()
    + classifier.withFlowSelector(concurrencyLimiterFlowSelector)
    + classifier.withRules({
      user_type: rule.new()
                 + rule.withExtractor(extractor.new()
                                      + extractor.withFrom('request.http.headers.user-type')),
    }),
  ],
  concurrencyLimiter+: {
    alerterChannels: ['service1-demo-app'],
    timeoutFactor: 0.5,
    defaultWorkloadParameters: {
      priority: 20,
    },
    workloads: [
      workload.new()
      + workload.withWorkloadParameters(workloadParameters.withPriority(50))
      // match the label extracted by classifier
      + workload.withLabelMatcher(labelMatcher.withMatchLabels({ user_type: 'guest' })),
      workload.new()
      + workload.withWorkloadParameters(workloadParameters.withPriority(200))
      // match the http header directly
      + workload.withLabelMatcher(labelMatcher.withMatchLabels({ 'http.request.header.user_type': 'subscriber' })),
    ],
  },
  components: [
    component.new()
    + component.withDecider(
      decider.new()
      + decider.withOperator('lt')
      + decider.withInPorts({ lhs: port.withSignalName('LOAD_MULTIPLIER'), rhs: port.withConstantSignal(1.0) })
      + decider.withOutPorts({ output: port.withSignalName('IS_BOT_ESCALATION') })
      + decider.withTrueFor('30s')
    ),
    component.new()
    + component.withSwitcher(
      switcher.new()
      + switcher.withInPorts({
        switch: port.withSignalName('IS_BOT_ESCALATION'),
        on_true: port.withConstantSignal(0.0),
        on_false: port.withConstantSignal(10.0),
      })
      + switcher.withOutPorts({ output: port.withSignalName('RATE_LIMIT') })
    ),
    component.new()
    + component.withRateLimiter(
      rateLimiter.new()
      + rateLimiter.withFlowSelector(rateLimiterSelector)
      + rateLimiter.withInPorts({ limit: port.withSignalName('RATE_LIMIT') })
      + rateLimiter.withLimitResetInterval('1s')
      + rateLimiter.withLabelKey('http.request.header.user_id')
      + rateLimiter.withDynamicConfigKey('rate_limiter'),
    ),
  ],
}).policyResource;

{
  latencyGradientPolicy: policyResource,
}