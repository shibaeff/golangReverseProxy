package stats

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

type Stats struct {
	CpuPercent float64 `json:"cpu_percent"`
	MemPercent float64 `json:"mem_percent"`
}

type StatsApi interface {
	GetStats(w http.ResponseWriter, r *http.Request)
}

type statsApi struct {
}

func (s *statsApi) GetStats(w http.ResponseWriter, r *http.Request) {
	m, err := mem.VirtualMemory()
	if err != nil {
		log.Println(err)
	}
	c, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(Stats{
		CpuPercent: c[0],
		MemPercent: m.UsedPercent,
	})

	if err != nil {
		log.Println(err)
	}
}

func NewStatsApi() StatsApi {
	return &statsApi{}
}
