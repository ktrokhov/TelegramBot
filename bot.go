package main
import (
	"github.com/yanzay/tbot"
	"log"
	"os"
)
func main() {
	bot, err := tbot.NewServer(os.Getenv("1083637766:AAFUspaWiqsy0tdibeCMVOXhriqwFRanHcU"))
	if err != nil {
		log.Fatal(err)
	}
	bot.Handle("/answer", "42")
	bot.ListenAndServe()
}