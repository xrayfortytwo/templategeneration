test_run:
	rm -f ./tmp/*.java
	go install 
	templategeneration

clean:
	rm -f ./tmp/*.java
