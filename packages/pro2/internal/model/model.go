package model

import (
	"fmt"
	"go-project-demo/packages/pro2/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	db, err := gorm.Open(mysql.Open(fmt.Sprintf(s,
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	//db.SingularTable(true)
	//db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	//db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	//db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	//db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	//db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)
	//otgorm.AddGormCallbacks(db)
	return db, nil
}
