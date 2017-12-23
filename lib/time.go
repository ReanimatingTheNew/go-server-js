package lib

import (
	"time"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

type _time struct {
	runtime *goja.Runtime
	t       *time.Time
}

func (This *_time) stringFunc(call goja.FunctionCall) goja.Value {
	str := This.t.String()
	return This.runtime.ToValue(str)
}

func (This *_time) sleep(call goja.FunctionCall) goja.Value {
	d := call.Argument(0).ToInteger()
	<-time.After(time.Duration(d) * time.Millisecond)
	return nil
}

func NewTime(runtime *goja.Runtime, t *time.Time) *goja.Object {
	obj := &_time{
		runtime: runtime,
		t:       t,
	}

	o := runtime.NewObject()
	o.Set("string", obj.stringFunc)
	return o
}

func init() {
	require.RegisterNativeModule("time", func(runtime *goja.Runtime, module *goja.Object) {
		This := &_time{
			runtime: runtime,
		}
		o := module.Get("exports").(*goja.Object)
		o.Set("sleep", This.sleep)
	})
}