package laundry

import "net/http"

type laundry struct {
	r *router
}

type Config struct {
	BasketDir string
}

func (l *laundry) Run(addr string) error {
	return http.ListenAndServe(addr, l.r)
}

func New(c ...Config) *laundry {

	l := &laundry{
		r: &router{
			handlers: make(map[string]map[string]handler),
		},
	}

	var baskets []basket
	if len(c) > 0 {
		baskets = loadBasket(c[0].BasketDir)
	} else {
		baskets = loadBasket("./basket")
	}

	for _, basket := range baskets {
		for _, endpoint := range basket.Endpoints {
			l.r.HandleFunc(endpoint.Method, endpoint.Pattern, basket)
		}
	}

	return l
}
