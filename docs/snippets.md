# go代码判断

~~~go
//数据库连接自动分析匹配

// 匹配的连接如: mysql:test:test888@tcp(127.0.0.1:3306)/go_gotms
// It uses user passed database configuration.
  if in.Link != "" {
    var (
      tempGroup = gtime.TimestampNanoStr()
      match, _  = gregex.MatchString(`([a-z]+):(.+)`, in.Link)
    )
    if len(match) == 3 {
      gdb.AddConfigNode(tempGroup, gdb.ConfigNode{
        Type: gstr.Trim(match[1]),
        Link: gstr.Trim(match[2]),
      })
      if db, err = gdb.Instance(tempGroup); err != nil {
        mlog.Debugf(`database initialization failed: %+v`, err)
      }
    } else {
      mlog.Fatalf(`invalid database configuration: %s`, in.Link)
    }
  } else {
    db = g.DB(in.Group)
  }
  if db == nil {
    mlog.Fatal(`database initialization failed, may be invalid database configuration`)
  }



//获取所有数据库表
var tableNames []string
  if in.Tables != "" {
    tableNames = gstr.SplitAndTrim(in.Tables, ",")
  } else {
    tableNames, err = db.Tables(context.TODO())
    if err != nil {
      mlog.Fatalf("fetching tables failed: %+v", err)
    }
  }

// Table excluding.
  if in.TablesEx != "" {
    array := garray.NewStrArrayFrom(tableNames)
    for _, v := range gstr.SplitAndTrim(in.TablesEx, ",") {
      array.RemoveValue(v)
    }
    tableNames = array.Slice()
  }

// 获取指定表的所有字段信息
// Generating table data preparing.
  fieldMap, err := db.TableFields(ctx, in.TableName)
  if err != nil {
    mlog.Fatalf(`fetching tables fields failed for table "%s": %+v`, in.TableName, err)
  }

//合并字符串
path := gfile.Join(dirPathDao, "internal", fileName+".go")


// formatComment formats the comment string to fit the golang code without any lines.
func formatComment(comment string) string {
  comment = gstr.ReplaceByArray(comment, g.SliceStr{
    "\n", " ",
    "\r", " ",
  })
  comment = gstr.Replace(comment, `\n`, " ")
  comment = gstr.Trim(comment)
  return comment
}


// map查找,忽略符号
// MapPossibleItemByKey tries to find the possible key-value pair for given key ignoring cases and symbols.
//
// Note that this function might be of low performance.
func MapPossibleItemByKey(data map[string]interface{}, key string) (foundKey string, foundValue interface{}) {
  if len(data) == 0 {
    return
  }
  if v, ok := data[key]; ok {
    return key, v
  }
  // Loop checking.
  for k, v := range data {
    if EqualFoldWithoutChars(k, key) {
      return k, v
    }
  }
  return "", nil
}
// EqualFoldWithoutChars checks string `s1` and `s2` equal case-insensitively,
// with/without chars '-'/'_'/'.'/' '.
func EqualFoldWithoutChars(s1, s2 string) bool {
  return strings.EqualFold(RemoveSymbols(s1), RemoveSymbols(s2))
}
//移除符号
// RemoveSymbols removes all symbols from string and lefts only numbers and letters.
func RemoveSymbols(s string) string {
  var b = make([]rune, 0, len(s))
  for _, c := range s {
    if c > 127 {
      b = append(b, c)
    } else if (c >= '0' && c <= '9') || (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
      b = append(b, c)
    }
  }
  return string(b)
}
~~~


~~~go

