# MiRolls deploy tutorial

> Author not native English speaker, so some sentience maybe has Garammatical, please excuse me, sorry!

> MiRolls use gin and vue frame build, if you can write code, you can change the source code

**If you need build the *Mirolls* you need have a normal the server manage experience. You must can insert the ".sql" file, and must use nginx/apache... to bind Your domain name. **

### Use Normal theme build website

1. Download [Releases@latest](https://github.com/liangmiQwQ/MiRolls-/releases), Choose your server version(Such as Linux arm64.zip or Windows amd64.zip and more)

   ![image](https://img.lmfans.cn/i/2023/01/26/10zj0ty.png)

2. Open config/config.yaml file, and change this file (follow this form)

  | Key                | Value           | Remark                                                       |
  | ------------------ | --------------- | ------------------------------------------------------------ |
  | Port(server)       | Number          | Server port(Such as 127.0.0.1:2202)                          |
  | Static(server)     | String,FilePath | Theme folder(default vite/dist)                              |
  | Username(database) | String          | If you'll build a MiRolls website, you must be a database. Enter your database username |
  | Password(database) | String          | Enter your database password                                 |
  | Protocol(database) | String          | Use tcp without special requirements                         |
  | Host(database)     | String          | If you use the remote database, else enter localhost         |
  | Post(database)     | Number          | If you change database post                                  |
  | Database(database) | String          | Database name                                                |
  | Name(Site)         | String          | Your website name(such as MiRolls)                           |
  | Link(Site)         | String          | Your website domain(IP)                                      |
  | MainColor(Site)    | String          | Theme color, can use"rgb()"function, or use "#xxxxx" or all colors that can be inserted into css |
  | icp(Site)          | String          | icp备案，其他国家的人貌似没有这个需求，就当是放在页脚的一个提示语！icp(A Policy for Chinese Sites) You can be put on footer, For example, write a promotion for this site here |

3. After configuration, you need import the sql file(database.sql)

   ![image](https://img.lmfans.cn/i/2023/01/26/126azd4.png)

   ***Tips: This Step is very important.If you cannot import, please use some server manage tools(such as cPanel and Plesk)***

4. Open shell(cmd), run MiRolls!

```shell
$ cd /mirolls/website/mirolls
$ ./mirolls
// server running now
```

```cmd
cd c:/User/mirolls/Desktop/server/
./mirolls.exe
```



***EnJoy the MiRolls***