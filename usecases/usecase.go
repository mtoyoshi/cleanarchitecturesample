package usecases

type UseCase interface {
	DoRun(param interface{}) UseCaseResult
}

type UseCaseResult struct {
	Value interface{}
	Error error
}

func (r UseCaseResult) HasError() bool {
	return r.Error != nil
}

type InputPort struct {
	UseCase UseCase
}

func (inputPort InputPort) Run(param interface{}, outputPort OutputPort) {
	outputPort.Run(inputPort.UseCase.DoRun(param))
}

type OutputPort interface {
	Run(result UseCaseResult)
}
