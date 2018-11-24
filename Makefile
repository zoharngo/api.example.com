GCC=go
GCMD=run
GPATH=main.go
GPORT=3000

run:
	$(GCC) $(GCMD) $(GPATH) -port=${GPORT}
