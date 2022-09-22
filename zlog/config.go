package zlog

import "github.com/rs/zerolog"

// Config - Configuration for logging
type Config struct {

	// caller set console color
	NoColor bool

	// caller show caller logging
	NoCaller bool

	// Level defines log levels.
	Level zerolog.Level

	// Enable console logging
	ConsoleLoggingEnabled bool

	// EncodeLogsAsJSON makes the log framework log JSON
	EncodeLogsAsJSON bool

	// FileLoggingEnabled makes the framework log to a file
	// the fields below can be skipped if this value is false!
	FileLoggingEnabled bool

	// Directory to log to to when filelogging is enabled
	Directory string

	// Filename is the name of the logfile which will be placed inside the directory
	Filename string

	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated. It defaults to 100 megabytes.
	MaxSize int `json:"maxsize" yaml:"maxsize"`

	// MaxAge is the maximum number of days to retain old log files based on the
	// timestamp encoded in their filename.  Note that a day is defined as 24
	// hours and may not exactly correspond to calendar days due to daylight
	// savings, leap seconds, etc. The default is not to remove old log files
	// based on age.
	MaxAge int `json:"maxage" yaml:"maxage"`

	// MaxBackups is the maximum number of old log files to retain.  The default
	// is to retain all old log files (though MaxAge may still cause them to get
	// deleted.)
	MaxBackups int `json:"maxbackups" yaml:"maxbackups"`
}

// Option - option config logging
type Option func(cfg *Config)

func (opts *Config) loadOptions(options ...Option) {
	for _, option := range options {
		option(opts)
	}
}

// WithCaller - Set Caller (defautl: false)
func WithCaller(show bool) func(c *Config) {
	return func(c *Config) {
		c.NoCaller = !show
	}
}

// WithColor - Set Caller (defautl: false)
func WithColor(show bool) func(*Config) {
	return func(c *Config) {
		c.NoColor = !show
	}
}

// WithLevel - Set Caller (defautl: 0)
func WithLevel(lvl zerolog.Level) func(*Config) {
	return func(c *Config) {
		c.Level = lvl
	}
}
