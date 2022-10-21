package req

import (
	"github.com/asmcos/requests"
)

func Waze() (string, error) {
	base_url := "https://www.waze.com/row-rtserver/web/TGeoRSS"
	p := requests.Params{
		"bottom": "-33.62888289909286",
		"top":    "-33.27586339379096",
		"left":   "-71.01147079467775",
		"right":  "-70.38154220581056",
		"ma":     "200",
		"mj":     "200",
		"mu":     "20",
		"types":  "alerts, traffic, users",
	}

	resp, err := requests.Get(base_url, p)

	if err != nil {
		return "", err
	}

	return resp.Text(), nil
}
