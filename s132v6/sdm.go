package sd

import (
	"device/arm"
)

func assertHandler() {
	panic("SoftDevice assert")
}

// Clock source configuration for Enable. Take a look at DefaultClockSource for
// a default that works in most cases.
type ClockSourceConfig struct {
	Source          ClockSource
	RCCalibInterval uint8         // for RC clock source: calibration interval in 0.25s
	RCCalibTemp     uint8         // for RC clock source: how often to calibrate when temp didn't change
	Accuracy        ClockAccuracy // external crystal accuracy
}

// DefaultClockSource is a good default clock source. Use this one if you don't
// know which one to use.
var DefaultClockSource = &ClockSourceConfig{
	Source:   ClockSourceXtal,
	Accuracy: ClockAccuracy250PPM,
}

type ClockSource uint8

// Clock source to use. The recommended source is ClockSourceXtal, or if you
// don't have a low-frequency oscillator the internal R/C oscillator. There is
// no reason to use the synthesized clock source: it leaves the high-frequency
// clock running continuously.
const (
	ClockSourceRC    ClockSource = 0 // internal RC oscillator
	ClockSourceSynth ClockSource = 1 // synthesized clock
	ClockSourceXtal  ClockSource = 2 // crystal oscillator
)

type ClockAccuracy uint8

// Clock accuracy used by the lower stack to compute timing windows.
const (
	ClockAccuracy250PPM ClockAccuracy = 0
	ClockAccuracy500PPM ClockAccuracy = 1
	ClockAccuracy150PPM ClockAccuracy = 2
	ClockAccuracy100PPM ClockAccuracy = 3
	ClockAccuracy75PPM  ClockAccuracy = 4
	ClockAccuracy50PPM  ClockAccuracy = 5
	ClockAccuracy30PPM  ClockAccuracy = 6
	ClockAccuracy20PPM  ClockAccuracy = 7
)

// Enable configures the SoftDevice.
func Enable(clockSource *ClockSourceConfig) Error {
	err := arm.SVCall2(svc_SD_SOFTDEVICE_ENABLE, clockSource, assertHandler)
	return Error(err)
}

// Disable deconfigures the SoftDevice.
func Disable() Error {
	err := arm.SVCall0(svc_SD_SOFTDEVICE_DISABLE)
	return Error(err)
}

// IsEnabled returns whether the SoftDevice is enabled.
func IsEnabled() bool {
	var enabled uint8
	arm.SVCall1(svc_SD_SOFTDEVICE_IS_ENABLED, &enabled)
	return enabled != 0
}
