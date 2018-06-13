package panda

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/W1llyu/ourcrawler/collector"
	"github.com/W1llyu/ourcrawler/downloader"
	"github.com/W1llyu/ourcrawler/livesource"
	"github.com/W1llyu/ourcrawler/util"
	"strings"
)

type Processor struct{}

func (p *Processor) Process(response *downloader.Response) (collector.ResultItems, error) {
	resultItem := collector.ResultItems{
		Request: response.Request,
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(response.Body))
	if err != nil {
		return resultItem, err
	}
	var liveRooms []livesource.LiveRoom
	doc.Find("li.video-list-item").Each(func(i int, s *goquery.Selection) {
		liveRoom := processEach(s)
		if liveRoom.AudienceCount != 0 {
			liveRooms = append(liveRooms, liveRoom)
		}
	})
	livesource.PackResultItem(&resultItem, liveRooms, strings.TrimSpace(doc.Find("h3.header-title-secondname").Text()), "熊猫TV")
	return resultItem, nil
}

func processEach(s *goquery.Selection) livesource.LiveRoom {
	liveRoom := livesource.LiveRoom{}
	liveRoom.Title = strings.TrimSpace(s.Find("span.video-title").Text())
	liveRoom.Anchor = strings.TrimSpace(s.Find("span.video-nickname").Text())
	liveRoom.AudienceCount = util.ZhNumToEn(strings.TrimSpace(s.Find("span.video-number").Text()))
	return liveRoom
}
