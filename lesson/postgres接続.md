package main

import (
	"database/sql"
	// "fmt"
	"log"

	_ "github.com/lib/pq" // installする postgressとの接続用ドライバー
)

var Db *sql.DB
var err error

type Person struct {
	Name string
	Age  int
}

func main() {
	Db, err = sql.Open("postgres", "user=okuyama dbname=testdb password=0121Apple sslmode=disable")
	if err != nil {
		log.Panicln(err)
	}
	defer Db.Close()

	/*
	//C
	cmd := "INSERT INTO persons (name, age) VALUES ($1, $2)"

	_, err = Db.Exec(cmd, "Nancy", 20)

	if err != nil {
		log.Fatalln(err)
	}
	*/

/*
	//R
	cmd := "SELECT * FROM persons where age = $1"
	//QueryRow 1レコード取得
	row := Db.QueryRow(cmd, 20)
	var p Person
	err = row.Scan(&p.Name, &p.Age) //&p(ポインタのP) 変数pにスキャンしたdataを代入して表示する
	if err != nil {
		//データがなかったら
		if err == sql.ErrNoRows {
			log.Println("No row")
		//それ以外のエラー
		} else {
			log.Println(err)
		}
	}
	fmt.Println(p.Name, p.Age) //Nancy 20

	cmd = "SELECT * FROM persons"
	// Queryは条件に合うデータをすべて取得するメソッドです。
	rows, _ := Db.Query(cmd)
	defer rows.Close()

	// 取得したデータを格納するためのPerson型のスライスを用意します。
	var pp []Person
	//取得したデータをループでスライスに追加 for rows.Next()
	for rows.Next() {
		var p Person

		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Println(err)
		}
		// 取得したデータをスライスに追加します。
		pp = append(pp, p)
	}

	//まとめてエラーハンドリングver
	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}

	// 取得したデータを表示します。
	for _, p := range pp {
		fmt.Println(p.Name, p.Age) //Nancy 20

	}
}
*/


	// //U
	// cmd := "UPDATE persons SET age = $1 WHERE name = $2"
	// _, err := Db.Exec(cmd, 25, "Nancy")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	//D
	cmd := "DELETE FROM persons WHERE name = $1"
	_, err = Db.Exec(cmd, "Nancy")
	if err != nil {
		log.Fatalln(err)
	}

}


//psql -U postgres　対話的ターミナルに入る
// ¥l データベース一覧を表示 \qで抜ける
// postgres=# \q でぬける

//create user okuyama with password '??????????'; ユーザー作成
//postgres=# create user okuyama with password '?????????';
//CREATE ROLE ユーザー作成成功


//select usename from pg_user; user一覧を見る
// postgres=# select usename from pg_user;
//  usename
// ----------
//  postgres
//  workuser
//  okuyama
// (3 行)

//postgres=# create database testdb owner okuyama; データベース作成コマンド
// CREATE DATABASE

//go run main.goで接続

//$ psql testdb そうするとこのコマンドでtestdbに入れるようになる
// ユーザー okuyama のパスワード:
// psql (15.3)
// "help"でヘルプを表示します。

// testdb=>
// testdb=> \d テーブル一覧が見れる
// リレーションが見つかりませんでした。　//まだテーブルがないという意味
// testdb=>

// postgresのテーブルは　psql.sqlファイルを自分で作って実行することで作られる

// ## テーブルはexample_psql.sqlを自分で作成しコマンドを実行
//example_psql.sqlの中身
// create table persons (　　　　　　　
//     name     varchar(255),
//     age      integer
// );
//$ psql -U okuyama -f example_psql.sql -d testdb;
// ユーザー okuyama のパスワード:
// CREATE TABLE　成功

//テーブル確認
// $ psql testdb
// ユーザー okuyama のパスワード:
// psql (15.3)
// "help"でヘルプを表示します。

// testdb=> \d
//             リレーション一覧
//  スキーマ |  名前   |  タイプ  | 所有者
// ----------+---------+----------+---------
//  public   | persons | テーブル | okuyama    //テーブル作成成功を確認
// (1 行)

// $ psql testdb
// ユーザー okuyama のパスワード:
// psql (15.3)
// "help"でヘルプを表示します。

// C　実行後(go run main.go)のテーブル確認
// testdb=> select * from persons;
//  name  | age
// -------+-----
//  Nancy |  20
// (1 行)

// database削除　drop database testdb;
// postgres=# drop database testdb;
// DROP DATABASE

// user 削除　drop user okuyama;
// postgres=# drop user okuyama;
// DROP ROLE
// postgres=# select usename from pg_user;
//  usename
// ----------
//  postgres
//  workuser
// (2 行)
