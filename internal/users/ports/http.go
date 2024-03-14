package ports

import (
	"common/server/binder"
	"net/http"
	"users/application"
)

type HttpServer struct {
	app    *application.App
	Binder binder.Binder
}

func NewHttpServer(app *application.App) *HttpServer {
	return &HttpServer{app: app, Binder: binder.JSONBinder{}}
}

func (h HttpServer) UserLogin(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h HttpServer) UserLogout(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h HttpServer) UserRegister(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h HttpServer) UserApplyRegister(w http.ResponseWriter, r *http.Request, params UserApplyRegisterParams) {
	//user, err := h.app.Queries.GetUserInfo.Handle(r.Context(), query.GetUserInfo{Email: params.Email})
	//if err != nil {
	//	httperr.Error(w, e.Wrap(e.QueryHandlerErr, err))
	//	return
	//}

}

func (h HttpServer) GetUserInfo(w http.ResponseWriter, r *http.Request, uid int) {
	//TODO implement me
	panic("implement me")
}

func (h HttpServer) ResetPassword(w http.ResponseWriter, r *http.Request, uid string) {
	//TODO implement me
	panic("implement me")
}

func (h HttpServer) ModifyUserName(w http.ResponseWriter, r *http.Request, uid string) {
	//TODO implement me
	panic("implement me")
}

func (h HttpServer) ModifyUserRoles(w http.ResponseWriter, r *http.Request, uid string) {
	//TODO implement me
	panic("implement me")
}

func (h HttpServer) UserVerify(w http.ResponseWriter, r *http.Request, uid string) {
	//TODO implement me
	panic("implement me")
}
