module github.com/mshogin/randomtrader

go 1.13

require (
	github.com/golangci/golangci-lint v1.24.0 // indirect
	github.com/gostaticanalysis/analysisutil v0.6.1 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lestrrat-go/strftime v1.0.3 // indirect
	github.com/nishanths/exhaustive v0.1.0 // indirect
	github.com/sonatard/noctx v0.0.1 // indirect
	github.com/stretchr/testify v1.6.1
	github.com/thrasher-corp/gocryptotrader v0.0.0-20200512041844-bfab151e9295
	golang.org/x/sys v0.0.0-20200930185726-fdedc70b468f
	golang.org/x/tools v0.0.0-20201204222352-654352759326 // indirect
	google.golang.org/grpc v1.34.0
)

replace github.com/thrasher-corp/gocryptotrader v0.0.0-20200512041844-bfab151e9295 => github.com/mshogin/gocryptotrader v0.0.0-20201206003423-342747266688
