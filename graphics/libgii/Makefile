# $NetBSD: Makefile,v 1.36 2018/01/14 14:58:39 rillig Exp $

DISTNAME=	libgii-1.0.2.src
PKGREVISION=	1
PKGNAME=	${DISTNAME:S/.src//}
CATEGORIES=	graphics
MASTER_SITES=	http://www.ggi-project.org/ftp/ggi/v2.2/
MASTER_SITES+=	${MASTER_SITE_SOURCEFORGE:=ggi/}
EXTRACT_SUFX=	.tar.bz2

MAINTAINER=	pkgsrc-users@NetBSD.org
HOMEPAGE=	http://www.ggi-project.org/
COMMENT=	General Input Interface - API for all possible input sources

WRKSRC=			${WRKDIR}/${PKGNAME_NOREV}
USE_LIBTOOL=		yes
TEST_TARGET=		check

GNU_CONFIGURE=		yes
CONFIGURE_ARGS+=	--with-x
CONFIGURE_ARGS+=	--sysconfdir=${PKG_SYSCONFDIR}

PKG_SYSCONFSUBDIR=	ggi
EGDIR=			${PREFIX}/share/examples/ggi
CONF_FILES=		${EGDIR}/libgii.conf ${PKG_SYSCONFDIR}/libgii.conf
MAKE_DIRS+=		${PKG_SYSCONFDIR}/filter
CONF_FILES+=		${EGDIR}/filter/keytrans ${PKG_SYSCONFDIR}/filter/keytrans
CONF_FILES+=		${EGDIR}/filter/mouse ${PKG_SYSCONFDIR}/filter/mouse

PTHREAD_OPTS+=		native

.include "../../mk/pthread.buildlink3.mk"

PLIST_VARS+=		ggtick
.if empty(PTHREAD_TYPE:Mnative)
PLIST.ggtick=		yes
.endif

BUILDLINK_DEPMETHOD.libXt?=	build

.include "../../mk/dlopen.buildlink3.mk"
.include "../../x11/libX11/buildlink3.mk"
.include "../../x11/libXt/buildlink3.mk"
.include "../../x11/libXxf86dga/buildlink3.mk"
.include "../../mk/bsd.pkg.mk"
