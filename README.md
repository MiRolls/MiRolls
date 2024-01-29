# MiRolls Deployment Documentation

> Author not native English speaker, so some sentience maybe has Grammatical error, please excuse me, sorry!

> MiRolls use gin and vue frame build, if you can write code, you can change the source code

**If you need build the *Mirolls* you need have a normal the server manage experience. You must be able to insert the ".sql" file, and must use nginx/apache... to bind Your domain name.**

### Build with the default theme

1. Download [Releases@latest](https://github.com/liangmiQwQ/MiRolls-/releases), Choose your server version(Such as Linux arm64.zip or Windows amd64.zip and more)

2. Open config/config.yaml file, and change this file (follow this form)

  | Key              | Value           | Remark                                                                                                                                         |
  | ------------------ |------------------------------------------------------------------------------------------------------------------------------------------------| ------------------------------------------------------------ |
  | Port(server)     | Number          | Server port(Such as 127.0.0.1:2202)                                                                                                            |
  | Static(server)   | String,FilePath | Theme folder(default vite/dist)                                                                                                                |
  | Username(database) | String          | If you'll build a MiRolls website, you must be a database. Enter your database username                                                        |
  | Password(database) | String          | Enter your database password                                                                                                                   |
  | Protocol(database) | String          | Use tcp without special requirements                                                                                                           |
  | Host(database)   | String          | If you use the remote database, else enter localhost                                                                                           |
  | Port(database)     | Number          | If you change database port                                                                                                                    |
  | Database(database) | String          | Database name                                                                                                                                  |
  | Name(Site)       | String          | Your website name(such as MiRolls)                                                                                                             |
  | Link(Site)       | String          | Your website domain(IP)                                                                                                                        |
  | Logo(Site)      | String          | Favicon and logo.png                                                                                                                           |
  | MainColor(Site)  | String          | Theme color, can use"rgb()"function, or use "#xxxxx" or all colors that can be inserted into css                                               |
  | icp(Site)        | String     | icp备案，其他国家的人貌似没有这个需求，就当是放在页脚的一个提示语！icp(A Policy for Chinese Sites) You can be put on footer, For example, write a promotion for this site here |
  | Lang(Site) | String | Language, only support English and Chinese(en and zh)                                                                                          |
  | NeedIcp(Site) | Number | 0=not need icp, 1= need icp |

3. After configuration, you need import the sql file(database.sql)

   ***Tips: This Step is very important.If you cannot import, please use some server manage tools(such as cPanel and Plesk)***

4. Open shell(cmd powershell), run MiRolls!

```shell
$ cd /mirolls/website/mirolls
$ chmod +x mirolls
$ ./mirolls
```

```cmd
cd c:/User/mirolls/Desktop/server/
mirolls.exe
```

License

[APACHE LICENSE 2.0](https://apache.org/licenses/LICENSE-2.0)

Copyright © 2023-to date,Liangmi

***EnJoy the MiRolls***
