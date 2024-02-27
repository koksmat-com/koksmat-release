<#---
title: GitHub
tag: github
---#>

param (
    $kitchenname = "koksmat-release"
)
$inputFile = join-path  $env:KITCHENROOT  $kitchenname ".koksmat/koksmat.json"

if (Test-Path -Path $inputFile ) {
    $json = Get-Content $inputFile | ConvertFrom-Json
} else {
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

gh release create $version --title $version --notes "Release $version" --target main