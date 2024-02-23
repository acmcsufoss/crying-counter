import sqlite3


class Database:

    def __init__(self, dbfile: str):
        self._conn = sqlite3.connect(dbfile)
        self._cursor = self._conn.cursor()

    def createTable(self):
        #to create user_counts table
        self.execute('CREATE TABLE IF NOT EXISTS Crying (user_id INTEGER PRIMARY KEY, count INTEGER DEFAULT 0)')

    #same deal here don't remember the syntax
    #replaces value, incorporate into the count_message function of HelloDiscordBot
    def increment_count(self, user_id: int, count: int):
        self.execute('INSERT OR REPLACE INTO Crying (user_id, count) VALUES (?,?)', (user_id, count))
        self.commit()

    def __enter__(self):
        return self

    def __exit__(self, exc_type, exc_val, exc_tb):
        self.close()

    @property
    def connection(self) -> sqlite3.Connection:
        return self._conn

    @property
    def cursor(self) -> sqlite3.Cursor:
        return self._cursor

    def commit(self) -> None:
        self.connection.commit()

    def close(self, commit=True) -> None:
        if commit:
            self.commit()
        self.connection.close()

    def execute(self, sql, params=()) -> None:
        self.cursor.execute(sql, params)

    #from my understanding retrives the result of the most reset query
    def fetchone(self) -> tuple:
        return self.cursor.fetchone()

    def fetchall(self) -> list[tuple]:
        return self.cursor.fetchall()