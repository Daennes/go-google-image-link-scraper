package imagelinkscraper

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/customsearch/v1"
	"google.golang.org/api/option"
)

// ImageLinkScraper ImageLinkScraper struct
type ImageLinkScraper struct {
	apiKey   string
	engineID string
	cse      *customsearch.CseService
}

// New new ImageLinkScraper
func New(APIKey string, EngineID string) *ImageLinkScraper {
	svc, err := customsearch.NewService(context.Background(), option.WithAPIKey(APIKey))
	if err != nil {
		log.Fatal(err)
	}
	return &ImageLinkScraper{
		apiKey:   APIKey,
		engineID: EngineID,
		cse:      svc.Cse,
	}
}

// Query query
func (s *ImageLinkScraper) Query(query string) []string {
	imageUrls := make([]string, 0)
	for i := 0; i < 10; i++ {
		resp, err := s.cse.List(query).Cx(s.engineID).SearchType("image").Start(int64(i*10 + 1)).Do()
		if err != nil {
			log.Fatal(err)
		}

		for _, result := range resp.Items {
			imageUrls = append(imageUrls, result.Link)
		}
		fmt.Printf("Page %d/%d: %d Images\n", i+1, 10, len(resp.Items))
	}
	return imageUrls
}
