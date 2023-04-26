package electronic

type Phone interface {
	Brand() string
	Model() string
	Type() string
}

type StationPhone interface {
	ButtonsCount() int
}

type Smartphone interface {
	OS() string
}

func NewApplePhone(model string) applePhone {
	var brand string
	brand = "Apple"
	return applePhone{basePhone: basePhone{brand: brand, model: model}}
}

func NewAndroidPhone(brand string, model string) androidPhone {
	return androidPhone{basePhone: basePhone{brand: brand, model: model}}
}

func NewRadioPhone(brand string, model string, buttonsCount int) radioPhone {
	return radioPhone{basePhone: basePhone{brand: brand, model: model}, buttonsCount: buttonsCount}
}

type basePhone struct {
	brand string
	model string
}

func (p *basePhone) Brand() string {
	return p.brand
}

func (p *basePhone) Model() string {
	return p.model
}

type applePhone struct {
	basePhone
}

func (p *applePhone) Type() string { return "smartphone" }

func (p *applePhone) OS() string { return "ios" }

type androidPhone struct {
	basePhone
}

func (p *androidPhone) Type() string { return "smartphone" }

func (p *androidPhone) OS() string { return "android" }

type radioPhone struct {
	basePhone
	buttonsCount int
}

func (p *radioPhone) Type() string { return "station" }

func (p *radioPhone) ButtonsCount() int {
	return p.buttonsCount
}
