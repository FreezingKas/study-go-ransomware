module encrypter

go 1.17

replace fileutils => ../fileutils

require fileutils v0.0.0-00010101000000-000000000000

require github.com/joho/godotenv v1.4.0

require (
	github.com/cretz/bine v0.2.0 // indirect
	github.com/gen2brain/go-libtor v1.2.0
	golang.org/x/crypto v0.0.0-20220313003712-b769efc7c000 // indirect
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2 // indirect
	golang.org/x/sys v0.0.0-20220310020820-b874c991c1a5 // indirect
)
