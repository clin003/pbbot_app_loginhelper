#!/bin/bash

VERSION=0.0.4
APPNAME=pbbot_app_loginhelper
echo "package constvar" > ./pkg/constvar/version.go
echo "const(APP_NAME = \"${APPNAME}\"" >> ./pkg/constvar/version.go
echo "APP_VERSION = \"${VERSION}\")" >> ./pkg/constvar/version.go
go fmt ./pkg/constvar

echo "#!/bin/bash" > _sh_docker_build_push
echo "#sudo docker image ls" >> _sh_docker_build_push
echo "#sudo docker container prune -f" >> _sh_docker_build_push
echo "#sudo docker image prune -af" >> _sh_docker_build_push
echo "#sudo docker system prune -af" >> _sh_docker_build_push

echo "curl -o cacert.pem https://curl.se/ca/cacert.pem" >> _sh_docker_build_push
echo "sudo docker build -t baicailin/${APPNAME}:v${VERSION} ." >> _sh_docker_build_push
echo "sudo docker image ls" >> _sh_docker_build_push
echo "echo \"--------------------->\"" >> _sh_docker_build_push
echo "echo \"--------------------->push ${APPNAME}:v${VERSION}\"" >> _sh_docker_build_push
echo "sudo docker push baicailin/${APPNAME}:v${VERSION}" >> _sh_docker_build_push

echo "sudo docker tag baicailin/${APPNAME}:v${VERSION} baicailin/${APPNAME}:latest" >> _sh_docker_build_push
echo "sudo docker image ls" >> _sh_docker_build_push
echo "echo \"--------------------->\"" >> _sh_docker_build_push
echo "echo \"--------------------->push ${APPNAME}:latest\"" >> _sh_docker_build_push
echo "sudo docker push baicailin/${APPNAME}:latest" >> _sh_docker_build_push

#echo "echo \"--------------------->tencentyun push ${APPNAME}:latest\"" >> _sh_docker_build_push
#echo "sudo docker tag baicailin/${APPNAME}:latest hkccr.ccs.tencentyun.com/baicai_dev/${APPNAME}:latest" >> _sh_docker_build_push
#echo "sudo docker push hkccr.ccs.tencentyun.com/baicai_dev/${APPNAME}:latest" >> _sh_docker_build_push

echo "sudo docker image ls" >> _sh_docker_build_push


#git init #
git add .
git commit -m "v${VERSION} debug"
#git remote add gitee git@gitee.com:lyhuilin/${APPNAME}.git #
#git remote add github git@github.com:clin003/${APPNAME}.git #
#git branch -M main #
git push -u gitee main
git push -u github main
git tag "v${VERSION}"
git push --tags  -u github main
git push --tags  -u gitee main
git remote -v