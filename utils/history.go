package utils

import "time"

type HistoryRequest struct {
	IdBreakfast   string
	IdLunch       string
	IdDinner      string
	TotalProtein  int
	TotalKalori   int
	TanggalDibuat time.Time
}
