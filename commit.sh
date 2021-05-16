#!/bin/bash

git config --global user.name "$(git --no-pager log --format=format:'%an' -n 1)"
git config --global user.email "$(git --no-pager log --format=format:'%ae' -n 1)"
cd /github/workspace
git add README.md
git commit -m $INPUT_COMMIT-MSG
git push
cd ../..