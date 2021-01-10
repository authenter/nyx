package db

import (
	"github.com/authenter/nyx/pkg/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connection represents the database Connection with gorm
// and holds other things like state and logger
type Connection struct {
	*gorm.DB
	state  State
	driver Driver
	dsn    string
	logger log.Logger
}

// New returns a configured Connection object
func New(driver Driver, dsn string, logger log.Logger) *Connection {
	var connection Connection
	connection.driver = driver
	connection.logger = logger
	connection.dsn = dsn

	connection.setState(StateCreated)
	return &connection
}

// Open opens the connection to the database
func (c *Connection) Open() error {
	var err error
	c.setState(StateOpening)

	gormConfig := &gorm.Config{
		Logger: logger.Discard,
	}

	switch c.driver {
	case DriverMySQL:
		c.DB, err = gorm.Open(mysql.Open(c.dsn), gormConfig)
	default:
		return ErrUnknownDriver
	}
	if err != nil {
		c.logger.Fatal(err)
	}

	c.setState(StateConnected)
	c.logger.Info("connection to database established")

	return nil
}

// Close closing the connection to the database
func (c *Connection) Close() {
	c.setState(StateClosing)
	c.Close()
	c.setState(StateClosed)
	c.logger.Info("connection to database closed")
}

// State returns state of database Connection
func (c *Connection) State() State {
	return c.state
}

// Driver return Driver that used by database Connection
func (c *Connection) Driver() Driver {
	return c.driver
}

func (c *Connection) setState(state State) {
	c.state = state
	c.logger = c.logger.WithFields(log.Fields{"driver": c.driver, "state": state.String()})
}
