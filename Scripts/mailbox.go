package main
import (
	"log"
	"os"
	"fmt"

	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-imap"
)


func main() {

	const mailServer string = "REPLACE_WITH_SERVER_FQDN:993"

	user := os.Args[1]
	pass := os.Args[2]

	c, _ := client.DialTLS(mailServer, nil)

	if err := c.Login(user, pass); err != nil {
		log.Fatal(err)
	}

	_, _ = c.Select("INBOX", false)

	criteria := imap.NewSearchCriteria()

	criteria.WithoutFlags = []string{"\\Seen"}
	uids, _ := c.Search(criteria)

	fmt.Print(len(uids))

	defer c.Logout()

}