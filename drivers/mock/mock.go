package mock

import (
	"os"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/emccode/rexray/core"
)

const (
	// MockOSDriverName is the name of a mock OS driver used primarily for
	// testing.
	MockOSDriverName = "mockOSDriver"

	// MockVolDriverName is the name of a mock volume driver used primarily for
	// testing.
	MockVolDriverName = "mockVolumeDriver"

	// MockStorDriverName is the name of a mock storage driver used primarily
	// for testing.
	MockStorDriverName = "mockStorageDriver"

	// BadMockOSDriverName is the name of a mock OS driver used primarily for
	// testing.
	BadMockOSDriverName = "badMockOSDriver"

	// BadMockVolDriverName is the name of a mock volume driver used primarily
	// for testing.
	BadMockVolDriverName = "badMockVolumeDriver"

	// BadMockStorDriverName is the name of a mock storage driver used primarily
	// for testing.
	BadMockStorDriverName = "badMockStorageDriver"
)

func init() {
	v := os.Getenv("REXRAY_MOCKDRIVERS")
	log.WithField("REXRAY_MOCKDRIVERS", v).Debug("got REXRAY_MOCKDRIVERS")

	if b, err := strconv.ParseBool(v); !b || err != nil {
		log.Debug("not registering mock drivers")
		return
	}

	RegisterMockDrivers()
}

// RegisterMockDrivers registers the mock drivers.
func RegisterMockDrivers() {
	log.Debug("registering mock drivers")
	core.RegisterDriver(MockOSDriverName, newOSDriver)
	core.RegisterDriver(MockVolDriverName, newVolDriver)
	core.RegisterDriver(MockStorDriverName, newStorDriver)
}

// RegisterBadMockDrivers registers the bad mock drivers.
func RegisterBadMockDrivers() {
	log.Debug("registering bad mock drivers")
	core.RegisterDriver(BadMockOSDriverName, newBadOSDriver)
	core.RegisterDriver(BadMockVolDriverName, newBadVolDriver)
	core.RegisterDriver(BadMockStorDriverName, newBadStorDriver)
}
