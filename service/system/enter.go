package system

type JwtService struct{}

type ServiceGroup struct {
	JwtService
	NicknameService
	FileService
}
