// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/comisiones_estudio_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/tipo_solicitud",
			beego.NSInclude(
				&controllers.TipoSolicitudController{},
			),
		),

		beego.NSNamespace("/solicitud",
			beego.NSInclude(
				&controllers.SolicitudController{},
			),
		),

		beego.NSNamespace("/comision",
			beego.NSInclude(
				&controllers.ComisionController{},
			),
		),

		beego.NSNamespace("/detalle_solicitud",
			beego.NSInclude(
				&controllers.DetalleSolicitudController{},
			),
		),

		beego.NSNamespace("/historico_estado_solicitud",
			beego.NSInclude(
				&controllers.HistoricoEstadoSolicitudController{},
			),
		),

		beego.NSNamespace("/estado_solicitud",
			beego.NSInclude(
				&controllers.EstadoSolicitudController{},
			),
		),

		beego.NSNamespace("/rol_usuario",
			beego.NSInclude(
				&controllers.RolUsuarioController{},
			),
		),

		beego.NSNamespace("/observacion",
			beego.NSInclude(
				&controllers.ObservacionController{},
			),
		),

		beego.NSNamespace("/documento_solicitud",
			beego.NSInclude(
				&controllers.DocumentoSolicitudController{},
			),
		),

		beego.NSNamespace("/tipo_documento_solicitud",
			beego.NSInclude(
				&controllers.TipoDocumentoSolicitudController{},
			),
		),

		beego.NSNamespace("/estado_documento",
			beego.NSInclude(
				&controllers.EstadoDocumentoController{},
			),
		),

		beego.NSNamespace("/historico_estado_comision",
			beego.NSInclude(
				&controllers.HistoricoEstadoComisionController{},
			),
		),

		beego.NSNamespace("/estado_comision",
			beego.NSInclude(
				&controllers.EstadoComisionController{},
			),
		),

		beego.NSNamespace("/documento_comision",
			beego.NSInclude(
				&controllers.DocumentoComisionController{},
			),
		),

		beego.NSNamespace("/tipo_documento_comision",
			beego.NSInclude(
				&controllers.TipoDocumentoComisionController{},
			),
		),

		beego.NSNamespace("/seguimiento",
			beego.NSInclude(
				&controllers.SeguimientoController{},
			),
		),

		beego.NSNamespace("/tipo_seguimiento",
			beego.NSInclude(
				&controllers.TipoSeguimientoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
