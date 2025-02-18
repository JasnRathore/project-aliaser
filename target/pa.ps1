param(
    [string]$command,
	[string]$name,
	[string]$location
)

Import-Module ".\lib.psm1"

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
