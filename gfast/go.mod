module gfast

require (
	github.com/casbin/casbin/v2 v2.31.9
	github.com/ethereum/go-ethereum v1.14.5
	github.com/go-sql-driver/mysql v1.7.0
	github.com/goflyfox/gtoken v1.4.5
	github.com/gogf/gcache-adapter v0.1.2
	github.com/gogf/gf v1.16.6
	github.com/gogf/swagger v1.3.0
	github.com/mervick/aes-everywhere/go/aes256 v0.0.0-20220722115308-bf92805fe805
	github.com/mojocn/base64Captcha v1.3.1
	github.com/mssola/user_agent v0.5.1
	github.com/pquerna/otp v1.4.0
	github.com/shirou/gopsutil v3.21.6+incompatible
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
	github.com/tencentyun/cos-go-sdk-v5 v0.7.27
	github.com/xuri/excelize/v2 v2.8.1
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.10
)

require (
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/Knetic/govaluate v3.0.1-0.20171022003610-9aa49832a739+incompatible // indirect
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/boombuler/barcode v1.0.1 // indirect
	github.com/clbanning/mxj v1.8.5-0.20200714211355-ff02cfb8ea28 // indirect
	github.com/fatih/color v1.16.0 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/gomodule/redigo v2.0.0+incompatible // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/grokify/html-strip-tags-go v0.0.0-20190921062105-daaa06bf1aaf // indirect
	github.com/holiman/uint256 v1.2.4 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/mozillazg/go-httpheader v0.3.0 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/richardlehane/mscfb v1.0.4 // indirect
	github.com/richardlehane/msoleps v1.0.3 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/xuri/efp v0.0.0-20231025114914-d1ff6096ae53 // indirect
	github.com/xuri/nfp v0.0.0-20230919160717-d98342af3f05 // indirect
	go.opentelemetry.io/otel v1.0.0-RC2 // indirect
	go.opentelemetry.io/otel/trace v1.0.0-RC2 // indirect
	golang.org/x/crypto v0.22.0 // indirect
	golang.org/x/image v0.14.0 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/goflyfox/gtoken v1.4.5 => github.com/tiger1103/gtoken v1.4.8

go 1.21

toolchain go1.21.4
