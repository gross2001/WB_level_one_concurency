package main

// using adapter pattern for middlewares https://medium.com/@matryer/writing-middleware-in-golang-and-how-go-makes-it-so-much-fun-4375c1246e81

type Target interface {
	SomeFuncFromFamousLib() string
}

type WantToUse struct {
}

func (w WantToUse) AlreadyExistedFunc() string {
	return "AlreadyExistedFunc"
}

type Adapter struct {
	*WantToUse
}

func (a Adapter) SomeFuncFromFamousLib() string {
	return a.AlreadyExistedFunc()
}

func NewAdapter(w *WantToUse) Target {
	return Adapter{
		WantToUse: w,
	}
}
