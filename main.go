package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.ExecPath(`C:\Program Files (x86)\Microsoft\Edge\Application\msedge.exe`), //edge browser  you can use chrome or etc
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-gpu", false),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	url := "#" //group url or user url

	jsSetText := `
		var el = document.querySelector("#editable-message-text");
		if(el){
			el.innerText = "#"; 
			el.dispatchEvent(new Event('input', {bubbles: true}));
		}
	`
// el.innerText = "#";  you can write your text for spam
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(`#editable-message-text`, chromedp.ByQuery),
	)

	if err != nil {
		log.Fatal(err)
	}

	for {

		err := chromedp.Run(ctx,

			chromedp.Evaluate(jsSetText, nil),

			chromedp.Sleep(1*time.Second),

			chromedp.Click(`div[aria-label="send-button"]`, chromedp.ByQuery),
		)

		if err != nil {
			log.Println("error:", err)
		} else {
			log.Println("sended")
		}

		time.Sleep(5 * time.Minute)  //timer for spam
	}
}
