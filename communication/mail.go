package mail

import (
	"fmt"
	"log"

	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func SendMailConfirmRegister(userMail string) (bool) {

    godotenv.Load()
    htmlBody := `
    <html>
        <body>
            <p>Clique no link e confirme o registro <a href="https://www.exemplo.com">confirmar</a>.</p>
        </body>
    </html>s
    `
    email := gomail.NewMessage()
    email.SetHeader("From", os.Getenv("MAIL_USER"))
    email.SetHeader("To", userMail)
    email.SetHeader("Subject", "Confirmacao de Registro")
    email.SetBody("text/html", htmlBody)
    mailPort, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))

    d := gomail.NewDialer(os.Getenv("MAIL_HOST"), mailPort, os.Getenv("MAIL_USER"), os.Getenv("MAIL_PASS"))

    if err := d.DialAndSend(email); err != nil {
        log.Fatal(err)
        fmt.Printf("error %s", err.Error())
        return false
    }

    return true
}