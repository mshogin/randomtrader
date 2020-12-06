module github.com/mshogin/randomtrader

go 1.13

require (
	github.com/stretchr/testify v1.6.1
	github.com/thrasher-corp/gocryptotrader v0.0.0-20200512041844-bfab151e9295
	google.golang.org/grpc v1.34.0
)

replace github.com/thrasher-corp/gocryptotrader v0.0.0-20200512041844-bfab151e9295 => github.com/mshogin/gocryptotrader v0.0.0-20201206003423-342747266688
