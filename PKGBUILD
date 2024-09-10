# Maintainer: Cookie Engineer <@cookiengineer>

pkgname=pacman-backup
pkgver=r16.7a7623c
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
	cd "${srcdir}/${_gitname}";
	bash build.sh;
}
package() {
	cd "${srcdir}/${_gitname}";
	install -Dm755 "build/pacman-backup_linux_amd64" "$pkgdir/usr/bin/pacman-backup";
}

