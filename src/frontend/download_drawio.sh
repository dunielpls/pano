#!/bin/bash

main() {
    drawio_version="15.8.1"
    drawio_repo_url="https://github.com/jgraph/drawio.git"

    if [ -d ./build/drawio ] || [ -d ./drawio_repo ]
    then
        echo "download_drawio.sh: Removing old directories."
        rm -r ./build/drawio ./drawio_repo
    fi

    echo "download_drawio.sh: Cloning repository."
    git clone --depth 1 --branch "v$drawio_version" "$drawio_repo_url" ./drawio_repo > /dev/null 2>&1

    echo "download_drawio.sh: Copying frontend code and removing unneeded directories."
    mkdir ./build/drawio
    cp -r ./drawio_repo/src/main/webapp/* ./build/drawio/
    rm -rf ./build/drawio/{META,WEB}-INF ./drawio_repo
}

main
