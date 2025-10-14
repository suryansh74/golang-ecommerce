server:
	powershell -Command '$$env:APP_ENV = "dev"; go run .'
