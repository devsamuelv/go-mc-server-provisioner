# Setting Up The Storage Server

- Getting Started
    - [ ] Getting The System Ready
        - [ ] Install All The Drives Needed
        - [ ] Format the Drives with fdisk using the linux or windows file system

    - [ ] Install 🐋 Docker
        - [ ] Remove Old Versions `sudo apt-get remove docker docker-engine docker.io containerd runc`
        - [ ] Update Linux `sudo apt-get update`
        - [ ] Allow Apt To Use Https `sudo apt-get install \ apt-transport-https \ ca-certificates \ curl \ gnupg-agent \ software-properties-common`
        - [ ] Add 🐋 Dockers Repo With GPG `curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -`
        - [ ] Check The Finger Print `sudo apt-key fingerprint 0EBFCD88`
        - [ ] Make Sure The Finger Print is `9DC8 5822 9FC7 DD38 854A  E2D8 8D81 803C 0EBF CD88`
        - [ ] Add 🐋 Docker Repo `sudo add-apt-repository \ "deb [arch=amd64] https://download.docker.com/linux/ubuntu \ $(lsb_release -cs) \ stable"`
        - [ ] Update Linux `sudo apt-get update`
        - [ ] Install 🐋 Docker Engine `sudo apt-get install docker-ce docker-ce-cli containerd.io`
        - [ ] Make Sure 🐋 Docker is Working by running 🐋 Docker Hello World `sudo docker run hello-world`