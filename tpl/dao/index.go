/*
* This is auto-generated by GoTMS Gen Tool
* @Author: Tekin
* @Date:   2022-04-16 07:48:22
* @Last Modified by:   tekintian
* @Last Modified time: 2022-05-23 23:13:17
*/

package dao

import (
	"{TplImportPrefix}/app/system/service/internal/dao/internal"
)

// internal{TplTableNameCamelCase}Dao is internal type for wrapping internal DAO implements.
type internal{TplTableNameCamelCase}Dao = *internal.{TplTableNameCamelCase}Dao

// {TplTableNameCamelLowerCase}Dao is the data access object for table {TplTableName}.
// You can define custom methods on it to extend its functionality as you wish.
type {TplTableNameCamelLowerCase}Dao struct {
	internal{TplTableNameCamelCase}Dao
}

var (
	// {TplTableNameCamelCase} is globally public accessible object for table {TplTableName} operations.
	{TplTableNameCamelCase} = {TplTableNameCamelLowerCase}Dao{
		internal.New{TplTableNameCamelCase}Dao(),
	}
)

// Fill with you ideas below.
