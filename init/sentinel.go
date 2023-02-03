package init

import (
	"log"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/config"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/alibaba/sentinel-golang/logging"
)

func init() {
	//初始化 sentinel
	conf := config.NewDefaultConfig()
	conf.Sentinel.Log.Logger = logging.NewConsoleLogger("cb-integration-normal")
	conf.Sentinel.Log.Dir = "."
	err := sentinel.InitWithConfig(conf)
	if err != nil {
		log.Fatal("err", err)
	}

	//go语言实现暂未发现appLimit即限制客户端每秒请求次数
	if _, err := flow.LoadRules([]*flow.Rule{
		{
			Resource:               "GET:/welcome",
			MetricType:             flow.QPS,
			Count:                  1,
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject,
		},
	}); err != nil {
		log.Fatalf("Unexpected error: %+v", err)
		return
	}

}
