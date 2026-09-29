package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

func writeJSONReport(pages map[string]PageData, filename string) error {
	if len(pages) == 0 {
		fmt.Printf("no data to write report")
		return nil
	}

	var keys []string
	for k := range pages {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var sortedData []PageData
	for _, k := range keys {
		sortedData = append(sortedData, pages[k])
	}

	data, err := json.MarshalIndent(sortedData, "", " ")
	if err != nil {
		return fmt.Errorf("couldnt parse data")
	}

	err = os.WriteFile(filename, data, 0o644)
	if err != nil {
		return fmt.Errorf("couldnt write data")
	}
	return nil
}
