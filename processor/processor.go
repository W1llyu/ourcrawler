package processor

import (
	"github.com/W1llyu/ourcrawler/collector"
	"github.com/W1llyu/ourcrawler/downloader"
)

type Processor interface {
	Process(response *downloader.Response) (collector.ResultItems, error)
}
