package model

import (
	"fmt"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

//公共模型
type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

//创建一个DB实例，
func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	//链接mysql
	db, err := gorm.Open(databaseSetting.DBType, fmt.Sprintf(s,
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	))
	if err != nil {
		return nil, err
	}
	//开启日志
	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.DB().SetMaxOpenConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)

	//将下面的三个Callback方法进行回调注册，
	db.SingularTable(true)
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)

	//返回
	return db, nil
}

//新增行为的回调
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	// HasError check if there are any error
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		// FieldByName find `gorm.Field` with field name or db name
		//FieldByName(name string) (field *Field, ok bool)
		//获取当前是否包含所需的字段
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank { //得知该字段的值是否为空
				//为空，重新设值
				//(field *Field) Set(value interface{}) (err error)
				_ = createTimeField.Set(nowTime)
			}
		}
		//
		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				_ = modifyTimeField.Set(nowTime)
			}
		}
	}

}

//更新行为的回调
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	//通过调用 scope.Get("gorm:update_column") 去获取当前设置了标识 gorm:update_column 的字段属性
	if _, ok := scope.Get("gorm:update_column"); !ok {
		//不存在，也就是没有自定义设置 update_column，那么将会在更新回调内设置默认字段 ModifiedOn 的值为当前的时间戳。
		_ = scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

//删除行为的回调
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string

		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")
		isDelField, hasIsDelField := scope.FieldByName("IsDel")

		if !scope.Search.Unscoped && hasDeletedOnField && hasIsDelField {
			//存在执行软删除
			now := time.Now().Unix()
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v,%v=%v%v%v",
				scope.QuotedTableName(),            //返回引用的表名，获取当前所引用的表名，并回调一系列方法针对SQL语句的组成部分进行处理和转移
				scope.Quote(deletedOnField.DBName), // Quote used to quote string to escape them for database
				scope.AddToVars(now),               //AddToVars添加值作为sql的变量，用于防止sql注入
				scope.Quote(isDelField.DBName),     //
				scope.AddToVars(1),
				addExtraSpaceIfExist(scope.CombinedConditionSql()), //调用 scope.CombinedConditionSql 方法完成 SQL 语句的组装
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
