# Maintainer: MoriGM
pkgname=dreamdump
pkgver=0.0.0-rc1
pkgrel=1
pkgdesc='A programm to dump the dreamcasts HD Area of a gd-rom'
arch=(x86_64)
url='https://github.com/MoriGM/dreamdump/'
source=("${pkgname}::git+file://${startdir}/")
license=(MIT)

sha256sums=('SKIP')

depends=('zlib' 'glibc')
makedepends=('go')

build() {
cd $pkgname
go build .
}

package() {
install -D -m755 $pkgname/dreamdump -t ${pkgdir}/usr/bin/
}
