package stats

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_statsApi_GetStats(t *testing.T) {
	api := NewStatsApi()
	req, err := http.NewRequest(http.MethodGet, "/stats", nil)
	assert.NoError(t, err)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.GetStats)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	body := rr.Body.Bytes()
	stats := &Stats{}
	err = json.Unmarshal(body, stats)
	assert.NoError(t, err)
	assert.NotZero(t, stats.CpuPercent)
	assert.NotZero(t, stats.MemPercent)
}
