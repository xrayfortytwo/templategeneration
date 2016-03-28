test_run:
	rm -f ./tmp/*.java
	go install 
	templategeneration
