package api

type Plugin interface {
	Iniciar()
	Coletar() (err error)
}
