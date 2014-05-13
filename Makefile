APPNAME:=quint

default: clean
	script/release

clean:
	rm -f ./$(APPNAME).*
