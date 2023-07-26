## DB初期設定
```
package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// database + sqlite3
//テーブル作成

var DB *sql.DB //グローバルな変数として定義

func main() {

	Db, _ := sql.Open("sqlite3", "./example.sql") //現在のディレクトリにexample.sqlがなければ作成する

	defer Db.Close()

	cmd := `CREATE TABLE IF NOT EXISTS persons(
				name STRING,
				age INT)`  //``で書くことで改行して書ける //もしpersonsテーブルがなければ作成すると指定している

	_, err := Db.Exec(cmd) //Execでコマンドを実行する　結果は使わないので_にしている

	if err != nil {
		log.Fatalln(err)
	}
}
//go run main.goで実行する
// sqlite3 example.sql　を実行すると  example.sqlの中に入ることができる
// .tableで テーブルを確認できる
// sqlite> .table
// persons
// .exit でぬける

```

## dataの追加
```
package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// database + sqlite3
//CURD処理

var DB *sql.DB //グローバルな変数として定義

func main() {
	Db, _ := sql.Open("sqlite3", "./example.sql") //現在のディレクトリにexample.sqlがなければ作成する
	defer Db.Close()

	//データの追加
	cmd := "INSERT INTO persons (name, age) VALUES (?, ?)"

	_, err := Db.Exec(cmd, "tarou", 20) //Execでコマンドを実行する　結果は使わないので_にしている

	if err != nil {
		log.Fatalln(err)
	}
}
//go run main.goで実行する
// sqlite3 example.sql　を実行すると  example.sqlの中に入ることができる
// SELECT * FROM persons; でテーブル内のすべてのデータを取り出すことができる
// sqlite> SELECT * FROM persons;
// tarou|20
```

##　データの更新
```
package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// database + sqlite3
//CURD処理

var DB *sql.DB //グローバルな変数として定義

func main() {
	Db, _ := sql.Open("sqlite3", "./example.sql") //現在のディレクトリにexample.sqlがなければ作成する
	defer Db.Close()

	//データの更新
	cmd := "UPDATE persons SET age = ? WHERE name = ?"

	_, err := Db.Exec(cmd, 25,  "tarou") //age = 25 where name "tarou"

	if err != nil {
		log.Fatalln(err)
	}
}
//go run main.goで実行する
// sqlite3 example.sql　を実行すると  example.sqlの中に入ることができる
// SELECT * FROM persons; でテーブル内のすべてのデータを取り出すことができる
// sqlite> SELECT * FROM persons;
// tarou|25  20から25に更新されていることが確認
```

## 特定のデータを取得
```
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// database + sqlite3
//CURD処理

var DB *sql.DB //グローバルな変数として定義

type Person struct {
	Name string
	Age int
}

func main() {
	Db, _ := sql.Open("sqlite3", "./example.sql") //現在のディレクトリにexample.sqlがなければ作成する
	defer Db.Close()

	// //データの更新
	// cmd := "INSERT INTO persons (name, age) VALUES (?, ?)"

	// _, err := Db.Exec(cmd, "hanako", 19)

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	//特定のデータを取得
	cmd := "SELECT * FROM persons where age = ?"
	//QueryRow 1レコード取得
	row := Db.QueryRow(cmd, 25)
	var p Person
	err := row.Scan(&p.Name, &p.Age) //&p(ポインタのP) 変数pにスキャンしたdataを代入して表示する
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No row")
		} else {
			log.Println(err)
		}
	}
	fmt.Println(p.Name, p.Age) //tarou 25

}


//go run main.goで実行する
// sqlite3 example.sql　を実行すると  example.sqlの中に入ることができる
// SELECT * FROM persons; でテーブル内のすべてのデータを取り出すことができる
// sqlite> SELECT * FROM persons;
// tarou|25
// hanako|19
```

## 複数のデータを取得
```
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Personはデータベース内のpersonsテーブルの行に対応するデータを格納します。
type Person struct {
	Name string
	Age  int
}

func main() {
	// SQLiteデータベースに接続します。
	// "./example.sql"は現在のディレクトリに"example.sql"がなければ作成します。
	Db, err := sql.Open("sqlite3", "./example.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer Db.Close()

	// データベースから複数のデータを取得するためのSQLコマンドを作成します。
	cmd := "SELECT * FROM persons"
	// Queryは条件に合うデータをすべて取得するメソッドです。
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 取得したデータを格納するためのPerson型のスライスを用意します。
	var pp []Person
	for rows.Next() {
		var p Person
		// rows.Scanメソッドを使ってデータを取得し、変数pに格納します。
		// &p.Nameと&p.AgeはそれぞれNameとAgeフィールドのアドレスを指定しており、
		// Scanメソッドは取得した値をそれらのアドレスに書き込みます。
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Println(err)
		}
		// 取得したデータをスライスに追加します。
		pp = append(pp, p)
	}

	// 取得したデータを表示します。
	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
		// 出力例:
		// tarou 25
		// hanako 19
	}
}

/*

rows.Next()は、database/sqlパッケージで提供される*Rows（行の集まり）型のメソッドです。これを使うことで、データベースから取得した結果セット内の次の行に移動することができます。

*Rows型は、データベースからのクエリ結果やトランザクションによるデータの取得などで使用されます。*Rows型は、通常、Db.QueryやTx.Queryなどのメソッドを使用して取得されます。そして、その後にrows.Next()を呼び出すことで、結果セットの次の行を読み込むことができます。

rows.Next()メソッドは、次の行が存在する場合にtrueを返し、次の行がない場合（結果セットの末尾に到達した場合）にfalseを返します。また、rows.Scan()メソッドを使用して、実際にその行のデータを取得することができます。

以下に再度示したサンプルコードの一部で、rows.Next()の使用例を示します：

// ...
rows, err := Db.Query(cmd)
if err != nil {
    log.Fatal(err)
}
defer rows.Close()

var pp []Person
for rows.Next() {
    var p Person
    err := rows.Scan(&p.Name, &p.Age)
    if err != nil {
        log.Println(err)
    }
    pp = append(pp, p)
}
// ...
このコードでは、rows.Next()メソッドを使って結果セット内の次の行がある限り、forループが続きます。rows.Scan()を使用して、Person型の変数pにデータを取得し、それをスライスppに追加しています。このようにして、結果セット内のすべての行のデータを取得してppに格納します。
*/
```

## データの削除
```
package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)


var Db *sql.DB

func main() {
	Db, _ := sql.Open("sqlite3", "./example.sql")
	defer Db.Close()

	//データの削除
	cmd := "DELETE FROM persons WHERE name = ?"
	_, err := Db.Exec(cmd, "hanako")
	if err != nil {
		log.Fatalln(err)
	}

}

// $ sqlite3 example.sql  削除前の確認
// SQLite version 3.42.0 2023-05-16 12:36:15
// Enter ".help" for usage hints.
// sqlite> SELECT * FROM persons;
// tarou|25
// hanako|19
// sqlite> .exit


// go run main.go　で実行
// 削除後
//$ sqlite3 example.sql
// SQLite version 3.42.0 2023-05-16 12:36:15
// Enter ".help" for usage hints.
// sqlite> SELECT * FROM persons;
// tarou|25　hanakoのデータ削除に成功

```
