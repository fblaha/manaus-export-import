#!/bin/bash
pushd  "$(dirname "$0")"

# TODO make parametrized
ls -t export*.zip | tail -n +2 | xargs -I {} rm {}

go install ../...
~/go/bin/mns-export

popd
