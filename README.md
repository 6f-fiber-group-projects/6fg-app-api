# github.com/6f-fiber-group-projects/6fg-app-api

## migration
* dev
  ```
  goose postgres "host=db user=root password=root dbname=api sslmode=disable" up
  ```
  ```
  heroku run cd app/migration && goose up
  ```