remoteRepo=git@github.com:gotamer/go.git
subDir=/src/net/mail

git clone -n --depth=1 --filter=tree:0 $remoteRepo
cd go
git sparse-checkout set --no-cone $subDir
git checkout
mv ../coSubDir.sh ./
git add --sparse coSubDir.sh
