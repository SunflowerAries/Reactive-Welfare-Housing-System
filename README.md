# Reactive-Welfare-Housing-System

写代码前，请先简单了解一下 

- [Go 的包管理](https://zhuanlan.zhihu.com/p/60703832)，目前可以通过 go install 将 go.mod 中的依赖包安装
- [Go 的 sql1](http://go-database-sql.org/overview.html)，[Go 的 sql2](https://xuchao918.github.io/2019/06/13/Go%E6%93%8D%E4%BD%9CMySql%E6%95%B0%E6%8D%AE%E5%BA%93%E7%9A%84%E6%96%B9%E5%BC%8F/)
- [Protobuf](https://juejin.im/post/5d81bb5cf265da03ae78ab7b) https://developers.google.com/protocol-buffers/docs/gotutorial
- [博客文档](https://blog.oklahome.net/)

## 参与开发

需要手动安装 [protobuf](https://grpc.io/docs/quickstart/go/) 

建议新建一个用户，方式为

```sql
-- 创建用户
CREATE USER 'housing'@'localhost' IDENTIFIED WITH mysql_native_password BY 'housing@2020';
-- 授予权限
GRANT all privileges on *.* to 'housing'@'localhost';
```

需要手动创建一个 `mysql` 的数据库

```sql
create database housing
```

并在数据库中创建表：

```sql
use housing
CREATE TABLE family (
	id INT AUTO_INCREMENT PRIMARY KEY,
    family_number INT NOT NULL,
    incomes INT NOT NULL,
    house_id INT
);

CREATE TABLE user (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    family_id INT NOT NULL,
    FOREIGN KEY (family_id) REFERENCES family(id),
    incomes INT NOT NULL,
    INDEX(name),
    INDEX(family_id)
);
```

为了便于开发维护，建议开发每个新的功能点时都新建分支。

### 写代码之前

- 无论你现在在哪个分支，先切换到 `master` 分支 `git checkout master`
- 确保你的本地代码是最新的 `git fetch origin` 然后 `git rebase origin/master`
- 创建一个新的分支，分支名应该与你准备开发的内容有关，如 `git checkout -b update-readme-contributing` （**注意** 请尽量准确描述你工作的内容，并且在分支被合并后及时删除分支）
- 在新的分支进行开发，并 `commit` 代码，`commit` 之前须 `git status` 确保不会提交不相关的文件

### 提交代码之前

- 由于你开发期间 `master` 可能有过更新，所以请再确保 `git` 历史是最新的
  - 更新仓库 `git fetch origin`
  - 载入 `master` 的更新 `git rebase origin/master`
    - 如果代码可以无法 `fast-forward`，说明你修改过的代码在你开发期间被其他人修改过，你需要手动解决冲突
    - 对于每一个有冲突的代码文件，删除冲突提示，并修改代码，然后 `git add path/to/conflict/file`
    - 都解决完后继续 rebase `git rebase --continue`
- 上传你的分支 `git push --set-upstream origin name-of-your-branch`
- 到本页面手动发起 pull request，并描述清楚你的改动