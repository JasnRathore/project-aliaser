param(
    [string]$command,
	[string]$name,
	[string]$location
)
$scriptDir = $PSScriptRoot
$modulePath = Join-Path $scriptDir 'lib.psm1'
Import-Module $modulePath

switch ($command) {
	"add" {
		addAlias $name $location 
	}
	"delete" {
		deleteAlias $name
	}
	"list" {
		listAliases
	}
	default {
		changeDirectory $command		
	}
}
