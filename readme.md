# cosmosdemo

## hello

* https://docs.ignite.com/guide/hello/express

```
~/bin/shell/ignite scaffold query say-hello name --response name


#	修改：     docs/static/openapi.yml
#	修改：     proto/cosmosdemo/cosmosdemo/query.proto
#	修改：     x/cosmosdemo/client/cli/query.go
#	新文件：   x/cosmosdemo/client/cli/query_say_hello.go
#	新文件：   x/cosmosdemo/keeper/query_say_hello.go
#	修改：     x/cosmosdemo/types/query.pb.go
#	修改：     x/cosmosdemo/types/query.pb.gw.go

~/bin/shell/ignite chain serve
/Users/hosea/go/bin/cosmosdemod q cosmosdemo say-hello Hosea

```