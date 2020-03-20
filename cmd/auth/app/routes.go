package app

import (
	"auth-service/pkg/core/token"
	"auth-service/pkg/mux/middleware/authenticated"
	"auth-service/pkg/mux/middleware/jwt"
	"auth-service/pkg/mux/middleware/logger"
	"reflect"
)

func (s *Server) InitRoutes() {
	s.router.POST(
		"/api/tokens",
		s.handleCreateToken(),
		//tokens.HandleCreateToken(s),
		logger.Logger("TOKEN"),
	)
	s.router.POST(
		"/api/newUser",
		s.handleAddUser(),
		logger.Logger("Registration"),
	)
	// /api/users/me
	// golang нельзя reflect.TypeOf(token.Payload)
	// golang нельзя reflect.TypeOf((*token.Payload)(nil))
	s.router.GET(
		"/api/users/me",
		s.handleProfile(),
		authenticated.Authenticated(jwt.IsContextNonEmpty),
		jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), s.secret),
		logger.Logger("USERS"),
	)
	s.router.GET(
		"/api/rooms/list",
		s.handleGetAllRooms(),
		logger.Logger("USERS"),
	)
	s.router.DELETE(
		"/api/users/1",
		s.handleProfile(),
		authenticated.Authenticated(jwt.IsContextNonEmpty),
		jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), s.secret),
		logger.Logger("USERS"),
	)
}
