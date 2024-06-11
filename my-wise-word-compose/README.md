A. Prepare a single database "{yourname}_access", single table "access_log" with single column named "timestamp" with data type timestamp or datetime
B. Create or use existing golang project that serve http and connect to postgres (or any) databases.
Path "/", print html text of your favorite wise word
Also when access to "/" print to console using fmt.Println with text "Ouch!"
Path "/ping" -> Do ping to database and show the result (success connect or failed) in browser
Also when access to "/ping" Insert or record "NOW()" or timestamp whenever /ping accessed. E.g., new record current timestamp 2024-01-31 19:55:55
Expose to port 78
C. Create Dockerfile to build and run your golang project
D. Create docker-compose.yml and instruct to run 2 containers
Golang Project (give name: "go_service"), expose container port 78 to host port 9911 Postgres image postgres:latest (give name: "database")
E. Make sure the golang project can connect to postgresql and postgresql data is persistent
network name="netwwc" and volume name = "volwwc”
You can check via pgadmin and expose the port 5432 (bebas) to host.
F. Run Docker compose with name "my-wise-word-compose" Make sure you can access http://localhost:9911 and show a wise word
Make sure you can access http://localhost:9911/ping and show Success (with success ping to database) and new data inserted in your database
G. Check logs of your each container in "my-wise-word-compose" H. Masukkan files & screenshot hasil ke repository github kalian {username}/learning-docker Folder my-wise-word-compose
