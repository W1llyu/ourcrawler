package collector

type ConsoleCollector struct {
}

func NewConsoleCollector() *ConsoleCollector {
	return &ConsoleCollector{}
}

func (f *ConsoleCollector) Process(resultItems ResultItems) error {
	return nil
}
