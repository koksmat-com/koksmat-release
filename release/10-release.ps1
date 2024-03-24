<#---
title: GitHub
tag: github
api: post
---#>

param (
    $kitchenname = "nexiintra-profile"
)

$inputFile = join-path  $env:KITCHENROOT  $kitchenname ".koksmat/koksmat.json"

set-location (join-path  $env:KITCHENROOT  $kitchenname)

if (Test-Path -Path $inputFile ) {
    $json = Get-Content $inputFile | ConvertFrom-Json
}
else {
    $version = @{
        major = 0
        minor = 0
        patch = 1
        build = 0
    }
    $json = @{
        version = $version
    }
}

$version = "v$($json.version.major).$($json.version.minor).$($json.version.patch).$($json.version.build)"

gh release create $version  --generate-notes  

#gh release create v0.0.3.canary  --target canary  --generate-notes