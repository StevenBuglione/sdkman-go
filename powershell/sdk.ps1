param(
    [Parameter(Position = 0,
            ValueFromRemainingArguments = $true)]
    [String[]]$commands
)

$exePath = Join-Path -Path $PSScriptRoot -ChildPath "sdk.exe"

& $exePath $commands
$exitCode = $LASTEXITCODE

# Only Update If Exit Code is 100, 100 = SuccessRefresh Powershell Session Environment
if ($exitCode -eq 100) {
    Update-SessionEnvironment
}


function Get-EnvironmentVariableNames([System.EnvironmentVariableTarget] $Scope) {
    switch ($Scope) {
        'User' { Get-Item 'HKCU:\Environment' | Select-Object -ExpandProperty Property }
        'Machine' { Get-Item 'HKLM:\SYSTEM\CurrentControlSet\Control\Session Manager\Environment' | Select-Object -ExpandProperty Property }
        'Process' { Get-ChildItem Env:\ | Select-Object -ExpandProperty Key }
        default { throw "Unsupported environment scope: $Scope" }
    }
}

function Get-EnvironmentVariable([string] $Name, [System.EnvironmentVariableTarget] $Scope) {
    [Environment]::GetEnvironmentVariable($Name, $Scope)
}

function Update-SessionEnvironment {
    Write-Debug "Running 'Update-SessionEnvironment' - Updating the environment variables for the session."

    #ordering is important here, $user comes after so we can override $machine
    'Machine', 'User' |
            % {
                $scope = $_
                Get-EnvironmentVariableNames -Scope $scope |
                        % {
                            Set-Item "Env:$($_)" -Value (Get-EnvironmentVariable -Scope $scope -Name $_)
                        }
            }

    #Path gets special treatment b/c it munges the two together
    $paths = 'Machine', 'User' |
            % {
                (Get-EnvironmentVariable -Name 'PATH' -Scope $_) -split ';'
            } |
            Select -Unique
    $Env:PATH = $paths -join ';'
}

