#!/bin/sh -l

echo "---- github workspace"
echo ${GITHUB_REF#refs/heads/}
echo "------"

if [ $GITHUB_ACTIONS == "true" ] ; then
  source=${SOURCE:-.}
  cd ${GITHUB_WORKSPACE}/${source}

  git config --global user.name "$(git --no-pager log --format=format:'%an' -n 1)"
  git config --global user.email "$(git --no-pager log --format=format:'%ae' -n 1)"
  git remote set-url origin https://x-access-token:$GITHUB_TOKEN@github.com/$GITHUB_REPOSITORY
  git checkout "${GITHUB_REF#refs/heads/}"
  git add README.md
  git commit -m "go-badges update"
  git push
  echo "commit complete"
fi