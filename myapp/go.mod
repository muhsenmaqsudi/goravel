module myapp

go 1.21.5

replace github.com/muhsenmaqsudi/goravel => ../goravel

require github.com/muhsenmaqsudi/goravel v0.0.0-00010101000000-000000000000

require (
	github.com/go-chi/chi/v5 v5.0.11 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
)
