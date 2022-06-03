# tuiter

# github
pasos:
- agrego en mi proyecto el main.go y lo seteo con el package y func main
- go mod init
- git init
- git add .
- git commit -m "first commit"
- git remote add origin https://github.com/benjacifre10/tuiter.git
- git push -u origin master

# heroku
pasos:
- voy al root de mi proyecto
- heroku login
- voy a heroku y le agrego en settings -> Addbuildpack -> go
- heroku git:remote -a tuitorheroku

- git add .
- git commit -am "make it better"
- git push heroku master
