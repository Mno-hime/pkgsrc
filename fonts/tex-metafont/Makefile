# $NetBSD: Makefile,v 1.12 2018/06/03 01:30:21 markd Exp $

DISTNAME=	metafont
# should be this instead:
#PKGNAME=	tex-${DISTNAME}-2.7182818.20140107
PKGNAME=	tex-${DISTNAME}-2014
PKGREVISION=	1
CATEGORIES=	fonts
TEXLIVE_REV=	35821

MAINTAINER=	minskim@NetBSD.org
HOMEPAGE=	http://ctan.org/pkg/metafont
COMMENT=	METAFONT system for specifying fonts
#LICENSE=	knuth

DEPENDS+=	tex-kpathsea-[0-9]*:../../print/tex-kpathsea
DEPENDS+=	web2c-[0-9]*:../../print/web2c

TEX_FORMATS=	"mf mf-nowin - -translate-file=cp227.tcx mf.ini\\n"

.include "../../print/texlive/package.mk"
.include "../../mk/bsd.pkg.mk"
