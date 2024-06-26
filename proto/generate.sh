# cd proto
# buf generate
# cd ..

# cp -r github.com/wfblockchain/noble-cctp/* ./
# rm -rf github.com

set -eo pipefail
echo "Generating gogo proto code"
cd proto
proto_dirs=$(find ./ -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  for file in $(find "${dir}" -maxdepth 1 -name '*.proto'); do
    if grep "option go_package" $file &> /dev/null ; then
      buf generate --template buf.gen.yaml $file
    fi
  done
done
cd ..

cp -r github.com/wfblockchain/noble-cctp/* ./
rm -rf github.com
