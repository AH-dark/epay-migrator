package actions

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/star-horizon/epay-database-mingrator/internal/conf"
	"github.com/star-horizon/epay-database-mingrator/internal/utils"
	"github.com/star-horizon/epay-database-mingrator/model"
)

var defaultDatabaseConfig = map[string]string{
	"version":              "2024",
	"admin_user":           "admin",
	"admin_pwd":            "123456",
	"admin_paypwd":         "123456",
	"homepage":             "0",
	"sitename":             "聚合易支付",
	"title":                "聚合易支付 - 行业领先的免签约支付平台",
	"keywords":             "聚合易支付,支付宝免签约即时到账,财付通免签约,微信免签约支付,QQ钱包免签约,免签约支付",
	"description":          "聚合易支付是XX公司旗下的免签约支付产品，完美解决支付难题，一站式接入支付宝，微信，财付通，QQ钱包,微信wap，帮助开发者快速集成到自己相应产品，效率高，见效快，费率低！",
	"orgname":              "XX公司",
	"kfqq":                 "123456789",
	"template":             "index1",
	"pre_maxmoney":         "1000",
	"blockname":            "云盘|网盘|Q币",
	"blockalert":           "温馨提醒该商品禁止出售，如有疑问请联系网站客服！",
	"settle_open":          "1",
	"settle_type":          "1",
	"settle_money":         "30",
	"settle_rate":          "0.5",
	"settle_fee_min":       "0.1",
	"settle_fee_max":       "20",
	"settle_alipay":        "1",
	"settle_wxpay":         "1",
	"settle_qqpay":         "1",
	"settle_bank":          "0",
	"transfer_alipay":      "0",
	"transfer_wxpay":       "0",
	"transfer_qqpay":       "0",
	"transfer_name":        "聚合易支付",
	"transfer_desc":        "聚合易支付自动结算",
	"login_qq":             "0",
	"login_alipay":         "0",
	"login_alipay_channel": "0",
	"login_wx":             "0",
	"login_wx_channel":     "0",
	"reg_open":             "1",
	"reg_pay":              "1",
	"reg_pre_uid":          "1000",
	"reg_pre_price":        "5",
	"verifytype":           "1",
	"test_open":            "1",
	"test_pre_uid":         "1000",
	"mail_cloud":           "0",
	"mail_smtp":            "smtp.qq.com",
	"mail_port":            "465",
	"mail_name":            "",
	"mail_pwd":             "",
	"sms_api":              "0",
	"captcha_open":         "1",
	"captcha_id":           "",
	"captcha_key":          "",
	"onecode":              "1",
	"recharge":             "1",
	"pageordername":        "1",
	"notifyordername":      "1",
}

func createDatabaseConfig(ctx context.Context, configs map[string]string, db *gorm.DB) error {
	db = db.WithContext(ctx)

	for k, v := range configs {
		logger := logrus.WithContext(ctx).WithField("key", k).WithField("val", v)
		logger.Debug("create default config")

		// check if config exists
		if err := db.Model(&model.Config{}).
			Where("k = ?", k).
			First(&model.Config{}).
			Error; err == nil {
			logger.Debug("config already exists")
			continue
		} else if err != gorm.ErrRecordNotFound {
			logger.WithError(err).Error("check default config failed")
			return err
		}

		// create config
		if err := db.
			Model(&model.Config{}).
			Create(&model.Config{
				Key: k,
				Val: v,
			}).Error; err != nil {
			logger.WithError(err).Error("create default config failed")
			return err
		}
	}

	return nil
}

func RunMigrate(ctx context.Context, db *gorm.DB) error {
	db = db.WithContext(ctx)

	logrus.WithContext(ctx).Info("auto migrate database tables")
	if err := db.AutoMigrate(
		&model.AlipayRisk{},
		&model.Anounce{},
		&model.Batch{},
		&model.Channel{},
		&model.Config{},
		&model.Domain{},
		&model.Group{},
		&model.Log{},
		&model.Order{},
		&model.Plugin{},
		&model.Record{},
		&model.RegCode{},
		&model.Risk{},
		&model.Roll{},
		&model.Settle{},
		&model.Type{},
		&model.User{},
		&model.Weixin{},
	); err != nil {
		logrus.WithContext(ctx).WithError(err).Error("auto migrate database tables failed")
		return err
	}

	logrus.WithContext(ctx).Info("create default config")
	if err := createDatabaseConfig(ctx, defaultDatabaseConfig, db); err != nil {
		logrus.WithContext(ctx).WithError(err).Panic("create default config failed")
		return err
	}

	logrus.WithContext(ctx).Info("init app config")
	initData := map[string]string{
		"syskey":  lo.If(lo.IsNotEmpty(conf.AppConfig.SysKey), conf.AppConfig.SysKey).Else(utils.RandString(32)),
		"build":   time.Now().Format("2006-01-02"),
		"cronkey": lo.If(lo.IsNotEmpty(conf.AppConfig.CronKey), conf.AppConfig.CronKey).Else(strconv.Itoa(utils.RandInt(100000, 999999))),
	}
	if err := createDatabaseConfig(ctx, initData, db); err != nil {
		logrus.WithContext(ctx).WithError(err).Panic("create app init config failed")
		return err
	}

	fmt.Printf(`
System Key: %s
Build Time: %s
Cron Key: %s

Admin Username: admin
Admin Password: 123456
`, initData["syskey"], initData["build"], initData["cronkey"])

	return nil
}
