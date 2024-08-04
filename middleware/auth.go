package middleware

var ActiveUser string

func Auth(username string) {
	ActiveUser = username
}
