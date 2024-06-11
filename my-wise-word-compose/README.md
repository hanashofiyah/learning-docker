1. Prepare a single database "{yourname}_access", single table "access_log" with single column named "timestamp" with data type timestamp or datetime
![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/6576473e-e5eb-482d-83a4-b108fd7e551e)
![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/e8b5a75d-d7da-4d13-ab3a-f2f31e5b8b62)

3. Create or use existing golang project that serve http and connect to postgres (or any) databases.
- Path "/", print html text of your favorite wise word
    - Also when access to "/" print to console using fmt.Println with text "Ouch!"
- Path "/ping"-> Do ping to database and show the result (success connect or failed) in browser
- Expose to port 78

3. Create Dockerfile to build and run your golang project
  
4. Create docker-compose.yml and instruct to run 2 containers
- Golang Project (give name: "go_service"), expose container port 78 to host port 9911
- Postgres image postgres:latest (give name: "database")

5. Make sure the golang project can connect to postgresql and postgresql data is persistent
- network name="netwwc" and volume name = "volwwc”
- You can check via pgadmin and expose the port 5432 (bebas) to host.

6. Run Docker compose with name "my-wise-word-compose"
- Make sure you can access http://localhost:9911 and show a wise word
- Make sure you can access http://localhost:9911/ping and show Success (with success ping to database) and new data inserted in your database

7. Check logs of your each container in "my-wise-word-compose"
