#include <iostream>
#include <libpq-fe.h>

void checkConnection(PGconn* conn) {
    if (PQstatus(conn) != CONNECTION_OK) {
        std::cerr << "Connection to database failed: " << PQerrorMessage(conn) << std::endl;
        PQfinish(conn);
        exit(1);
    }
}

void fetchUsers(PGconn* conn) {
    const char* query = "SELECT id, username, email, created_at FROM users";
    PGresult* res = PQexec(conn, query);

    if (PQresultStatus(res) != PGRES_TUPLES_OK) {
        std::cerr << "Failed to execute query: " << PQerrorMessage(conn) << std::endl;
        PQclear(res);
        PQfinish(conn);
        exit(1);
    }

    int rows = PQntuples(res);
    int cols = PQnfields(res);

    // Print column headers
    for (int col = 0; col < cols; ++col) {
        std::cout << PQfname(res, col) << "\t";
    }
    std::cout << std::endl;

    // Print rows
    for (int row = 0; row < rows; ++row) {
        for (int col = 0; col < cols; ++col) {
            std::cout << PQgetvalue(res, row, col) << "\t";
        }
        std::cout << std::endl;
    }

    PQclear(res);
}

int main() {
    const char* conninfo = "host=localhost dbname=yourdb user=youruser password=yourpassword";
    PGconn* conn = PQconnectdb(conninfo);

    checkConnection(conn);

    std::cout << "Connected to the database successfully!" << std::endl;
    std::cout << "Fetching users:" << std::endl;

    fetchUsers(conn);

    PQfinish(conn);
    std::cout << "Connection closed." << std::endl;

    return 0;
}
