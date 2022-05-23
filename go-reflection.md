

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


