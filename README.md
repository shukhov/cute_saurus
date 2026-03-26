# CuteSaurus
YTsaurus data operation API over standard client

This client created for ETL and Analytics tasks.


#### Before start of work You must initialize client

```go
package main
import "github.com/shukhov/cute_saurus/client"
import "github.com/shukhov/cute_saurus/tools/selector"

func main() {
	_, err := client.NewClient("ytsaurus.cluster.ru", "***", "hahn")
	if err != nil {
		log.Fatal(err)
	}
}
```

#### Start YQL query and track it
YQL runner wraps your query into subquery and write it to tmp directory.

```go
package main

import (
	"github.com/shukhov/cute_saurus/client"
	"github.com/shukhov/cute_saurus/tools/mem_storage"
	"github.com/shukhov/cute_saurus/tools/selector"
    "log"
)

func main() {
	cl, err := client.NewClient("ytsaurus.cluster.ru", "***", "hahn")
	if err != nil {
		log.Fatal(err)
	}

	sel, err := selector.NewSelect(cl, "SELECT * `//home/ods/market/data/clients/2026-03-05`", &ctx)
	if err != nil {
		log.Fatal(err)
	}
	_, err = sel.WaitStatus(true)
	if err != nil {
		log.Fatal(err)
	}

	res, err := sel.Result()
	if err != nil {
		log.Fatal(res)
	}

	_, err = res.ReadTableJson("market_clients")
	if err != nil {
		log.Fatalf("err: %+v", err)
	}

	tbl, err = mem_storage.ConnectTable("market_clients")
	if err != nil {
		log.Fatalf("err: %+v", err)
	}
}
```