##### protobuf tool to collect and merge all changes!
```go
person := pbgen.NewPersonSync()
person.SetName("proto sync")
person.SetAge(1)
personProto := &pbgen.Person{}
person.MergeDirtyToPb(personProto)
fmt.Println(protojson.Format(personProto))
```
###### output:
```json
{
  "age": 1,
  "name": "proto sync"
}
```