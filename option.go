package html2pdf

import (
	"github.com/chromedp/cdproto/page"
	"time"
)

type options struct {
	//ctx context.Context

	// Sleep is an empty action that calls time.Sleep with the specified duration.
	//
	// Note: this is a temporary action definition for convenience, and will likely
	// be marked for deprecation in the future, after the remaining Actions have
	// been able to be written/tested.
	sleep time.Duration

	printToPDFParams *PrintToPDFParams
}

// Option represents the Html2image options
type Option interface {
	Apply(o *options)
}

// OptionFunc is a function that configures a Html2image.
type OptionFunc func(*options)

// Apply calls f(client)
func (f OptionFunc) Apply(o *options) {
	f(o)
}

// WithLandscape paper orientation. Defaults to false.
func WithLandscape(landscape bool) Option {
	return OptionFunc(func(o *options) {
		o.printToPDFParams.Landscape = landscape
	})
}

// WithDisplayHeaderFooter display header and footer. Defaults to false.
func WithDisplayHeaderFooter(displayHeaderFooter bool) Option {
	return OptionFunc(func(o *options) {
		o.printToPDFParams.DisplayHeaderFooter = displayHeaderFooter
	})
}

// WithPrintBackground print background graphics. Defaults to false.
func WithPrintBackground(printBackground bool) Option {
	return OptionFunc(func(o *options) {
		o.printToPDFParams.PrintBackground = printBackground
	})
}

// WithScale scale of the webpage rendering. Defaults to 1.
func WithScale(scale float64) Option {
	return OptionFunc(func(o *options) {
		o.printToPDFParams.Scale = scale
	})
}

// WithPaperWidth paper width in inches. Defaults to 8.5 inches.
func WithPaperWidth(paperWidth float64) Option {
	return OptionFunc(func(o *options) {
		o.printToPDFParams.PaperWidth = paperWidth
	})
}

// WithPaperHeight paper height in inches. Defaults to 11 inches.
func WithPaperHeight(paperHeight float64) Option {
	return OptionFunc(func(o *options) {
		o.printToPDFParams.PaperHeight = paperHeight
	})
}

// WithMarginTop top margin in inches. Defaults to 1cm (~0.4 inches).
func WithMarginTop(marginTop float64) Option {
	return OptionFunc(func(o *options) {
		o.printToPDFParams.MarginTop = marginTop
	})
}

// WithMarginBottom bottom margin in inches. Defaults to 1cm (~0.4 inches).
func WithMarginBottom(marginBottom float64) Option {
	return OptionFunc(func(o *options) {
		o.printToPDFParams.MarginBottom = marginBottom
	})
}

// WithMarginLeft left margin in inches. Defaults to 1cm (~0.4 inches).
func WithMarginLeft(marginLeft float64) Option {
	return OptionFunc(func(o *options) {
		o.printToPDFParams.MarginLeft = marginLeft
	})
}

// WithMarginRight right margin in inches. Defaults to 1cm (~0.4 inches).
func WithMarginRight(marginRight float64) Option {
	return OptionFunc(func(o *options) {
		o.printToPDFParams.MarginRight = marginRight
	})
}

// WithPageRanges paper ranges to print, e.g., '1-5, 8, 11-13'. Defaults to
// the empty string, which means print all pages.
func WithPageRanges(pageRanges string) Option {
	return OptionFunc(func(o *options) {
		o.printToPDFParams.PageRanges = pageRanges
	})
}

// WithIgnoreInvalidPageRanges whether to silently ignore invalid but
// successfully parsed page ranges, such as '3-2'. Defaults to false.
func WithIgnoreInvalidPageRanges(ignoreInvalidPageRanges bool) Option {
	return OptionFunc(func(o *options) {
		o.printToPDFParams.IgnoreInvalidPageRanges = ignoreInvalidPageRanges
	})
}

// WithHeaderTemplate HTML template for the print header. Should be valid
// HTML markup with following classes used to inject printing values into them:
// - date: formatted print date - title: document title - url: document location
// - pageNumber: current page number - totalPages: total pages in the document
// For example, <span class=title></span> would generate span containing the
// title.
func WithHeaderTemplate(headerTemplate string) Option {
	return OptionFunc(func(o *options) {
		o.printToPDFParams.HeaderTemplate = headerTemplate
	})
}

// WithFooterTemplate HTML template for the print footer. Should use the same
// format as the headerTemplate.
func WithFooterTemplate(footerTemplate string) Option {
	return OptionFunc(func(o *options) {
		o.printToPDFParams.FooterTemplate = footerTemplate
	})
}

// WithPreferCSSPageSize whether or not to prefer page size as defined by
// css. Defaults to false, in which case the content will be scaled to fit the
// paper size.
func WithPreferCSSPageSize(preferCSSPageSize bool) Option {
	return OptionFunc(func(o *options) {
		o.printToPDFParams.PreferCSSPageSize = preferCSSPageSize
	})
}

type PrintToPDFTransferMode string

// PrintToPDFTransferMode values.
const (
	PrintToPDFTransferModeReturnAsBase64 PrintToPDFTransferMode = "ReturnAsBase64"
	PrintToPDFTransferModeReturnAsStream PrintToPDFTransferMode = "ReturnAsStream"
)

// WithTransferMode return as stream.
func WithTransferMode(transferMode PrintToPDFTransferMode) Option {
	return OptionFunc(func(o *options) {
		o.printToPDFParams.TransferMode = page.PrintToPDFTransferMode(transferMode)
	})
}
