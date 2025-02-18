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
    
    $output = .\fa.exe "add" "$name" "$location"
    echo $output
}

function checkAlias {
    param (
        [string]$name
    )
    $output = .\fa.exe "check" "$name"    
    $output
}

function deleteAlias {
    param (
        [string]$name
    )
    $output = .\fa.exe "delete" "$name"    
    echo $output
}

function listAliases {
    $output = .\fa.exe "list" | ConvertFrom-Json
    echo $output
}

function changeDirectory {
    param (
        [string]$name
    )
    $output = .\fa.exe "cd" "$name" | ConvertFrom-Json
    cd $output.location
}
