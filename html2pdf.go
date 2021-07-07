package html2pdf

import (
	"context"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

type PrintToPDFParams struct {
	*page.PrintToPDFParams
}

//NewDefaultPrintToPDFParams default pdf params
func NewDefaultPrintToPDFParams() *PrintToPDFParams {
	return &PrintToPDFParams{
		&page.PrintToPDFParams{
			Landscape:               DefaultLandscape,
			DisplayHeaderFooter:     DefaultDisplayHeaderFooter,
			PrintBackground:         DefaultPrintBackground,
			Scale:                   DefaultScale,
			PaperWidth:              DefaultPaperWidth,
			PaperHeight:             DefaultPaperHeight,
			MarginTop:               DefaultMarginTop,
			MarginBottom:            DefaultMarginBottom,
			MarginLeft:              DefaultMarginLeft,
			MarginRight:             DefaultMarginRight,
			PageRanges:              DefaultPageRanges,
			IgnoreInvalidPageRanges: DefaultIgnoreInvalidPageRanges,
			HeaderTemplate:          DefaultHeaderTemplate,
			FooterTemplate:          DefaultFooterTemplate,
			PreferCSSPageSize:       DefaultPreferCSSPageSize,
		},
	}
}

type PDFer interface {
	Bytes() []byte
	GetConvertElapsed() time.Duration
}

type pdf struct {
	convertElapsed time.Duration
	buf            []byte
}

func PrintToPDF(ctx context.Context, url string, opts ...Option) (PDFer, error) {
	reply := &pdf{}

	start := time.Now()
	defer func() {
		reply.convertElapsed = time.Since(start)
	}()

	o := &options{
		printToPDFParams: NewDefaultPrintToPDFParams(),
	}

	for _, opt := range opts {
		opt.Apply(o)
	}

	ctx, cancel := chromedp.NewContext(ctx)
	defer cancel()

	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(time.Duration(o.sleep)*time.Millisecond),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			reply.buf, _, err = o.printToPDFParams.Do(ctx)
			return err
		}),
	); err != nil {
		return nil, err
	}

	return reply, nil
}

func (p *pdf) Bytes() []byte {
	return p.buf
}

func (p *pdf) GetConvertElapsed() time.Duration {
	return p.convertElapsed
}
