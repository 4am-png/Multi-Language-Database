import psycopg2
from psycopg2 import sql

def connect_to_db():
    try:
        connection = psycopg2.connect(
            host="localhost",
            database="yourdb",
            user="youruser",
            password="yourpassword"
        )
        return connection
    except psycopg2.Error as e:
        print(f"Error connecting to the database: {e}")
        exit(1)

def fetch_users(connection):
    try:
        with connection.cursor() as cursor:
            query = sql.SQL("SELECT id, username, email, created_at FROM users")
            cursor.execute(query)
            rows = cursor.fetchall()
            if rows:
                # Print column headers
                print(f"{'ID':<5}{'Username':<20}{'Email':<30}{'Created At':<20}")
                print("=" * 75)
                # Print rows
                for row in rows:
                    print(f"{row[0]:<5}{row[1]:<20}{row[2]:<30}{row[3]:<20}")
            else:
                print("No users found.")
    except psycopg2.Error as e:
        print(f"Error executing query: {e}")

def main():
    connection = connect_to_db()
    print("Connected to the database successfully!")
    print("Fetching users:")
    fetch_users(connection)
    connection.close()
    print("Connection closed.")

if __name__ == "__main__":
    main()
