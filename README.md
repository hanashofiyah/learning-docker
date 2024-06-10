# learning-docker

- Install Docker
  https://docs.docker.com/get-docker/
  ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/00bda554-4aef-448a-8611-5a3cdcf26421)

- Run a new container
  1. Melalui CMD
  **docker run -d -p 8080:80 --name welcome1 docker/welcome-to-docker**
  ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/7e345194-50c6-409b-8a5b-8258fa9603e8)
  2. Melalui Docker Windows
  ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/2c44221c-45ed-40f6-9978-6a1e7df997b6)

- Try in your browser to access apps inside welcome-to-docker image
  http://localhost:8080
  ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/7e7b1678-a9fd-4c28-ad05-8a04a0695ee8)

- See logs of running welcome1 container
  **docker logs -f --tail 10 welcome1**
  ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/bf718638-b8b4-41bb-a156-6dfbd3da3da8)

- Execute a command inside running welcome1 container
  **cat usr/share/nginx/html/index.html**
  ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/fffed33d-3887-4851-b078-22f0ba20a940)

- Stop container and see latest logs
  **docker stop welcome1**
  ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/d23a23d5-3529-4872-98a6-187099898db0)
  **docker logs -f --tail 10 welcome1**
  ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/5c1ca496-1869-43b9-9999-4068348c8a96)





