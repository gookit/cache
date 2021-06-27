package cache

import (
	"encoding/json"
	"errors"

	"github.com/gookit/gsr"
)

type (
	// MarshalFunc define
	MarshalFunc func(v interface{}) ([]byte, error)

	// UnmarshalFunc define
	UnmarshalFunc func(data []byte, v interface{}) error
)

// data (Un)marshal func
var (
	Marshal   MarshalFunc   = json.Marshal
	Unmarshal UnmarshalFunc = json.Unmarshal

	errNoMarshal = errors.New("must set Marshal func")
	errNoUnmarshal = errors.New("must set Unmarshal func")
)

// Option struct
type Option struct {
	Debug bool
	// Encode (Un)marshal save data
	Encode bool
	Logger gsr.Printer
	// Prefix key prefix
	Prefix string
}

/*************************************************************
 * base driver
 *************************************************************/

// BaseDriver struct
type BaseDriver struct {
	opt Option
	// last error
	lastErr error
}

// WithDebug add option: debug
func WithDebug(debug bool) func(opt *Option) {
	return func (opt *Option) {
		opt.Debug = debug
	}
}

// WithEncode add option: encode
func WithEncode(encode bool) func(opt *Option) {
	return func (opt *Option) {
		opt.Encode = encode
	}
}

// WithPrefix add option: prefix
func WithPrefix(prefix string) func(opt *Option) {
	return func (opt *Option) {
		opt.Prefix = prefix
	}
}

// WithOptions for driver
func (l *BaseDriver) WithOptions(optFns ...func(option *Option)) {
	for _, optFn := range optFns {
		optFn(&l.opt)
	}
}

// MustMarshal cache value
func (l *BaseDriver) MustMarshal(val interface{}) ([]byte, error) {
	if Marshal != nil {
		return nil, errNoMarshal
	}
	return Marshal(val)
}

// Marshal cache value
func (l *BaseDriver) Marshal(val interface{}) (interface{}, error) {
	if Marshal != nil {
		return Marshal(val)
	}

	return val, nil
}

// MustUnmarshal cache value
func (l *BaseDriver) MustUnmarshal(bts []byte, ptr interface{}) error {
	if Unmarshal != nil {
		return errNoUnmarshal
	}
	return Unmarshal(bts, ptr)
}

// Unmarshal cache value
func (l *BaseDriver) Unmarshal(val []byte, err error) interface{} {
	if err != nil {
		l.SetLastErr(err)
		return nil
	}

	var newV interface{}
	if Unmarshal != nil {
		err := Unmarshal(val, &newV)
		l.SetLastErr(err)
		return newV
	}

	return val
}

// GetAs get cache value and unmarshal as ptr.
func (l *BaseDriver) GetAs(key string, ptr interface{}) error {
	// TODO bts, err := c.Get(key)
	// if Unmarshal != nil {
	// 	err := Unmarshal(bts, ptr)
	// 	return err
	// }

	return errNoUnmarshal
}

// Key cache key build
func (l *BaseDriver) Key(key string) string {
	if l.opt.Prefix != "" {
		return l.opt.Prefix + key
	}
	return key
}

// Debugf print an debug message
func (l *BaseDriver) Debugf(format string, v ...interface{}) {
	if l.opt.Debug && l.opt.Logger != nil {
		l.opt.Logger.Printf(format, v...)
	}
}

// Logf print an log message
func (l *BaseDriver) Logf(format string, v ...interface{}) {
	if l.opt.Logger != nil {
		l.opt.Logger.Printf(format, v...)
	}
}

// SetLastErr save last error
func (l *BaseDriver) SetLastErr(err error) {
	if err != nil {
		l.lastErr = err
		l.Logf("redis error: %s\n", err.Error())
	}
}

// LastErr get
func (l *BaseDriver) LastErr(key string) error {
	return l.lastErr
}

// IsDebug get
func (l *BaseDriver) IsDebug() bool {
	return l.opt.Debug
}
