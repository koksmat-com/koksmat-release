<#---
title: Minor
tag: minor
api: post
---#>
param (
    $kitchenname = "nexiintra-profile"
)


$inputFile = join-path  $env:KITCHENROOT  $kitchenname ".koksmat/koksmat.json"

if (Test-Path -Path $inputFile ) {
    $json = Get-Content $inputFile | ConvertFrom-Json
    $json.version.minor++
    $json.version.build++
    $json.version.patch = 0

    Write-Host "Minor version bumped"
} else {
    $version = @{
        major = 0
        minor = 1
        patch = 0
        build = 1
    }
    $json = @{
        version = $version
    }
        


    Write-Host "New koksmat.json file will be created "
}
$json  | ConvertTo-Json -Depth 10 | Out-File -FilePath $inputFile -Encoding utf8NoBOM

Write-Host "New version $($json.version.major).$($json.version.minor).$($json.version.patch).$($json.version.build)"

