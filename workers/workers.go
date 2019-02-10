package workers

import (
	"fmt"

	"github.com/olegsobchuk/go-health/configs"
	"github.com/olegsobchuk/go-health/services/checker"
)

// StartProcesses generates gorutines to check sources
func StartProcesses() {
	availableSources := configs.KVClient.SMembers(configs.AvailableSources).Val()
	for _, ID := range availableSources {
		go func(ID string) {
			url := configs.KVClient.LRange(ID, 1, -1).Val()[0]
			rez, err := checker.Check(url)
			fmt.Println(ID, url, rez, err)
		}(ID)
	}
}
