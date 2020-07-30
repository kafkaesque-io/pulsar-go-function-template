# Pulsar Golang Function template

### go.mod
To include this module in any existing go.mod
```
replace github.com/apache/pulsar/pulsar-function-go => github.com/kafkaesque-io/pulsar-go-function-template
```

Or run go get
```
go get github.com/kafkaesque-io/pulsar-go-function-template
```
## Example

An example is given by this file `./src/trigger-only.go`

`./script/build.sh` builds a binary that can be uploaded as Pulsar function.