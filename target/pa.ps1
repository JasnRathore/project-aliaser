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
        break
    }
    { $_ -in @("delete", "dl") } {
        deleteAlias $name
        break
    }
    { $_ -in @("list", "ls") } {
        listAliases
        break
    }
    "" {
        $exe = getExe
        & $exe 
        $file = GetMidFile
        
        if (Test-Path $file -PathType Leaf) {
            $content = Get-Content $file -Raw
            if (-not [string]::IsNullOrEmpty($content)) {
                $data = $content | ConvertFrom-Json
                switch ($data.command) {
                    "cd" {
                        changeDirectory $data.name
                    }
                }
            }
            Clear-Content -Path $file
        }
        break  # Explicitly break after handling empty command
    }
    default {
        if (-not [string]::IsNullOrEmpty($command)) {
            changeDirectory $command
        }
        break
    }
}