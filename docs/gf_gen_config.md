# 代码生成配置文件说明

修改 importPrefix 为你的dao包的导入前缀, 一般为你的模块名称, 如: gotms

~~~yaml

# 获取当前所有数据库表的信息  
# select * from information_schema.tables WHERE table_schema = (select database());
#  gf gen dao -p ./model -c config.yaml
gfcli:
  gen:
    dao:
      # 修改数据库连接为你自己的信息
    - link:     "mysql:test:test888@tcp(127.0.0.1:3306)/go_gotms"
      # 将所有要生成的表名称放到这里,多个逗号分隔
      # 指定当前数据库中需要执行代码生成的数据表。如果为空，表示数据库的所有表都会生成。
      tables:   ""
      # Tables Excluding，指定当前数据库中需要排除代码生成的数据表。 如: gen_table
      tablesEx:   ""
      path:   "./app/admin" # 代码生成后保存路径, 需要先创建 mkdir -p ./app/admin
      # json输出的命名分隔,  CamelLower 小驼峰(默认)  Snake 小写下划线
      jsonCase: "CamelLower"
       # 在数据库配置中的数据库分组名称。只能配置一个名称。数据库在配置文件中的分组名称往往确定之后便不再修改。
      group: "default"
      # 生成数据库对象及文件的前缀，以便区分不同数据库或者不同数据库中的相同表名，防止数据表同名覆盖。 如: sys_,user_
      prefix: "sys_" # 
      # 删除数据表的指定前缀名称。多个前缀以,号分隔。 如: gf_,sys_
      removePrefix: "sys_"
      # 当数据表字段类型为JSON类型时，代码生成的属性类型使用*gjson.Json类型。默认 false
      gJsonSupport: true
      # 当数据表字段类型为时间类型时，代码生成的属性类型使用标准库的time.Time而不是框架的*gtime.Time类型。 默认 false
      stdTime: false
      # 每次生成dao代码时是否重新生成覆盖dao/internal目录外层的文件。注意dao/internal目录外层的文件可能由开发者自定义扩展了功能，覆盖可能会产生风险。 默认 false
      overwriteDao: false
      # 用于指定生成Go文件的import路径前缀。特别是针对于不是在项目根目录下使用gen dao命令，或者想要将代码文件生成到自定义的其他目录，这个时候配置该参数十分必要。
      # 默认是 使用当前目录下的 go.mod文件中的配置, 如果需要将代码放到其他地方使用,则这里就必须要指定了!
      importPrefix: "gotms"
      # 用于指定是否为数据模型结构体属性增加desription的标签，内容为对应的数据表字段注释。 默认 false
      descriptionTag: false
      # 用于指定是否关闭数据模型结构体属性的注释自动生成，内容为数据表对应字段的注释。 默认 false
      noModelComment: false

~~~






