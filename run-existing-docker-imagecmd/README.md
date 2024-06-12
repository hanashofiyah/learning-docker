1. Install Docker.
   **https://docs.docker.com/get-docker/**
   ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/eba72915-fc13-437b-b45f-2caf829b1ca0)
   ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/e77d5ca0-5390-4b8f-8586-a24e86ac66e7)

2. Run a new Container.
   via CMD
   **docker run -d -p 8080:80 --name welcome1 docker/welcome-to-docker**
   ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/3996bb52-888f-4ee1-8f1e-ad7a731d3a25)

   via Docker Dekstop
    ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/0fafd271-f9e5-499b-a942-f63056a0de6d)

3. Try in your browser to access.
   - Open URL **http://localhost:8080**
  ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/c8a01f15-6b57-4dc7-bc44-9bbde8bf05eb)
  
4. See logs of running welcome1 container.
   Command is used to list all running Docker containers: **docker container ls**
   ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/bed88769-6834-469a-8fbf-89d5838ef6a4)
   Command is used to view the logs of a specific Docker container: **docker logs -f --tail 10 welcome1**
   ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/81a753f2-ac3e-46a0-9337-e2f8b94320a8)

5. execute a command inside running welcome1 container.
  **docker exec welcome1 cat ../usr/share/nginx/html/index.html**
   ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/9889b5af-7817-4872-8f7b-737464d7c81b)


6. Stop Container and see latest logs.
   Command is used for stop container: **docker stop welcome1**
   ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/d841aae5-71de-46a7-bbb7-2c607d685bd6)
   Command is used for see latest logs: **docker logs -f --tail 10 welcome1**
   ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/2839ae67-bbcf-4748-9441-9f194a2b745f)
