package translations

import (
	_ "golang.org/x/text/message"
)

//go:generate gotext -srclang=en-GB update -out=catalog.go -lang=en-GB,ja-JP github.com/mrtnhwttktc/kinto-cli/cmd/cli github.com/mrtnhwttktc/kinto-cli/cmd/utils
