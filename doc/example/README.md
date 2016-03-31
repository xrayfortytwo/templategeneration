#Small example of usage

This directory contains a small example of how the application can be used.

It is possible to build a binary and go to the src dir of a java project. Then execute:

    templategeneration -filename=<path to *.g4 file>

The grammar will be processed and the simple java source files will be written into a directory called tmp.
These classes can than be added, thus code can be produced.

Example of a simple generated java class:

```java
package tmp;

public class Transition {

    String event;
    String action;
    String stateid;
}
```

Example of finally generated code in this case FSML:
```
initial state state_0 { 
    event0_0 / action0_1 -> state_1 ; 
    event0_1 ; 
    event0_2 / action0_11 -> state_1 ; 
  } 
state state_1 { 
    event1_1 / action1_2 -> state_2 ; 
    event1_2 -> state_3 ; 
  } 
state state_2 { 
    event2_1 / action2_1 -> state_2 ; 
    event2_2 -> state_3 ; 
  } 
```


[Some other examples](https://github.com/xrayfortytwo/templategeneration/tree/master/doc/example_output)
