# $NetBSD: Makefile,v 1.2 2018/07/19 14:24:10 jaapb Exp $

GITHUB_PROJECT=	ocsigen-i18n
DISTNAME=	${GITHUB_PROJECT}-3.1.0
PKGREVISION=	1
CATEGORIES=	www
MASTER_SITES=	${MASTER_SITE_GITHUB:=besport/}

MAINTAINER=	jaapb@NetBSD.org
HOMEPAGE=	https://github.com/besport/ocsigen-i18n/
COMMENT=	i18n package for Ocsigen
LICENSE=	gnu-lgpl-v2.1

USE_LANGUAGES=	# none
USE_TOOLS=	gmake
BUILD_TARGET=	build
MAKE_ENV=	bindir=${DESTDIR}${PREFIX}/bin
INSTALLATION_DIRS+=	bin

OCAML_USE_FINDLIB=	yes
OCAML_FINDLIB_DIRS=	# empty

.include "../../mk/ocaml.mk"

.if ${OCAML_USE_OPT_COMPILER} == "yes"
MAKE_ENV+=	OPTNESS=native
.else
MAKE_ENV+=	OPTNESS=byte
.endif

.include "../../mk/bsd.pkg.mk"
