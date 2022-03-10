#!/usr/bin/env bash
#表示有报错即退出 跟set -e含义一样
set -o errexit
#执行脚本的时候，如果遇到不存在的变量，Bash 默认忽略它 ,跟 set -u含义一样
set -o nounset
# 只要一个子命令失败，整个管道命令就失败，脚本就会终止执行
set -o pipefail

#kubebuilder项目的MODULE
MODULE=elasticgpu.io/elastic-gpu

#api包
APIS_PKG=api

# group-version such as cronjob:v1
GROUP=elasticgpu.io
VERSION=v1alpha1
GROUP_VERSION=$GROUP:$VERSION

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
CODEGEN_PKG="${SCRIPT_ROOT}/vendor/k8s.io/code-generator"

# kubebuilder2.3.2版本生成的api目录结构code-generator无法直接使用
rm -rf "${SCRIPT_ROOT}/${APIS_PKG}/${GROUP}" && mkdir -p "${SCRIPT_ROOT}/${APIS_PKG}/${GROUP}" && cp -r "${SCRIPT_ROOT}/${APIS_PKG}/${VERSION}" "${SCRIPT_ROOT}/${APIS_PKG}/${GROUP}"

# generate the code with:
# --output-base    because this script should also be able to run inside the vendor dir of
#                  k8s.io/kubernetes. The output-base is needed for the generators to output into the vendor dir
#                  instead of the $GOPATH directly. For normal projects this can be dropped.
#client,informer,lister(注意: code-generator 生成的deepcopy不适配 kubebuilder 所生成的api)
bash "${CODEGEN_PKG}"/generate-groups.sh "client,informer,lister" \
  ${MODULE} \
  ${MODULE}/${APIS_PKG} \
  ${GROUP_VERSION} \
  --go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt \
  --output-base "${SCRIPT_ROOT}"

rm -rf "${SCRIPT_ROOT}/clientset" && rm -rf "${SCRIPT_ROOT}/informers" && rm -rf "${SCRIPT_ROOT}/listers"
cp -r "${SCRIPT_ROOT}/${MODULE}/." "${SCRIPT_ROOT}"
rm -rf "${SCRIPT_ROOT}/${GROUP}"
