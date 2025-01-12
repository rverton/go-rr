package frontend

import (
	_ "embed"
)

//go:embed build/client/index.html
var IndexHTML string
