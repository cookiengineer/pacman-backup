#!/bin/bash

GO="$(which go 2> /dev/null)";
ROOT="$(pwd)";



build() {

	local errors=0;

	IFS=$'\n' read -d "" -ra variants <<< "${1// /$'\n'}"; unset IFS;

	for variant in ${variants[@]}; do

		go_os="${variant%%:*}";
		go_arch="${variant##*:}";

		if [[ ! -d "${ROOT}/build/${go_os}" ]]; then
			mkdir -p "${ROOT}/build/${go_os}";
		fi;

		if [[ "${go_os}" == "windows" ]]; then
			cd "${ROOT}/source";
			env CGO_ENABLED=0 GOOS="${go_os}" GOARCH="${go_arch}" ${GO} build -o "${ROOT}/build/${go_os}/pacman-backup-${go_arch}.exe" "${ROOT}/source/cmds/pacman-backup/main.go";
		else
			cd "${ROOT}/source";
			env CGO_ENABLED=0 GOOS="${go_os}" GOARCH="${go_arch}" ${GO} build -o "${ROOT}/build/${go_os}/pacman-backup-${go_arch}" "${ROOT}/source/cmds/pacman-backup/main.go";
		fi;

		if [[ $? == 0 ]]; then
			echo -e "- Build ${go_tags}: ${go_os} / ${go_arch} [\e[32mok\e[0m]";
		else
			echo -e "- Build ${go_tags}: ${go_os} / ${go_arch} [\e[31mfail\e[0m]";
			errors=$((errors+1));
		fi;

	done;

	if [[ ${errors} == 0 ]]; then
		return 0;
	fi;

	return 1;

}



build "linux:amd64";

