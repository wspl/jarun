rsrc -manifest jarun.manifest -o rsrc.syso
go build -ldflags="-H windowsgui -s -w"
upx393w\upx.exe -9 jarun.exe