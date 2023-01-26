package mysql

import (
	"database/sql"

	"github.com/ardianeffendi/snippetbox/pkg/models"
)

// Define a SnippetModel type which wraps a sql.DB connection pool.
type SnippetModel struct {
	DB *sql.DB
}

// This will insert a new snippet into the database.
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	// Write the SQL statement we want to execute. It's split over two lines
	// for readability (the reason being why it's surrounded with backquotes
	// instead of normal double quotes).
	stmt := `INSERT INTO snippets (title, content, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	// Use the exec() method on the embedded connection pool to execute the statement.
	// The first parameter is the SQL statement, followed by the title, content
	// and expiry values for the placeholder parameters. This method returns
	// a sql.Result object, which containts some bacic information about what
	// happened when the statement was executed.
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	// Use the LastInsertId() method on the result object to get the ID of our
	// newly inserted recrod in the snippets table.
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// The ID returned has the tyep int64, so we convert it to an int type
	// before returning
	return int(id), nil
}

// This will return a specific snippet based on its id.
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// This will return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
