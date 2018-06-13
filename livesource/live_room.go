package livesource

import (
	"sort"
	"github.com/W1llyu/ourcrawler/collector"
	"time"
)

type LiveRoom struct {
	Title string `json:"title"`
	Anchor string `json:"anchor"`
	AudienceCount int `json:"audience_count"`
}

type By func(l1, l2 *LiveRoom) bool

func (by By) Sort(liveRooms []LiveRoom) {
	ps := &LiveRoomSorter{
		liveRooms: liveRooms,
		by:      by,
	}
	sort.Sort(ps)
}

type LiveRoomSorter struct {
	liveRooms []LiveRoom
	by      func(l1, l2 *LiveRoom) bool
}

func (s *LiveRoomSorter) Len() int {
	return len(s.liveRooms)
}

func (s *LiveRoomSorter) Swap(i, j int) {
	s.liveRooms[i], s.liveRooms[j] = s.liveRooms[j], s.liveRooms[i]
}

func (s *LiveRoomSorter) Less(i, j int) bool {
	return s.by(&s.liveRooms[i], &s.liveRooms[j])
}

func PackResultItem(resultItem *collector.ResultItems, liveRooms []LiveRoom, game, source string) {
	sortFunc := func(l1, l2 *LiveRoom) bool {
		return l1.AudienceCount > l2.AudienceCount
	}
	By(sortFunc).Sort(liveRooms)
	if 10 < len(liveRooms) {
		liveRooms = liveRooms[0: 10]
	}
	totalCount := 0
	for _, liveRoom := range liveRooms {
		totalCount += liveRoom.AudienceCount
	}
	resultItem.Put("source", source)
	resultItem.Put("game", game)
	resultItem.Put("time", time.Now().Format("2006-01-02 15:04:05"))
	resultItem.Put("total_count", totalCount)
	resultItem.Put("list", liveRooms)
}