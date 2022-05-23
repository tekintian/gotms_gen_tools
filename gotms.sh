#!/bin/bash
# gotms dao do entity make tool
# 
CUR_DIR=`PWD`

GOTMS_PATH="/Users/Tekin/Downloads/go/gotms"

echo "$CUR_DIR/app dir delete"
rm -rf $CUR_DIR/app

echo "gf gen dao"
# gen the dao do entity
cd $CUR_DIR/bin
./gf gen dao  -c $CUR_DIR/bin/config.yaml

echo "$GOTMS_PATH gotms entity dao do make"
rm -rf $GOTMS_PATH/app/system/model/entity
rm -rf $GOTMS_PATH/app/system/service/internal

echo "mv gen  entity dao do  to gotms path $GOTMS_PATH "
mv $CUR_DIR/app/system/model/entity $GOTMS_PATH/app/system/model/entity
mv $CUR_DIR/app/system/service/internal $GOTMS_PATH/app/system/service/internal

rm -rf $CUR_DIR/app
echo "all ok!"
