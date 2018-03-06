for updating the protobuff run the following command

```bash
protoc -I. ./*.proto --go_out=plugins=grpc,:.
```
