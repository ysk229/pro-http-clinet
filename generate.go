package main

//win protoc-gen-validate.exe
//linux protoc-gen-validate
//go:generate protoc  --plugin="protoc-gen-validate.exe"  --proto_path=.  --proto_path=./third_party    --go_out=paths=source_relative:./pb   --validate_out=paths=source_relative,lang=go:./pb ./proto/api/weather/v1/*.proto
