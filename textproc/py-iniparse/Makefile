# $NetBSD: Makefile,v 1.10 2018/07/03 05:03:36 adam Exp $

DISTNAME=	iniparse-0.4
PKGNAME=	${PYPKGPREFIX}-${DISTNAME}
CATEGORIES=	python textproc
MASTER_SITES=	http://iniparse.googlecode.com/files/

MAINTAINER=	pkgsrc-users@NetBSD.org
HOMEPAGE=	https://pypi.python.org/pypi/iniparse/
COMMENT=	Better INI parser for Python
LICENSE=	mit AND python-software-foundation

USE_LANGUAGES=	# none

PYTHON_VERSIONS_INCOMPATIBLE=	34 35 36 37

post-install:
	${CHMOD} -x ${DESTDIR}${PREFIX}/share/doc/iniparse-0.4/index.html

.include "../../lang/python/distutils.mk"
.include "../../mk/bsd.pkg.mk"
