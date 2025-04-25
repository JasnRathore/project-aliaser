function GetExe {
    $scriptDir = $PSScriptRoot
    $exePath = Join-Path $scriptDir 'fa.exe'
    $exePath 
}

function GetMidFile {
    $scriptDir = $PSScriptRoot
    $exePath = Join-Path $scriptDir 'mid_file.json'
    $exePath 
}
function addAlias {
    param (
        [string]$name,
        [string]$location
    )
    $validPath = Test-Path $location -PathType Container

    $doesExist = checkAlias $name
    
    if ($name -eq "") {
        echo "Name Can Not be Empty"
        return
    }
    if ($doesExist -eq "true") {
        echo "Name Already Exists"
        return
    }
    if ($location -eq ".") {
        $location = pwd
    } elseif ($location -eq "" -or -not $validPath) {
        echo "Enter Valid Location "
        return
    }
    
	$exe = getExe
    $output = & $exe "add" "$name" "$location"
    echo $output
}

function checkAlias {
    param (
        [string]$name
    )
	$exe = getExe
    $output = & $exe "check" "$name"    
    $output
}

function deleteAlias {
    param (
        [string]$name
    )
	$exe = getExe
    $output = & $exe "delete" "$name"    
    echo $output
}

function listAliases {
	$exe = getExe
    $output = & $exe "list" | ConvertFrom-Json
    echo $output
}

function changeDirectory {
    param (
        [string]$name
    )
	$exe = getExe
    $output = & $exe "cd" "$name" | ConvertFrom-Json
    cd $output.location
}
