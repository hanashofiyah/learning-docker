1. Create a new Golang project that serve http
  - Path "/", do print string/html text of your favorite wise word
  - When you access path "/", also print to log/console using fmt.Println with text "Ouch!" Start HTTP on port 77
    ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/0d70e673-68f4-4c41-a3fa-824a576c91a5)


2. Create file "AUTHORS.md" and "LINKS.md" within same folder with "Dockerfile
   - Edit "AUTHORS.md" file and fill it with your first name or your fake name
     ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/7b906fbd-b3e6-4daa-9809-622658f13668)

   - Edit "LINKS.md" file and fill it with your github profile link
     ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/62565226-30d7-4eb4-b98c-d1a5a48b3d1c)


3. Create Dockerfile to build and run your golang project
    - Use Base Image "golang:1.21"
    - Set WORKDIR with /myapp
    - RUN some command "go version" after WORKDIR
    - COPY "AUTHORS.md" to image BEFORE run build golang (go build) COPY "LINKS.md" to image BEFORE run build golang (go build)
    - Make golang build output with name "my-go-app" and make sure it will run correctly
      ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/5d09aa4d-7ca2-4876-b5b0-08382bfd8627)


4. Build Dockerfile image with name "my-go-app:v2"
   ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/9d63e9fd-6c9f-496d-aa14-78da6670409b)


6. Run the image into container with name "go-app" and expose to port host 5555

7. Access your golang inside docker via localhost:5555 and see logs of your container "go-app"
