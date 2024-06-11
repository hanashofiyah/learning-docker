1. Hapus container postgres yang lama
   ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/05340da8-9586-4b3c-b1e3-3191d74edd91)

2. Jalankan postgres dengan nama container "my-postgres-{username}" dengan env dan volume "my-pg-volume-{username}"
   **docker run -d --name my-postgres-hana -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -v my-pg-volume-hana:/var/lib/postgresql/data -p 5431:5432 postgres**
   ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/259e273f-1b8f-474f-a597-b7c9885fe7e8)
   ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/e37076e8-6d7d-4136-9d3d-c693089fb941)

4. Pastikan bisa sukses terkoneksi
   ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/4e29bf6f-78d8-4986-9856-e0065ab1deb8)
   ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/8bb63815-6b44-45af-a14c-4741bac96d16)
   ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/61dd820f-71a3-4898-bd9b-2b357b8e4027)

6. Buat table baru via pgadmin dengan nama github "{username}"
   ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/0e9b5250-9814-464c-b970-7bcb55012fe2)
   ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/a2aae009-75be-4508-87fc-4478a951b7b4)


8. Stop dan hapus container postgres lama "my-postgres-{username}"
   ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/f1696420-ff9b-438a-baed-5198e9391d23)

9. Jalankan ulang docker postgres dengan nama baru "my-postgres-v2-{username}" env dan volume diatas
   **docker run -d --name my-postgres-hana -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -v my-pg-volume-hana:/var/lib/postgresql/data -p 5431:5432 postgres**
   ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/9a403654-c28b-43f8-8800-69b139986f47)
   ![image](https://github.com/hanashofiyah/learning-docker/assets/104729134/bd4327cb-c258-4688-9249-1d20b74359de)

