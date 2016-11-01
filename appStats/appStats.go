package appStats

import (
    //"fmt"
    "time"
    "sort"
    //"strings"
    "github.com/cloudfoundry/sonde-go/events"
    "github.com/kkellner/cloudfoundry-top-plugin/util"
)

type LogMetric struct {
  OutCount uint64
  ErrCount uint64
}

type dataSlice []*AppStats

type AppStats struct {
  AppUUID     *events.UUID
  AppId       string
  AppName     string
  SpaceName    string
  OrgName      string
  EventCount  uint64

  responseL60Time    *util.AvgTracker
  AvgResponseL60Time float64
  EventL60Rate       int

  responseL10Time    *util.AvgTracker
  AvgResponseL10Time float64
  EventL10Rate       int

  responseL1Time    *util.AvgTracker
  AvgResponseL1Time float64
  EventL1Rate       int

  //EventResTime float64
  //EventTime    int64
  Event2xxCount uint64
  Event3xxCount uint64
  Event4xxCount uint64
  Event5xxCount uint64
  ContainerMetric []*events.ContainerMetric
  LogMetric []*LogMetric
  NonContainerOutCount uint64
  NonContainerErrCount uint64
}


func NewAppStats(appId string) *AppStats {
	stats := &AppStats{}
  stats.AppId = appId
  stats.responseL60Time = util.NewAvgTracker(time.Minute)
  stats.responseL10Time = util.NewAvgTracker(time.Second * 10)
  stats.responseL1Time = util.NewAvgTracker(time.Second)
  return stats

}


// Take the stats map and generated a reverse sorted list base on attribute X
func getStats(statsMap map[string]*AppStats) []*AppStats {
  s := make(dataSlice, 0, len(statsMap))
  for _, d := range statsMap {
      s = append(s, d)
  }
  sort.Sort(sort.Reverse(s))
  return s
}


// Len is part of sort.Interface.
func (d dataSlice) Len() int {
    return len(d)
}

// Swap is part of sort.Interface.
func (d dataSlice) Swap(i, j int) {
    d[i], d[j] = d[j], d[i]
}

// Less is part of sort.Interface. We use count as the value to sort by
func (d dataSlice) Less(i, j int) bool {
    if (d[i].EventCount == d[j].EventCount) {
      return d[i].AppName >= d[j].AppName
    }
    return d[i].EventCount <= d[j].EventCount
}


/*
func mainX() {
    m := map[string]*appStats.AppStats {
        "x": {"x", "x", 0, 0, 0 , 0 ,0 },
        "y": {"y", "y", 9, 0, 0 , 0 ,0 },
        "z": {"z", "z", 7, 0, 0 , 0 ,0 },
        "a": {"z", "a", 5, 0, 0 , 0 ,0 },
        "b": {"z", "b", 3, 0, 0 , 0 ,0 },
        "c": {"z", "c", 10, 0, 0 , 0 ,0 },
        "d": {"z", "d", 1, 0, 0 , 0 ,0 },
        "e": {"z", "e", 15, 0, 0 , 0 ,0 },
    }

    s := make(appStats.DataSlice, 0, len(m))

    for _, d := range m {
        s = append(s, d)
    }

    //sort.Sort(s)
    sort.Sort(sort.Reverse(s))

    for _, d := range s {
        fmt.Printf("%+v\n", *d)
    }
}

*/
