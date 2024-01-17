module github.com/thalescpl-io/k8s-kms-plugin

go 1.21

// TODO replace packages :
//   - gose
//   - crypto11
require (
	github.com/ThalesIgnite/crypto11 v1.2.3
	github.com/ThalesIgnite/gose v0.8.7
	github.com/golang/protobuf v1.4.2
	github.com/google/uuid v1.1.2
	github.com/grpc-ecosystem/grpc-gateway v1.15.2
	github.com/infobloxopen/atlas-app-toolkit v0.22.1
	github.com/keepeye/logrus-filename v0.0.0-20190711075016-ce01a4391dd1
	github.com/miekg/pkcs11 v1.0.3-0.20190429190417-a667d056470f
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/protoc-gen-go-json v0.0.0-20200917194518-364b693410ae
	github.com/sirupsen/logrus v1.9.3
	github.com/spf13/cobra v1.7.0
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.8.4
	golang.org/x/sync v0.4.0
	golang.org/x/tools v0.14.0
	google.golang.org/grpc v1.33.0
	google.golang.org/protobuf v1.25.0
	k8s.io/apiserver v0.19.2
)

require (
	github.com/cilium/ebpf v0.11.0 // indirect
	github.com/cosiner/argv v0.1.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/derekparker/trie v0.0.0-20230829180723-39f4de51ef7d // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-delve/delve v1.22.0 // indirect
	github.com/go-delve/gore v0.11.6 // indirect
	github.com/go-delve/liner v1.2.3-0.20220127212407-d32d89dd2a5d // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b // indirect
	github.com/google/go-dap v0.11.0 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mitchellh/mapstructure v1.1.2 // indirect
	github.com/pelletier/go-toml v1.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cast v1.3.0 // indirect
	github.com/spf13/jwalterweatherman v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/thales-e-security/pool v0.0.2 // indirect
	go.starlark.net v0.0.0-20231101134539-556fd59b42f6 // indirect
	golang.org/x/arch v0.6.0 // indirect
	golang.org/x/exp v0.0.0-20230224173230-c95f2b4c22f2 // indirect
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/net v0.16.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	gopkg.in/ini.v1 v1.51.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/ThalesIgnite/gose v0.8.7 => github.com/IceManGreen/gose v0.0.0-20210519194136-6d77caabfe3c

replace github.com/ThalesIgnite/crypto11 v1.2.3 => github.com/IceManGreen/crypto11 v0.0.0-20210909225015-a81014c7c410
