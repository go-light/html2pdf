package html2pdf

import (
	"context"
	"io/ioutil"
	"testing"
)

func TestPrintToPDF(t *testing.T) {
	ret, err := PrintToPDF(context.Background(), "https://github.com/go-light/html2pdf")
	if err != nil {
		t.Error(err)
		return
	}

	err = ioutil.WriteFile("test.pdf", ret.Bytes(), 0644)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestWithPrintToPDFParams(t *testing.T) {
	ret, err := PrintToPDF(context.Background(), "https://github.com/go-light/html2pdf",
		WithPaperWidth(20.1),
		WithPaperHeight(20.1),
	)
	if err != nil {
		t.Error(err)
		return
	}

	err = ioutil.WriteFile("test.pdf", ret.Bytes(), 0644)
	if err != nil {
		t.Error(err)
		return
	}
}
