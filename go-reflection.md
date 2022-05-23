



```
package main

import (
	"reflect"
	. "test/social_network"
	"testing"
)

func TestFill(t *testing.T) {
	tt := reflect.TypeOf(Post{})
	v := reflect.New(tt)
	initializeStruct(tt, v.Elem())
	c := v.Interface().(*Post)

	mmap := make(map[string]interface{}, 0)
	mmap["text"] = "Hello"
	mmap["creator"] = map[string]interface{}{"username": "per"}

	FillStruct(c, mmap, "json")

}

func initializeStruct(t reflect.Type, v reflect.Value) {
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		ft := t.Field(i)
		switch ft.Type.Kind() {
		case reflect.Map:
			f.Set(reflect.MakeMap(ft.Type))
		case reflect.Slice:
			f.Set(reflect.MakeSlice(ft.Type, 0, 0))
		case reflect.Chan:
			f.Set(reflect.MakeChan(ft.Type, 0))
		case reflect.Struct:
			initializeStruct(ft.Type, f)
		case reflect.Ptr:
			fv := reflect.New(ft.Type.Elem())
			initializeStruct(ft.Type.Elem(), fv.Elem())
			f.Set(fv)
		default:
		}
	}
}


const (
	tag = "json"
)

func FillStruct(container interface{}, dict map[string]interface{}, tag string) {
	cv := reflect.ValueOf(container)

	fillStruct(cv, dict, tag)
}

func fillBool(v reflect.Value, x interface{}) {
	v.SetBool(x.(bool))
}

func fillInt(v reflect.Value, x interface{}) {
	xv := reflect.ValueOf(x)

	if v.Type() != xv.Type() {
		xv = xv.Convert(v.Type())
	}

	v.SetInt(xv.Int())
}

func fillUint(v reflect.Value, x interface{}) {
	xv := reflect.ValueOf(x)

	if v.Type() != xv.Type() {
		xv = xv.Convert(v.Type())
	}

	v.SetUint(xv.Uint())
}

func fillString(v reflect.Value, x interface{}) {
	v.SetString(x.(string))
}

func fillSlice(v reflect.Value, s []interface{}, tag string) {
	nv := reflect.MakeSlice(v.Type(), len(s), cap(s))
	kind := nv.Type().Elem().Kind()

	for i := 0; i < len(s); i++ {
		fv := nv.Index(i)
		sx := s[i]
		shunt(kind, fv, sx, tag)
	}

	v.Set(nv)
}

func fillMap(v reflect.Value, d map[string]interface{}, tag string) {
	nv := reflect.MakeMapWithSize(v.Type(), len(d))

	kt := nv.Type().Key()
	ktk := kt.Kind()
	vt := nv.Type().Elem()
	vtk := vt.Kind()

	for dk, dx := range d {
		kv := reflect.Indirect(reflect.New(kt))
		shunt(ktk, kv, dk, tag)

		xv := reflect.Indirect(reflect.New(vt))
		shunt(vtk, xv, dx, tag)

		nv.SetMapIndex(kv, xv)
	}

	v.Set(nv)
}

func fillStruct(v reflect.Value, d map[string]interface{}, tag string) {
	if reflect.Ptr == v.Kind() {
		v = v.Elem()
	}
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fv := v.Field(i)
		ft := t.Field(i)
		dx, ok := d[ft.Tag.Get(tag)]
		if ok && fv.CanSet() {
			shunt(fv.Kind(), fv, dx, tag)
		}
	}
}

func shunt(kind reflect.Kind, v reflect.Value, x interface{}, tag string) {
	switch kind {
	case reflect.Bool:
		fillBool(v, x)
	case reflect.Int:
		fillInt(v, x)
	case reflect.Int8:
		fillInt(v, x)
	case reflect.Int16:
		fillInt(v, x)
	case reflect.Int32:
		fillInt(v, x)
	case reflect.Int64:
		fillInt(v, x)
	case reflect.Uint:
		fillUint(v, x)
	case reflect.Uint8:
		fillUint(v, x)
	case reflect.Uint16:
		fillUint(v, x)
	case reflect.Uint32:
		fillUint(v, x)
	case reflect.Uint64:
		fillUint(v, x)
	case reflect.Map:
		fillMap(v, x.(map[string]interface{}), tag)
	case reflect.Ptr:
		fillStruct(v, x.(map[string]interface{}), tag)
	case reflect.Slice:
		fillSlice(v, x.([]interface{}), tag)
	case reflect.String:
		fillString(v, x)
	case reflect.Struct:
		fillStruct(v, x.(map[string]interface{}), tag)
	}
}


```


[Stack OF](https://stackoverflow.com/questions/64138199/how-to-set-a-struct-member-that-is-a-pointer-to-an-arbitrary-value-using-reflect)

```
type MyStruct struct {
    SomeIntPtr    *int
    SomeStringPtr *string
}

var ms MyStruct

// Set int pointer
{
    var i interface{} = 3 // of type int

    f := reflect.ValueOf(&ms).Elem().FieldByName("SomeIntPtr")
    x := reflect.New(f.Type().Elem())
    x.Elem().Set(reflect.ValueOf(i))
    f.Set(x)
}

// Set string pointer
{
    var i interface{} = "hi" // of type string

    f := reflect.ValueOf(&ms).Elem().FieldByName("SomeStringPtr")
    x := reflect.New(f.Type().Elem())
    x.Elem().Set(reflect.ValueOf(i))
    f.Set(x)
}

fmt.Println("ms.SomeIntPtr", *ms.SomeIntPtr)
fmt.Println("ms.SomeStringPtr", *ms.SomeStringPtr)
```



```
// assign "value" to "field":
if isPointer {
    x := reflect.New(field.Type().Elem())
    x.Elem().Set(reflect.ValueOf(value))
    field.Set(x)
} else {
    field.Set(reflect.ValueOf(value)) // works
}
```



[Stack](https://stackoverflow.com/questions/7850140/how-do-you-create-a-new-instance-of-a-struct-from-its-type-at-run-time-in-go)
```
package main

import (
    "fmt"
    "reflect"
)

type Config struct {
    Name string
    Meta struct {
        Desc string
        Properties map[string]string
        Users []string
    }
}

func initializeStruct(t reflect.Type, v reflect.Value) {
  for i := 0; i < v.NumField(); i++ {
    f := v.Field(i)
    ft := t.Field(i)
    switch ft.Type.Kind() {
    case reflect.Map:
      f.Set(reflect.MakeMap(ft.Type))
    case reflect.Slice:
      f.Set(reflect.MakeSlice(ft.Type, 0, 0))
    case reflect.Chan:
      f.Set(reflect.MakeChan(ft.Type, 0))
    case reflect.Struct:
      initializeStruct(ft.Type, f)
    case reflect.Ptr:
      fv := reflect.New(ft.Type.Elem())
      initializeStruct(ft.Type.Elem(), fv.Elem())
      f.Set(fv)
    default:
    }
  }
}

func main() {
    t := reflect.TypeOf(Config{})
    v := reflect.New(t)
    initializeStruct(t, v.Elem())
    c := v.Interface().(*Config)
    c.Meta.Properties["color"] = "red" // map was already made!
    c.Meta.Users = append(c.Meta.Users, "srid") // so was the slice.
    fmt.Println(v.Interface())
}

```


