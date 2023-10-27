package laundry

import "net/http"

type laundry struct {
	r *router
}

func (l *laundry) Run(addr string) error {
	return http.ListenAndServe(addr, l.r)
}

func New(basketDir string) *laundry {
	l := &laundry{
		r: &router{
			handlers: make(map[string]map[string]handler),
		},
	}

	baskets := loadBasket(basketDir)

	for _, basket := range baskets {
		for _, endpoint := range basket.Endpoints {
			l.r.HandleFunc(endpoint.Method, endpoint.Pattern, basket.Target.IP+basket.Target.Port)
		}
	}

	return l
}
