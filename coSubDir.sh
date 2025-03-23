
# https://stackoverflow.com/questions/600079/how-do-i-clone-a-subdirectory-only-of-a-git-repository/52269934#52269934

remoteRepo=git@github.com:gotamer/go.git
subDir=/src/net/mail

git clone -n --depth=1 --filter=tree:0 $remoteRepo
cd go
git sparse-checkout set --no-cone $subDir
git checkout
mv ../coSubDir.sh ./
git add --sparse coSubDir.sh
