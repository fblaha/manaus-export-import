#!/bin/bash
pushd  "$(dirname "$0")"

# TODO make parametrized
ls export*.zip -t | tail -n +2 | xargs rm --

~/go/bin/mns-export

popd