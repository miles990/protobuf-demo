protoc.exe --csharp_out=. protobuf.proto
protoc.exe --go_out=. protobuf.proto

move "Protobuf.cs" "../unity/Assets"
move "protobuf.pb.go" "../server/protobuf"


pause