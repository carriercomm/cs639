LBITS := $(shell getconf LONG_BIT)
ifeq ($(LBITS),64)
gc=6g
gl=6l
su=6
else
gc=8g
gl=8l
su=8
endif

all: put sfscat get sfsls SFShell

put: put.$(su) ../client/client.go
	$(gl) -o put put.$(su)

put.$(su): put.go
	$(gc) put.go

sfscat: sfscat.$(su) ../client/client.go
	$(gl) -o sfscat sfscat.$(su)

sfscat.$(su): sfscat.go
	$(gc) sfscat.go

SFShell: SFShell.$(su) ../client/client.go
	$(gl) -o SFShell SFShell.$(su)

SFShell.$(su): SFShell.go
	$(gc) SFShell.go

get: get.$(su) ../client/client.go
	$(gl) -o get get.$(su)

get.$(su): get.go
	$(gc) get.go
	
	
sfsls: sfsls.$(su) ../client/client.go
	$(gl) -o sfsls sfsls.$(su)

sfsls.$(su): sfsls.go
	$(gc) sfsls.go

clean:
	-rm -f *.$(su) put sfscat get sfsls SFShell

clean-all: clean
