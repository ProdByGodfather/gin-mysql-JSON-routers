# gin-mysql-JSON-routers

A simple project with delete, single view, global view and edit apis connected to mysql database

## how to work

1. first you need change your `mysql` connection string.
    - line 26:
    ```go
    var err error
    var db *sql.DB
    db, err = sql.Open("mysql", "[mysql username]:[mysql password]@tcp(localhost:3306)/[your db name]")
    ```

2. Create a table in your database with these columns:
   - id (`INTEGER` `PRIMARY KEY` `NOT NULL` `AUTO ENCREMENT`)
   - name (`VARCHAR`)
   - price (`FLOAT`)
   - description(`VARCHAR` OR `TEXT`)
   
     warning: Choose the name of your database preferably `products`
   