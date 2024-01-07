.PHONY: build clean

build: clean
	@PowerShell -File scripts/build.ps1
	@mkdir -p bin
	@cp powershell/sdk.ps1 bin/sdk.ps1

clean:
	@if [ -d "bin" ]; then rm -Rf bin; fi