func generateDaoInternal(
  in cGenDaoInternalInput,
  tableNameCamelCase, tableNameCamelLowerCase, importPrefix string,
  dirPathDao, fileName string,
  fieldMap map[string]*gdb.TableField,
) {
  path := gfile.Join(dirPathDao, "internal", fileName+".go")

  // getTplDaoInternalContent("") 获取的是预先定义的模板的内容,参考
  // /cmd/gf/internal/consts/consts_gen_dao_template_dao.go
  // 引用路径: github.com/gogf/gf/cmd/gf/v2/internal/consts
  modelContent := gstr.ReplaceByMap(getTplDaoInternalContent(""), g.MapStrStr{
    tplVarImportPrefix:            importPrefix,
    tplVarTableName:               in.TableName,
    tplVarGroupName:               in.Group,
    tplVarTableNameCamelCase:      tableNameCamelCase,
    tplVarTableNameCamelLowerCase: tableNameCamelLowerCase,
    tplVarColumnDefine:            gstr.Trim(generateColumnDefinitionForDao(fieldMap)),
    tplVarColumnNames:             gstr.Trim(generateColumnNamesForDao(fieldMap)),
  })
  modelContent = replaceDefaultVar(in, modelContent)
  if err := gfile.PutContents(path, strings.TrimSpace(modelContent)); err != nil {
    mlog.Fatalf("writing content to '%s' failed: %v", path, err)
  } else {
    utils.GoFmt(path)
    mlog.Print("generated:", path)
  }
}
// 获取 dao内部 模板内容
func getTplDaoInternalContent(tplDaoInternalPath string) string {
  if tplDaoInternalPath != "" {
    return gfile.GetContents(tplDaoInternalPath)
  }
  return consts.TemplateDaoDaoInternalContent
}


// generateColumnDefinitionForDao generates and returns the column names definition for specified table.
func generateColumnDefinitionForDao(fieldMap map[string]*gdb.TableField) string {
  var (
    buffer = bytes.NewBuffer(nil)
    array  = make([][]string, len(fieldMap))
    names  = sortFieldKeyForDao(fieldMap)
  )
  for index, name := range names {
    var (
      field   = fieldMap[name]
      comment = gstr.Trim(gstr.ReplaceByArray(field.Comment, g.SliceStr{
        "\n", " ",
        "\r", " ",
      }))
    )
    array[index] = []string{
      "    #" + gstr.CaseCamel(field.Name),
      " # " + "string",
      " #" + fmt.Sprintf(`// %s`, comment),
    }
  }
  tw := tablewriter.NewWriter(buffer)
  tw.SetBorder(false)
  tw.SetRowLine(false)
  tw.SetAutoWrapText(false)
  tw.SetColumnSeparator("")
  tw.AppendBulk(array)
  tw.Render()
  defineContent := buffer.String()
  // Let's do this hack of table writer for indent!
  defineContent = gstr.Replace(defineContent, "  #", "")
  buffer.Reset()
  buffer.WriteString(defineContent)
  return buffer.String()
}



// generateColumnNamesForDao generates and returns the column names assignment content of column struct
// for specified table.
func generateColumnNamesForDao(fieldMap map[string]*gdb.TableField) string {
  var (
    buffer = bytes.NewBuffer(nil)
    array  = make([][]string, len(fieldMap))
    names  = sortFieldKeyForDao(fieldMap)
  )
  for index, name := range names {
    field := fieldMap[name]
    array[index] = []string{
      "            #" + gstr.CaseCamel(field.Name) + ":",
      fmt.Sprintf(` #"%s",`, field.Name),
    }
  }
  tw := tablewriter.NewWriter(buffer)
  tw.SetBorder(false)
  tw.SetRowLine(false)
  tw.SetAutoWrapText(false)
  tw.SetColumnSeparator("")
  tw.AppendBulk(array)
  tw.Render()
  namesContent := buffer.String()
  // Let's do this hack of table writer for indent!
  namesContent = gstr.Replace(namesContent, "  #", "")
  buffer.Reset()
  buffer.WriteString(namesContent)
  return buffer.String()
}



