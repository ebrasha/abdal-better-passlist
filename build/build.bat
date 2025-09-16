@echo off
REM -------------------------------------------------------------------
REM Programmer       : Ebrahim Shafiei (EbraSha)
REM Email            : Prof.Shafiei@Gmail.com
REM -------------------------------------------------------------------
setlocal
set CURRENT_DIR=%CD%
cd /d %CURRENT_DIR%
cd ..

 

REM پاکسازی هر syso باقی‌مانده
if exist rsrc.syso del /f /q rsrc.syso

REM ---------- Windows amd64 ----------
echo [win-amd64] build...
set GOOS=windows
set GOARCH=amd64
set CGO_ENABLED=0
rsrc -arch %GOARCH% -ico icon.ico -o rsrc.syso || goto :fail
timeout /t 3 >nul

go build -ldflags="-s -w" -o abdal-better-passlist-windows-amd64.exe  || goto :fail
del /f /q rsrc.syso
timeout /t 3 >nul

REM ---------- Windows 386 ----------
echo [win-386] build...
set GOOS=windows
set GOARCH=386
set CGO_ENABLED=0
rsrc -arch %GOARCH% -ico icon.ico -o rsrc.syso || goto :fail
timeout /t 3 >nul

go build -ldflags="-s -w" -o abdal-better-passlist-windows-386.exe || goto :fail
del /f /q rsrc.syso
timeout /t 3 >nul


REM ---------- Windows arm64 ----------
echo [win-arm64] build...
set GOOS=windows
set GOARCH=arm64
set CGO_ENABLED=0
rsrc -arch %GOARCH% -ico icon.ico -o rsrc.syso || goto :fail
timeout /t 3 >nul

go build -ldflags="-s -w" -o abdal-better-passlist-windows-arm64.exe  || goto :fail
del /f /q rsrc.syso
timeout /t 3 >nul


echo Building for Linux amd64 ...
set GOOS=linux
set GOARCH=amd64
go build -o abdal-better-passlist-linux-amd64 main.go || goto :fail

echo Building for Linux arm64 ...
set GOOS=linux
set GOARCH=arm64
go build -o abdal-better-passlist-linux-arm64 main.go || goto :fail

echo Building for macOS amd64 ...
set GOOS=darwin
set GOARCH=amd64
go build -o abdal-better-passlist-macos-amd64 main.go || goto :fail

echo Building for macOS arm64 ...
set GOOS=darwin
set GOARCH=arm64
go build -o abdal-better-passlist-macos-arm64 main.go || goto :fail



echo OK.
pause
endlocal
exit /b 0

:fail
echo Build failed.
if exist rsrc.syso del /f /q rsrc.syso
pause
exit /b 1
