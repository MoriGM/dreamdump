# Maintainer: MoriGM
pkgname=dreamdump
pkgver=0.0.0
pkgrel=1
pkgdesc='A programm to dump the dreamcasts HD Area of a gd-rom'
arch=(x86_64)
url='https://github.com/MoriGM/dreamdump/'
source=(https://github.com/MoriGM/dreamdump/archive/refs/heads/main.zip)
license=(MIT)

sha256sums=('SKIP')

depends=('zlib' 'glibc')
makedepends=('go')

build() {
cd $pkgname-main
go build .
}

package() {
install -D -m755 $pkgname-main/dreamdump -t ${pkgdir}/usr/bin/
}
