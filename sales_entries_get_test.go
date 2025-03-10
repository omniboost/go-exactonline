package exact_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestSalesEntriesGet(t *testing.T) {
	params := client.SalesEntry.NewSalesEntriesGetParams()
	params.Filter.Set("EntryNumber eq 2100")
	// params.Filter.Set("YourRef eq 'AMSVO-2100'")
	resp, err := client.SalesEntry.SalesEntriesGet(params, nil)
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
