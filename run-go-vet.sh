#!/usr/bin/env bash

# golint 代码规范检测
if golint >/dev/null 2>&1; then  # 检测是否安装
	lint_errors=false
	for file in ${gofiles[@]} ; do
		lint_result="$(golint $file)" # run golint
		if test -n "$lint_result" ; then
			echo "golint '$file':\n$lint_result"
			lint_errors=true
			has_errors=1
		fi
	done
	if [ $lint_errors = true ] ; then
		echo "\n"
	fi
else
	echo 'Error: golint not install. Run: "go get -u github.com/golang/lint/golint"' >&2
	exit 1
fi
