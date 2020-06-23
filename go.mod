module github.com/mshogin/randomtrader

go 1.13

require (
	github.com/stretchr/testify v1.6.1
	github.com/thrasher-corp/gocryptotrader v0.0.0-20200512041844-bfab151e9295
	google.golang.org/grpc v1.33.2
)

// replace github.com/thrasher-corp/gocryptotrader v0.0.0-20200416225906-c0d2ac5e51f5 => /home/mshogin/go/src/github.com/thrasher-corp/gocryptotrader

replace github.com/thrasher-corp/gocryptotrader v0.0.0-20200512041844-bfab151e9295 => /home/mshogin/gocryptotrader
