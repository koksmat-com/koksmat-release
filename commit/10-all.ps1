<#---
title: All
tag: all
api: post
---

https://git-scm.com/docs/git-status

#>
param (
    $kitchenname = "azure-users",
    $commitmessage = "more"

)

push-location
$gitpath = (join-path  $env:KITCHENROOT  $kitchenname)
set-location $gitpath

koksmat trace log "updating git repository $gitpath"w
git add -u
git commit -m $commitmessage 

koksmat trace log "pushing changes to git"
git push
Pop-Location

