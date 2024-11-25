# GO.AWS_EC2

## Project Overview

This project is a Go application designed for deployment on an AWS EC2 instance. It utilizes **PM2** for process management and **GitHub Actions** for continuous deployment. The workflow ensures that the application is automatically deployed and running in the background whenever new changes are pushed to the `main` branch.

---

## Features

- **Go Application**: A robust and scalable application written in Go.
- **PM2 Integration**: Efficient process management for seamless background execution.
- **CI/CD Workflow**: Automated deployments using GitHub Actions.

---

## Requirements

### Local Environment
- Go 1.18 or later
- Git

### Server Environment (EC2 Instance)
- PM2 installed globally:
  ```bash
  npm install pm2 -g
  ```
- Go runtime installed on the server:
  ```bash
  sudo yum install golang -y
  ```

---

## Setup

### 1. Clone the Repository
Clone the repository onto your local machine or directly on your EC2 instance:
```bash
git clone https://github.com/your-repo/GO.AWS_EC2.git
cd GO.AWS_EC2
```

### 2. Install Dependencies
Ensure all Go modules are initialized:
```bash
go mod tidy
```

### 3. PM2 Configuration

To run locally:
```go
go run main.go
```

To run the Go application with PM2:
```bash
pm2 start go --name "my-app" -- run main.go
pm2 save
```

---

## GitHub Actions Workflow

### Workflow Description

The included GitHub Actions workflow automatically deploys the Go application to your EC2 instance when code is pushed to the `main` branch.

#### Key Steps:
1. **Check Out the Code**: Fetches the latest code from the repository.
2. **Deploy to EC2**: Uses the `easingthemes/ssh-deploy` action to transfer files to the EC2 instance.
3. **Execute Remote Commands**:
   - Stops any running instance of the application using PM2.
   - Pulls the latest code from the `main` branch.
   - Restarts the application in the background with PM2.

### Workflow File
```yaml
name: Deploy to EC2

on:
  push:
    branches:
      - main

jobs:
  deploy:
    runs-on: ubuntu-latest

    steps:
    - name: Check out the code
      uses: actions/checkout@v3

    - name: Deploy to EC2
      uses: easingthemes/ssh-deploy@main
      env:
        SSH_PRIVATE_KEY: ${{ secrets.EC2_SSH_KEY }}
        REMOTE_HOST: ${{ secrets.HOST_DNS }}
        REMOTE_USER: ${{ secrets.EC2_USER }}
        TARGET: ${{ secrets.TARGET_DIR }}

    - name: Executing remote ssh commands using ssh key
      uses: appleboy/ssh-action@master
      with:
        host: ${{ secrets.HOST_DNS }}
        username: ${{ secrets.EC2_USER }}
        key: ${{ secrets.EC2_SSH_KEY }}
        script: |
          pm2 delete my-app || echo "No existing process found"
          cd GO.AWS_EC2
          git pull origin main
          pm2 start go --name "my-app" -- run main.go
          pm2 save
          echo "Application is running in the background."
```

---

## How to Use

### Local Testing
1. Run the Go application locally:
   ```bash
   go run main.go
   ```
2. Test endpoints or functionalities to ensure proper behavior.

### Deployment
1. Push changes to the `main` branch:
   ```bash
   git add .
   git commit -m "Update application"
   git push origin main
   ```
2. GitHub Actions will automatically deploy the application to the EC2 instance.

### Verify the Deployment
1. Log in to your EC2 instance:
   ```bash
   ssh -i your-key.pem ec2-user@your-ec2-ip
   ```
2. Check the PM2 process list:
   ```bash
   pm2 list
   ```
3. View application logs (optional):
   ```bash
   pm2 logs my-app
   ```

---

## Troubleshooting

### GitHub Actions Deployment Fails
- Confirm that the correct secrets are set up in the repository settings (`EC2_SSH_KEY`, `HOST_DNS`, `EC2_USER`, and `TARGET_DIR`).

## Take a look on my deploy
You can see my project running on aws ec2 by this url: http://3.209.237.116:8080/

The project is running on an AWS Academy environment, so if you got no response for the web app it's probably because the aws academy environment isn't running.

## License
This project is licensed under the MIT License :D, so do what you need with the code.