package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/sclevine/agouti"
)

func main() {
	driver := agouti.ChromeDriver(agouti.Browser("chrome"))
	if err := driver.Start(); err != nil {
		log.Fatalf("failed to start driver:%v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage()
	if err != nil {
		log.Fatalf("failed to open driver:%v", err)
	}

	if err := page.Navigate("https://work.e-typing.ne.jp/e-typing_pro/user/"); err != nil {
		log.Fatalf("failed to navigate:%v", err)
	}

	id := page.FindByID("user_id")
	pass := page.FindByID("password")
	err = godotenv.Load("login.env")
	if err != nil {
		log.Fatalf("failed to loading .env file:%v", err)
	}
	id.Fill(os.Getenv("ID"))
	pass.Fill(os.Getenv("PASS"))
	if err := page.FindByID("login_btn").Submit(); err != nil {
		log.Fatalf("failed to login:%v", err)
	}

	if err := page.FindByID("skilcheck_btn").Click(); err != nil {
		log.Fatalf("failed to click:%v", err)
	}

	time.Sleep(time.Second)

	if err := page.FindByID("typing_app_frame").SwitchToFrame(); err != nil {
		log.Fatalf("failed to swich frame:%v", err)
	}
	if err := page.FindByXPath("//*[@id=\"start_btn\"]").Click(); err != nil {
		log.Fatalf("failed to click:%v", err)
	}

	time.Sleep(time.Second)

	if err := page.FindByXPath("/html/body").SendKeys(" "); err != nil {
		log.Fatalf("failed to send space%v", err)
	}

	time.Sleep(time.Second * 3)
}
