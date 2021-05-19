#!/bin/sh -l

echo "---- github workspace"
echo $GITHUB_REPOSITORY
echo "------"

if [ $GITHUB_ACTIONS == "true" ] ; then
  source=${SOURCE:-.}
  git remote -v

  git clone $GITHUB_REPOSITORY local_repo
  rm local_repo/README.md
  mv README.md local_repo/README.md
  cd local_repo

  git config --global user.name "$(git --no-pager log --format=format:'%an' -n 1)"
  git config --global user.email "$(git --no-pager log --format=format:'%ae' -n 1)"
  git remote set-url origin https://x-access-token:$GITHUB_TOKEN@github.com/$GITHUB_REPOSITORY
  # git checkout "${GITHUB_REF#refs/heads/}"
  git add README.md
  git commit -m "go-badges update"
  git push --verbose
  echo "commit complete"
fi
