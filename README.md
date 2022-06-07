# tuiter

# github
pasos:
- agrego en mi proyecto el main.go y lo seteo con el package y func main
- go mod init
- go mod tidy
- git init
- git add .
- git commit -m "first commit"
- git remote add origin https://github.com/benjacifre10/tuiter.git
- git push -u origin master

# heroku
pasos:
- voy al root de mi proyecto
- heroku login
- voy a heroku y le agrego en settings -> Addbuildpack -> go (unica vez)
- heroku git:remote -a tuitorheroku (unica vez)
- git push heroku master

# packages
pasos:
- comando go get + package a instalar
