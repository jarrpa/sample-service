package authorization

type authService struct{}

func NewService() Service { return &authService{} }
