package queries

const (
	CreateUserQuery = `CREATE TABLE IF NOT EXISTS users(
		user_id SERIAL PRIMARY KEY,
		username VARCHAR(50) UNIQUE,
		password_hash TEXT 
	);`

	CreateGoalQuery = `CREATE TABLE IF NOT EXISTS goal(
		goal_id SERIAL PRIMARY KEY,
		name TEXT,
		image_path TEXT UNIQUE,
		current_amount INT,
		required_amount INT,
		user_id INT,

		FOREIGN KEY(user_id)
		REFERENCES users(user_id)
	);`

	CreateWalletQuery = `CREATE TABLE IF NOT EXISTS wallet(
		wallet_id SERIAL PRIMARY KEY,
		amount INT,
		user_id INT UNIQUE,

		FOREIGN KEY(user_id)
		REFERENCES users(user_id)
	);`

	CreateSummaryQuery = `CREATE TABLE IF NOT EXISTS summary(
		summary_id SERIAL PRIMARY KEY,
		date_start DATE,
		date_end DATE,
		spent INT,
		earned INT,
		wallet_id INT,

		FOREIGN KEY(wallet_id)
		REFERENCES wallet(wallet_id)
	);`

	CreateCategoryQuery = `CREATE TABLE IF NOT EXISTS category(
		category_id SERIAL PRIMARY KEY,
		name VARCHAR(30)
	);`

	CreateExpenseQuery = `CREATE TABLE IF NOT EXISTS expense(
		expense_id SERIAL PRIMARY KEY,
		amount INT,
		date DATE,
		category_id INT,
		wallet_id INT,

		FOREIGN KEY (category_id) 
		REFERENCES category(category_id),

		FOREIGN KEY (wallet_id)
		REFERENCES wallet(wallet_id)
	);`

	CreateIncomeQuery = `CREATE TABLE IF NOT EXISTS income(
		income_id SERIAL PRIMARY KEY,
		amount INT,
		date DATE,
		category_id INT,
		wallet_id INT,

		FOREIGN KEY (category_id) 
		REFERENCES category(category_id),

		FOREIGN KEY (wallet_id)
		REFERENCES wallet(wallet_id)
	);`
)
