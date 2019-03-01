package validator_test

import (
	"fmt"
	"testing"

	"github.com/bytedance/go-tagexpr/validator"
)

func Test(t *testing.T) {
	var vd = validator.New("vd")

	type InfoRequest struct {
		Name string `vd:"($=='Alice') && regexp('\\w')"`
		Age  int    `vd:"$>0"`
		Addr string `vd:"{@:$=='shenzhen'&&(Age)$==18}{msg:'addr should be shenzhen'}"`
	}
	// case1: all validation passed
	info1 := &InfoRequest{Name: "Alice", Age: 18, Addr: "shenzhen"}
	t.Log(vd.Validate(info1))

	// case2: addr not pass
	info2 := &InfoRequest{Name: "Alice", Age: 18, Addr: "beijing"}
	t.Log(vd.Validate(info2))

	// case3: name & addr not pass
	info3 := &InfoRequest{Name: "Bob", Age: 18, Addr: "beijing"}
	t.Log(vd.Validate(info3))

	type A struct {
		A int `vd:"$<0||$>=100"`
	}
	a := &A{107}
	t.Log(vd.Validate(a) == nil)

	type B struct {
		B string `vd:"len($)>1 && regexp('^\\w*$')"`
	}
	b := &B{"abc"}
	t.Log(vd.Validate(b) == nil)

	type C struct {
		C bool `vd:"{@:(S.A)$>0 && !$}{msg:'C must be false when S.A>0'}"`
		S *A
	}
	c := &C{C: true, S: a}
	t.Log(vd.Validate(c))

	type D struct {
		d []string `vd:"{@:len($)>0 && $[0]=='D'} {msg:sprintf('Invalid d: %v',$)}"`
	}
	d := &D{d: []string{"x", "y"}}
	t.Log(vd.Validate(d))

	type E struct {
		e map[string]int `vd:"len($)==$['len']"`
	}
	e := &E{map[string]int{"len": 2}}
	t.Log(vd.Validate(e))

	// Customizes the factory of validation error.
	vd.SetErrorFactory(func(fieldSelector, msg string) error {
		return fmt.Errorf(`{"succ":false, "error":"invalid parameter: %s"}`, fieldSelector)
	})

	type F struct {
		f struct {
			g int `vd:"$%3==0"`
		}
	}
	f := &F{}
	f.f.g = 10
	t.Log(vd.Validate(f))

	// Output:
	// true
	// true
	// true
	// C must be false when S.A>0
	// Invalid d: [x y]
	// Invalid parameter: e
	// {"succ":false, "error":"invalid parameter: f.g"}
}
