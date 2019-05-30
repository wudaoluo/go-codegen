# go-codegen

### 使用方式
    > 修改配置文件
    default:
      logfile: log/codegen.log
    
    mysql:
      DBuser : root
      DBpasswd : ******
      DBaddr : 127.0.0.1
      DBport : 3306
      DBname : codegen
      DBmaxconn : 100
      DBidleconn : 80

   ```
   ./codegen -h
   
   Usage:
     codegen [command]
   
   Available Commands:
     help        Help about any command
     mongo       A brief description of your command
     mysql       生成mysql语句和markdown文档
     update      
   获取最新二进制程序
     version     
   打印当前版本
   
   Flags:
         --config string    config file (default "codegen.yaml")
         --debug            开启debug模式 开启后不生成任何文件
     -h, --help             help for go-codegen
     -o, --outPath string   指定生成的文件路径 (default ".")
   
   Use "codegen [command] --help" for more information about a command.
   
   ```
   
   ### 生成mysql链接
    codegen mysql --add init -o 指定目录
   
   ### 生成表结构和curd语句
    codegen mysql --add 表名 -o 指定目录
    
   ### 生成表文档
    codegen mysql --doc -o 指定目录
    
![avatar](https://github.com/wudaoluo/go-codegen/blob/master/mysql_doc.jpg)
    

    
    
