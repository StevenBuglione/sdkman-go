package service

type IUseHandlerService interface {
	Use(args []string) error
}

type UseHandlerService struct {
	PowershellService IPowershellService
}

func NewUseHandlerService() IUseHandlerService {
	return &UseHandlerService{
		PowershellService: NewPowershellService(),
	}
}

func (u *UseHandlerService) Use(args []string) error {
	err := u.PowershellService.ExecuteJavaPathUpdate(args)
	if err != nil {
		return err
	}
	return nil
}
