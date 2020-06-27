package apis

import (
	"PowerTranslate/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type TransData struct {
	Q           string `json:"input"`
	TransRouter string `json:"translate_route"`
}

func TransApi(c *gin.Context) {
	data := &TransData{}
	err := c.ShouldBindJSON(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	routes := strings.Split(data.TransRouter, "-")
	if len(routes) > 10 {
		c.JSON(http.StatusForbidden, nil)
		return
	}
	TransRet := data.Q
	for i := 1; i < len(routes); i++ {
		ret, err := utils.BaiduTransAPI(TransRet, routes[i-1], routes[i])
		if err != nil || ret["error_code"] != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"trans_result": ret, "error": err})
			return
		}
		transResults := ret["trans_result"].([]interface{})
		var dsts []string
		for _, v := range transResults {
			transRes := v.(map[string]interface{})
			dsts = append(dsts, transRes["dst"].(string))
		}
		TransRet = strings.Join(dsts, "\n")
	}
	c.JSON(http.StatusOK, map[string]interface{}{"data": TransRet, "code": 2000})
}
