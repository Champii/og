#!/bin/bash

set -e

BRANCH=`git branch | grep '*'`

if [[ "$BRANCH" != "* develop" ]]; then
  echo "Must be launched from develop"
  exit 1
fi

TYPE=patch

if [[ "$1" != "" ]]; then
  TYPE=$1
fi

VERSION=`git describe --abbrev=0`

VERSION_BITS=(${VERSION//./ })

VNUM1=${VERSION_BITS[0]}
VNUM2=${VERSION_BITS[1]}
VNUM3=${VERSION_BITS[2]}

if [[ "$TYPE" == "patch" ]]; then
  VNUM3=$((VNUM3+1))
elif [[ "$TYPE" == "minor" ]]; then
  VNUM2=$((VNUM2+1))
  VNUM3=0
elif [[ "$TYPE" == "major" ]]; then
  VNUM1=$((VNUM1+1))
  VNUM2=0
  VNUM3=0
else
  echo "Bad semver action: $TYPE"
  exit 1
fi

NEW_TAG="$VNUM1.$VNUM2.$VNUM3"

echo "Preparing next version: $NEW_TAG"

sed -i -e "s/DEV: Current version/$NEW_TAG: Current version/g" CHANGELOG.md

sed -i -e "s/DEV/$NEW_TAG/g" README.md

sed -i -e "s/DEV/$NEW_TAG/g" cli.go

git commit -am "Version tag $NEW_TAG"

git push origin develop

echo "Releasing $NEW_TAG to master"

git checkout master

git pull origin develop

git push origin master

git tag -a $NEW_TAG -m "$NEW_TAG"

git push --tags

git checkout develop

echo "Release of $NEW_TAG OK"

echo "Restoring DEV tag in develop"

sed -i -e "s/$NEW_TAG: Current version/DEV: Current version\\n\\n## $NEW_TAG/g" CHANGELOG.md

sed -i -e "s/$NEW_TAG/DEV/g" README.md

sed -i -e "s/$NEW_TAG/DEV/g" cli.go

git commit -am "Restoring DEV version"

git push origin develop

echo "Ok"
