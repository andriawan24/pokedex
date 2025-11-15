package commands

import "fmt"

func getMaps(cfg *config) error {
	locationResponse, err := cfg.client.ListLocation(cfg.nextURL)
	if err != nil {
		return err
	}

	cfg.nextURL = locationResponse.Next
	cfg.prevURL = locationResponse.Previous

	for _, location := range locationResponse.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func getPrevMaps(cfg *config) error {
	if cfg.prevURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	locationResponse, err := cfg.client.ListLocation(cfg.prevURL)
	if err != nil {
		return err
	}

	cfg.nextURL = locationResponse.Next
	cfg.prevURL = locationResponse.Previous

	for _, location := range locationResponse.Results {
		fmt.Println(location.Name)
	}

	return nil
}
