#!/bin/bash

GO="$(which go 2> /dev/null)";
ROOT="$(pwd)";

build() {

	local go_os="${1}";
	local go_arch="${2}";

	if [[ ! -d "${ROOT}/build" ]]; then
		mkdir -p "${ROOT}/build";
	fi;

	if [[ "${go_os}" == "windows" ]]; then
		cd "${ROOT}/source";
		env CGO_ENABLED=0 GOOS="${go_os}" GOARCH="${go_arch}" ${GO} build -o "${ROOT}/build/pacman-backup-${go_os}_${go_arch}.exe" "${ROOT}/source/cmds/pacman-backup/main.go";
	else
		cd "${ROOT}/source";
		env CGO_ENABLED=0 GOOS="${go_os}" GOARCH="${go_arch}" ${GO} build -o "${ROOT}/build/pacman-backup-${go_os}_${go_arch}" "${ROOT}/source/cmds/pacman-backup/main.go";
	fi;

	if [[ $? == 0 ]]; then
		echo -e "- Build ${go_os} / ${go_arch}: [\e[32mok\e[0m]";
		return 1;
	else
		echo -e "- Build ${go_os} / ${go_arch}: [\e[31mfail\e[0m]";
		return 0;
	fi;

}



build "linux" "amd64";

