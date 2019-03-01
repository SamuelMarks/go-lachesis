#!/usr/bin/env sh

mkdir -p "$GOPATH/src/github.com/Fantom-foundation/go-lachesis" /cp_bin /bin
git clone 'https://github.com/SamuelMarks/lachesis' "$GOPATH/src/github.com/Fantom-foundation/go-lachesis"
cd "$GOPATH/src/github.com/Fantom-foundation/go-lachesis"
glide install
cd "$GOPATH/src/github.com/Fantom-foundation/go-lachesis/cmd/lachesis"
go build -ldflags "-linkmode external -extldflags -static -s -w" -a main.go
mv "$GOPATH/src/github.com/Fantom-foundation/go-lachesis/cmd/lachesis/main" /cp_bin/lachesis

apk --no-cache add libc-dev cmake
git clone https://github.com/SamuelMarks/docker-static-bin /build/docker-static-bin
mkdir /build/docker-static-bin/cmake-build-release
cd    /build/docker-static-bin/cmake-build-release
TEST_ENABLED=0 cmake -DCMAKE_BUILD_TYPE=Release ..
cd /build/docker-static-bin/cmd
gcc copy.c      -o "/cp_bin/copy"      -Os -static -Wno-implicit-function-declaration
gcc env.c       -o "/cp_bin/env"       -Os -static -Wno-implicit-function-declaration
gcc list.c      -o "/cp_bin/list"      -Os -static
gcc crappy_sh.c -o "/cp_bin/crappy_sh" -Os -static -Wno-implicit-function-declaration -Wno-int-conversion -I./../cmake-build-release
strip -s /cp_bin/crappy_sh /cp_bin/copy /cp_bin/env /cp_bin/list /cp_bin/lachesis
