#!/bin/bash
# Build gosdk project for publish
# Author: Baidu BCE SDK Team

WORKROOT=$(pwd)
cd $WORKROOT
OUTPUT=$WORKROOT/output
if [ -d $OUTPUT ]; then
	rm -rf $OUTPUT
fi

DEST=$OUTPUT/github.com/baidubce/baiducloud-go-sdk
mkdir -p $DEST
cp -Rf ./bce $DEST
cp -Rf ./core $DEST
cp -Rf ./sample $DEST
cp -Rf ./services $DEST
echo "Build gosdk success"
