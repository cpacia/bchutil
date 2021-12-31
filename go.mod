module github.com/jchavannes/bchutil

replace (
	github.com/jchavannes/btcd => ../btcd
	github.com/jchavannes/btclog => ../btclog
	github.com/jchavannes/btcutil => ../btcutil
)

go 1.16

require (
	github.com/btcsuite/btcd v0.22.0-beta
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce
	github.com/davecgh/go-spew v1.1.1
	github.com/jchavannes/btcd v0.0.0-20211229221630-1e34441017f4
	golang.org/x/crypto v0.0.0-20211215153901-e495a2d5b3d3
)
