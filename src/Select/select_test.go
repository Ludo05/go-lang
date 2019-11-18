package selector

import (
	"fmt"
	"testing"
)

	func TestSelect(t *testing.T) {
		t.Run("Test Select", func(t *testing.T){
			fastUrl := "https://www.google.com"
			slowURL := "https://www.facebook.com"

			got := Race(fastUrl, slowURL)
			want := slowURL

			if got != want {
				fmt.Printf("Got %v wanted %v", got,want)
			}
		})
	}


