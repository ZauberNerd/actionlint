package main

import (
	"fmt"
	"io/ioutil"
	"reflect"

	"github.com/rhysd/actionlint"
)

func toMap(input interface{}) map[string]interface{} {
  out := make(map[string]interface{})
  value := reflect.ValueOf(input)
  if value.Kind() == reflect.Ptr {
    value = value.Elem()
  }

  for i := 0; i < value.NumField(); i++ {
    out[value.Type().Field(i).Name] = value.Field(i).Interface()
  }

  return out
}

var input []byte
//export prepareInput
func prepareInput(len int) *byte {
        fmt.Println("prepareInput", len)
        input = make([]byte, len)
        fmt.Println("return", &input[0])
        return &input[0]
}

//export runActionlint
func runActionlint() error {
        opts := actionlint.LinterOptions{}
        linter, err := actionlint.NewLinter(ioutil.Discard, &opts)
        if err != nil {
                return err
        }

        errs, err := linter.Lint("test.yml", input, nil)
        if err != nil {
                return err
        }

        fmt.Println("errors:", len(errs))
        for _, err := range errs {
                fmt.Println(err)
        }

        return nil
        // ret := make([]interface{}, 0, len(errs))
        // for _, err := range errs {
        //         ret = append(ret, toMap(*err))
        // }

        // return ret, nil
}

func main() {
        fmt.Println("Hello from wasm")
}
