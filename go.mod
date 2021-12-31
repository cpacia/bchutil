module github.com/jchavannes/bchutil

replace (
	github.com/jchavannes/btcd => ../btcd
	github.com/jchavannes/btcutil => ../btcutil
)

go 1.16

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/jchavannes/btcd v0.0.0-20211231102001-e1d2e062e663
	github.com/jchavannes/btcutil v1.0.3-0.20211231102310-a1560bb282e9
	golang.org/x/crypto v0.0.0-20211215153901-e495a2d5b3d3
)
