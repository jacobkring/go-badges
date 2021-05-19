#!/bin/sh -l

echo "---- github workspace"
echo ${GITHUB_REF#refs/heads/}
echo "------"

if [ $GITHUB_ACTIONS == "true" ] ; then
  source=${SOURCE:-.}
  git remote -v

  cd ${GITHUB_WORKSPACE}/${source}
  
  git remote -v
  git config --global user.name "$(git --no-pager log --format=format:'%an' -n 1)"
  git config --global user.email "$(git --no-pager log --format=format:'%ae' -n 1)"
  git remote set-url origin https://x-access-token:$GITHUB_TOKEN@github.com/$GITHUB_REPOSITORY
  git remote -v
  # git checkout "${GITHUB_REF#refs/heads/}"
  git add README.md
  git commit -m "go-badges update"
  git push --verbose
  echo "commit complete"
fi
