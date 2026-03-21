# YTMeteor
YTsaurus data operation API over standard client

This client created for ETL and Analytics tasks.


#### Before start of work You must initialize client

```go
package main
import "github.com/shukhov/ytapi/client"

func main() {
	client, err := client.NewClient("ytsaurus.cluster.ru", "***", "hahn")
	if err != nil {
		log.Fatal(err)
	}
}
```

#### Start YQL query and track it
YQL runner wraps your query into subquery and write it to tmp directory.

```go

```