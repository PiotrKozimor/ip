set -ex
go build .
cont=$(buildah from scratch)
buildah copy $cont ip /bin/ip
buildah copy $cont /lib64/libpthread.so.0 /lib64/libpthread.so.0
buildah copy $cont /lib64/libc.so.6 /lib64/libc.so.6
buildah copy $cont /lib64/ld-linux-x86-64.so.2 /lib64/ld-linux-x86-64.so.2
buildah config --entrypoint '["/bin/ip"]'  $cont
buildah commit $cont ip
echo "👌 Tag nad push latest"
buildah tag ip docker.io/narciarz96/ip:latest
buildah push docker.io/narciarz96/ip:latest
buildah rm $cont