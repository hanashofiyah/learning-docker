1. Create a new Golang project that serve http
  - Path "/", do print string/html text of your favorite wise word
  - When you access path "/", also print to log/console using fmt.Println with text "Ouch!" Start HTTP on port 77
  ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/09964305-4026-4418-afc6-8c0bd704882c)

2. Create file "AUTHORS.md" and "LINKS.md" within same folder with "Dockerfile
   - Edit "AUTHORS.md" file and fill it with your first name or your fake name
     ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/7b906fbd-b3e6-4daa-9809-622658f13668)

   - Edit "LINKS.md" file and fill it with your github profile link
    ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/28a2ef0c-ee62-47d0-af7b-94cc1bf2fd82)

3. Create Dockerfile to build and run your golang project
    - Use Base Image "golang:1.21"
    - Set WORKDIR with /myapp
    - RUN some command "go version" after WORKDIR
    - COPY "AUTHORS.md" to image BEFORE run build golang (go build) COPY "LINKS.md" to image BEFORE run build golang (go build)
    - Make golang build output with name "my-go-app" and make sure it will run correctly
      ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/c97d2e77-f277-4496-8cef-d0c3adf95008)

4. Build Dockerfile image with name "my-go-app:v2"
   **docker run --name go-app -p 5555:77 my-go-app:v2**
   ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/359a112f-916c-4522-962f-f3b8aa908765)

5. Run the image into container with name "go-app" and expose to port host 5555
   **docker build -t my-go-app:v2 .**
   ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/08bf768b-1aab-4249-8ecf-e9384ada3f56)
    ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/be7e98cb-2984-45c4-88d1-24d9acd588d3)


6. Access your golang inside docker via localhost:5555 and see logs of your container "go-app"
   ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/176b88e7-ca8f-40d6-9d3d-4f52b7d2fa4e)

