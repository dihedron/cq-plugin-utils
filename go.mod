module github.com/dihedron/cq-plugin-utils

go 1.22.0

require (
	github.com/apache/arrow/go/v15 v15.0.0
	github.com/cloudquery/plugin-sdk/v4 v4.32.0
	github.com/gobwas/glob v0.2.3
	github.com/rs/zerolog v1.32.0
	github.com/thoas/go-funk v0.9.3
)

require (
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/google/flatbuffers v23.5.26+incompatible // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1
	github.com/klauspost/cpuid/v2 v2.2.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/zeebo/xxh3 v1.0.2 // indirect
	golang.org/x/exp v0.0.0-20240222234643-814bf88cf225 // indirect
	golang.org/x/mod v0.16.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/tools v0.19.0 // indirect
	golang.org/x/xerrors v0.0.0-20231012003039-104605ab7028 // indirect
)

replace github.com/thoas/go-funk v0.9.3 => github.com/dihedron/go-funk v0.0.0-20230503154649-f530b38601cc
