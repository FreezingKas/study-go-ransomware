module attacker

go 1.17

require github.com/cretz/bine v0.2.0

require (
	github.com/ebfe/keccak v0.0.0-20150115210727-5cc570678d1b // indirect
	github.com/gen2brain/go-libtor v1.2.0 // indirect
	github.com/paxos-bankchain/moneroutil v0.0.0-20170611151923-33d7e0c11a62 // indirect
	golang.org/x/crypto v0.0.0-20220313003712-b769efc7c000 // indirect
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2 // indirect
	golang.org/x/sys v0.0.0-20220310020820-b874c991c1a5 // indirect
)

replace monero => ../monero

require (
	berty.tech/go-libtor v1.0.385
	github.com/mattn/go-sqlite3 v1.14.12
	monero v0.0.0-00010101000000-000000000000
)
