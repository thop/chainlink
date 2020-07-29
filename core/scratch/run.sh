#/bin/bash


set -ex

here="$(dirname $0)"

for pkg in EACAggregatorProxy; do
    lowercase="$(echo $pkg | tr '[:upper:]' '[:lower:]')"
    target_dir="$here/gethwrappers/$lowercase"
    mkdir -p $target_dir
    abigen -sol "$here/../../evm-contracts/src/v0.6/$pkg.sol" -pkg "$lowercase" \
           -out "$target_dir/$lowercase.go"
done

for pkg in MockV3Aggregator; do
    lowercase="$(echo $pkg | tr '[:upper:]' '[:lower:]')"
    target_dir="$here/gethwrappers/$lowercase"
    mkdir -p $target_dir
    abigen -sol "$here/../../evm-contracts/src/v0.6/tests/$pkg.sol" -pkg "$lowercase" \
           -out "$target_dir/$lowercase.go"
done

go test -v
