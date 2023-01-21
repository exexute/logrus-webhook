package logrusWebhook

import (
	"fmt"
	"github.com/aliyun/aliyun-log-go-sdk/producer"
	"sync"
	"time"
)

type SlsWriter struct {
	Config   *SlsConfig
	Producer *producer.Producer
	mu       sync.Mutex
}

func RegisterSlsWriter(config *SlsConfig) (w *SlsWriter, err error) {
	w = &SlsWriter{Config: config}

	producerConfig := producer.GetDefaultProducerConfig()
	producerConfig.Endpoint = config.EndPoint
	producerConfig.AccessKeyID = config.AccessKeyID
	producerConfig.AccessKeySecret = config.AccessKeySecret

	w.mu.Lock()
	defer w.mu.Unlock()
	producerInstance := producer.InitProducer(producerConfig)
	producerInstance.Start()
	w.Producer = producerInstance
	return
}

func (w *SlsWriter) Write(msg string) error {
	log := producer.GenerateLog(uint32(time.Now().Unix()), map[string]string{"content": msg})
	err := w.Producer.SendLog(w.Config.Project, w.Config.LogStore, w.Config.Topic, "ast-app", log)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
