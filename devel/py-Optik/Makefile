# $NetBSD: Makefile,v 1.22 2018/07/03 05:03:10 adam Exp $

DISTNAME=	optik-1.5.3
PKGNAME=	${PYPKGPREFIX}-${DISTNAME:S/optik/Optik/}
PKGREVISION=	1
CATEGORIES=	devel python
MASTER_SITES=	${MASTER_SITE_SOURCEFORGE:=optik/}

MAINTAINER=	pkgsrc-users@NetBSD.org
HOMEPAGE=	http://optik.sourceforge.net/
COMMENT=	Command line parsing library for Python
LICENSE=	modified-bsd

USE_LANGUAGES=	#none

PYTHON_VERSIONS_INCOMPATIBLE=	34 35 36 37

do-test:
	cd ${WRKSRC} && ${SETENV} ${TEST_ENV} ${PYTHONBIN} test/test_optik.py

.include "../../lang/python/distutils.mk"
.include "../../mk/bsd.pkg.mk"
