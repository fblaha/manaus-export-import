#!/bin/bash

pushd  "$(dirname "$0")/backup" || exit

# TODO make parametrized
ls -t export*.zip | tail -n +4 | xargs -I {} rm {}

~/go/bin/mns-export

popd || exit
