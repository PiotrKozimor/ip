set -ex
CGO_ENABLED=0 go build .
cont=$(buildah from scratch)
buildah copy $cont ip /bin/ip
buildah config --entrypoint '["/bin/ip"]' --port 80 --port 443 $cont
buildah commit $cont ip
echo "👌 Tag nad push latest"
buildah tag ip docker.io/narciarz96/ip:latest
buildah push docker.io/narciarz96/ip:latest
buildah rm $cont