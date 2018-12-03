package sd

import (
	"device/arm"
)

// EnableBLE enables the BLE stack. It must be called after the SoftDevice
// itself has been enabled.
func EnableBLE(app_ram_base uintptr) (uintptr, Error) {
	err := arm.SVCall1(svc_SD_BLE_ENABLE, &app_ram_base)
	return app_ram_base, Error(err)
}
