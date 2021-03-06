package zhscript

import (
	"reflect"
)

type code_call___ struct {
	code_z___
}

var call_t_ reflect.Value
func call_cfg__(i interface{}) {
	call_t_ = reflect.ValueOf(i)
}

var callcache_ map[string] reflect.Value = make(map[string] reflect.Value)

func call__(args *Args___, qv *Qv___, buf2 *Buf___) (goto1 *Goto___, err1 *Errinfo___) {
	m, ok := callcache_[args.Src]
	if !ok {
		m = call_t_.MethodByName(args.Src)
		ok = m.IsValid()
		if !ok {
			err1 = New_errinfo__(args.Src, Errs_.Fail)
			return
		}
		callcache_[args.Src] = m
	}

	a := make([]reflect.Value, 1 + len(args.A))
	a[0] = reflect.ValueOf(qv)
	for i, s := range args.A {
		a[i + 1] = reflect.ValueOf(s)
	}
	defer func() {
		if err := recover(); err != nil {
			err1 = New_errinfo__(err.(string), Errs_.Fail)
		}
	}()
	for i, v := range m.Call(a) {
		switch i {
		case 0:
			goto1 = v.Interface().(*Goto___)
		case 1:
			err1 = v.Interface().(*Errinfo___)
		default:
			buf2.get__(i - 2).WriteString(v.String())
		}
	}
	return
}

func (this *buf_codes___) add_call__(kw *Keyword___) *codes___ {
	v := code_call___{code_z___{code_i___:code_i___{kw}, codes:new(codes___)}}
	this.add__(&v)
	return v.codes
}
