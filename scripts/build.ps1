# PowerShell
$ErrorActionPreference = "Stop"
$go = "go"
$buildDir = "bin/"
$projectName = "sdk.exe"

# Ensure Go is available
if (-Not (Get-Command $go -ErrorAction SilentlyContinue)) {
    Write-Output "Please install Go"
    Exit 1
}

# Ensure build directory
if(-Not(Test-Path -Path $buildDir )){
    New-Item -ItemType directory -Path $buildDir
}

cd $buildDir

$build = "$go build -o $projectName ../cmd/sdk"

try {
    Invoke-Expression $build
} catch {
    Write-Output "Build failed: $_"
    Exit 1
}

Write-Output "Build complete"