package pkg_test

import (
	"testing"

	"github.com/mserebryaakov/course-golang/BinanceRequest/pkg"
	"github.com/mserebryaakov/course-golang/BinanceRequest/postgresql"
)

func TestRequest(t *testing.T) {
	st := postgresql.New()

	err := st.Open()
	if err != nil {
		t.Fatal("Error open database")
	}

	defer st.Close()

	data, err := pkg.RunGet()
	if err != nil {
		t.Fatal("Error INSERT")
	}
	for _, value := range *data {
		st.InsertBinanceData(value.Symbol, value.Price)
	}
}
