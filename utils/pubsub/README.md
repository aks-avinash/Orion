# pubsub
`import "github.com/carousell/Orion/utils/pubsub"`

* [Overview](#pkg-overview)
* [Imported Packages](#pkg-imports)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>

## <a name="pkg-imports">Imported Packages</a>

- [cloud.google.com/go/pubsub](https://godoc.org/cloud.google.com/go/pubsub)
- [github.com/afex/hystrix-go/hystrix](https://godoc.org/github.com/afex/hystrix-go/hystrix)
- [github.com/carousell/Orion/utils/executor](./../executor)
- [github.com/carousell/Orion/utils/log](./../log)
- [github.com/carousell/Orion/utils/pubsub/message_queue](./message_queue)
- [github.com/carousell/Orion/utils/spanutils](./../spanutils)

## <a name="pkg-index">Index</a>
* [type Config](#Config)
* [type Service](#Service)
  * [func NewPubSubService(config Config) Service](#NewPubSubService)

#### <a name="pkg-files">Package files</a>
[pubsub.go](./pubsub.go) 

## <a name="Config">type</a> [Config](./pubsub.go#L16-L23)
``` go
type Config struct {
    Key                    string
    Project                string
    Enabled                bool
    Timeout                int
    BulkPublishConcurrency int
    Retries                int
}

```
Config is the config for pubsub

## <a name="Service">type</a> [Service](./pubsub.go#L26-L31)
``` go
type Service interface {
    PublishMessage(ctx context.Context, topic string, data []byte, waitSync bool) (*goPubSub.PublishResult, error)
    BulkPublishMessages(ctx context.Context, topic string, data [][]byte, waitSync bool)
    SubscribeMessages(ctx context.Context, subscribe string, subscribeFunction messageQueue.SubscribeFunction) error
    Close()
}
```
Service is the interface implemented by a pubsub service

### <a name="NewPubSubService">func</a> [NewPubSubService](./pubsub.go#L42)
``` go
func NewPubSubService(config Config) Service
```
NewPubSubService build and returns an pubsub service handler

- - -
Generated by [godoc2ghmd](https://github.com/GandalfUK/godoc2ghmd)