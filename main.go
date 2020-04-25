package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"gotutorial/util"
	"log"
	"strconv"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	sdStr := flag.String("sd", "20200101", "start date")
	edStr := flag.String("ed", "20200331", "end date")
	headlessOpt := flag.Bool("headless", true, "if start chrome in headless mode; true by default")

	flag.Parse()

	fmt.Println("Start Date:", *sdStr)
	fmt.Println("End Date:", *edStr)
	fmt.Println("Headless Mode:", *headlessOpt)

	ctx := context.Background()

	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", *headlessOpt),
		chromedp.Flag("hide-scrollbars", false),
		chromedp.Flag("mute-audio", false),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}

	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	c, cc := chromedp.NewExecAllocator(ctx, options...)
	defer cc()

	ctx, cancel := chromedp.NewContext(c)
	defer cancel()

	var jsonString string

	baseURL := "https://bet.hkjc.com/marksix/getJSON.aspx?"
	url := fmt.Sprintf("%ssd=%s&ed=%s&sb=0", baseURL, *sdStr, *edStr)

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.InnerHTML("pre", &jsonString, chromedp.ByQuery),
	)

	if err != nil {
		log.Fatal(err)
	}

	var results []map[string]interface{}
	json.Unmarshal([]byte(jsonString), &results)

	var rows [][]string
	rows = util.Parse(results)

	fileName := "sd" + *sdStr + "_" + "ed" +
		*edStr + "_" + strconv.FormatInt(time.Now().Unix(), 10) + ".csv"
	util.Write(fileName, rows)
}
