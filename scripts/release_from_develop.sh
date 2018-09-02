#!/bin/bash

BRANCH=`git branch | grep '*'`

if [[ "$BRANCH" != "* develop" ]]; then
  echo "Must be launched from develop"
  exit
fi

VERSION=`git describe --abbrev=0`

echo "Releasing $VERSION to master"

git checkout master

if [[ "$?" != "0" ]]; then
  echo "Error"
  exit
fi

git pull origin develop

if [[ "$?" != "0" ]]; then
  echo "Error"
  exit
fi

git push origin master

if [[ "$?" != "0" ]]; then
  echo "Error"
  exit
fi

git tag -a $VERSION -m "$VERSION"

if [[ "$?" != "0" ]]; then
  echo "Error"
  exit
fi

git push --tags

if [[ "$?" != "0" ]]; then
  echo "Error"
  exit
fi

git checkout develop

if [[ "$?" != "0" ]]; then
  echo "Error"
  exit
fi


echo "Release of $VERSION OK"

VERSION_BITS=(${VERSION//./ })

VNUM1=${VERSION_BITS[0]}
VNUM2=${VERSION_BITS[1]}
VNUM3=${VERSION_BITS[2]}
VNUM3=$((VNUM3+1))

NEW_TAG="$VNUM1.$VNUM2.$VNUM3"

echo "Preparing next version: $NEW_TAG"

sed -i -e "s/$VERSION/$NEW_TAG/g" README.md

if [[ "$?" != "0" ]]; then
  echo "Error"
  exit
fi


sed -i -e "s/$VERSION/$NEW_TAG/g" cli.go

if [[ "$?" != "0" ]]; then
  echo "Error"
  exit
fi


git commit -am "Preparing new tag $NEW_TAG"

if [[ "$?" != "0" ]]; then
  echo "Error"
  exit
fi

git push origin develop

if [[ "$?" != "0" ]]; then
  echo "Error"
  exit
fi

echo "Ok"
