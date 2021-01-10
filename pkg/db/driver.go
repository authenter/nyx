package db

// Driver defines database Drivers
type Driver string

const (
	// DriverMySQL means that database will use gorm mysql driver
	DriverMySQL Driver = "mysql"
)

// ParseDriver takes a string driver and returns the Driver constant
func ParseDriver(driver string) (Driver, error) {
	switch driver {
	case "mysql":
		return DriverMySQL, nil
	default:
		return "", ErrUnknownDriver
	}
}
