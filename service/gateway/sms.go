package gateway

import (
	"errors"
	"self-discipline/global"
	"self-discipline/model/user"
	"self-discipline/utils"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

type SmsService struct{}

//插入短信
func (s *SmsService) CreateSms(sms user.PhoneSms) (errMsg string, err error) {

	if num, err := s.remarkSms(sms.Phone); err != nil {
		if num >= 10 {
			errMsg = err.Error()
		}
		return errMsg, err
	}
	sms.Code = utils.RandomStr(1, 6)

	//短信接口
	/* url := "http://www.xxxsms.com/"
	form := make(httpURL.Values)
	form["code"] = []string{sms.Code}
	body, err := httpclient.Get(url, form,
		httpclient.WithTTL(time.Second*5),
		httpclient.WithHeader("Authorization", "demoGetAuthorization")
	) */
	err = global.DB.Create(&sms).Error
	return errMsg, err
}

//redis记录1小时手机号短信次数
func (*SmsService) remarkSms(phone string) (num int, err error) {

	rkey := utils.MergeStr([]string{"self:", phone})
	recode, err := global.REDIS.Get(rkey).Result()
	if err != nil {
		if err != redis.Nil {
			return num, err
		}
		recode = "0"
	}

	num, err = strconv.Atoi(recode)
	if err != nil {
		return num, err
	}
	if num >= 10 {
		return num, errors.New("1小时内限制10条短信")
	}

	if err := global.REDIS.Incr(rkey).Err(); err != nil {
		return num, err
	}
	if recode == "" {
		if err := global.REDIS.Expire(rkey, 3600*time.Second).Err(); err != nil {
			return num, err
		}
	}
	return num, nil
}

//查找短信
func (s *SmsService) FindSmsByPhone(sms user.PhoneSms) (user.PhoneSms, error) {

	err := global.DB.Last(&sms).Error
	return sms, err
}