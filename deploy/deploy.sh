#!/usr/bin/env bash
set -e

# usage: ./deploy.sh dev master 32d1f89fb6f65cc9d7a0d3573ff2c8217fa58c58

PARENT_PATH=$( cd "$(dirname "${BASH_SOURCE}")" ; pwd -P )

AWS_ACCOUNT_ID=827562370231
EB_BUCKET=elasticbeanstalk-us-west-2-827562370231

ENV_NAME=wfe-${1:-dev}
BRANCH=${2:-$(git rev-parse --abbrev-ref HEAD)}
SHA1=${3:-$(git rev-parse HEAD)}

VERSION=$BRANCH-$SHA1
ZIP=$VERSION.zip

aws configure set default.region us-west-2

# Authenticate against our Docker registry
eval $(aws ecr get-login | sed -e 's/-e none//g')

# Build assets
make assets

# Build the wfe binaries
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64
export GO_FLAGS="-installsuffix cgo"
make release VERSION=$VERSION

# Push the Docker image
docker tag wfe:$VERSION $AWS_ACCOUNT_ID.dkr.ecr.us-west-2.amazonaws.com/wfe:$VERSION
docker push $AWS_ACCOUNT_ID.dkr.ecr.us-west-2.amazonaws.com/wfe:$VERSION

# Replace the <AWS_ACCOUNT_ID> with the real ID
sed -i='' "s/<AWS_ACCOUNT_ID>/$AWS_ACCOUNT_ID/" "$PARENT_PATH/Dockerrun.aws.json"
# Replace the <NAME> with the real name
sed -i='' "s/<NAME>/wfe/" "$PARENT_PATH/Dockerrun.aws.json"
# Replace the <TAG> with the real version number
sed -i='' "s/<TAG>/$VERSION/" "$PARENT_PATH/Dockerrun.aws.json"

# Zip up the Dockerrun file (feel free to zip up an .ebextensions directory with it)
cd "$PARENT_PATH" && zip -r $ZIP Dockerrun.aws.json .ebextensions

aws s3 cp $ZIP s3://$EB_BUCKET/$ZIP

# Create a new application version with the zipped up Dockerrun file
aws elasticbeanstalk create-application-version --application-name wfe \
    --version-label $VERSION --source-bundle S3Bucket=$EB_BUCKET,S3Key=$ZIP

# Update the environment to use the new application version
aws elasticbeanstalk update-environment --environment-name $ENV_NAME \
      --version-label $VERSION