#!/bin/bash

main() {
    drawio_version="15.8.1"
    drawio_repo_url="https://github.com/jgraph/drawio.git"

    while true
    do
        case $answer in
            [Yy]*)
                if [ -d ./web/drawio ] || [ -d ./web/drawio_repo ]
                then
                    echo "Removing old directories."
                    rm -r ./web/drawio ./web/drawio_repo
                fi

                echo "Cloning repository."
                git clone --depth 1 --branch "v$drawio_version" "$drawio_repo_url" ./web/drawio_repo > /dev/null 2>&1

                echo "Copying frontend code and removing unneeded directories."
                mkdir ./web/drawio
                cp -r ./web/drawio_repo/src/main/webapp/* ./web/drawio/
                rm -rf ./web/drawio/{META,WEB}-INF ./web/drawio_repo

                exit;;
            *) echo "Exiting."; exit;;
        esac
    done
}

main
