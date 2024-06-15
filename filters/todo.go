package filters

import "github.com/rapidstellar/pagi"

type FTodo struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
	pagi.CommonTimeFilters
}
