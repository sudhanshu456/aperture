package components

import (
	"math"

	"go.uber.org/fx"

	policylangv1 "github.com/fluxninja/aperture/v2/api/gen/proto/go/aperture/policy/language/v1"
	"github.com/fluxninja/aperture/v2/pkg/config"
	"github.com/fluxninja/aperture/v2/pkg/notifiers"
	"github.com/fluxninja/aperture/v2/pkg/policies/controlplane/iface"
	"github.com/fluxninja/aperture/v2/pkg/policies/controlplane/runtime"
	"github.com/fluxninja/aperture/v2/pkg/policies/controlplane/runtime/tristate"
)

// Integrator is a component that accumulates sum of signal every tick.
type Integrator struct {
	sum float64
}

// Name implements runtime.Component.
func (*Integrator) Name() string { return "Integrator" }

// Type implements runtime.Component.
func (*Integrator) Type() runtime.ComponentType { return runtime.ComponentTypeSignalProcessor }

// ShortDescription implements runtime.Component.
func (in *Integrator) ShortDescription() string {
	return ""
}

// IsActuator implements runtime.Component.
func (*Integrator) IsActuator() bool { return false }

// NewIntegrator creates an integrator component.
func NewIntegrator(initialValue float64) runtime.Component {
	integrator := &Integrator{
		sum: initialValue,
	}
	return integrator
}

// NewIntegratorAndOptions creates an integrator component and its fx options.
func NewIntegratorAndOptions(integratorProto *policylangv1.Integrator, _ runtime.ComponentID, _ iface.Policy) (runtime.Component, fx.Option, error) {
	initialValue := integratorProto.GetInitialValue()
	return NewIntegrator(initialValue), fx.Options(), nil
}

// Execute implements runtime.Component.Execute.
func (in *Integrator) Execute(inPortReadings runtime.PortToReading, tickInfo runtime.TickInfo) (runtime.PortToReading, error) {
	inputVal := inPortReadings.ReadSingleReadingPort("input")
	resetVal := inPortReadings.ReadSingleReadingPort("reset")
	if tristate.FromReading(resetVal).IsTrue() {
		in.sum = 0
	} else if inputVal.Valid() {
		in.sum += inputVal.Value()

		maxVal := inPortReadings.ReadSingleReadingPort("max")
		if maxVal.Valid() {
			in.sum = math.Min(in.sum, maxVal.Value())
		}

		minVal := inPortReadings.ReadSingleReadingPort("min")
		if minVal.Valid() {
			in.sum = math.Max(in.sum, minVal.Value())
		}
	}

	return runtime.PortToReading{
		"output": []runtime.Reading{runtime.NewReading(in.sum)},
	}, nil
}

// DynamicConfigUpdate is a no-op for Integrator.
func (in *Integrator) DynamicConfigUpdate(event notifiers.Event, unmarshaller config.Unmarshaller) {}
