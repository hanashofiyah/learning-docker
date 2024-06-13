A. Prepare a single database "{username}_access",
  A. single table "{username}_access_log"
  B. with single column named "timestamp" with data type timestamp or datetime
  ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/0b8796f7-886a-4e3b-93c4-bd6792339f23)
  ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/c563d6ab-41aa-4f7f-86c4-88291241fa48)
  ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/90a738be-6dc2-45cc-a9eb-e85a791fa6a9)

B. Use existing golang project (from Project 1) that serve http and try to connect to postgres (or any) databases with additional:
  - Path "/ping", Do ping to database and show the result
  - Output or return to browser as string or html Success = "Pong!", Failed = "Ouch!”
    ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/1f9fe2e5-27a0-44b1-b5d0-0a86a4312730)
  - Print to console or log Success = "Ping successful", Failed = "Ping failed"
  - Insert or record a new data timestamp with "NOW()" or current timestamp whenever /ping accessed. E.g., new record current timestamp 2024-01-31 19:55:55
  - Start HTTP on port 77 (or 78 if your chrome blocked the port)

C. Create Dockerfile to build and run your golang project
  ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/9ae19a1f-f3ef-4761-8213-cd39fa578c10)

D. Create docker-compose.yml and instruct to run 2 containers
  - Golang Project (container name: "gosvc_container"), expose container port 77 (or 78) to host port 4331 Postgres image postgres:latest (container name: "dbsvc_container")
  - Pastikan env postgres database "{username}_access"

E. Make sure the golang project can connect to postgresql and postgresql data is persistent 
  - network name = "net_mycompose_{username}" and volume name = "vol_mycompose_{username}"
  - You can check via pgadmin and expose the port 5432 (bebas) to host.

F. Run Docker compose with name "project-my-compose" 
  - Make sure you can access http://localhost:4331 and show a wise word
  - Make sure you can access http://localhost:4331/ping and show Success (with success ping to database) and new data inserted in your database

G. Check logs of your each container in "project-my-compose"

H. Masukkan semua files project & screenshot hasilnya ke repository github kalian
{username}/learning-docker Folder project-my-compose