// getJsonTagFromCase call gstr.Case* function to convert the s to specified case.
func getJsonTagFromCase(str, caseStr string) string {
  switch gstr.ToLower(caseStr) {
  case gstr.ToLower("Camel"):
    return gstr.CaseCamel(str)

  case gstr.ToLower("CamelLower"):
    return gstr.CaseCamelLower(str)

  case gstr.ToLower("Kebab"):
    return gstr.CaseKebab(str)

  case gstr.ToLower("KebabScreaming"):
    return gstr.CaseKebabScreaming(str)

  case gstr.ToLower("Snake"):
    return gstr.CaseSnake(str)

  case gstr.ToLower("SnakeFirstUpper"):
    return gstr.CaseSnakeFirstUpper(str)

  case gstr.ToLower("SnakeScreaming"):
    return gstr.CaseSnakeScreaming(str)
  }
  return str
}


// 模型字段创建
// 
// generateStructFieldForModel generates and returns the attribute definition for specified field.
func generateStructFieldDefinition(field *gdb.TableField, in generateStructDefinitionInput) []string {
  var (
    typeName string
    jsonTag  = getJsonTagFromCase(field.Name, in.JsonCase)
  )
  t, _ := gregex.ReplaceString(`\(.+\)`, "", field.Type)
  t = gstr.Split(gstr.Trim(t), " ")[0]
  t = gstr.ToLower(t)
  switch t {
  case "binary", "varbinary", "blob", "tinyblob", "mediumblob", "longblob":
    typeName = "[]byte"

  case "bit", "int", "int2", "tinyint", "small_int", "smallint", "medium_int", "mediumint", "serial":
    if gstr.ContainsI(field.Type, "unsigned") {
      typeName = "uint"
    } else {
      typeName = "int"
    }

  case "int4", "int8", "big_int", "bigint", "bigserial":
    if gstr.ContainsI(field.Type, "unsigned") {
      typeName = "uint64"
    } else {
      typeName = "int64"
    }

  case "real":
    typeName = "float32"

  case "float", "double", "decimal", "smallmoney", "numeric":
    typeName = "float64"

  case "bool":
    typeName = "bool"

  case "datetime", "timestamp", "date", "time":
    if in.StdTime {
      typeName = "time.Time"
    } else {
      typeName = "*gtime.Time"
    }
  case "json", "jsonb":
    if in.GJsonSupport {
      typeName = "*gjson.Json"
    } else {
      typeName = "string"
    }
  default:
    // Automatically detect its data type.
    switch {
    case strings.Contains(t, "int"):
      typeName = "int"
    case strings.Contains(t, "text") || strings.Contains(t, "char"):
      typeName = "string"
    case strings.Contains(t, "float") || strings.Contains(t, "double"):
      typeName = "float64"
    case strings.Contains(t, "bool"):
      typeName = "bool"
    case strings.Contains(t, "binary") || strings.Contains(t, "blob"):
      typeName = "[]byte"
    case strings.Contains(t, "date") || strings.Contains(t, "time"):
      if in.StdTime {
        typeName = "time.Time"
      } else {
        typeName = "*gtime.Time"
      }
    default:
      typeName = "string"
    }
  }

  var (
    tagKey = "`"
    result = []string{
      "    #" + gstr.CaseCamel(field.Name),
      " #" + typeName,
    }
    descriptionTag = gstr.Replace(formatComment(field.Comment), `"`, `\"`)
  )

  result = append(result, " #"+fmt.Sprintf(tagKey+`json:"%s"`, jsonTag))
  result = append(result, " #"+fmt.Sprintf(`description:"%s"`+tagKey, descriptionTag))
  result = append(result, " #"+fmt.Sprintf(`// %s`, formatComment(field.Comment)))

  for k, v := range result {
    if in.NoJsonTag {
      v, _ = gregex.ReplaceString(`json:".+"`, ``, v)
    }
    if !in.DescriptionTag {
      v, _ = gregex.ReplaceString(`description:".*"`, ``, v)
    }
    if in.NoModelComment {
      v, _ = gregex.ReplaceString(`//.+`, ``, v)
    }
    result[k] = v
  }
  return result
}

~~~
