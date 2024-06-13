A. Prepare a single database "{username}_access",
  A. single table "{username}_access_log"
  B. with single column named "timestamp" with data type timestamp or datetime
  ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/0b8796f7-886a-4e3b-93c4-bd6792339f23)
  ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/c563d6ab-41aa-4f7f-86c4-88291241fa48)

B. Use existing golang project (from Project 1) that serve http and try to connect to postgres (or any) databases with additional:
  - Path "/ping", Do ping to database and show the result
    ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/440bc9f8-319c-4e3a-9829-990cbec9a61c)
    ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/a6d6a0bb-7c20-4c20-97e9-8d87cf947d45)

  - Output or return to browser as string or html Success = "Pong!", Failed = "Ouch!”
  - Print to console or log Success = "Ping successful", Failed = "Ping failed"
    
  - Insert or record a new data timestamp with "NOW()" or current timestamp whenever /ping accessed. E.g., new record current timestamp 2024-01-31 19:55:55
    ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/ae13630b-7870-4ff0-82fe-6e991f0380dd)
    
  - Start HTTP on port 77 (or 78 if your chrome blocked the port)
    ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/9513337c-1d2d-4d26-872b-b9ea7839e337)

C. Create Dockerfile to build and run your golang project
**dockerfile**
![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/512d014f-1c59-4bf7-baff-c694b337bb37)

**build**
![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/71ae471d-fe71-45dc-924d-42c45b6298d2)

**run**
![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/b5363b30-fb65-4bfc-9f80-1c65f2c076d5)

D. Create docker-compose.yml and instruct to run 2 containers
  - Golang Project (container name: "gosvc_container"), expose container port 77 (or 78) to host port 4331 Postgres image postgres:latest (container name: "dbsvc_container")
    ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/8f52f8fd-1d83-40f5-aebd-c2fee89248af)

  - Pastikan env postgres database "{username}_access"
    ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/76d40a0d-420b-4dc0-b97e-d219befc495a)

E. Make sure the golang project can connect to postgresql and postgresql data is persistent 
  - network name = "net_mycompose_{username}" and volume name = "vol_mycompose_{username}"
    ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/b2f645ad-1d4c-4186-a9c7-13ce5bd4bf6a)

  - You can check via pgadmin and expose the port 5432 (bebas) to host.
    ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/90a738be-6dc2-45cc-a9eb-e85a791fa6a9)

F. Run Docker compose with name "project-my-compose" 
  - Make sure you can access http://localhost:4331 and show a wise word
   ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/d9a1f621-3e05-4d0a-8396-97c311cc1a1e)
  - Make sure you can access http://localhost:4331/ping and show Success (with success ping to database) and new data inserted in your database
    ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/7d1f6361-7a6d-47fd-8525-88fc0f27474f)

G. Check logs of your each container in "project-my-compose"
  ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/87be23c9-e436-4f12-95e2-88d278800bf8)
  ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/485c0012-e9cf-41f8-89bc-9e2352841124)
