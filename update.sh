#!/bin/sh

git submodule foreach git pull

# Collect paths of current script
origin_dir=$(pwd)
script_path=`realpath $0`
base_dir=$(dirname $script_path)

# Collect full path of the submodule and base directory
cd $base_dir
base_dir=$(pwd)
protos_dir="$base_dir/POGOProtos"

# Compile protobuffers and clean submodule
cd $protos_dir
current_commit=$(git rev-parse HEAD)
python ./compile_single.py -l=go --out_path=$base_dir
rm -rf "$protos_dir/tmp"

# Return back to original directory
cd $origin_dir

git add -A
git commit -m "Update POGOProtos $current_commit"
