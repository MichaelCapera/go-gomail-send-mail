# Init module

 go mod init send_email_gomail

 # Install dependencies

 go get gopkg.in/gomail.v2

 ## Configure de server EMAIL

 d := gomail.NewDialer("smtp.gmail.com", 587, "andres@gmail.com", "xxxx")

 ## Run project

 go run main.go