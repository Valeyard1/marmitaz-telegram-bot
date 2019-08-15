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
		log.Printf("Link found: %s\n", e.Attr("href"))

		// Testing with another site because this one is always open, so that I can work on that
		if e.Text == "Pedidos Açaí" || e.Attr("href") == "pedidos_acai.php" {
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
