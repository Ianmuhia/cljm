package dbrepo

const (
	CheckUserExists         = `SELECT id from users WHERE email = $1`
	LoginQuery              = `SELECT * from users WHERE email = $1`
	CreateUserQuery         = `INSERT INTO users(id,name,password,email) VALUES (DEFAULT, $1 , $2, $3);`
	UpdateUserPasswordQuery = `UPDATE users SET password = $2 WHERE id = $1`
)

//CreateUsersTable //creates users table

func (m *PostgresDBRepo) CreateUsersTable() {
	m.DB.Query(`
		CREATE TABLE IF NOT EXISTS users( id serial PRIMARY KEY, name VARCHAR (100) NOT NULL, password VARCHAR (355) NOT NULL, email VARCHAR (355) UNIQUE NOT NULL, created_on TIMESTAMP NOT NULL default current_timestamp,updated_at TIMESTAMP NOT NULL default current_timestamp )`,
	)
}

