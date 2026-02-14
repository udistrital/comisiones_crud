package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:AnexoComisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:AnexoComisionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:AnexoComisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:AnexoComisionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:AnexoComisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:AnexoComisionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:AnexoComisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:AnexoComisionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:AnexoComisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:AnexoComisionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:AnexoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:AnexoSolicitudController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:AnexoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:AnexoSolicitudController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:AnexoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:AnexoSolicitudController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:AnexoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:AnexoSolicitudController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:AnexoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:AnexoSolicitudController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:ComisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:ComisionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:ComisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:ComisionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:ComisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:ComisionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:ComisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:ComisionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:ComisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:ComisionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:DetalleSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:DetalleSolicitudController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:DetalleSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:DetalleSolicitudController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:DetalleSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:DetalleSolicitudController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:DetalleSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:DetalleSolicitudController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:DetalleSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:DetalleSolicitudController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:EstadoComisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:EstadoComisionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:EstadoComisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:EstadoComisionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:EstadoComisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:EstadoComisionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:EstadoComisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:EstadoComisionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:EstadoComisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:EstadoComisionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:EstadoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:EstadoDocumentoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:EstadoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:EstadoDocumentoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:EstadoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:EstadoDocumentoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:EstadoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:EstadoDocumentoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:EstadoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:EstadoDocumentoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:EstadoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:EstadoSolicitudController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:EstadoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:EstadoSolicitudController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:EstadoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:EstadoSolicitudController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:EstadoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:EstadoSolicitudController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:EstadoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:EstadoSolicitudController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:HistoricoEstadoComisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:HistoricoEstadoComisionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:HistoricoEstadoComisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:HistoricoEstadoComisionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:HistoricoEstadoComisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:HistoricoEstadoComisionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:HistoricoEstadoComisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:HistoricoEstadoComisionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:HistoricoEstadoComisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:HistoricoEstadoComisionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:HistoricoEstadoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:HistoricoEstadoSolicitudController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:HistoricoEstadoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:HistoricoEstadoSolicitudController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:HistoricoEstadoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:HistoricoEstadoSolicitudController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:HistoricoEstadoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:HistoricoEstadoSolicitudController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:HistoricoEstadoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:HistoricoEstadoSolicitudController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:ObservacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:ObservacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:ObservacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:ObservacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:ObservacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:ObservacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:ObservacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:ObservacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:ObservacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:ObservacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:RolUsuarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:RolUsuarioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:RolUsuarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:RolUsuarioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:RolUsuarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:RolUsuarioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:RolUsuarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:RolUsuarioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:RolUsuarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:RolUsuarioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:SeguimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:SeguimientoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:SeguimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:SeguimientoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:SeguimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:SeguimientoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:SeguimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:SeguimientoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:SeguimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:SeguimientoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:SolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:SolicitudController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:SolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:SolicitudController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:SolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:SolicitudController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:SolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:SolicitudController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:SolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:SolicitudController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoDocumentoComisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoDocumentoComisionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoDocumentoComisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoDocumentoComisionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoDocumentoComisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoDocumentoComisionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoDocumentoComisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoDocumentoComisionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoDocumentoComisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoDocumentoComisionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoDocumentoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoDocumentoSolicitudController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoDocumentoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoDocumentoSolicitudController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoDocumentoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoDocumentoSolicitudController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoDocumentoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoDocumentoSolicitudController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoDocumentoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoDocumentoSolicitudController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoSeguimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoSeguimientoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoSeguimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoSeguimientoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoSeguimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoSeguimientoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoSeguimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoSeguimientoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoSeguimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoSeguimientoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoSolicitudController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoSolicitudController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoSolicitudController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoSolicitudController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/comisiones_estudio_crud/controllers:TipoSolicitudController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
