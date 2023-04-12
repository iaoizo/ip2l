# ip解析
```
ip2location
```

# Install
```
go get  github.com/iaoizo/ip2lo@v1.0.2
```

# Usage
```
ip2l := ip2lo.NewIp2L()
err := ip2l.LoadDbFile("IPV6-COUNTRY-REGION-CITY-LATITUDE-LONGITUDE-ISP.BIN")
parse, err := ip2l.Parse("127.0.0.1")
fmt.Printf("%#v", parse)

//ip2l.Result{Country_long:"China", Country_short:"CN", City:"Chengdu", Latitude:30.66667, Longitude:104.06667, Isp:"ChinaNet Sichuan Province Network"}
```