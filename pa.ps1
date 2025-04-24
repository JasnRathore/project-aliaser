param(
    [string]$command,
	[string]$name,
	[string]$location
)
$scriptDir = $PSScriptRoot
$modulePath = Join-Path $scriptDir 'lib.psm1'
Import-Module $modulePath

switch ($command) {
    { $_ -in @("add", "ad") } {
        addAlias $name $location
    }
    { $_ -in @("delete", "dl") } {
        deleteAlias $name
    }
    { $_ -in @("list", "ls") } {
        listAliases
    }
	"" {
		Write-Host "Invalid Input" -ForegroundColor Red
		Write-Host "Enter Command or Alias" -ForegroundColor Yellow
	}
    default {
        changeDirectory $command
    }
}
