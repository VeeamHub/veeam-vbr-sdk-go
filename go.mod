module github.com/veeamhub/veeam-vbr-sdk-go/v2

go 1.21.3

require (
	github.com/deepmap/oapi-codegen v1.16.2
	github.com/oapi-codegen/runtime v1.1.0
	github.com/pb33f/libopenapi v0.13.21
)

// Replace until https://github.com/oapi-codegen/runtime/pull/12 will be fixed
replace github.com/gomarkdown/markdown v0.0.0-20230716120725-531d2d74bc12 => github.com/gomarkdown/markdown v0.0.0-20230922105210-14b16010c2ee

require (
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/dprotaso/go-yit v0.0.0-20220510233725-9ba8df137936 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/google/uuid v1.4.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/nxadm/tail v1.4.11 // indirect
	github.com/vmware-labs/yaml-jsonpath v0.3.2 // indirect
	golang.org/x/exp v0.0.0-20231206192017-f3f8817b8deb // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/sync v0.5.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	gopkg.in/check.v1 v1.0.0-20200902074654-038fdea0a05b // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
