<#---
title: web
output: web.yaml
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

$SPAUTH_TENANTID=$env:SPAUTH_TENANTID
$SPAUTH_CLIENTID=$env:SPAUTH_CLIENTID
$SPAUTH_CLIENTSECRET=$env:SPAUTH_CLIENTSECRET

$repo = gh repo view --json nameWithOwner  | convertfrom-json

$image = "ghcr.io/$($repo.nameWithOwner)-web:$($version)"

$yaml = @"
apiVersion: apps/v1
kind: Deployment
metadata:
  name: $kitchenname
spec:
  selector:
    matchLabels:
      app: $kitchenname
  replicas: 1
  template:
    metadata:
      labels:
        app: $kitchenname
    spec: 
      containers:
      - name: $kitchenname
        image: $image
        ports:
          - containerPort: 3001
        env:
        - name: SPAUTH_TENANTID
          value: $SPAUTH_TENANTID
        - name: SPAUTH_CLIENTID
          value: $SPAUTH_CLIENTID
        - name: SPAUTH_CLIENTSECRET
          value: $SPAUTH_CLIENTSECRET

      
---
apiVersion: v1
kind: Service
metadata:
  name: $kitchenname
  labels:
    app: $kitchenname
    service: $kitchenname
spec:
  ports:
  - name: http
    port: 5301
    targetPort: 3001
  selector:
    app: $kitchenname

"@



$result = join-path  $env:WORKDIR  "web.yaml"
$yaml  | Out-File -FilePath $result -Encoding utf8NoBOM
