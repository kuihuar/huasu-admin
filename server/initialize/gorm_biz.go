package initialize

import (
	"github.com/kuihuar/huasu-admin/server/global"
	"github.com/kuihuar/huasu-admin/server/model/huasu"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(huasu.CarouselManage{}, huasu.Notice{})
	if err != nil {
		return err
	}
	return nil
}
