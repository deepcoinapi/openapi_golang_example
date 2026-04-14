package rebate

import (
	"apiRequest/consts"
	"apiRequest/signature"
	"apiRequest/structs"
	"encoding/json"
	"fmt"
)

type RebateCtrl struct {
	env structs.Env
}

func NewRebateCtrl(env *structs.Env) *RebateCtrl {
	return &RebateCtrl{
		env: *env,
	}
}

// RebateConfig 返佣比例查询
func (t *RebateCtrl) RebateConfig() {
	requestURL := fmt.Sprintf(t.env.Url+consts.REBATE_CONFIG+"?uid=%d", 36007196)
	requestPath := fmt.Sprintf(consts.REBATE_CONFIG+"?uid=%d", 36007196)
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &t.env)
}

// SetRebateConfig 返佣比例设置
func (t *RebateCtrl) SetRebateConfig() {
	type setRebateConfigRequest struct {
		Uid  int `json:"uid"`
		Rate int `json:"rate"`
	}

	req := &setRebateConfigRequest{
		Uid:  36023405, // 下级 uid
		Rate: 10,       // 返佣比例
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	requestURL := t.env.Url + consts.REBATE_CONFIG
	signature.DoHttp(requestURL, consts.HTTP_METHOD_POST, consts.REBATE_CONFIG, string(requestBody), &t.env)
}

// AgentsUsers 合伙人查询
func (t *RebateCtrl) AgentsUsers() {
	requestURL := fmt.Sprintf(t.env.Url+consts.AGENTS_USERS+"?uid=%d&startTime=%d&endTime=%d", 36007196, 1700000000, 1800000000)
	requestPath := fmt.Sprintf(consts.AGENTS_USERS+"?uid=%d&startTime=%d&endTime=%d", 36007196, 1700000000, 1800000000)
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &t.env)
}

// AgentsRebateList 返佣明细查询
func (t *RebateCtrl) AgentsRebateList() {
	requestURL := fmt.Sprintf(t.env.Url+consts.AGENTS_REBATE_LIST+"?uid=%d&type=%d&startTime=%d&endTime=%d&pageNum=%d&pageSize=%d", 36007196, 0, 1776081649, 1776210092, 1, 100)
	requestPath := fmt.Sprintf(consts.AGENTS_REBATE_LIST+"?uid=%d&type=%d&startTime=%d&endTime=%d&pageNum=%d&pageSize=%d", 36007196, 0, 1776081649, 1776210092, 1, 100)
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &t.env)
}

// AgentsRebates 返佣汇总查询
func (t *RebateCtrl) AgentsRebates() {
	requestURL := fmt.Sprintf(t.env.Url+consts.AGENTS_REBATES+"?uid=%d&type=%d&startTime=%d&endTime=%d", 36007196, 0, 1776081649, 1776210092)
	requestPath := fmt.Sprintf(consts.AGENTS_REBATES+"?uid=%d&type=%d&startTime=%d&endTime=%d", 36007196, 0, 1776081649, 1776210092)
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &t.env)
}
