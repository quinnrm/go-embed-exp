package embedData

import "embed"

//go:embed hello/hello.txt
//go:embed data
var EFS embed.FS
