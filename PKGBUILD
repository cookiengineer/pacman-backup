# Maintainer: Cookie Engineer <@cookiengineer>

pkgname=pacman-backup
pkgver=r20.d746588
pkgrel=1
pkgdesc='Pacman Backup tool for off-the-grid updates via portable USB sticks or (mesh) LAN networks.'
arch=('i686' 'x86_64' 'armv6h' 'armv7h' 'aarch64')
makedepends=('git' 'go')
url='https://github.com/cookiengineer/pacman-backup'
license=('GPL')
provides=('pacman-backup')
source=('git+https://github.com/cookiengineer/pacman-backup.git')
sha256sums=('SKIP')

_gitname="pacman-backup"

pkgver() {
	cd "${srcdir}/${_gitname}"
	printf "r%s.%s" "$(git rev-list --count HEAD)" "$(git rev-parse --short HEAD)"
}

build() {
	cd "${srcdir}/${_gitname}/source";
	env CGO_ENABLED=0 GOOS="linux" go build -o "${srcdir}/pacman-backup.bin" "${srcdir}/${_gitname}/source/cmds/pacman-backup/main.go";
}
package() {
	cd "${srcdir}/${_gitname}";
	install -Dm755 "${srcdir}/pacman-backup.bin" "${pkgdir}/usr/bin/pacman-backup";
}

