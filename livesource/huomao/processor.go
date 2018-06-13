package huomao

import (
	"github.com/W1llyu/ourcrawler/collector"
	"github.com/W1llyu/ourcrawler/downloader"
	"github.com/W1llyu/ourcrawler/livesource"
	"github.com/W1llyu/ourjson"
)

type Processor struct{}

func (p *Processor) Process(response *downloader.Response) (collector.ResultItems, error) {
	resultItem := collector.ResultItems{
		Request: response.Request,
	}
	var liveRooms []livesource.LiveRoom
	json, err := ourjson.ParseObject(response.Body)
	if err != nil {
		return resultItem, err
	}
	for _, val := range json.GetJsonObject("data").GetJsonArray("channelList").Values() {
		liveRoom := processEach(val.JsonObject())
		if liveRoom.AudienceCount != 0 {
			liveRooms = append(liveRooms, liveRoom)
		}
	}
	livesource.PackResultItem(&resultItem, liveRooms, "", "火猫TV")
	return resultItem, nil
}

func processEach(val *ourjson.JsonObject) livesource.LiveRoom {
	liveRoom := livesource.LiveRoom{}
	liveRoom.Title, _ = val.GetString("channel")
	liveRoom.Anchor, _ = val.GetString("nickname")
	liveRoom.AudienceCount, _ = val.GetInt("originviews")
	return liveRoom
}
