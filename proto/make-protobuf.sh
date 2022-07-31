
for x in **/*.proto; do
  echo $x
  protoc --go-grpc_out=paths=source_relative:. --go_out=paths=source_relative:. $x;
done