package sub

import (
	"encoding/json"
	"runtime"

	"github.com/nats-io/nats.go"
	"github.com/ubulllka/wbL0/internal/cache"
	"github.com/ubulllka/wbL0/internal/config"
	"github.com/ubulllka/wbL0/internal/db"
	"github.com/ubulllka/wbL0/internal/logger"
	"github.com/ubulllka/wbL0/internal/models"
)

func Connect() {
	confNuts := config.CONFIG.Nats
	logg := logger.GetLogger()

	nc, _ := nats.Connect(confNuts.URL)
	js, err := nc.JetStream()

	if err != nil {
		logg.Fatal(err)
	}
	
	js.Subscribe(confNuts.StreamSubjects, func(msg *nats.Msg) {
		msg.Ack()
		var order models.Order
		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			logg.Fatal(err)
		}
		logg.Infof("%s service subscribes from subject:%s\n", confNuts.Subscriber, msg.Subject)
		logg.Infof("OrderID:%s", order.ID)

		db.AddDbOrder(order)
		cache.AddCacheOrder(order)

	}, nats.Durable(confNuts.Subscriber), nats.ManualAck())

	logg.Info("Subscribe successfully")
	runtime.Goexit()

}
