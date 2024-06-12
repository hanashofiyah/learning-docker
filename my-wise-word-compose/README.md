1. Prepare a single database "{yourname}_access", single table "access_log" with single column named "timestamp" with data type timestamp or datetime
    ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/6576473e-e5eb-482d-83a4-b108fd7e551e)
    ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/e8b5a75d-d7da-4d13-ab3a-f2f31e5b8b62)
    ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/f4315c48-3f77-4f71-b9f4-3abc368de99c)


3. Create or use existing golang project that serve http and connect to postgres (or any) databases.
- Path "/", print html text of your favorite wise word
    - Also when access to "/" print to console using fmt.Println with text "Ouch!"
    ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/11534a13-ae67-48ec-a718-d28ce7d73c7a)
- Path "/ping"-> Do ping to database and show the result (success connect or failed) in browser
  ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/2967c050-40bd-4a78-992f-1bf7955b81fd)
- Expose to port 78
  ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/7c647408-a723-43f0-9d1d-46a101e0d0c7)


3. Create Dockerfile to build and run your golang project
    ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/5ec3da9c-79e2-476d-b721-100f631bf19c)
  
4. Create docker-compose.yml and instruct to run 2 containers

    - Golang Project (give name: "go_service"), expose container port 78 to host port 9911
  ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/1c26435f-e3f3-4f4c-bcb0-28e821e6b1ae)
    - Postgres image postgres:latest (give name: "database")
  ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/8f7a439b-2bf4-4204-bb06-f7d36462c291)

    Command for make container : **docker compose up**
    **Result from container in Docker Dekstop**
    ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/75ba7fe7-920e-4075-99de-65d5df56efd8)



5. Make sure the golang project can connect to postgresql and postgresql data is persistent
    - network name="netwwc" and volume name = "volwwc”
  ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/c92a9804-f603-4dfb-a335-6e8b1870e8ac)

6. Run Docker compose with name "my-wise-word-compose"
   ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/d7f05cc6-7261-4bb0-8e13-a8fa4ec80934)

    - Make sure you can access http://localhost:9911 and show a wise word
   ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/aa6448cc-4d14-4216-9e1c-bdf016bb9f18)

    - Make sure you can access http://localhost:9911/ping and show Success (with success ping to database) and new data inserted in your     database
    ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/f4a03da1-1a5d-40a3-af4c-e694ae146480)


8. Check logs of your each container in "my-wise-word-compose"
   **go-service**
   ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/90cd6f4f-f68c-4ea5-9166-479fa1df729e)
   **database**
   ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/2d89f8a1-f1aa-4491-8842-6f841b8616e9)

