package exact_test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestFinancialGLAccountsGet(t *testing.T) {
	params := client.Financial.NewGLAccountsGetParams()
	params.Filter.Set(fmt.Sprintf("Code eq '%s'", "1790"))
	resp, err := client.Financial.GLAccountsGet(params, nil)
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
