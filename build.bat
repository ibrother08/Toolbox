@REM windres.exe -i ./winRes/app.rc -o defaultRes_windows_amd64.syso -F pe-x86-64

SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build -ldflags="-s -w -H windowsgui" -tags tempdll -buildmode=exe -o Toolbox.exe

upx -9 Toolbox.exe
pause