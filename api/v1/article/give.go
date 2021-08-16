package article

import "github.com/gin-gonic/gin"

// @Tags Base
// @Summary 点赞
// @Accept application/x-www-form-urlencoded
// @Produce  application/json
// @Param data formData articleReq.Issue true "用户标识, 内容"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /v1/article/issue [post]
func (h *Handler) Give(ctx *gin.Context) {

}
