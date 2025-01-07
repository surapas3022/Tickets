# Tickets
Learning golang modules


#generate go.mod

#go mod init "domainName"
#e.g.
go mod init github.com/surapas3022/Tickets   #Or the name your project is case example with Tickets



#command run at project path
go run main.go


#Clean the Go Module Cache:
go clean -modcache


#Download All Dependencies:
go mod tidy


#Verify Version 
go version



#%.2f = decimal 2 positions


#========================
#for call use mod

#the name your project
go mod init callmodule   

#in main.go file in call use module 
import (
	"github.com/surapas3022/Tickets/movie"
)  
#Or the name your project is case example with Tickets
go mod tidy