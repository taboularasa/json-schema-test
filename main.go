package main

import(
  "os"
  "fmt"
  "strings"
  "github.com/santhosh-tekuri/jsonschema"
  _ "github.com/santhosh-tekuri/jsonschema/httploader"
)

func main() {
  data := `
  {
      "title": "Person",
      "type": "object",
      "properties": {
          "firstName": {
              "type": "string"
          },
          "lastName": {
              "type": "string"
          },
          "age": {
              "description": "Age in years",
              "type": "integer",
              "minimum": 0
          }
      },
      "required": ["firstName", "lastName"]
  }
  `
  url := "sch.json"
  compiler := jsonschema.NewCompiler()
  if err := compiler.AddResource(url, strings.NewReader(data)); err != nil {
      fmt.Println(err)
  }

  schema, err := compiler.Compile(url)
  if err != nil {
    fmt.Println(err)
  }

  f, err := os.Open("doc.json")
  if err != nil {
    fmt.Println(err)
  }
  defer f.Close()

  if err = schema.Validate(f); err != nil {
    fmt.Println(err)
  }

  fmt.Println("hahahah")
}
