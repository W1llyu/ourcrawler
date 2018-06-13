package processor

import (
	"github.com/W1llyu/ourcrawler/downloader"
	"github.com/W1llyu/ourcrawler/collector"
)

type Processor interface {
	Process(response *downloader.Response) (collector.ResultItems, error)
}
