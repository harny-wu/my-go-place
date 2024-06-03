package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/tidwall/gjson"
)

var jsonData = []byte(`{"name": "Alice", "age": 30, "city": "New York", "country": "USA"}`)

func BenchmarkEncodingJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var data map[string]interface{}
		if err := json.Unmarshal(jsonData, &data); err != nil {
			b.Fatal(err)
		}
		_ = data["name"].(string)
		_ = data["age"].(float64)
	}
}

func BenchmarkGJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		name := gjson.GetBytes(jsonData, "name").String()
		age := gjson.GetBytes(jsonData, "age").Int()
		_ = name
		_ = age
	}
}

func main() {
	resultEncodingJSON := testing.Benchmark(BenchmarkEncodingJSON)
	fmt.Printf("encoding/json: %s\n", resultEncodingJSON.String())

	resultGJSON := testing.Benchmark(BenchmarkGJSON)
	fmt.Printf("gjson: %s\n", resultGJSON.String())
}
