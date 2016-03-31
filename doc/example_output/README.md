#Other examples:

##Calculation language example:

###Grammar:
```
grammar Calc ;

@header {
    // calc is for calculating

}

calc : expre* EOF ;
expre : num0 op num1 (op num2)? '=' num3 ;
num0 : NUMBER ;
num1 : NUMBER ;
num2 : NUMBER ;
num3 : NUMBER ;
ob : OP ;
OP : (*|+|-|) ;
NUMBER : [0-9]+ ;

```
###Generated Template:

```
calc(calc) ::= "<\n><calc.expre:{z|<expre(z)>}><\n>"
expre(expre) ::= "<\n><expre.num0><expre.op><expre.num1><if(expre.num2)><expre.op><expre.num2><endif> = <expre.num3>"
num0(num0) ::= "<\n><num0.ID>"
num1(num1) ::= "<\n><num1.ID>"
num2(num2) ::= "<\n><num2.ID>"
num2(num3) ::= "<\n><num3.ID>"
ob(ob) ::= "<\n><ob.ID>"
```

###Generated classes:

```java
package tmp;

public class Expre {

    String num0;
    String op;
    String num1;
    String num2;
    String num3;

}

```


```java
package tmp;

public class Calc {

    ArrayList<Expre> expre;

}
```

##Hello language example:

###Grammar:

```
grammar Hello ;

subscript : descript* EOF ;
descript : 'surname' '=' surname ',' (id'.'company '@' email)? ;
company : NAME ;
email : NAME
name : NAME ;
surname : NAME ;
NAME : [a-z]+ ;
```

###Generated template:

```
subscript(subscript) ::= "<\n><subscript.descript:{z|<descript(z)>}><\n>"
descript(descript) ::= "<\n> surname  = <descript.surname> , <if(descript.email)><descript.id> . <descript.company> @ <descript.email><endif>"
company(company) ::= "<\n><company.ID>"
surname(surname) ::= "<\n><surname.ID>"
```

###Generated classes:
```java
public class Subscript {

    ArrayList<Descript> descript;

}
```

```java
package tmp;

public class Descript {

    String surname;
    String id;
    String company;
    String email;

}
```
