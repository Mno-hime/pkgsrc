# $NetBSD: buildlink3.mk,v 1.2 2018/07/19 15:15:27 jaapb Exp $

BUILDLINK_TREE+=	ocaml-opam-file-format

.if !defined(OCAML_OPAM_FILE_FORMAT_BUILDLINK3_MK)
OCAML_OPAM_FILE_FORMAT_BUILDLINK3_MK:=

BUILDLINK_API_DEPENDS.ocaml-opam-file-format+=	ocaml-opam-file-format>=2.0.0beta
BUILDLINK_ABI_DEPENDS.ocaml-opam-file-format+=	ocaml-opam-file-format>=2.0.0beta3nb4
BUILDLINK_PKGSRCDIR.ocaml-opam-file-format?=	../../misc/ocaml-opam-file-format
.endif	# OCAML_OPAM_FILE_FORMAT_BUILDLINK3_MK

BUILDLINK_TREE+=	-ocaml-opam-file-format
