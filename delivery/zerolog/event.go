package zero

import (
	"fmt"
	"github.com/nooize/go-logz"
	"github.com/rs/zerolog"
)

type zeroEvent struct {
	event *zerolog.Event
}

func (e *zeroEvent) Enabled() bool {
	return e.event.Enabled()
}

func (e *zeroEvent) Discard() logz.Event {
	e.event = e.event.Discard()
	return e
}

func (e *zeroEvent) With(key string, v interface{}) logz.Event {
	switch v.(type) {
	case string:
		e.event = e.event.Str(key, v.(string))
	case fmt.Stringer:
		e.event = e.event.Str(key, v.(fmt.Stringer).String())
	case int:
		e.event = e.event.Int(key, v.(int))
	case int8:
		e.event = e.event.Int8(key, v.(int8))
	case int16:
		e.event = e.event.Int16(key, v.(int16))
	case int32:
		e.event = e.event.Int32(key, v.(int32))
	case int64:
		e.event = e.event.Int64(key, v.(int64))
	case float32:
		e.event = e.event.Float32(key, v.(float32))
	case float64:
		e.event = e.event.Float64(key, v.(float64))
	}
	return e
}

func (e *zeroEvent) Err(err error) logz.Event {
	e.event = e.event.Err(err)
	return e
}

func (e *zeroEvent) Send(f string, i ...interface{}) {
	e.event.Msgf(f, i...)
}
