package ws

import (
	"apiRequest/consts"
	"apiRequest/signature"
	"apiRequest/structs"
	"fmt"
)

type PrivateWsCtrl struct {
	env structs.Env
}

func NewPrivateWsCtrl(env *structs.Env) *PrivateWsCtrl {
	return &PrivateWsCtrl{
		env: *env,
	}
}

func (t *PrivateWsCtrl) GetListenKey() {
	requestURL := fmt.Sprintf(t.env.Url + consts.ListenKey)
	requestPath := fmt.Sprintf(consts.ListenKey)

	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &t.env)
}

func (t *PrivateWsCtrl) Extend() {
	requestURL := fmt.Sprintf(t.env.Url+consts.ExtendListenKey+"?listenKey=%s", "a04291a7ba8e6cb6e695763b6b0c5132")
	requestPath := fmt.Sprintf(consts.ExtendListenKey+"?listenKey=%s", "a04291a7ba8e6cb6e695763b6b0c5132")
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &t.env)
}
