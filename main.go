package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	ip2location "github.com/ip2location/ip2location-go"
	"github.com/valyala/fasthttp"
)

func main() {
	port := os.Getenv("PORT")

	log.Printf("starting up [port: %s]\n", port)

	h := requestHandler
	if err := fasthttp.ListenAndServe(fmt.Sprintf(":%s", port), h); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}

func ip6ToCountry(ip string) (string, error) {
	db, err := ip2location.OpenDB("./data/IP2LOCATION-LITE-DB1.IPV6.BIN")
	if err != nil {
		return "", err
	}

	defer db.Close()

	res, err := db.Get_country_short(ip)
	if err == nil {
		return res.Country_short, nil
	}

	return "", errors.New("country not found")
}

func ip4ToCountry(ip string) (string, error) {
	db, err := ip2location.OpenDB("./data/IP2LOCATION-LITE-DB1.BIN")
	if err != nil {
		return "", err
	}

	defer db.Close()

	res, err := db.Get_country_short(ip)
	if err == nil {
		return res.Country_short, nil
	}

	return "", errors.New("country not found")
}

func requestHandler(ctx *fasthttp.RequestCtx) {

	// fmt.Fprintf(ctx, "Requested path is %q\n", ctx.Path())

	path := ctx.Path()
	country := ""

	if len(path) > 1 {
		path = path[1 : len(path)-1]
		if strings.Count(string(path), ".") == 4 {
			country, _ = ip4ToCountry(string(path))
		} else {
			country, _ = ip6ToCountry(string(path))
		}

		if len(country) > 2 {
			country = ""
		}
	}

	fmt.Fprintf(ctx, country)
}
