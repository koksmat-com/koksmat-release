<#---
title: Minor
tag: patch
api: post
---#>
param (
    $kitchenname = "koksmat-release"
)

$inputFile = join-path  $env:KITCHENROOT  $kitchenname ".koksmat/koksmat.json"

if (Test-Path -Path $inputFile ) {
    $json = Get-Content $inputFile | ConvertFrom-Json

    $json.version.patch++
    $json.version.build++
    Write-Host "Patch version bumped"
}
else {
    $version = @{
        major = 0
        minor = 0
        patch = 1
        build = 1
    }
    $json = @{
        version = $version
    }

    Write-Host "New koksmat.json file will be created "
}
$json  | ConvertTo-Json -Depth 10 | Out-File -FilePath $inputFile -Encoding utf8NoBOM

Write-Host "New version $($json.version.major).$($json.version.minor).$($json.version.patch).$($json.version.build)"