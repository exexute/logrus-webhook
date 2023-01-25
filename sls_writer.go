package logrusWebhook

import (
	"fmt"
	sls "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/aliyun/aliyun-log-go-sdk/producer"
	"github.com/gogo/protobuf/proto"
	"sync"
	"time"
)

type SlsWriter struct {
	Config       *SlsConfig
	Client       sls.ClientInterface
	BatchSize    int
	SupportBatch bool
	Logs         []*sls.Log
	mu           sync.Mutex
}

func RegisterSlsWriter(config *SlsConfig) (w *SlsWriter, err error) {
	w = &SlsWriter{Config: config}

	w.mu.Lock()
	defer w.mu.Unlock()
	w.BatchSize = config.BatchSize
	w.SupportBatch = config.BatchSize > 0
	w.Client = sls.CreateNormalInterface(config.EndPoint, config.AccessKeyID, config.AccessKeySecret, "")
	return
}

func (w *SlsWriter) Write(msg string) error {
	w.mu.Lock()
	log := producer.GenerateLog(uint32(time.Now().Unix()), map[string]string{"content": msg})
	w.Logs = append(w.Logs, log)
	if !w.SupportBatch || len(w.Logs) == w.BatchSize {
		logGroup := &sls.LogGroup{
			Source: proto.String("127.0.0.1"),
			Logs:   w.Logs,
		}

		// PutLogs API Ref: https://intl.aliyun.com/help/doc-detail/29026.htm
		err := w.Client.PutLogs(w.Config.Project, w.Config.LogStore, logGroup)
		if err == nil {
			fmt.Println("PutLogs success")
		} else {
			fmt.Printf("PutLogs fail, err: %s\n", err)
		}

		w.Logs = nil
	}
	defer w.mu.Unlock()

	return nil
}
