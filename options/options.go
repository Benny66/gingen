package options

import (
	"github.com/spf13/pflag"
)

type Options interface {
	Lookup(key string) bool
	GetInt(key string, def int) (int, error)
	GetInt8(key string, def int8) (int8, error)
	GetInt16(key string, def int16) (int16, error)
	GetInt32(key string, def int32) (int32, error)
	GetInt64(key string, def int64) (int64, error)
	GetFloat32(key string, def float32) (float32, error)
	GetFloat64(key string, def float64) (float64, error)
	GetBool(key string, def bool) (bool, error)
	GetString(key string, def string) string
}

type flagOptions struct {
	flags *pflag.FlagSet
}

func NewFlagOptions(flags *pflag.FlagSet) Options {
	return &flagOptions{
		flags: flags,
	}
}

func (o *flagOptions) Lookup(key string) bool {
	flag := o.flags.Lookup(key)
	return flag != nil
}

func (o *flagOptions) GetInt(key string, def int) (int, error) {
	if o.Lookup(key) {
		return o.flags.GetInt(key)
	}
	return def, nil
}

func (o *flagOptions) GetInt8(key string, def int8) (int8, error) {
	if o.Lookup(key) {
		return o.flags.GetInt8(key)
	}
	return def, nil
}

func (o *flagOptions) GetInt16(key string, def int16) (int16, error) {
	if o.Lookup(key) {
		return o.flags.GetInt16(key)
	}
	return def, nil
}

func (o *flagOptions) GetInt32(key string, def int32) (int32, error) {
	if o.Lookup(key) {
		return o.flags.GetInt32(key)
	}
	return def, nil
}

func (o *flagOptions) GetInt64(key string, def int64) (int64, error) {
	if o.Lookup(key) {
		return o.flags.GetInt64(key)
	}
	return def, nil
}

func (o *flagOptions) GetFloat32(key string, def float32) (float32, error) {
	if o.Lookup(key) {
		return o.flags.GetFloat32(key)
	}
	return def, nil
}

func (o *flagOptions) GetFloat64(key string, def float64) (float64, error) {
	if o.Lookup(key) {
		return o.flags.GetFloat64(key)
	}
	return def, nil
}

func (o *flagOptions) GetBool(key string, def bool) (bool, error) {
	if o.Lookup(key) {
		return o.flags.GetBool(key)
	}
	return def, nil
}

func (o *flagOptions) GetString(key string, def string) string {
	if o.Lookup(key) {
		v, _ := o.flags.GetString(key)
		return v
	}
	return def
}
