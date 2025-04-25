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
    	$exe = getExe
        & $exe
        $file = GetMidFile
        $data = Get-Content $file -Raw | ConvertFrom-Json
        switch ($data.command) {
            "cd" {
               changeDirectory $data.name
            }
        }
        Clear-Content -Path $file       
	}
    default {
        changeDirectory $command
    }
}
