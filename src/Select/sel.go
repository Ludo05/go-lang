package selector

import (
	"net/http"
	"time"
)

func Race(url, urlTwo string)  string {
	startUrl := time.Now()
	http.Get(url)
	endUrl := time.Since(startUrl)

	startUrlTwo := time.Now()
	http.Get(urlTwo)
	endUrlTwo := time.Since(startUrlTwo)

	if endUrl < endUrlTwo {
		return url
	}
	return urlTwo
}
