set -ex
go build .
cont=$(buildah from gcr.io/distroless/base-debian11)
buildah copy $cont ip /bin/ip
buildah config --entrypoint '["/bin/ip"]' $cont
buildah commit --format docker $cont ip
echo "👌 Tag nad push latest"
buildah tag ip docker.io/narciarz96/ip:latest
buildah push docker.io/narciarz96/ip:latest