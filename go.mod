module github.com/mshogin/randomtrader

go 1.13

require (
	cloud.google.com/go/storage v1.27.0
	github.com/fastly/go-utils v0.0.0-20180712184237-d95a45783239 // indirect
	github.com/jehiah/go-strftime v0.0.0-20171201141054-1d33003b3869 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lestrrat-go/strftime v1.0.3 // indirect
	github.com/stretchr/testify v1.8.1
	github.com/tebeka/strftime v0.1.5 // indirect
	github.com/thrasher-corp/gocryptotrader v0.0.0-20200512041844-bfab151e9295
	google.golang.org/api v0.103.0
	google.golang.org/grpc v1.53.0
)

replace github.com/thrasher-corp/gocryptotrader v0.0.0-20200512041844-bfab151e9295 => github.com/mshogin/gocryptotrader v0.0.0-20201206003423-342747266688
