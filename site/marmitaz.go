package site

import (
	"log"
	"regexp"

	"github.com/gocolly/colly"
)

// TemperoDeMaeIsOpen returns a bool if the restaurant is open or not
func TemperoDeMaeIsOpen() (bool, error) {

	exist := false

	c := colly.NewCollector(
		colly.URLFilters(
			// Visit only urls which belongs to the site
			regexp.MustCompile("https://marmitaz\\.pushsistemas\\.com\\.br/.*"),
		),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if e.Text == "Tempero de Mãe" || e.Attr("href") == "cardapio_mae.php?r=Tempero de Mãe" {
			exist = true
		}
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c.Visit("https://marmitaz.pushsistemas.com.br/")
	if err != nil {
		return false, err
	}

	return exist, nil
}
