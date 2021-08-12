package user

type JwtService struct{}
type BaseService struct{}

type ServiceGroup struct {
	JwtService
	BaseService
}
