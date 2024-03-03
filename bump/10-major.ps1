<#---
title: Minor
tag: major
api: post
---#>
param (
    $kitchenname = "nexiintra-profile"
)


$inputFile = join-path  $env:KITCHENROOT  $kitchenname ".koksmat/koksmat.json"

if (Test-Path -Path $inputFile ) {
    $json = Get-Content $inputFile | ConvertFrom-Json
    $json.version.minor = 0
    $json.version.major++
    $json.version.patch = 0
    $json.version.build++
    Write-Host "Major version bumped"
} else {
    $version = @{
        major = 1
        minor = 0
        patch = 0
        build = 1
    }
    $json = ${
        version = $version
    }
        
    Write-Host "New koksmat.json file will be created "
}
$json  | ConvertTo-Json -Depth 10 | Out-File -FilePath $inputFile -Encoding utf8NoBOM

Write-Host "New version $($json.version.major).$($json.version.minor).$($json.version.patch).$($json.version.build)"

